<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-08 16:34:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 12:57:11
 * @FilePath: /pinkmoe_index/layouts/home.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup>
let isMobile = false
if (process.client)
  isMobile = !!/(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile)/i.test(navigator.userAgent)
</script>

<template>
  <main v-if="!isMobile" class="bg-slate-100 text-gray-600 dark:bg-gray-900 dark:text-gray-200 font-sans">
    <Header />
    <slot />
    <Footer />
  </main>
  <Mobile v-else />
</template>
