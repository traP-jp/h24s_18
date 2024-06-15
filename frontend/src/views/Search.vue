<script setup lang="ts">
import { ref } from "vue";
import Tagbutton from "../components/Tagbutton.vue";
import UserIcon from "../components/UserIcon.vue";

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
    <div class="search-title" v-if="$route.query.q">
      <strong># {{ $route.query.q }}</strong
      >&nbsp;の検索結果
    </div>
    <div class="user-list" v-for="user in userData.slice()" :key="user.name">
      <div class="user">
        <div class="username">
          <RouterLink :to="{ name: 'User', params: { id: user.name } }"
            ><UserIcon :userId="user.name" :size="90"
          /></RouterLink>

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
.search-title {
  font-size: 24px;
  margin: 20px 0;
  strong {
    font-size: 32px;
    color: #005bac;
  }
}
.user-list {
  max-width: 800px;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  border-top: 1px solid #e2e5e9;
  margin: 0 auto;
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
