import { gql } from 'graphql-tag';


export const MUTATION = gql`
  mutation {
    validate {
      id
      login
      role
    }
  }
`;

export const Role = gql`
  enum Role {
    ADMIN
    USER
  }
`;
