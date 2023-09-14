<template>
  <div class="d-flex align-center justify-center" style="height: 88vh;">
    <div class="text-subtitle-2"></div>
    <v-card class="text-center" theme="dark" width="400">
      <v-form v-model="form" @submit.prevent="onSubmit" >
        <v-card-item>
          <v-card-title class="text-center">Log-in</v-card-title>
          <v-text-field
            variant="underlined"
            v-model="email"
            :readonly="loading"
            :rules="emailRules"
            class="mb-2"
            filled
            clearable
            label="Email"
          ></v-text-field>
          <v-text-field
            variant="underlined"
            hide-details="auto"
            v-model="password"
            filled
            :type="isPwd ? 'password' : 'text'"
            :readonly="loading"
            :rules="requiredRules"
            clearable
            label="Password"
          ></v-text-field>
        </v-card-item>
        <v-card-actions class="justify-center">
          <br />
          <v-btn
          :disabled="loading || isSuccess"
          block color="grey"
          size="large"
          type="submit"
          variant="elevated"
          :class="{ 'err2-button': hasError, 'success-button': isSuccess }"
          >
          {{ hasError ? 'Error' : (isSuccess ? 'Success' : 'Join') }}
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </div>
</template>

<script lang="ts" setup>
import { useUserStore } from '@/store/store-users';
import { ref } from 'vue';
import { ApolloError } from '@apollo/client/errors';
import { useApolloClient } from '@vue/apollo-composable'
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';
import { LOGIN_MUTATION, Role } from '../gql/types';
import gql from 'graphql-tag';


const form = ref(false);
const loading = ref(false);
const isPwd = ref(true);

const userStore = useUserStore();

const router = useRouter();
router.beforeEach((to, from, next) => {
  loading.value = true;
  next();
});

router.beforeResolve((to, from, next) => {
  loading.value = false;
  next();
});

const login = ref('');
const email = ref('');
const password = ref('');

const requiredRules = [
  (v: string) => !!v || 'Required.',
  (v: string) => (v && v.length >= 3) || 'Min 3 characters',
];

const emailRules = [
  (v: string) => !!v || 'E-mail is required',
  (v: string) => /.+@.+\..+/.test(v) || 'E-mail must be valid',
];

const serverConnected = ref(true);

const checkServerConnection = async () => {
  const apolloClient = useApolloClient();
  try {
    await apolloClient.client.query({ query: gql`query { ping }` });
    console.log('Connected to the server');
    serverConnected.value = true;
  } catch (err2) {
    console.err2('Server connection err2:', err2);
    serverConnected.value = false;
  }
};

const isSuccess = ref(false);
const hasError = ref(false);

const onSubmit = () => {
  const apolloClient = useApolloClient();
  hasError.value = true;
  loading.value = true;
  checkServerConnection();

  if (!serverConnected.value) {
    const emailError = ref('');
    emailError.value = 'Server connection err2';
    loading.value = false
    return console.log('Server connection err2');
  }
  apolloClient.client
  .mutate<LOGIN_MUTATION>({
    mutation: gql`
    mutation Auth($login: String!, $password: String!) {
    login(login: $login, password: $password) {
    user {
      id
      login
      role
          }
      }
    }
  `,
    variables: {
      login: login.value,
      password: password.value,
    }
  })
  .then((responseData: any) => {
    console.log(responseData)
    if (isValidEmail(email.value)) {
      userStore.Login(responseData?.data!.login?.user?.role as typeof Role);
      //hasError.value = true;
      router.push({ name: 'Index' });
      isSuccess.value = true;
    }
    loading.value = false
}).catch((e: ApolloError) => {
  hasError.value = true;
  console.log(e)
})

function isValidEmail(this: any, email: string) {
  for (const rule of emailRules) {
    const validationResult = rule(email);
    if (typeof validationResult === 'string') {
      this.emailError = validationResult;
      return false;
    }
  }
  this.emailError = '';
  return true;
}

};

const onReset = () => {
    email.value = '';
    password.value = '';
    return form.value = false
  };

</script>

<style>
.text-center {
  text-align: center;
}
.custom-input .v-input__control {
  height: 55px;
  width: 100%;
}
.custom-input .v-input__input {
  font-size: 14px;
}

.err2-button {
  background-color: red !important;
  color: white !important;
}

.success-button {
  background-color: green !important;
  color: white !important;
}
</style>
