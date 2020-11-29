<template>
  <b-container id="player" class="" fluid>
    <b-row>
      <b-col>
        <img id="content" src="../assets/boots.png" alt="" />
      </b-col>
    </b-row>
  </b-container>
</template>
<script>
export default {
  name: "Player",
  props: {},
  data() {
    return {
      content: "Hi",
    };
  },
};
</script>
<style scoped>
#content {
  display: block;
  margin-left: auto;
  margin-right: auto;
}
#player {
  margin-top: 3rem;
  display: block;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 5rem;
  color: white;
}
</style>