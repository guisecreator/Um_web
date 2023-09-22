<template>
  <div>
    <v-app style="background-color: #303030;">
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
    const initResult = await userStore.InitStore()
    if (!initResult){
      redirect({name: 'login'})
    } else{
      redirect({name: 'index'})
    }
  },
}
</script>
