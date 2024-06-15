<script setup lang="ts">
import { ref } from "vue";
import Tagbutton from "../components/Tagbutton.vue";

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

const tags = ref<string[]>(["23M", "aaa", "ハッカソンなう"]);
</script>

<template>
  <div>
    <h1>ユーザーページ</h1>
    <h2>{{ $route.params.id }}</h2>
    <img
      :src="`https://q.trap.jp/api/v3/public/icon/${$route.params.id}`"
      alt="アイコン"
      id="icon"
    />
    <p>
      <img src="../assets/x.svg" alt="アイコン" id="snsIcon">
      <a href="https://x.com/ayase_lab">{{ $route.params.id }}のツイッター</a>
    </p>
    <p><img src="../assets/traQ.svg" alt="アイコン" id="snsIcon">
    <a :href="`https://q.trap.jp/channels/gps/times/${$route.params.id}`"
      >{{ $route.params.id }}のtraqtimes</a
    ></p>
    <div v-for="tag in tags.slice()" :key="tag">
      <Tagbutton :tag="tag" />
    </div>
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
  </div>
</template>

<style scoped>
#icon {
  width: 90px;
  border-radius: 50%;
}
#snsIcon {
  width: 20px;
}
</style>
