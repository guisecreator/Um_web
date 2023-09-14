<template>
  <v-spacer></v-spacer>
    <div class="d-flex justify-center" style="height: 90vh;">
      <div class="text-subtitle-2"></div>
      <v-card  theme="dark" width="1000" height="700">
        <v-progress-linear
        v-show="loading"
        :active="loading"
        :indeterminate="loading"
        absolute
        bottom
        color="deep-purple-accent-4"
      ></v-progress-linear>
        <div class="text-h5 my-234 text-center">Control Panel</div>
        <v-card-actions class="justify-end">
          <br />
          <v-btn
            color="primary"
            size="small"
            type="submit"
            class="mr-2 mr-auto "
            variant="elevated"
            icon="mdi-checkbox-multiple-marked"
            @click="toggleCheckboxes"
          ></v-btn>
          <v-btn
          color="red"
          size="small"
          type="submit"
          class=" mr-2 fixed-button"
          variant="elevated"
          :disabled="!isDeleteSelectedButtonActive"
          icon="mdi-delete-outline"
          v-show="showCheckboxes"
          @click="deleteSelectedUsers"
          ></v-btn>
          <SaveButton v-model:saveData="rows" @update:saveData="onSaved"></SaveButton>
          <v-dialog v-model="createDialog" persistent width="1024">
            <template v-slot:activator="{ props }">
                <v-btn
                color="success"
                size="small"
                type="submit"
                variant="elevated"
                icon="mdi-plus"
                @click="createDialog = true"
                v-bind="props"
              ></v-btn>
            </template>
          <v-card>
            <v-card-title class="d-flex">
              <span class="text-h5 flex-grow-1">Create new user</span>
              <v-btn icon dark @click="createDialog = false" class="ml-auto">
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-row>
                  <v-col cols="12">
                  </v-col>

                  <v-col cols="12">
                  <v-text-field
                        label="Login*"
                        required
                        v-model="newRow.fields.name.value"
                        color="grey"
                        hide-details="auto"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                      </v-col>
                      <v-col cols="12">
                  <v-text-field
                        label="Email*"
                        required
                        :rules="emailRules"
                        v-model="newRow.fields.email.value"
                        color="grey"
                        hide-details="auto"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                      </v-col>
                      <v-col cols="12">
                      <v-select
                        v-model="newRow.fields.role.value"
                        :items="Object.values(Role as any)"
                        label="Role*"
                        required
                      ></v-select>
                      </v-col>

                </v-row>
              </v-container>
              <small>*indicates required field</small>
            </v-card-text>
            <v-spacer></v-spacer>
            <v-card-actions class="d-flex justify-center">
              <v-col cols="12" sm="6" md="4">
                <v-btn
                block rounded="xl"
                variant="outlined"
                color="success"
                role="link"
                type="submit"
                @click="createUser"
                >Create
              </v-btn>
            </v-col>
          </v-card-actions>
        </v-card>
        </v-dialog>
        <v-btn
        color="grey"
        size="small"
        type="submit"
        variant="elevated"
        icon="mdi-cog"
        role="link"
        dark
        @click="navigateToSettings"
        >
      </v-btn>
    </v-card-actions>
    <v-table
    height="550px"
    class ="table1"
    bordered
    title="Пользователи"
    :columns="columns"
    :rows="rows"
    table-class="table"
    :loading="loading"
     >
    <thead>
      <tr>
        <th class="text-left">Email</th>
        <th class="text-left">Role</th>
        <th class="text-right">Change</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in rows" :key="item.id">
        <td>{{ item.fields.email }}</td>
        <td>{{ item.fields.role }}</td>
        <td>
          <v-dialog v-model="editDialog" persistent width="1024">
              <template v-slot:activator="{ props }">
                <v-btn
                color="grey"
                size="small"
                type="submit"
                variant="elevated"
                icon="mdi-pencil"
                role="link"
                dark
                @click="editUserData()"
                >Edit
              </v-btn>
            </template>
            <v-card>
              <v-card-title class="d-flex">
                <span class="text-h5 flex-grow-1">Edit user</span>
                <v-btn icon dark @click="editDialog = false" class="ml-auto">
                  <v-icon>mdi-close</v-icon>
                </v-btn>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-row>

                    <v-col cols="12">
                  <v-text-field
                        label="Login*"
                        required
                        v-model="newRow.fields.name.value"
                        color="grey"
                        hide-details="auto"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                      </v-col>
                      <v-col cols="12">
                  <v-text-field
                        label="Email*"
                        required
                        :rules="emailRules"
                        v-model="newRow.fields.email.value"
                        color="grey"
                        hide-details="auto"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                      </v-col>
                      <v-col cols="12">
                      <v-select
                        v-model="newRow.fields.role.value"
                        :items="Object.values(Role as any)"
                        label="Role*"
                        required
                      ></v-select>
                      </v-col>

                  <v-col cols="12">
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>
            <v-spacer></v-spacer>
            <v-card-actions class="d-flex justify-center">
              <v-col cols="12" sm="6" md="4">
              <v-btn
                color="red"
                block rounded="xl"
                variant="outlined"
                role="link"
                type="submit"
                @click="deleteUser"
                >Delete User
              </v-btn>
            </v-col>
              <v-col cols="12" sm="6" md="4">
                <v-btn
                color="success"
                block rounded="xl"
                variant="outlined"
                role="link"
                type="submit"
                >Save
              </v-btn>
              </v-col>
            </v-card-actions>
          </v-card>
        </v-dialog>
        </td>
        <td>
        <v-checkbox
        v-model="item.selected"
        v-if="showCheckboxes"
        ></v-checkbox>
        </td>
      </tr>
    </tbody>
  </v-table>
