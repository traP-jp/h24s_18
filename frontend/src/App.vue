<script setup lang="ts">
import { RouterView } from "vue-router";
import Header from "./components/Header.vue";
</script>

<template>
  <div>
    <Header />
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
main {
  width: 100%;
  margin-top: 30px;
}
</style>
