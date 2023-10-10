import { createRouter, createWebHistory, RouteRecordRaw  } from 'vue-router'
import { useUserStore } from '@/store/store-users';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    children: [
      {
        path: '/',
        name: 'Login',
        component: () => import('../views/login.vue'),
        meta: {
          requiresAuth: false,
          title: 'Login'
        }
      },
      {
        path: '/Dashboard',
        name: 'Index',
        component: () => import('../views/dashboard.vue'),
        meta: { requiresAuth: true,
          title: 'Dashboard'
        }
      },
      {
        path: '/Settings',
        name: 'Settings',
        component: () => import('../views/settings.vue'),
        meta: { requiresAuth: true,
          title: 'Settings'
        }
      },
    ],
  },
  {
    path: '/:catchAll(.*)*',
    component: () => import('../views/not_found.vue'),
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

  const userStore = useUserStore();

  if (to.meta.requiresAuth) {
    if (!userStore.isLoggedIn) {
      next({ name: 'Login' });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router
