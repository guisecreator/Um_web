export type LOGIN_MUTATION = {
  login: {
    user: {
      id: string;
      email: string;
      password: string;
      role: string;
      createAt: string;
      updateAt: string;
      deletedAt: string;
    };
    info: {
      Token: string;
    }
  };
};

export enum Role {
  MODERATOR = 'MODERATOR',
  ADMIN = 'ADMIN',
  USER = 'USER',
}

export type MUTATION = {
  validate: {
    id: string;
    email: string;
    role: string;
  };
  userNew: {
    id: string;
    email: string;
    password: string;
    role: string;
  }[];
  userDelete: {
    id: string;
    email: string;
    role: string;
  }[];
  userUpdate: {
    id: string;
    email: string;
    role: string;
  }[];
  userSelect: {
    id: string;
    email: string;
    role: string;
  }[];
};

export type MutationVariables = {};

export type MutationResult = {
  data: MUTATION;
};

export type User = {
  id: string;
  role: Role;
  password: string;
  email: string;
};

