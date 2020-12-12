package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/lestrrat-go/strftime"
	"github.com/yargevad/filepathx"
)

const (
	//INTERVAL - seconds between pictures
	INTERVAL = 5
	//PATH - where images are stored
	PATH = "/home/pi/recordings"
)

func ginMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

// take a picture
func capture(f *strftime.Strftime, liveFeed chan string) {
	log.Println("Running capture.sh...")
	hour, _, _ := time.Now().Clock()
	filename := fmt.Sprintf("%s/%02d/%s.jpg", PATH, hour, f.FormatString(time.Now()))
	out, err := exec.Command("/home/pi/capture/capture.sh", filename).CombinedOutput()
	if strings.TrimSpace(string(out)) != "" {
		log.Println(string(out))
	}
	if err != nil {
		log.Println(err)
	}
	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
		// do this in prod
		liveFeed <- filename
	}
}

func loop(interval time.Duration, liveFeed chan string) {
	ticker := time.NewTicker(interval)
	f, err := strftime.New("%Y_%m_%d_%H_%M_%S")
	if err != nil {
		panic(err)
	}
	for {
		<-ticker.C
		//spawn a process to capture the picture in the background
		go capture(f, liveFeed)
	}
}

func sendFeedUpdates(liveFeed chan string, sio *socketio.Server) {
	for {
		path := <-liveFeed
		sio.BroadcastToRoom("", "users", "live", path)
	}
}

func web() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join("users")
		return nil
	})
	// channel for new images from the camera
	liveFeed := make(chan string, 0)
	// start taking pictures
	//go loop(INTERVAL*time.Second, liveFeed)
	// make the socket.io handler for these messages
	go sendFeedUpdates(liveFeed, server)
	go server.Serve()
	defer server.Close()

	r := gin.Default()
	//r.Use(cors.Default())
	// r.Use(ginMiddleware("http://localhost:3000"))

	//serve vue content
	r.Static("/vue", "../dist/")

	// Sanity GET request
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// hostname
	r.GET("/hostname", func(c *gin.Context) {
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"hostname": name,
		})
	})
	// Get a list of images on the filesystem, newest first
	r.GET("/images", func(c *gin.Context) {
		files, err := filepathx.Glob("/home/pi/recordings/*/*.jpg")
		if err != nil {
			panic(err)
		}
		type Match struct {
			name string
			date time.Time
		}

		matches := make([]*Match, len(files))
		for i, f := range files {
			path, file := filepath.Split(f)
			if strings.Contains(path, "tmp") {
				//skip tmp pictures
				continue
			}
			base := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(file, "mon_", ""), ".jpg", ""), "~", "")
			t1, err := time.Parse("2006_01_02_15_04_05", base)
			if err != nil {
				panic(err)
			}
			matches[i] = &Match{f, t1}
		}

		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			panic(err)
		}
		time.Local = loc

		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].date.After(matches[j].date)
		})

		results := make([]string, len(files))
		for i, m := range matches {
			results[i] = m.name
		}

		c.JSON(200, gin.H{
			"images": results,
		})
	})
	// Request to download a specific image
	r.GET("/image", func(c *gin.Context) {
		//get the image they want from the query string
		name, keyExists := c.GetQuery("name")
		if keyExists {
			// read the image from the filesystem
			data, err := ioutil.ReadFile(name)
			if err != nil {
				c.JSON(500, gin.H{
					"error": fmt.Sprintf("Something went wrong when reading the image: %s", err),
				})
			}
			c.Data(200, "image/jpg", data)
		} else {
			c.JSON(404, gin.H{
				"error": "name was not passed",
			})
		}
	})
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	//block on http serve
	r.Run(":8090")
}

func main() {
	//set up and block on web server forever
	web()
}
