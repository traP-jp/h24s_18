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
  background-color: #e2e5e9;
  border: 2px solid #d0d0d0;
  border-radius: 18px;
  padding: 4px 15px;
  margin: 5px;
  cursor: pointer;
  font: inherit;
  font-size: 16px !important;
  font-weight: bold;
  color: inherit;
}
.tag-button:hover {
  background-color: #d0d0d0;
}
</style>
