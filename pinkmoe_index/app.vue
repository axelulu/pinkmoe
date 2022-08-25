<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:53:10
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 15:22:36
 * @FilePath: /pinkmoe_index/app.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script setup>
import { useAppStore } from './store'
const key = computed(() => route.name !== undefined ? route.name + +new Date() : route + +new Date())

const { siteBasic } = useAppStore()
useHead({
  link: [
    {
      rel: 'icon',
      href: siteBasic?.icon,
    },
  ],
  meta: [
    {
      name: 'og:title',
      content: computed(() => `${siteBasic?.title} - ${siteBasic?.desc}`),
    },
    { name: 'og:url', content: computed(() => siteBasic?.url) },
    { name: 'og:site_name', content: computed(() => siteBasic?.keyword) },
    { name: 'description', content: computed(() => siteBasic?.desc) },
  ],
})
</script>

<template>
  <NuxtLayout name="home">
    <NuxtPage :page-key="key" />
  </NuxtLayout>
</template>

<style>
html, body , #__nuxt{
  height: 100vh;
  margin: 0;
  padding: 0;
}

html.dark {
  background: #222;
  color: white;
}
</style>
