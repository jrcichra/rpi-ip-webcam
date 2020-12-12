<template>
  <b-container id="player" class="" fluid>
    <b-row>
      <b-col>
        <img
          id="content"
          :src="`http://secpi:8090/image?name=${picture}`"
          alt=""
          v-bind:class="{ zoomed: clicked, hovered: hover }"
          @click="wasClicked"
          @mouseenter="hover = true"
          @mouseleave="hover = false"
        />
      </b-col>
    </b-row>
  </b-container>
</template>
<script>
export default {
  name: "Player",
  props: {
    picture: String,
  },
  data() {
    return {
      clicked: false,
      hover: false,
    };
  },
  methods: {
    wasClicked: function (event) {
      this.clicked = !this.clicked;
    },
  },
};
</script>
<style scoped>
#content {
  width: 40rem;
  height: auto;
}
#player {
  margin-top: 1rem;
  margin-left: 5rem;
  display: block;
  margin-bottom: 1rem;
  color: white;
  height: 30%;
}
.zoomed {
  position: absolute;
  zoom: 200%;
  z-index: 5;
  left: 4rem;
}
.hovered {
  border: 0.2rem solid azure !important;
}
</style>