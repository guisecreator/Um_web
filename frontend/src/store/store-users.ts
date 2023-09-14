import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useApolloClient } from '@vue/apollo-composable';
import { MUTATION } from '@/gql/types';
import { Role } from '@/gql/types';
import gql from 'graphql-tag';

export const useUserStore = defineStore('user', ()  => {
  const isLoggedIn = ref(false);
  const Roles = ref<typeof Role>();

  async function InitStore(): Promise<boolean> {
    let responseData: any;

    const { resolveClient } = useApolloClient('MUTATION');
    const client = await resolveClient();

    try {
      responseData = await client.mutate<MUTATION>({
        mutation: gql`mutation Mutation{
      validate{
        id
        email
        login
        role
      }
    }`})
    } catch (error) {
      console.log('не авторизован');
      isLoggedIn.value = false;
      return false;
    }

    isLoggedIn.value = true;
    //TODO: converted from Roles.value = response.data?.validate?.role || null;
    typeof Roles.value ; responseData.data?.validate?.role || null;
    return true;
  }

  function Login(role: typeof Role): void {
    isLoggedIn.value = true;
    Roles.value = role;
  }

  function Logout(): void {
    isLoggedIn.value = false;
    Roles.value = undefined;
  }

  return { isLoggedIn, Roles, Login, Logout, InitStore };
});
