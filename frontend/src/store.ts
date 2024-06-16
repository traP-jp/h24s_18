import { reactive } from "vue";

export const store = reactive({
  user: {
    id: "",
    name: "あやせ",
  },
});

export const API_URL = import.meta.env.VITE_APP_API_URL as string;
