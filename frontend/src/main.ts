/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'

const app = createApp(App)

registerPlugins(app)

app.mount('#app')

import { ApolloClients } from '@vue/apollo-composable'
import { apolloClient } from '@/apollo/apollo'
app.provide(ApolloClients, apolloClient);

