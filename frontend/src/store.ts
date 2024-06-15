import { reactive } from "vue";

export const store = reactive({
  user: {
    id: "",
    name: "あやせ",
  },
});

export const API_URL = "http://localhost:8080";
