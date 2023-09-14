// apollo.d.ts

import { ApolloClient, NormalizedCacheObject } from '@apollo/client/core'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $apollo: ApolloClient<NormalizedCacheObject>
  }
}

declare module 'vue' {
  interface App {
    $apollo: ApolloClient<NormalizedCacheObject>
  }
}
