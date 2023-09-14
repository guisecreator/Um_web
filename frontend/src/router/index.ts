import {
  createMemoryHistory,
  createRouter,
  createWebHashHistory,
  createWebHistory,
  RouteRecordRaw,
} from 'vue-router';

import routes from './routes';
import { useUserStore } from '@/store/store-users';

export default function route () {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : process.env.VUE_ROUTER_MODE === 'history'
    ? createWebHistory
    : createWebHashHistory;

  // const router = createRouter({
  //   history: createWebHistory(process.env.BASE_URL),
  //   routes,
  // })


  
  // const router = createRouter({
  //   scrollBehavior: () => ({ left: 0, top: 0 }),
  //   routes: routes as unknown as RouteRecordRaw[],
  //   history: createHistory(process.env.VUE_ROUTER_BASE),
  // });

  // const userStore = useUserStore();

  // router.beforeEach((to, from, next) => {
  //   document.title = (to.meta.title as string) || 'Default Title';

  //   const userStore = useUserStore();

  //   if (to.meta.requiresAuth) {
  //     if (!userStore.isLoggedIn) {
  //       next({ name: 'Login' });
  //     } else {
  //       next();
  //     }
  //   } else {
  //     next();
  //   }
  // });

  // return router;
}

//export default router