</v-card>
</div>
</template>

<script lang="ts" setup>
import 'custom-vue-scrollbar/dist/style.css';
import { ref, Ref, onMounted, computed  } from 'vue';
import { UserData, UserDataFields, UserField, UserDataKey} from '@/types/types';
import { useRouter } from 'vue-router';
import { User, Role } from '@/gql/types';
import gql from 'graphql-tag';
import { useApolloClient } from '@vue/apollo-composable'
import SaveButton from '@/components/users/SaveButton.vue';


const loading = ref(false);
const createDialog = ref(false);
const editDialog = ref(false);
const selected = ref(false);

const someData = ref<UserDataFields>();

const columns: Ref<UserDataFields[]> = ref([]);

// const roleValues: Role[] = Object.values<Role>(Role) as Role[];
const rows = ref<UserData[]>([])
// const rows = ref<UserData[]>([
//   {
//     id: '',
//     fields: {
//       name: { oldValue: '', value: '' },
//       email: { oldValue: '', value: '' },
//       role: { oldValue: Role.USER, value: Role.USER },
//     },
//     isNew: false,
//     isDelete: false,
//     selected: false,
//   },
// ]);
// const row = rows.value[0];

const router = useRouter();
router.beforeEach((to, from, next) => {
  loading.value = true;
  next();
});

router.beforeResolve((to, from, next) => {
  loading.value = false;
  next();
});

const showCheckboxes = ref(false);
function toggleCheckboxes() {
  showCheckboxes.value = !showCheckboxes.value;
}

const isDeleteSelectedButtonActive = computed(() => {
  return rows.value.some((item) => item.selected);
});

const isSelectedUser = () => {
  return selected.value
};

onMounted(() => {
  refreshUserData()
})

async function refreshUserData() {
  loading.value = true;
  const apolloClient = useApolloClient();
  try{
    const serverResult: any = await apolloClient.client.query({
    query: gql`
      query {
        users {
          id
          name
          email
          role
        }
      }
    `,
  }).then((result: any) => {
    rows.value = result.data.users;
  });

  if (!serverResult.data.users){
    console.log(serverResult.data.users);
    return;
  }

  const users = serverResult.data.users;
  if (!users){
    console.log(users);
    return;
  }
  rows.value = users.map((user: any) => {
    let newItem: UserData = {
      id: user.id,
      fields: {
        name: { oldValue: '', value: '' },
        email: { oldValue: '', value: '' },
        role: { oldValue: Role.USER, value: Role.USER },
      },
      isNew: false,
      isDelete: false,
      selected: false
    };
    Object.keys(user).forEach((key) => {
      const fieldKeys = key as keyof UserDataKey;
      if (fieldKeys in newItem.fields) {
        const val: any = users.item[fieldKeys];
        newItem.fields[fieldKeys] = val as UserField<typeof val>
          console.log(newItem.fields[fieldKeys]);
        }
      });

      return newItem;
    });
  }
  catch(e){
    loading.value = false;
    console.log(e);
  }
}

