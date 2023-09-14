<template>
  <div>
    <v-app style="background-color: #303030;">
      <!-- <router-link to="/">auth</router-link>
      <router-link to="/control">control</router-link>
      <router-link to="/settings">settings</router-link> -->
    <router-view></router-view>
  </v-app>
  </div>
</template>

<style>
.my-app {
  background-color: #303030;
}
</style>

<script lang="ts">
import { useUserStore } from '@/store/store-users';

export default {
  async preFetch({ redirect }: { redirect: any }) {
    const userStore = useUserStore()
    const initResult = await userStore.Init()
    if (!initResult){
      redirect({name: 'login'})
    } else{
      redirect({name: 'index'})
    }
  },
}
</script>
