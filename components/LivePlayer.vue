<template>
  <b-container id="player" class="" fluid>
    <b-row>
      <b-col>
        <img
          id="content"
          :src="`http://secpi.pk5001z:8090/image?name=${live}`"
          alt=""
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
</style>