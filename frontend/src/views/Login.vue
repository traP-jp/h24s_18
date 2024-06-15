<script setup lang="ts">
import router from "../router";
import { store, API_URL } from "../store";
import { onMounted } from "vue";

onMounted(async () => {
  // traQ認証の実装
  try {
    const res = await fetch(`${API_URL}/api/me`, {
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
    store.user.id = data.Id;
    store.user.name = data.Name;
    router.push({ name: "User", params: { id: store.user.id } });
  } catch (error) {
    console.error(error);
  }
});
</script>

<template>
  <div>
    <h1>認証ページ</h1>
    <a class="authorize_button" :href="`${API_URL}/api/oauth2/authorize`"
      >認証する</a
    >
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
