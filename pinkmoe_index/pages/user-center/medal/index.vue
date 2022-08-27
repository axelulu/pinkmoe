<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:11:15
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 18:14:52
 * @FilePath: /pinkmoe_index/pages/user-center/medal/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="UserCenterMedalIndex">
import UserCenterLayout from '/@/components/Usercenterlayout/index.vue'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'

definePageMeta({
  middleware: ['user-auth'],
})

const { siteBasic } = useAppStore()
useHead({
  title: '荣誉 - 用户中心',
  meta: [
    { name: 'og:type', content: 'medal' },
    {
      name: 'og:title',
      content: '荣誉 - 用户中心',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <UserCenterLayout>
    <div class="ml-6">
      <div
        class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <i class="inline-block i-material-symbols:bookmark" />
          <span class="ml-1 select-none">我的荣誉</span>
        </div>
        <div class="p-4">
          <div
            class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
          >
            <p>这里会展示我的荣誉！</p>
          </div>
        </div>
        <div class="flex flex-row items-center px-4 pb-4" />
      </div>
    </div>
  </UserCenterLayout>
</template>
