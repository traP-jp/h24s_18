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
  <button @click="navigateTo(tag)">{{ tag }}</button>
</template>

<style></style>
