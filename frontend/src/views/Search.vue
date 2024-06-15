<script setup lang="ts">
import { ref } from "vue";
import Tagbutton from "../components/Tagbutton.vue";

interface User {
  name: string;
  tags: string[]; // tags を文字列の配列として定義する
}

const userData = ref<User[]>([
  { name: "mumumu", tags: ["24B", "aaa", "ハッカソンなう"] }, // tags を配列で指定する
  { name: "aya_se", tags: ["23M"] },
  { name: "eru_o2lo", tags: ["24B"] },
]);
</script>

<template>
  <div>
    <h1>検索ページ</h1>
    <h2 v-if="$route.query.q">
      <div id="app"># {{ $route.query.q }}</div>
    </h2>
    <div class="user-list" v-for="user in userData.slice()" :key="user.name">
      <div class="user">
        <div class="username">
          <img
            :src="`https://q.trap.jp/api/v3/public/icon/${user.name}`"
            alt="アイコン"
            id="icon"
          />
          <strong>@{{ user.name }}</strong>
        </div>
        <div class="usertag">
          <div v-for="tag in user.tags.slice()" :key="tag" class="Tag">
            <Tagbutton :tag="tag" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
#icon {
  width: 90px;
  border-radius: 50%;
}

.user-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.user {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
  margin: 10px;
}
.username {
  display: flex;
  flex-wrap: wrap;
  width: 100px;
  align-items: center;
  justify-content: center;
}

.usertag {
  display: flex;
  flex-wrap: wrap;
  width: calc(100% - 140px);
  padding: 0 20px;
}
</style>
