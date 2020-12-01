<template>
  <b-container id="carousel" class="shadow" fluid @scroll="scrolled">
    <b-row>
      <Card
        v-for="picture in frame"
        :key="picture.id"
        :picture="picture"
        :selected="isSelected(picture)"
        @clicked="wasSelected(picture)"
      />
    </b-row>
  </b-container>
</template>
<script>
export default {
  name: "Carousel",
  data() {
    return {
      arrows: {
        left: 37,
        up: 38,
        right: 39,
        down: 40,
      },
      frame: [],
      selectedPicture: "",
      //only keep n image elements around at a time.
      //These are the current bounds that will change
      lowerBound: 0,
      upperBound: 100,
      step: 20,
    };
  },
  props: {
    pictures: Array,
  },
  watch: {
    pictures: function () {
      this.frame = this.pictures.slice(this.lowerBound, this.upperBound);
    },
  },
  mounted() {
    window.removeEventListener("keypress", this._keyListener);
    window.removeEventListener("keypress", this._keyListener);
    window.addEventListener("keydown", (e) => {
      console.log(e.keyCode);
      if (e.keyCode == this.arrows.right) {
        //find the index of who we are
        let us = this.pictures.indexOf(this.selectedPicture);
        //set selected to the next image in the list
        this.wasSelected(this.pictures[us + 1]);
      } else if (e.keyCode == this.arrows.left) {
        //find the index of who we are
        let us = this.pictures.indexOf(this.selectedPicture);
        //set selected to the previous image in the list
        this.wasSelected(this.pictures[us - 1]);
      }
    });
  },
  beforeDestroy() {
    window.removeEventListener("keydown", this._keyListener);
  },

  methods: {
    isSelected: function (picture) {
      return picture === this.selectedPicture;
    },
    wasSelected: function (picture) {
      this.selectedPicture = picture;
      this.$emit("changed", picture);
    },
    scrolled: function (asdf) {
      let currY = this.$el.scrollTop;
      let postHeight = this.$el.clientHeight;
      let scrollHeight = this.$el.firstChild.clientHeight;
      let percent = (currY / (scrollHeight - postHeight)) * 100;
      console.log(percent);
      if (percent >= 75.0) {
        this.scrolledDown();
      } else if (percent <= 15.0) {
        // this.scrolledUp();
      }
    },
    scrolledDown: function () {
      //change our view
      this.lowerBound += this.step;
      this.upperBound += this.step;
      this.frame = this.frame.concat(
        this.pictures.slice(this.lowerBound, this.upperBound)
      );
    },
    scrolledUp: function () {
      //change our view
      this.lowerBound -= this.step;
      this.upperBound -= this.step;
    },
  },
};
</script>
<style scoped>
#carousel {
  margin-top: 1rem;
  display: block;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 1rem;
  color: white;
  overflow-y: scroll;
  height: calc(100% - 38rem);
}
</style>