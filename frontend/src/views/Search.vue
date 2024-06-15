<script setup lang="ts">
import { ref } from "vue";
import router from "../router";

interface User {
  name: string;
  tags: string[]; // tags を文字列の配列として定義する
}

const userData = ref<User[]>([
  { name: "mumumu", tags: ["24B", "aaa", "ハッカソンなう"] }, // tags を配列で指定する
  { name: "aya_se", tags: ["23M"] },
  { name: "eru_o2lo", tags: ["24B"] },
]);

const navigateToTag = (tag: string) => {
  // クエリパラメータを指定してページ遷移する
  router.push({ name: "Search", query: { q: tag } });
};
</script>

<template>
  <div>
    <h1>検索ページ</h1>
    <h2 v-if="$route.query.q">
      <div id="app"># {{ $route.query.q }}</div>
    </h2>

    <li v-for="user in userData.slice()" :key="user.name">
      <div class="username">
        <h3>{{ user.name }}</h3>
        <img
          :src="`https://q.trap.jp/api/v3/public/icon/${user.name}`"
          alt="アイコン"
          id="icon"
        />
        <div v-for="tag in user.tags.slice()" :key="tag" class="tag">
          <button @click="navigateToTag(tag)">{{ tag }}</button>
        </div>
      </div>
    </li>
  </div>
</template>

<style scoped>
#icon {
  width: 90px;
  border-radius: 50%;
}

.username {
  display: inline-block;
}

.usertag {
  display: inline-block;
}
</style>
