import { createRouter, createWebHistory } from "vue-router";
import User from "../views/User.vue";
import Search from "../views/Search.vue";

const routes = [
  {
    path: "/users/:id",
    name: "User",
    component: User,
    props: true, // ルートパラメータをコンポーネントのプロパティとして渡す
  },
  {
    path: "/search",
    name: "Search",
    component: Search,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
