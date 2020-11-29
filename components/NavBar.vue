<template>
  <b-container fluid class="navbar shadow-lg rounded">
    <b-row>
      <b-col>
        <nuxt-link class="nolink" to="/">
          <h1 class="title">Raspberry Pi Camera Viewer</h1>
        </nuxt-link>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
export default {
  name: "NavBar",
  props: {
    msg: String,
  },
};
</script>

<style scoped>
.navbar {
  background-color: rgb(212, 38, 38);
  padding: 0.7rem;
}
.title {
  font-size: 2rem;
  margin-top: 0.6rem;
}
.nolink {
  color: black;
  text-decoration: none; /* no underline */
}
</style>
