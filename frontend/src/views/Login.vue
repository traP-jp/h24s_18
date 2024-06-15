<script setup lang="ts">
import { store } from "../store";
import router from "../router";
const login = () => {
  store.user.id = "aya_se";
  router.push({ name: "User", params: { id: store.user.id } });
};
</script>

<template>
  <div>
    <h1>認証ページ</h1>
    <button @click="login">認証する</button>
  </div>
</template>

<style scoped></style>
