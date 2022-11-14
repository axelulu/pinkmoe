<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 17:31:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-11-13 22:36:20
 * @FilePath: /pinkmoe_index/pages/author/:author/userinfo.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="AuthorAuthorUserinfo">
import { useHead } from '@vueuse/head'
import { getUserInfo } from '/@/api/user'
import { useAppStore } from '/@/store/modules/app'
const detail = ref()
const loading = ref(false)
function userInfos(res) {
  detail.value = [
    {
      key: '昵称',
      value: res?.nickName,
      icon: [],
    },
    {
      key: 'UID',
      value: res?.uuid,
      icon: [],
    },
    {
      key: '描述',
      value: res?.desc,
      icon: [],
    },
    {
      key: '积分',
      value: res?.credit,
      icon: [],
    },
    {
      key: '经验',
      value: res?.exp,
      icon: [],
    },
    {
      key: '性别',
      value: res?.sex === 1 ? '保密' : res?.sex === 2 ? '女' : '男',
      icon: [],
    },
  ]
}

onMounted(async () => {
  loading.value = true
  const route = useRoute()
  const userInfo = await getUserInfo({
    uuid: route.params.author,
  })
  console.log(userInfo)
  userInfos(userInfo)
  setTimeout(() => {
    loading.value = false
  }, 300)
})

const { siteBasic } = useAppStore()
useHead({
  title: '用户信息 - 用户主页',
  link: [
    { rel: 'icon', type: 'image/x-icon', href: `${siteBasic?.icon}` },
  ],
  meta: [
    { name: 'og:type', content: 'userInfo' },
    {
      name: 'og:title',
      content: '用户信息 - 用户主页',
    },
    {
      name: 'og:keywords',
      content: `${siteBasic?.keyword}`,
    },
    { name: 'og:description', content: siteBasic?.desc },
    { name: 'og:url', content: siteBasic?.url },
  ],
})

definePageMeta({
  layout: 'author',
})
</script>

<template>
  <Spin :show="loading" class="flex flex-wrap">
    <div class="mt-4 bg-white w-full dark:bg-gray-700 py-4 animate-fadeIn30 min-h-68">
      <div
        v-for="(item, index) in detail"
        :key="index"
        class="px-4 py-3 flex flex-row hover:bg-gray-100 dark:hover:bg-gray-800 duration-300"
      >
        <div class="text-xs text-gray-500 dark:text-gray-200 w-28 font-semibold">
          {{
            item.key
          }}
        </div>
        <div class="text-xs text-gray-500 dark:text-gray-200">
          <i
            v-if="item.icon.length > 0" :class="`text-gray-700 dark:text-gray-200 text-xs inline-block i-${item.icon}`"
          />
          <span class="ml-1">{{ item.value }}</span>
        </div>
      </div>
    </div>
  </Spin>
</template>
