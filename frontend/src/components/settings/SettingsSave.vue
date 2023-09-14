<template>
  <v-btn @click="SaveSettings" block rounded="xl" variant="outlined" role="link" type="submit">Save </v-btn>
<!--  v-bind="$props"  -->
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store/store-users';
import { load } from 'webfontloader';
import { ref } from 'vue';

const SaveSettings = async () => {
  const userStore = useUserStore();

}

</script>
