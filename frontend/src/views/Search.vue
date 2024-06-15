<script setup lang="ts">
import { useRoute } from "vue-router";

const route = useRoute();
const tag = route.query.q;
</script>

<template>
  <div>
    <h1>検索ページ</h1>

    <div id="app">
      {{ tag }}
    </div>
   
  </div>
</template>

<style scoped></style>
