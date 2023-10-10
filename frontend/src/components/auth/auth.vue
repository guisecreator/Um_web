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
          :class="{ 'error-button': hasError, 'success-button': isSuccess }"
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
import { ref, onMounted } from 'vue';
import { ApolloError } from '@apollo/client/errors';
import { useRouter } from 'vue-router';
import { LOGIN_MUTATION, Role } from '@/gql/types';
import gql from 'graphql-tag';
import Cookies from 'js-cookie';

import { useApolloClient } from '@vue/apollo-composable';
import { createApolloClient } from '@/apollo/apollo';

const apolloClient = useApolloClient();

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

const isSuccess = ref(false);
const hasError = ref(false);


async function onSubmit() {
  loading.value = true;
  const cookie = Cookies.get('auth_cookie');

  try {
    const responseData = await apolloClient.client.mutate<LOGIN_MUTATION>({
      mutation: gql`
        mutation Auth($email: String!, $password: String!) {
          login(email: $email, password: $password) {
            user {
              id
              email
              password
              createAt
              updateAt
              deletedAt
              role
            }
            info {
            Token
          }
        }
      }
      `,
      variables: {
        email: email.value,
        password: password.value,
      },
    });

    console.log(responseData);
    const role = responseData?.data?.login?.user?.role;
    userStore.Login(role as Role);

    const authToken = responseData?.data?.login?.info?.Token;
    console.log(authToken);

    const setCookie = Cookies.set('auth_cookie', authToken, {
      expires: 7, path: '/', sameSite: 'lax', secure: false });
    if (!setCookie) {
      hasError.value = true;
      loading.value = false;
      console.log("coockies not set");
      return;
    }

    router.push({ name: 'Index' });
    isSuccess.value = true;
  } catch (e) {
    console.error(e);
    hasError.value = true;
  } finally {
    loading.value = false;
}

function isValidEmail(this: any, email: string): boolean {
  if (!email) {
    this.emailError = 'E-mail is required';
    return false;
  }
  for (const rule of emailRules) {
    const validationResult = rule(email);
    if (typeof validationResult === 'string') {
      this.emailError = validationResult;
      return false;
    }
  }
  this.emailError = '';
  return true;
}};

const onReset = () => {
    email.value = '';
    password.value = '';
    return form.value = false
  };

</script>
