<script setup lang="ts">
import { ref } from "vue";
interface User {
  name: string;
  tag: string;
}
// ユーザーの配列の中にタグの配列をいれる
const userData = ref<User[]>([
  { name: "mumumu", tag: "24B" },
  { name: "aya_se", tag: "23M" },
  { name: "eru_o2lo", tag: "24B"}
]);
</script>

<template>
  <div>
    <h1>検索ページ</h1>
    <h2 v-if="$route.query.q">
      <div id="app"># {{ $route.query.q }}</div>
    </h2>
    
    <li v-for="user in userData.slice()" :key="user.name">
      <h3>{{ user.name }}</h3>
      <img
        :src="`https://q.trap.jp/api/v3/public/icon/${user.name}`"
        alt="アイコン"
        id="icon"
      />
    </li>
  </div>
</template>

<style scoped>
#icon {
  width: 90px;
  border-radius: 50%;
}
</style>
