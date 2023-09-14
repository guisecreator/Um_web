import { createRouter, createWebHistory, RouteRecordRaw  } from 'vue-router'
import { useUserStore } from '@/store/store-users';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    children: [
      {
        path: '/',
        name: 'Login',
        component: () => import('../views/AuthPage.vue'),
        meta: {
          requiresAuth: false,
          title: 'Login'
        }
      },
      {
        path: '/Control-Panel',
        name: 'Index',
        component: () => import('../views/PanelPage.vue'),
        meta: { requiresAuth: true,
          title: 'Control Panel'
        }
      },
      {
        path: '/Settings',
        name: 'Settings',
        component: () => import('../views/SettingsPage.vue'),
        meta: { requiresAuth: true,
          title: 'Settings'
        }
      },
    ],
  },
  {
    path: '/:catchAll(.*)*',
    component: () => import('../views/PageNotFound.vue'),
    meta: { requiresAuth: false,
      title: '404 Page Not Found'
    }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

router.beforeEach((to, from, next) => {
  document.title = (to.meta.title as string) || 'Default Title';
  next();
});

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

export default router
