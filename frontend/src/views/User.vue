<script setup lang="ts">
import Biography from "../components/Biography.vue";
import SampleButton from "../components/SampleButton.vue";
//import Biography from "../components/Biography.vue";一旦無視
import { ref } from "vue";

//プロフィールデータ
const biography = ref<string>("ここに自己紹介文が表示されます");
const editBiography = ref<string>("");
const isEditing = ref<boolean>(false);

//関数定義
const startEditing = () => {
  editBiography.value = biography.value; //現在のプロフィールを編集用に設定
  isEditing.value = true;
};

const saveBiography = () => {
  biography.value = editBiography.value; //編集内容を保存
  isEditing.value = false; //編集モードを終了
};

const cancelEditing = () => {
  isEditing.value = false; //編集モードを終了
};
</script>

<template>
  <div>
    <h1>ユーザーページ</h1>
    <h2>{{ $route.params.id }}</h2>
    <h3>Profile</h3>

    <!-profile表示->
    <div v-if="!isEditing">
      <p>{{ biography }}</p>
      <button @click="startEditing">編集</button>
    </div>

    <!-profile編集->
    <div v-else>
      <textarea v-model="editBiography"></textarea>
      <button @click="saveBiography">保存</button>
      <button @click="cancelEditing">キャンセル</button>
    </div>

    <h4>{{ biography }}</h4>

    <Biography />

    <SampleButton />
  </div>
</template>

<style scoped></style>