let newRow: UserData = {
  id: '',
  fields: {
    name: { oldValue: '', value: '' },
    email: { oldValue: '', value: '' },
    role: { oldValue: Role.USER, value: Role.USER },
  },
  isNew: false,
  isDelete: false,
  selected: false,
};

//fix: deleting newly entered values ​​from fields
function createUser() {
  try {
    loading.value = true;
    const row = rows.value[0];
    if (row && row.fields) {
      console.log(row.fields);
    } else {
      console.err2("row or row.fields is undefined");
    }

    const NewUser = {
      id: '',
      fields: {
        name: { oldValue: '', value: newRow.fields.name.value },
        email: { oldValue: '', value: newRow.fields.email.value },
        role: { oldValue: Role.USER, value: newRow.fields.role.value },
      },
      isNew: true,
      isDelete: false,
      selected: false,
    };

    const doNotRepeatUser = rows.value.every(
      (user) =>
        user.fields.name.value
        !== NewUser.fields.name.value ||
        user.fields.email.value
        !== NewUser.fields.email.value
    );

    if (doNotRepeatUser) {
      rows.value.push(NewUser);
    } else {
      console.err2("User with the same data already exists.");
    }

    newRow = {
      id: '',
      fields: {
        name: { oldValue: '', value: '' },
        email: { oldValue: '', value: '' },
        role: { oldValue: Role.USER, value: Role.USER },
      },
      isNew: false,
      isDelete: false,
      selected: false,
    };

    // newRow.fields.name.value = '';
    // newRow.fields.email.value = '';
    // newRow.fields.role.value = Role.USER;
    // newRow.isNew = false;

    createDialog.value = false;
    loading.value = false;
    return doNotRepeatUser;
  }
  catch (e) {
    console.err2(e);
    loading.value = false;
  }
}

//TODO: create edit
function editUserData() {
  loading.value = true;
  const oldRow = rows.value[0];
};

//TODO: create save edited user
// function saveEditedUserData(){
//   loading.value = true;
// }

//TODO: create delete
function deleteUser() {
  loading.value = true;
  // if (row.isNew) {
  //   console.log(row.fields);
  //   rows.value.splice(rowIndex, 1);
  // } else {
  //   row.isDelete = !row.isDelete
  //   console.err2("row has been deleted");
  // }
}

async function deleteSelectedUsers() {
  loading.value = true;
  const selectedRow = newRow.selected;

  if (selectedRow) {
    // rows.value[0].isDelete = !rows.value[0].isDelete
    rows.value = await rows.value.filter((user) => !selectedRow);
    console.log("row has been deleted");
  } else {
    rows.value.splice(0, 1)
  }
  loading.value = false
}

function checkNewUserRow(field: UserField<string>, isNew: boolean) {
  const checkNewUserRow = field.oldValue !== field.value
  if (checkNewUserRow){
    console.log(checkNewUserRow);
    return checkNewUserRow.valueOf()
  }
  const checkOnUserRow = checkNewUserRow
  && field && field.value
  && field.value !== field.oldValue && !isNew

  return checkOnUserRow
}

function navigateToSettings() {
  router.push({ name: 'Settings' });
}

function onSaved(n:User[]){
  console.dir(n)
  console.dir(rows.value)
}

const requiredRules = [
  (v: string) => !!v || 'Required.',
  (v: string) => (v && v.length >= 3) || 'Min 3 characters',
];

const emailRules = [
  (v: string) => !!v || 'E-mail is required',
  (v: string) => /.+@.+\..+/.test(v) || 'E-mail must be valid',
];
</script>

<style>
.dialog-bottom-transition-enter-active,
.dialog-bottom-transition-leave-active {
    transition: transform 0.2s ease-in-out;
  }
.container-wrapper {
  position: relative;
}

.selected {
    background-color: rgba(0, 0, 0, 0.1);
}

.scrollbar__thumb {
    background-color: rgb(97, 97, 97);
}

.color_form {
  background-color: #303030;
}


.v-table3 {
    padding: 0 150px;
    transition: height cubic-bezier(0.4, 0, 0.2, 1);
}

.my-234 {
    margin-top: 6px !important;
    margin-bottom: -38px !important;
}

.table1 {
    height: 550px;
    padding: 12px;
}
</style>
