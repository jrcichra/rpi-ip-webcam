<template>
  <b-container id="player" class="" fluid>
    <b-row>
      <b-col>
        <img
          id="content"
          :src="`http://secpi.pk5001z/image?name=${picture}`"
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
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50rem;
}
#player {
  margin-top: 1rem;
  display: block;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 1rem;
  color: white;
}
</style>