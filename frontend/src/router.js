import Vue from 'vue';
import VueRouter from 'vue-router';
import AppHeader from './layout/AppHeader.vue';
import AppFooter from './layout/AppFooter.vue';
import Page404 from './views/Page404.vue';
import Home from './views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    components: {
      header: AppHeader,
      default: Home,
      footer: AppFooter,
    },
  },
  {
    path: '*',
    name: '404',
    components: {
      header: AppHeader,
      default: Page404,
      footer: AppFooter,
    },
  },
];

const router = new VueRouter({
  routes,
});

export default router;
