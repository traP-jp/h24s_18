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
    <button class="authorize_button" @click="login">認証する</button>
  </div>
</template>

<style scoped>
.authorize_button {
  background-color: #005bac;
  color: white;
  font-weight: bold;
  font-size: 20px;
  width: 128px;
  padding: 10px 20px;
  border-radius: 10px;
  outline: none;
  border: none;
  cursor: pointer;
  &:hover {
    background-color: #0066cc;
  }
}
</style>
