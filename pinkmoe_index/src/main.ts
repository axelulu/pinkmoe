/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:03:13
 * @FilePath: /pinkmoe_index/src/main.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import 'virtual:windi-base.css';
import 'virtual:windi-components.css';
import 'virtual:windi-utilities.css';

import App from './App.vue';
import './index.css';
//改为element-plus默认UI
import 'element-plus/dist/index.css';

// FontAwesomeIcon
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import Search from '/@/components/Search/index';
import Login from '/@/components/Login/index';
import VipConfirm from '/@/components/Vipconfirm/index';
import ShopConfirm from '/@/components/Shopconfirm/index';
import GlobalNotice from '/@/components/GlobalNotice/index';
import Message from '/@/components/Message/index';

import './utils/fontawsome';
import VueLazyLoad from 'vue3-lazyload';
import 'moment/dist/locale/zh-cn';

// 支持SVG
import 'virtual:svg-icons-register';
import 'nprogress/nprogress.css';
import vue3PhotoPreview from 'vue3-photo-preview';
import 'vue3-photo-preview/dist/index.css';
import vueI18n from '/@/locales';
import routes from 'virtual:generated-pages';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import { useAppStore, useUserStore } from './store';
import { ViteSSG } from 'vite-ssg';
import VueAPlayer from 'vue3-aplayer';
import NPlayer from '@nplayer/vue';

// vitessg模式
ViteSSG(
  // the root component
  App,
  // vue-router options
  { routes },
  // function to have custom setups
  ({ app, router, initialState, isClient, onSSRAppRendered }) => {
    // install plugins etc.
    if (isClient) {
      router.beforeEach(async (_to, _from, next) => {
        NProgress.start();
        const { settings, siteBasic, siteFooter, userLevel } = useAppStore();
        if (!siteBasic || !siteFooter || !userLevel) {
          await settings();
        }
        const { info, userInfo, isLogin } = useUserStore();
        if (_to.meta.auth && !isLogin) {
          next('/');
        }
        if (!userInfo && isLogin) {
          await info();
        }
        next();
      });

      router.afterEach((_to) => {
        NProgress.done();
      });
    }

    // pinia
    const pinia = createPinia();
    app.use(pinia);
    if (isClient) {
      pinia.state.value = initialState.pinia || {};
    } else {
      onSSRAppRendered(() => {
        initialState.pinia = pinia.state.value;
      });
    }
    app
      .component('FontAwesomeIcon', FontAwesomeIcon)
      .component('aplayer', VueAPlayer)
      .use(VueLazyLoad, {
        preLoad: 1.3, //距离当前dom距离页面底部的高度
        error: 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7', //加载失败显示的图片
        loading: 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7', //加载中显示的图片
        attempt: 1, // 图片加载失败，最多重试的次数
      })
      .use(NPlayer)
      .use(vueI18n)
      .use(vue3PhotoPreview)
      .use(VipConfirm)
      .use(ShopConfirm)
      .use(GlobalNotice)
      .use(Search)
      .use(Login)
      .use(Message);
  },
);

// 默认模式
// createApp(App)
//   .component('font-awesome-icon', FontAwesomeIcon)
//   .use(VueLazyLoad, {
//     preLoad: 1.3, //距离当前dom距离页面底部的高度
//     error: 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7', //加载失败显示的图片
//     loading: 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7', //加载中显示的图片
//     attempt: 1, // 图片加载失败，最多重试的次数
//   })
//   .use(Search)
//   .use(Login)
//   .use(Message)
//   .use(router)
//   .use(piniaStore)
//   .use(vueI18n)
//   .use(vue3PhotoPreview)
//   .mount('#app');
