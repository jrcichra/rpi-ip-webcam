<template>
  <b-col @click="clicked" class="column">
    <img id="card" src="../assets/boots.png" alt="" />
  </b-col>
</template>
<script>
export default {
  name: "Card",
  props: {},
  data() {
    return {};
  },
  methods: {
    clicked: function (mouse) {
      console.log(`You clicked from within a card at:`);
      console.log(mouse);
    },
  },
};
</script>
<style scoped>
#card {
  display: block;
  width: 10rem;
}
.column {
  padding: 0.05rem;
  margin: 0;
  border: 0;
  flex-grow: unset;
}
</style>