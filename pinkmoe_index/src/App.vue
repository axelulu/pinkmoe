<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-08 16:34:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 11:55:12
 * @FilePath: /pinkmoe_vitesse/src/App.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script setup lang="ts">
// https://github.com/vueuse/head
// you can use this to manipulate the document head in any components,
// they will be rendered correctly in the html results with vite-ssg
import { useHead } from '@vueuse/head'
import { useAppStore } from './store/modules/app'
import { useSocketStore } from './store/modules/socket'
import './styles/client.css'
const route = useRoute()
let isMobile = false
if (!import.meta.env.SSR)
  isMobile = !!/(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile)/i.test(navigator.userAgent)

const { createWebSocket } = useSocketStore()

createWebSocket()
const i18n = useI18n()

const { siteBasic } = useAppStore()
useHead({
  title: computed(() => `${siteBasic?.title} - ${siteBasic?.desc}`),
  meta: [
    {
      name: 'og:title',
      content: computed(() => `${siteBasic?.title} - ${siteBasic?.desc}`),
    },
    { name: 'og:url', content: computed(() => siteBasic?.url) },
    { name: 'og:locale', content: computed(() => i18n.locale.value) },
    { name: 'og:site_name', content: computed(() => siteBasic?.keyword) },
    { name: 'description', content: computed(() => siteBasic?.desc) },
  ],
})
</script>

<template>
  <div v-if="!isMobile" class="bg-slate-100 text-gray-600 dark:bg-gray-900 dark:text-gray-200 font-sans">
    <Header />
    <RouterView :key="route.path" />
    <Footer />
  </div>
  <Mobile v-else />
</template>
