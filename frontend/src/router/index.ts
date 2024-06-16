import { createRouter, createWebHistory } from "vue-router";
import User from "../views/User.vue";
import Search from "../views/Search.vue";
import Login from "../views/Login.vue";
import { store } from "../store";

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
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach((to, _from, next) => {
  // 認証がない場合はloginページにリダイレクト
  if (to.name === "Login") {
    if (store.user.id !== "") {
      next({ name: "User", params: { id: store.user.id } });
    } else {
      next();
    }
  } else {
    if (store.user.id === "") {
      next({ name: "Login" });
    } else {
      next();
    }
  }
});

export default router;
