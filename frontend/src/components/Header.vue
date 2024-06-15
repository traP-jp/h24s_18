<script setup lang="ts">
import { RouterLink } from "vue-router";
import { store } from "../store";
import router from "../router";

const logout = () => {
  store.user.id = "";
  router.push({ name: "Login" });
};
</script>

<template>
  <div class="header">
    <div v-if="store.user.id">
      <RouterLink
        class="header-button"
        :to="{ name: 'User', params: { id: store.user.id } }"
        ><strong>traPハッカソン2024春18班</strong></RouterLink
      >
    </div>
    <div v-else>
      <RouterLink class="header-button" :to="{ name: 'Login' }"
        ><strong>traPハッカソン2024春18班</strong></RouterLink
      >
    </div>
    <div class="header-contents">
      <div v-if="store.user.id">
        <RouterLink
          class="header-button"
          :to="{ name: 'User', params: { id: store.user.id } }"
          >ユーザーページ</RouterLink
        >
        <RouterLink
          class="header-button"
          :to="{ name: 'Search', query: { q: 'SysAd' } }"
          >検索ページ</RouterLink
        >
        <button class="header-button" @click="logout">ログアウト</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.header {
  position: fixed;
  top: 0;
  left: 0;
  height: 40px;
  width: 100%;
  padding: 0 20px;
  border-bottom: 1px solid #ccc;
  padding: 10px 20px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-items: center;
}
.header-contents {
  display: flex;
  flex-wrap: wrap;
  justify-content: right;
  align-items: center;
  padding-right: 40px;
}
.header-button {
  color: #333;
  background: none;
  border: none;
  cursor: pointer;
  font-size: 16px;
  margin: 0 10px;
  padding: 0;
}
.header-button:hover {
  color: #005bac;
}
</style>
