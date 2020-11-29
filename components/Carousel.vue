<template>
  <b-container @scroll="move" id="carousel" class="shadow" fluid>
    <b-row>
      <Card v-for="card in numofcards" :key="card" />
    </b-row>
  </b-container>
</template>
<script>
export default {
  name: "Carousel",
  props: {},
  data() {
    return {
      numofcards: 50,
    };
  },
  methods: {},
};
</script>
<style scoped>
#carousel {
  margin-top: 5rem;
  display: block;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 5rem;
  color: white;
  overflow-y: scroll;
  height: 28.5rem;
}
</style>