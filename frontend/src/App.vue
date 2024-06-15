<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { store } from "./store";
</script>

<template>
  <div>
    <p><strong>Current route path:</strong> {{ $route.fullPath }}</p>
    <RouterLink :to="{ name: 'User', params: { id: store.user.id } }"
      >ユーザーページへ</RouterLink
    >
    <hr />
    <RouterLink :to="{ name: 'Search' }">検索ページへ</RouterLink>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped></style>
