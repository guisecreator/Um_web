<template>
  <v-spacer></v-spacer>
  <div class="d-flex justify-center" style="height: 90vh;">
    <div class="text-subtitle-2"></div>
    <v-card class="text-center" theme="dark" width="1000" height="700">
      <v-progress-linear
        v-show="loading"
        :active="loading"
        :indeterminate="loading"
        absolute
        bottom
        color="deep-purple-accent-4"
      ></v-progress-linear>
      <v-card-item>
        <div class="d-flex justify-start">
          <v-btn
            role="link"
            type="submit"
            @click="navigateToIndex"
          >Back</v-btn>
        </div>
        <div class="text-h5 my-4">Settings</div>
      </v-card-item>
      <v-list lines="two" subheader>
        <v-list-subheader>Personal settings</v-list-subheader>
        <div class="d-flex flex-column align-start">
          <v-btn variant="text" color="white" >Change Login</v-btn>
          <v-btn variant="text" color="white" >Change Password</v-btn>
          <v-btn variant="text" color="white" >Change Rights</v-btn>
          <v-btn variant="text" color="white" >Sign out</v-btn>
        </div>
      </v-list>
      <v-divider></v-divider>
      <v-list lines="two" subheader>
        <v-list-subheader>General</v-list-subheader>
        <div>Добавить выбор показа только почты или только логина</div>
      </v-list>
      <v-card-actions class="d-flex justify-center v-card-actions-button-save my-2342">
        <v-col cols="12" sm="6" md="4">
          <save-settings></save-settings>
        </v-col>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import SaveSettings from '@/components/settings/SettingsSave.vue';


const loading = ref(false);

const router = useRouter();
router.beforeEach((to, from, next) => {
  loading.value = true;
  next();
});

router.beforeResolve((to, from, next) => {
  loading.value = false;
  next();
});


//const userStore = useUserStore();
// const { user } = storeToRefs(userStore);


function navigateToIndex() {
  loading.value = true;
  router.push({ name: 'Index' });
  loading.value = false;
}

</script>

<style>
.v-card-actions-button-save {
    align-items: center;
    display: flex;
    flex: none;
    min-height: 353px;
    padding: 0.5rem;
}

.my-2342 {
    margin-top: 100px !important;
    margin-bottom: 28px !important;
}

.my-4 {
    margin-top: -34px !important;
    margin-bottom: 16px !important;
}
</style>
