<template>
  <b-col @click="clicked" class="column">
    <b-spinner id="spinner" v-if="loading" variant="danger" label="Spinning">
    </b-spinner>
    <b-img
      id="card"
      @load="loading = false"
      v-bind:class="{ selected: selected }"
      :src="`http://secpi.pk5001z:8090/image?name=${this.picture}`"
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
        blankColor: "#fff",
        // width: 50,
        // height: 60,
        class: "my-5",
        show: true,
      },
      loading: true,
    };
  },
  computed: {},
  watch: {
    picture: function () {
      this.loading = true;
    },
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
#spinner {
  position: absolute;
  margin: 3rem;
  margin-left: 4rem;
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