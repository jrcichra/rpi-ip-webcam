package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
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

// take a picture
func capture(f *strftime.Strftime) {
	hour, _, _ := time.Now().Clock()
	filename := fmt.Sprintf("%s/%d/%s.jpg", PATH, hour, f.FormatString(time.Now()))
	out, err := exec.Command("/home/pi/capture/capture.sh", filename).CombinedOutput()
	log.Println(string(out))
	if err != nil {
		log.Println(err)
	}
}

func loop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	f, err := strftime.New("%Y_%m_%d_%H_%M_%S")
	if err != nil {
		panic(err)
	}
	for {
		<-ticker.C
		//spawn a process to capture the picture in the background
		go capture(f)
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
		return nil
	})
	go server.Serve()
	defer server.Close()

	r := gin.Default()
	// Sanity GET request
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Get a list of images on the filesystem, newest first
	r.GET("/images", func(c *gin.Context) {
		matches, err := filepathx.Glob("/home/pi/recordings/*/*.jpg")
		if err != nil {
			panic(err)
		}
		matchesMtime := make([]time.Time, len(matches))
		for i, match := range matches {
			fileinfo, err := os.Stat(match)
			if err != nil {
				panic(err)
			}
			matchesMtime[i] = fileinfo.ModTime()
		}
		sort.SliceStable(matches, func(i, j int) bool {
			// After gives us newest first
			return matchesMtime[i].After(matchesMtime[j])
		})

		c.JSON(200, gin.H{
			"images": matches,
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
	//vue serve
	r.Static("/", "./vue/")
	//block on http serve
	r.Run(":8090")
}

func main() {
	// start taking pictures
	go loop(INTERVAL * time.Second)
	//set up and block on web server forever
	web()
}
