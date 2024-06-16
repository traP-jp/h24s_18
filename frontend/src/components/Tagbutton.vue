<script setup lang="ts">
import { useRouter } from "vue-router";
import { Tag } from "../store";

defineProps<{
  tag: Tag;
}>();
const router = useRouter();
const navigateTo = (tag: string) => {
  // クエリパラメータを指定してページ遷移する
  router.push({ name: "Search", query: { q: tag } });
};
</script>

<template>
  <button
    v-if="tag.isStarred"
    class="tag-button tag-button-starred"
    @click="navigateTo(tag.name)"
  >
    #&nbsp;{{ tag.name }}
  </button>
  <button v-else class="tag-button" @click="navigateTo(tag.name)">
    #&nbsp;{{ tag.name }}
  </button>
</template>

<style scoped>
.tag-button {
  background-color: #e2e5e9;
  border: 2px solid #d0d0d0;
  border-radius: 20px;
  padding: 4px 15px;
  margin: 5px;
  cursor: pointer;
  font: inherit;
  font-size: 14px !important;
  font-weight: bold;
  color: inherit;
  transition: background-color 0.1s, color 0.1s, border-color 0.1s;
}
.tag-button:hover {
  background-color: #005bac;
  color: white;
  border-color: #005bac;
}
.tag-button-starred {
  background-color: rgba(242, 186, 74, 0.2);
  border-color: rgb(242, 186, 74);
}
</style>
