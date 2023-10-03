export type LOGIN_MUTATION = {
  login: {
    user: {
      id: string;
      login: string;
      role: string;
    };
  };
};

export enum Role {
  ADMIN = 'ADMIN',
  USER = 'USER',
}

export type MUTATION = {
  validate: {
    id: string;
    login: string;
    role: string;
  };
  userNew: {
    id: string;
    login: string;
    role: string;
  }[];
  userDelete: {
    id: string;
    login: string;
    role: string;
  }[];
  userUpdate: {
    id: string;
    login: string;
    role: string;
  }[];
  userSelect: {
    id: string;
    login: string;
    role: string;
  }[];
};

export type MutationVariables = {};

export type MutationResult = {
  data: MUTATION;
};

export type User = {
  id: string;
  name: string;
  email: string;
  role: Role;
};


// const response = await this.$apollo.mutate({
//     mutation: gql`
//       mutation Login($email: String!, $password: String!) {
//         login(email: $email, password: $password) {
//         }
//       }
//     `,
//     variables: {
//       email: 'example@example.com',
//       password: 'password',
//     },
//   });
