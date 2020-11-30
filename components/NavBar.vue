<template>
  <b-container fluid class="navbar shadow-lg rounded">
    <b-row>
      <b-col>
        <nuxt-link class="nolink" to="/">
          <h1 class="title">{{ title }}</h1>
        </nuxt-link>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
export default {
  name: "NavBar",
  props: { title: String },
};
</script>

<style scoped>
.navbar {
  background-color: rgb(212, 38, 38);
  padding: 0.3rem;
  padding-right: 0;
  margin-right: 0;
  border-right: 0;
}
.title {
  font-size: 1.5rem;
  margin-top: 0.6rem;
  margin-left: 2rem;
}
.nolink {
  color: black;
  text-decoration: none; /* no underline */
}
</style>
