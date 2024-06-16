<script setup lang="ts">
import { useRouter } from "vue-router";
import { Tag } from "../store";
import { computed, defineProps } from "vue";

const props = defineProps<{
  tag: Tag;
  isEditing: boolean;
  onClick: (tag: Tag) => void;
}>();
const router = useRouter();
const onClick = (tag: Tag) => {
  if (props.isEditing) {
    props.onClick(tag);
  } else {
    router.push({ name: "Search", query: { q: tag.name } });
  }
};
const classNames = computed(() => {
  return {
    "tag-button": true,
    "tag-button-starred": props.tag.isStarred,
    "tag-button-editing": props.isEditing,
  };
});
</script>

<template>
  <button :class="classNames" @click="onClick(tag)">
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
.tag-button-editing:hover {
  background-color: #f2ba4a;
  border-color: #f2ba4a;
}
.tag-button-starred {
  background-color: rgba(242, 186, 74, 0.2);
  border-color: rgb(242, 186, 74);
}
</style>
