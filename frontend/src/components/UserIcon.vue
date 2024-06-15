<script setup lang="ts">
defineProps<{
  userId: string;
  size: number;
}>();
</script>
<template>
  <img
    :src="`https://q.trap.jp/api/v3/public/icon/${userId}`"
    alt="アイコン"
    :width="size"
    :height="size"
    class="user-icon"
  />
</template>

<style scoped>
.user-icon {
  border-radius: 50%;
}
</style>
