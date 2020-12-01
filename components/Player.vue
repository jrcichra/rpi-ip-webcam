<template>
  <b-container id="player" class="" fluid>
    <b-row>
      <b-col>
        <img
          id="content"
          :src="`http://secpi.pk5001z:8090/image?name=${picture}`"
          alt=""
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
    return {};
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
</style>