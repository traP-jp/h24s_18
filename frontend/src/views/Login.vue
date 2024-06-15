<script setup lang="ts">
import { store } from "../store";
import router from "../router";
import { API_URL } from "../const";
import axios from "axios";

const login = () => {
  store.user.id = "aya_se";
  router.push({ name: "User", params: { id: store.user.id } });
  return;
  // TODO: traQ認証の実装
  const searchParams = new URLSearchParams(window.location.search);
  console.log(searchParams.toString());
  if (searchParams.toString() === "") {
    axios
      .get(`${API_URL}/api/me`, {
        withCredentials: true,
        /* CORS回避 */
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": API_URL,
        },
      })
      .then((res) => {
        console.log(res.data);
        store.user.id = res.data.id;
        store.user.name = res.data.name;
        router.push({ name: "User", params: { id: store.user.id } });
      })
      .catch((err) => {
        console.log(err);
        axios
          .get(`${API_URL}/api/oauth2/authorize?${searchParams}`, {
            withCredentials: true,
            /* CORS回避 */
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": API_URL,
            },
          })
          .then((res) => {
            console.log(res);
            axios
              .get(`${API_URL}/api/me`, {
                withCredentials: true,
                /* CORS回避 */
                headers: {
                  "Content-Type": "application/json",
                  "Access-Control-Allow-Origin": API_URL,
                },
              })
              .then((res) => {
                console.log(res.data);
                store.user.id = res.data.id;
                store.user.name = res.data.name;
                router.push({ name: "User", params: { id: store.user.id } });
              });
          });
      });
  } else {
    axios
      .get(`${API_URL}/api/me`, {
        withCredentials: true,
        /* CORS回避 */
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": API_URL,
        },
      })
      .then((res) => {
        console.log(res.data);
        store.user.id = res.data.id;
        store.user.name = res.data.name;
        router.push({ name: "User", params: { id: store.user.id } });
      })
      .catch((err) => {
        console.log(err);
      });
  }
};
</script>

<template>
  <div>
    <h1>認証ページ</h1>
    <button @click="login">認証する</button>
  </div>
</template>

<style scoped></style>
