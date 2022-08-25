<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-08 16:34:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-23 21:52:08
 * @FilePath: /pinkmoe_index/src/layouts/home.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup>
let isMobile = false
if (!import.meta.env.SSR)
  isMobile = !!/(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile)/i.test(navigator.userAgent)
</script>

<template>
  <main v-if="!isMobile" class="bg-slate-100 text-gray-600 dark:bg-gray-900 dark:text-gray-200 font-sans">
    <Header />
    <RouterView />
    <Footer />
  </main>
  <Mobile v-else />
</template>
