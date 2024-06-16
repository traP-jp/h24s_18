<script setup lang="ts">
import { ref, onMounted, watch } from "vue";
import Tagbutton from "../components/Tagbutton.vue";
import UserIcon from "../components/UserIcon.vue";
import { useRoute } from "vue-router";
import { API_URL, Tag } from "../store";

interface User {
  id: string;
  tags: Tag[];
}

const userData = ref<User[]>([]);
const route = useRoute();
const userTag = ref<string>(route.query.q as string);

const getUserInfo = async () => {
  const res = await fetch(`${API_URL}/api/users?q=${userTag.value}`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": API_URL,
    },
  });
  if (!res.ok) {
    throw new Error("Failed to fetch user info");
  }
  const data = await res.json();
  console.log(data);
  userData.value = data.map((item: any) => ({
    id: item.User.Id,
    tags: item.Tags.map((tag: any) => {
      return { name: tag.Name, isStarred: tag.IsStarred };
    }),
  }));
};

onMounted(() => {
  getUserInfo();
});

watch(
  () => route.query.q,
  (newId: any) => {
    userTag.value = newId as string;
    getUserInfo();
  }
);
</script>

<template>
  <div>
    <div class="search-title" v-if="$route.query.q">
      <strong># {{ $route.query.q }}</strong
      >&nbsp;の検索結果
    </div>
    <div class="user-list" v-for="user in userData" :key="user.id">
      <div class="user">
        <div class="username">
          <RouterLink :to="{ name: 'User', params: { id: user.id } }"
            ><UserIcon :userId="user.id" :size="90"
          /></RouterLink>

          <strong>@{{ user.id }}</strong>
        </div>
        <div class="usertag">
          <div v-for="tag in user.tags.slice()" :key="tag.name" class="Tag">
            <Tagbutton :tag="tag" :isEditing="false" :onClick="() => {}" />
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
