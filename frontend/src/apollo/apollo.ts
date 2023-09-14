import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client/core';
import { useApolloClient } from '@vue/apollo-composable';

const httpLink = createHttpLink({
  uri: 'http://localhost:3020/graphql',
});

export const apolloClient = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
  connectToDevTools: true,
  credentials: 'include',
  defaultOptions: {
    query: {
      fetchPolicy: 'no-cache',
    },
  },
});

export const useApollo = () => useApolloClient(typeof apolloClient);
