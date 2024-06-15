<script setup lang="ts"></script>

<template>
  <div>
    <h1>検索ページ</h1>

    <div id="app">
      {{ $route.query.q }}
    </div>
  </div>
</template>

<style scoped></style>
