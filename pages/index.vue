<template>
  <div id="app">
    <NavBar />
    <Player :picture="picture" />
    <Carousel @changed="updatePlayer" :pictures="pictures" />
  </div>
</template>

<script>
import moment from "moment";
export default {
  name: "App",
  mounted() {
    this.getPictures();
  },
  data() {
    return {
      picture: "",
      pictures: [],
    };
  },
  methods: {
    updatePlayer(picture) {
      this.picture = picture;
    },
    getPictures() {
      this.$axios.get(`http://secpi.pk5001z/images`).then((response) => {
        console.log(response.data.images);
        this.pictures = response.data.images;
        this.picture = this.pictures[0];
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
</style>
