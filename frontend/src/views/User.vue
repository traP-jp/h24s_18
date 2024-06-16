<script setup lang="ts">
import Tagbutton from "../components/Tagbutton.vue";
import { ref, computed, watch } from "vue";
import { useRoute } from "vue-router";
import UserIcon from "../components/UserIcon.vue";

//プロフィールデータ
const route = useRoute();
const biography = ref<string>("ここに自己紹介文が表示されます");
const editBiography = ref<string>("");
const isEditing = ref<boolean>(false);
const userId = ref<string>(route.params.id as string);

watch(
  () => route.params.id,
  (newId: any) => {
    userId.value = newId as string;
  }
);

const tangs = ref<string>("");
const editTangs = ref<string>("");
const isEditingTangs = ref<boolean>(false);

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

//ここからtagns
const startEditingTangs = () => {
  editTangs.value = tangs.value; //現在のtagを編集用に設定
  isEditingTangs.value = true;
};

const saveTangs = () => {
  tangs.value = editTangs.value; //編集内容を保存
  isEditingTangs.value = false; //編集モードを終了
};

const cancelEditingTangs = () => {
  isEditingTangs.value = false; //編集モードを終了
};

const tags = ref<string[]>(["23M", "aaa", "ハッカソンなう"]);

//本人確認
import { store } from "../store";
const isMyPage = computed(() => {
  return store.user.id === userId.value;
});
const xId = ref<string>("ayase_lab");
const profileLinks = ref<any>([
  {
    name: "x",
    url: `https://x.com/${xId.value}`,
    displayUrl: `@${xId.value}`,
  },
  {
    name: "traQ",
    url: `https://q.trap.jp/channels/gps/times/${userId.value}`,
    displayUrl: `#gps/times/${userId.value}`,
  },
  {
    name: "crowi",
    url: `https://wiki.trap.jp/user/${userId.value}`,
    displayUrl: `user/${userId.value}`,
  },
  {
    name: "traP",
    url: `https://trap.jp/author/${userId.value}`,
    displayUrl: `author/${userId.value}`,
  },
]);

const getImageUrl = (name: string) => {
  return new URL(`../assets/${name}.svg`, import.meta.url).href;
};
</script>

<template>
  <div>
    <UserIcon :userId="userId" :size="120" />
    <h2>@{{ userId }}</h2>
    <div class="profile-links">
      <div class="profile-link" v-for="link in profileLinks" :key="link.name">
        <div class="profile-icon">
          <img
            class="profile-icon-image"
            :src="getImageUrl(link.name)"
            alt="icon"
          />
        </div>
        <a
          class="profile-url"
          :href="link.url"
          target="_blank"
          rel="noopener noreferrer"
          >{{ link.displayUrl }}</a
        >
      </div>
    </div>
    <div v-for="tag in tags.slice()" :key="tag">
      <Tagbutton :tag="tag" />
    </div>

    











    <!--tag表示-->
    <div v-if="!isEditingTangs">
      <p>{{ tangs }}</p>
      <button v-if="isMyPage" @click="startEditingTangs">tag追加</button>
    </div>

    <!--tag追加-->
    <div v-else>
      <textarea v-model="editTangs"></textarea>
      <button @click="saveTangs">保存</button>
      <button @click="cancelEditingTangs">キャンセル</button>
    </div>

    <!--profile表示-->
    <div v-if="!isEditing">
      <p>{{ biography }}</p>
      <button class="edit-button" v-if="isMyPage" @click="startEditing">
        編集
      </button>
    </div>

    <!--profile編集-->
    <div v-else>
      <div>
        <textarea v-model="editBiography"></textarea>
      </div>
      <button class="edit-button save-button" @click="saveBiography">
        保存
      </button>
      <button class="edit-button" @click="cancelEditing">キャンセル</button>
    </div>

   
  </div>
</template>

<style scoped>
.profile-links {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
.profile-link {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  margin: 10px;
}
.profile-icon {
  width: 24px;
  height: 24px;
  margin-right: 5px;
  background-color: #e2e5e9;
  border-radius: 50%;
  padding: 5px;
}
.profile-icon-image {
  width: 100%;
  height: 100%;
  border-radius: 50%;
}
.profile-url {
  color: #49535b;
}
.profile-url:hover {
  color: #0070f3;
}
.snsIcon {
  fill: currentColor;
  width: 20px;
}
textarea {
  width: 600px;
  height: 6%;
  border-radius: 40px;
  border: 2px solid #e2e5e9;
  margin: 10px;
  padding: 20px;
  color: inherit;
  font: inherit;
  font-size: 16px;
}
.edit-button {
  background-color: #e2e5e9;
  border: none;
  border-radius: 20px;
  padding: 8px 24px;
  margin: 5px;
  cursor: pointer;
  font: inherit;
  font-size: 16px !important;
  font-weight: bold;
  color: inherit;
}
.save-button {
  background-color: #005bac;
  color: white;
}
</style>
