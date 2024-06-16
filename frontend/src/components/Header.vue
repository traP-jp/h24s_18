<script setup lang="ts">
import { RouterLink } from "vue-router";
import { store } from "../store";
import UserIcon from "../components/UserIcon.vue";
</script>

<template>
  <div class="header">
    <RouterLink
      class="header-title"
      :to="
        store.user.id
          ? { name: 'User', params: { id: store.user.id } }
          : { name: 'Login' }
      "
      >traPハッカソン2024春18班</RouterLink
    >
    <div class="header-contents">
      <div v-if="store.user.id">
        <RouterLink :to="{ name: 'User', params: { id: store.user.id } }"
          ><UserIcon :userId="store.user.id" :size="45"
        /></RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.header {
  height: 50px;
  width: 100%;
  padding: 0;
  border-bottom: 1px solid #e2e5e9;
  padding: 10px 20px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-items: center;
}
.header-contents {
  display: flex;
  flex-wrap: wrap;
  justify-content: right;
  align-items: center;
  padding-right: 60px;
}
.header-title {
  color: #005bac;
  font-size: 20px !important;
  font-weight: bold;
  background: none;
  border: none;
  cursor: pointer;
  font-size: 16px;
  margin: 0 10px;
  padding: 0;
}
.header-title:hover {
  color: #0066cc;
}
</style>
