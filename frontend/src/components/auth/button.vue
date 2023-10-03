<template>
  <v-btn @click="LoginButton" :disabled="loading" v-if="!userStore.isLoggedIn" block color="grey" size="large" type="submit" variant="elevated">
    {{ loading ? 'Loading...' : 'Join' }}
  </v-btn>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store/store-users';
import { ref } from 'vue';

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false);

const LoginButton = async () => {
  if (!userStore.isLoggedIn) {
    loading.value = true;
    try {
      await router.push({ name: 'Index' });
    } catch (err2) {
      console.error('Error while navigating:', err2);
    } finally {
      loading.value = false;
    }
  }
}
</script>
