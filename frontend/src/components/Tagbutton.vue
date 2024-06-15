<script setup lang="ts">
import { useRouter } from "vue-router";

defineProps<{
  tag: string;
}>();
const router = useRouter();
const navigateTo = (tag: string) => {
  // クエリパラメータを指定してページ遷移する
  router.push({ name: "Search", query: { q: tag } });
};
</script>

<template>
  <button class="tag-button" @click="navigateTo(tag)">#&nbsp;{{ tag }}</button>
</template>

<style scoped>
.tag-button {
  background-color: #f0f0f0;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 5px 10px;
  margin: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
  font-family: "M PLUS 1p", sans-serif;
  font-size: 16px;
}
.tag-button:hover {
  background-color: #e0e0e0;
}
</style>
