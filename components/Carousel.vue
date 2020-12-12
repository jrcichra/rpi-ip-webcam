<template>
  <b-container id="carousel" class="shadow" fluid @scroll="scrolled">
    <b-row>
      <Card
        v-for="(picture, index) in frame"
        :key="index"
        :picture="picture"
        :selected="isSelected(index)"
        @clicked="wasSelected(index)"
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
      selectedIndex: "",
      //only keep n image elements around at a time.
      //These are the current bounds that will change
      frame: [],
      lowerBound: 0,
      upperBound: 40,
      step: 10,
    };
  },
  props: {
    pictures: Array,
  },
  watch: {
    pictures: function () {
      this.newFrame();
      this.selectedIndex = "";
      this.lowerBound = 0;
      this.upperBound = 40;
      //TODO: how to scroll to the top of the pane when we reset images
      // this.$el.scrollIntoView(true);
    },
  },
  mounted() {
    window.removeEventListener("keypress", this._keyListener);
    window.removeEventListener("keydown", this._keyListener);
    window.addEventListener("keydown", (e) => {
      console.log(e.keyCode);
      if (e.keyCode == this.arrows.right) {
        //set selected to the next image in the list
        this.wasSelected(this.selectedIndex + 1);
      } else if (e.keyCode == this.arrows.left) {
        //set selected to the previous image in the list
        this.wasSelected(this.selectedIndex - 1);
      }
    });
  },
  beforeDestroy() {
    window.removeEventListener("keydown", this._keyListener);
  },

  methods: {
    isSelected: function (index) {
      return index === this.selectedIndex;
    },
    wasSelected: function (index) {
      this.selectedIndex = index;
      this.$emit("changed", this.pictures[index]);
    },
    scrolled: function (asdf) {
      let currY = this.$el.scrollTop;
      let postHeight = this.$el.clientHeight;
      let scrollHeight = this.$el.firstChild.clientHeight;
      let percent = (currY / (scrollHeight - postHeight)) * 100;
      console.log(`percent: ${percent.toFixed(0)}, bottomed: ${this.bottomed}`);
      if (percent >= 90.0) {
        this.scrolledDown();
      }
    },
    newFrame: function () {
      this.frame = this.pictures.slice(this.lowerBound, this.upperBound);
    },
    downFrame: function () {},
    scrolledDown: function () {
      //update our pointers
      this.lowerBound += this.step;
      this.upperBound += this.step;
      // add this.step to the bottom
      //remove this.step from the top
      this.frame = this.frame.concat(
        this.pictures.slice(this.upperBound - this.step, this.upperBound)
      );
      // this.frame = this.frame.slice(this.step);
      // this.newFrame();
      //set the scroll to make them think they went down?
      // this.$el.scrollTop -= 100;

      console.log("wentdown");
    },
    scrolledUp: function () {
      //change our view
      if (this.lowerBound > 0) {
        this.lowerBound -= this.step;
        this.upperBound -= this.step;
        this.offset -= this.step;
        this.newFrame();
      }
      console.log("wentup");
      // this.frame.splice();
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
  content-visibility: auto;
}
</style>