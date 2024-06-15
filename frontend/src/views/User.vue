<script setup lang="ts">
import SampleButton from "../components/SampleButton.vue";
</script>

<template>
  <div>
    <h1>ユーザーページ</h1>
    <h3>{{ $route.params.id }}</h3>
    <SampleButton />
  </div>
</template>

<style scoped></style>
