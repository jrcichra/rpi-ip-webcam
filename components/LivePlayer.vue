<template>
  <b-container id="player" class="" fluid>
    <b-row>
      <b-col>
        <img
          id="content"
          :src="`http://secpi/image?name=${live}`"
          alt=""
          v-bind:class="{ zoomed: clicked, hovered: hover }"
          @click="wasClicked"
          @mouseenter="hover = true"
          @mouseleave="hover = false"
        />
        <p v-if="showLive" id="icon">🔴</p>
        <p id="label">Live</p>
      </b-col>
    </b-row>
  </b-container>
</template>
<script>
export default {
  name: "Player",
  props: {
    live: String,
  },
  data() {
    return {
      showLive: true,
      showLiveInterval: undefined,
      clicked: false,
      hover: false,
    };
  },
  created() {
    this.showLiveInterval = setInterval(() => {
      this.showLive = !this.showLive;
    }, 1000);
  },
  destroyed() {
    clearInterval(this.showLiveInterval);
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
  margin-left: 0;
  padding-left: 0;
  border-left: 0;
}
#player {
  margin-top: 1rem;
  display: block;
  margin-bottom: 1rem;
  margin-left: 4em;
  color: white;
}
#label,
#icon {
  position: absolute;
  top: 1rem;
  left: 3rem;
  font-family: fantasy;
  font-size: 1.2rem;
  color: red;
}
#label {
  left: 4.6rem;
  top: 1.05rem;
}
.zoomed {
  position: absolute;
  zoom: 200%;
  left: -22rem;
  z-index: 5;
}
.hovered {
  border: 0.2rem solid azure !important;
}
</style>