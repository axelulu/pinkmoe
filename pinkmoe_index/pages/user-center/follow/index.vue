<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:10:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-11-13 22:36:56
 * @FilePath: /pinkmoe_index/pages/user-center/follow/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="UserCenterFollowIndex">
import UserCenterLayout from '/@/components/Usercenterlayout/index.vue'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'

definePageMeta({
  middleware: ['user-auth'],
  layout: 'user-center',
})

const { siteBasic } = useAppStore()
useHead({
  title: '关注 - 用户中心',
  link: [
    { rel: 'icon', type: 'image/x-icon', href: `${siteBasic?.icon}` },
  ],
  meta: [
    { name: 'og:type', content: 'follow' },
    {
      name: 'og:title',
      content: '关注 - 用户中心',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <div class="ml-6">
    <div
      class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
    >
      <div class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
        <i class="inline-block i-fa6-solid:basket-shopping" />
        <span class="ml-1 select-none">我的关注</span>
      </div>
      <div class="p-4">
        <div
          class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
        >
          <p>这里会展示我的关注！</p>
        </div>
      </div>
      <div class="flex flex-row items-center px-4 pb-4" />
    </div>
  </div>
</template>
