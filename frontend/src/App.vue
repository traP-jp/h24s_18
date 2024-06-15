<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { store } from "./store";
import router from "./router";

const logout = () => {
  store.user.id = "";
  router.push({ name: "Login" });
};
</script>

<template>
  <div>
    <p><strong>ユーザーID：</strong> {{ store.user.id }}</p>
    <p><strong>パス：</strong> {{ $route.fullPath }}</p>
    <div v-if="store.user.id">
      <RouterLink :to="{ name: 'User', params: { id: store.user.id } }"
        >ユーザーページへ</RouterLink
      >
      <br />
      <RouterLink :to="{ name: 'Search' }">検索ページへ</RouterLink>
      <br />
      <button @click="logout">ログアウト</button>
    </div>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped></style>
