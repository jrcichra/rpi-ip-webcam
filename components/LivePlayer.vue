<template>
  <b-container id="player" class="" fluid>
    <b-row>
      <b-col>
        <img
          id="content"
          :src="`http://secpi.pk5001z/image?name=${live}`"
          alt=""
        />
        <p id="label">🔴 Live</p>
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
    return {};
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
  color: white;
}
#label {
  position: absolute;
  top: 1rem;
  left: 3rem;
  font-family: fantasy;
  font-size: 1.2rem;
  color: red;
}
</style>