<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:03:24
 * @FilePath: /pinkmoe_index/src/App.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <div v-if="!isMobile" class="bg-slate-100 dark:bg-gray-900 dark:text-gray-200 font-sans">
    <Header />
    <router-view :key="route.path" />
    <Footer />
  </div>
  <Mobile v-else />
</template>
<script setup lang="ts">
  import { useAppStore } from './store/modules/app';
  import { useHead } from '@vueuse/head';
  import { useSocketStore } from './store/modules/socket';
  const route = useRoute();
  const isMobile = /(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile)/i.test(navigator.userAgent)
    ? true
    : false;

  const { createWebSocket } = useSocketStore();

  createWebSocket();

  const { siteBasic } = useAppStore();
  useHead({
    meta: [
      { name: 'og:locale', content: 'zh_CN' },
      { name: 'og:site_name', content: siteBasic?.keyword },
      { name: 'description', content: siteBasic?.desc },
    ],
  });
</script>

<style></style>
