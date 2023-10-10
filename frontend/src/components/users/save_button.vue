<script lang="ts" setup>
import { useMutation } from '@vue/apollo-composable'
import { UserData } from '@/types/types'
import gql from 'graphql-tag';
import { MUTATION } from '@/gql/types';

const props = defineProps<{ saveData: UserData[] }>();
const emit = defineEmits(['update:saveData']);

const addUserData = async (data: UserData[]) => {
//   const newUsers = data.filter((value) => value.isNew).map((item) => ({
//     email: item.fields.email.value,
//     login: item.fields.name.value,
//     role: item.fields.role.value
//   }));

//   if (!newUsers.length) return data;

//   const { mutate: sendMessage } = useMutation<MUTATION>(gql`
//     mutation NewUser($users: [NewUserInput!]!){
//         userNew(users: $users){
//           id
//           login
//         }
//       }
//   `);

//   try {
//     const result = await sendMessage({
//       users: newUsers
//     });

//     if (result?.errors) {
//       result.errors.forEach((currentValue) => {
//         console.log(currentValue);
//       });
//       return data;
//     }

//     return data.map(elem => {
//       const user = result?.data?.userNew.find(val => elem.fields.name.value == val.login);
//       if (user) {
//         return {
//           ...elem,
//           fields: {
//             ...elem.fields,
//             login: {
//               oldValue: elem.fields.name.value,
//               value: elem.fields.name.value
//             }
//           },
//           id: user.id,
//           isNew: false
//         };
//       }
//       return elem;
//     });
//   } catch (e: any) {
//     console.log(e);
//   }
//   return data;
// }

// const updateUserData = async () => {
//   const updateUsers = props.saveData.filter((value) => {
//     for (const v of Object.values(value.fields))
//       if (v.oldValue != v.value)
//         return true;
//     return false;
//   })
//     .map((item) => ({
//       id: item.id,
//       login: item.fields.name.value,
//       role: item.fields.role.value
//     }))

//   if (updateUsers.length == 0) return props.saveData

//   const { mutate: sendMessage } = useMutation<MUTATION>(gql`
//     mutation UpdateUser($users: [UpdateUserInput!]!){
//         userUpdate(users: $users){
//           id
//           login
//           role
//         }
//       }
//   `)

//   try {
//     const result = await sendMessage({
//       users: updateUsers
//     })

//     if (result?.errors) {
//       result.errors.forEach(
//         (currentValue) => {
//           console.log(currentValue)
//         }
//       )
//       return props.saveData
//     }
//     return props.saveData.map(elem => {
//       const user = result?.data?.userUpdate.find(val => elem.id == val.id)
//       if (user)
//         return {
//           ...elem, fields: {
//             ...elem.fields, login: {
//               oldValue: elem.fields.name.value,
//               value: elem.fields.name.value
//             },
//             role: {
//               oldValue: elem.fields.role.value,
//               value: elem.fields.role.value
//             }
//           },
//         }
//       return elem;
//     })
//   }
//   catch (e: any) {
//     console.log(e)
//   }
//   return;
// }

// const deleteUserData = async () => {
//   const deleteUsers = props.saveData.filter((value) => !value.isDelete)
//     .map((item) => ({
//       id: item.id,
//       login: item.fields.name.value,
//       role: item.fields.role.value
//     }))

//   const { mutate: sendMessage } = useMutation<MUTATION>(gql`
//     mutation DeleteUser($users: [DeleteUserInput!]!){
//         userDelete(users: $users){
//           id
//           login
//           role
//         }
//       }
//   `)

//   try {
//     const result = await sendMessage({
//       users: deleteUsers
//     })

//     if (result?.errors) {
//       result.errors.forEach(
//         (currentValue) => {
//           console.log(currentValue)
//         }
//       )
//       return props.saveData
//     }
//     const deletedUser = result?.data?.userDelete;
//     if (deletedUser && deletedUser.length !== props.saveData.length) {
//       console.error("error: you can't delete all users");
//     }
//   } catch (e: any) {
//     console.log("An error occurred:", e)
//   }
// }

const saveUserData = async () => {
  const result = await addUserData(props.saveData);
  // await updateUserData();
  // await deleteUserData();
  emit('update:saveData', result)

  // if (props.saveData) {
  //   const result = await addUserData(props.saveData);
  //   emit('update:saveData', result);
  // } else {
  //   console.error('props.saveData is undefined or null');
  // }
}
</script>

<template>
  <v-btn
    variant="outlined"
    role="link"
    type="submit"
    label="Save"
    color="primary"
    size="small"
    class="mr-2"

  >Save</v-btn>
</template>
