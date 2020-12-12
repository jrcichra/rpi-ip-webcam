<template>
  <div id="app">
    <b-row>
      <NavBar :title="title" />
    </b-row>
    <b-row id="previews">
      <b-col cols="0"> </b-col>
      <b-col cols="6">
        <Player :picture="picture" />
      </b-col>
      <b-col cols="6">
        <LivePlayer :live="live" />
      </b-col>
    </b-row>
    <b-row>
      <b-col cols="1"></b-col>
      <b-col cols="0">
        <label id="datelabel" for="date">Start Date</label>
      </b-col>
      <b-col cols="3">
        <input
          type="datetime-local"
          v-model="datetime"
          id="datetime"
          :min="min"
          :max="max"
        />
        <b-button @click="findClosestImage"> Go </b-button>
      </b-col>
    </b-row>
    <Carousel @changed="updatePlayer" :add="live" :pictures="pictures" />
  </div>
</template>

<script>
import moment from "moment";
export default {
  name: "App",
  mounted() {
    this.getPictures();
    this.getHostname();
    this.socket = this.$nuxtSocket({
      channel: "/",
      cors: {
        origin: "*",
        methods: ["GET", "POST"],
      },
    });
    this.socket.on("live", (msg, cb) => {
      console.log(`got a live socket.io packet: ${msg}`);
      //make this the live image
      this.live = msg;
      //only add it to the existing view if the first picture matches the first in all pictures, meaning they're at the top of the list
      if (this.allPictures[0] == this.pictures[0]) {
        this.pictures.unshift(msg);
      }
      //regardless, add it to the pictures list. This should solve the previous condition if nothing changes
      this.allPictures.unshift(msg);
    });
  },
  data() {
    return {
      picture: "",
      pictures: [],
      allPictures: [],
      live: "",
      hostname: "",
      datetime: moment().format("YYYY-MM-DDTHH:mm"),
      refresh: 0,
    };
  },
  computed: {
    title: function () {
      return `Raspberry Pi Camera Viewer (${this.hostname})`;
    },
    min: function () {
      if (this.allPictures.length > 0) {
        return moment(
          this.basename(this.allPictures[this.allPictures.length - 1])
            .replace("mon_", "")
            .replace(".jpg", ""),
          "YYYY_MM_DD_HH_mm_ss"
        ).format("YYYY-MM-DDTHH:mm");
      } else {
        return "";
      }
    },
    max: function () {
      if (this.allPictures.length > 0) {
        return moment(
          this.basename(this.allPictures[0])
            .replace("mon_", "")
            .replace(".jpg", ""),
          "YYYY_MM_DD_HH_mm_ss"
        ).format("YYYY-MM-DDTHH:mm");
      } else {
        return "";
      }
    },
  },
  methods: {
    basename: function (str) {
      return str.substring(str.lastIndexOf("/") + 1);
    },
    findClosestImage: function () {
      for (let i = this.allPictures.length - 1; i > 0; i--) {
        //get the date for this picture based on the filename
        let base = this.basename(this.allPictures[i]);
        let datestr = base.replace("mon_", "").replace(".jpg", "");
        let date = moment(datestr, "YYYY_MM_DD_HH_mm_ss");
        if (date.isAfter(this.datetime)) {
          console.log(`I found ${date.toString()}`);
          this.pictures = this.allPictures.slice(
            i,
            this.allPictures.length - 1
          );
          return;
        }
      }
    },
    updatePlayer(picture) {
      this.picture = picture;
    },
    getPictures() {
      this.$axios.get(`http://secpi:8090/images`).then((response) => {
        console.log(response.data.images);
        this.pictures = response.data.images;
        this.allPictures = response.data.images;
        this.picture = this.pictures[0];
        this.live = this.pictures[0];
      });
    },
    getHostname() {
      this.$axios.get(`http://secpi:8090/hostname`).then((response) => {
        console.log(response.data.hostname);
        this.hostname = response.data.hostname;
      });
    },
  },
};
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #ffffff;
  height: 100%;
}
#previews {
  margin-left: 1%;
}
#datelabel {
  margin-top: 0.5rem;
}
</style>
