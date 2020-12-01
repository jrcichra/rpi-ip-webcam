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
    <Carousel @changed="updatePlayer" :pictures="pictures" />
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
      this.live = msg;
    });
  },
  data() {
    return {
      picture: "",
      pictures: [],
      live: "",
      hostname: "",
    };
  },
  computed: {
    title: function () {
      return `Raspberry Pi Camera Viewer (${this.hostname})`;
    },
  },
  methods: {
    updatePlayer(picture) {
      this.picture = picture;
    },
    getPictures() {
      this.$axios.get(`http://secpi.pk5001z:8090/images`).then((response) => {
        console.log(response.data.images);
        this.pictures = response.data.images;
        this.picture = this.pictures[0];
        this.live = this.pictures[0];
      });
    },
    getHostname() {
      this.$axios.get(`http://secpi.pk5001z:8090/hostname`).then((response) => {
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
  color: #1a222b;
  height: 100%;
}
#previews {
  margin-left: 1%;
}
</style>
