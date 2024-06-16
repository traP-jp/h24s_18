<script setup lang="ts">
import Tagbutton from "../components/Tagbutton.vue";
import { ref, computed, watch, onMounted } from "vue";
import { useRoute } from "vue-router";
import UserIcon from "../components/UserIcon.vue";
import { API_URL } from "../store";

//プロフィールデータ
const route = useRoute();
const userId = ref<string>(route.params.id as string);

const biography = ref<string>("");
const tags = ref<string[]>(["23M", "aaa", "ハッカソンなう"]);

const editBiography = ref<string>("");
const editTag = ref("");

const isEditing = ref<boolean>(false);

const xId = ref<string>("");

const getUserInfo = async () => {
  const res = await fetch(`${API_URL}/api/users/${userId.value}`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": API_URL,
    },
  });
  if (!res.ok) {
    throw new Error("Failed to fetch user info");
  }
  const data = await res.json();
  biography.value = data.Bio;
  xId.value = data.TwitterId;
  console.log(data);
};

onMounted(() => {
  getUserInfo();
});

watch(
  () => route.params.id,
  (newId: any) => {
    userId.value = newId as string;
    getUserInfo();
  }
);

//関数定義
const startEditing = () => {
  editBiography.value = biography.value; //現在のプロフィールを編集用に設定
  isEditing.value = true;
};

const addTag = () => {
  if (editTag.value === "") return;
  tags.value.push(editTag.value);
  editTag.value = "";
};

const deleteTag = (tag: string) => {
  tags.value = tags.value.filter((t) => t !== tag);
};

const cancelEditing = () => {
  isEditing.value = false; //編集モードを終了
};

const saveBiography = async () => {
  try {
    biography.value = editBiography.value; // 編集内容を保存
    isEditing.value = false; // 編集モードを終了
    await fetch(`${API_URL}/api/me`, {
      method: "PATCH",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": API_URL,
      },
      body: JSON.stringify({ Bio: biography.value }),
    });
    console.log("Biography saved");
  } catch (error) {
    console.error("Failed to save biography:", error);
  }
};

//本人確認
import { store } from "../store";
const isMyPage = computed(() => {
  return store.user.id === userId.value;
});
const profileLinks = computed(() => [
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

    <div class="tag-container">
      <div class="tag-content" v-for="tag in tags.slice()" :key="tag">
        <Tagbutton :tag="tag" />
        <button class="delete-button" v-if="isEditing" @click="deleteTag(tag)">
          ×
        </button>
      </div>
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
      <!--tag追加-->
      <label>
        新しいタグ名
        <input class="tag-input" v-model="editTag" type="text" />
      </label>
      <button class="edit-button save-button" @click="addTag">追加</button>
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
  border: 2px solid #d0d0d0;
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
.delete-button {
  display: flex;
  justify-content: center;
  align-items: center;
  border: none;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  font-size: 16px;
  cursor: pointer;
  background-color: #e2e5e9;
}
.tag-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  margin: 10px;
}
.tag-content {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  margin: 5px;
}
.tag-input {
  height: 36px;
  padding: 0 20px;
  color: inherit;
  font: inherit;
  border-radius: 18px;
  border: 2px solid #d0d0d0;
  margin: 0 10px 0 5px;
}
</style>
