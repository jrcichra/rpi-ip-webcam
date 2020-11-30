<template>
  <b-col @click="clicked" class="column">
    <b-img
      id="card"
      v-bind="mainProps"
      v-bind:class="{ selected: selected }"
      :src="`http://secpi.pk5001z/image?name=${picture}`"
    ></b-img>
  </b-col>
</template>
<script>
export default {
  name: "Card",
  props: {
    picture: String,
    selected: Boolean,
  },
  data() {
    return {
      mainProps: {
        center: false,
        fluidGrow: false,
        fluid: false,
        // blank: true,
        blankColor: "#bbb",
        // width: 50,
        // height: 60,
        class: "my-5",
        show: true,
      },
    };
  },
  methods: {
    clicked: function (mouse) {
      this.$emit("clicked");
    },
  },
};
</script>
<style scoped>
#card {
  display: block;
  width: 10rem;
  border: 0;
  margin-top: 0 !important;
  margin-bottom: 0 !important;
}
.column {
  padding: 0.05rem;
  margin: 0;
  border: 0;
  flex-grow: unset;
}
.selected {
  border: 0.2rem solid red !important;
}
</style>