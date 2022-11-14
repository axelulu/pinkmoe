<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 17:31:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-11-13 22:36:27
 * @FilePath: /pinkmoe_index/pages/author/:author/follow.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="AuthorAuthorFollow">
import AuthorLayout from '/@/components/Authorlayout/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import Spin from '/@/components/Spin/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import { useAuthorFollow } from '/@/hooks/author/follow'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'

const { authorFollowList, loading, hasMore, nextPage } = useAuthorFollow()

const { siteBasic } = useAppStore()
useHead({
  title: '用户关注 - 用户主页',
  link: [
    { rel: 'icon', type: 'image/x-icon', href: `${siteBasic?.icon}` },
  ],
  meta: [
    { name: 'og:type', content: 'follow' },
    {
      name: 'og:title',
      content: '用户关注 - 用户主页',
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
  <div class="flex justify-start flex-wrap mt-4 animate-fadeIn30">
    <Spin :show="loading" class="flex flex-wrap">
      <div
        v-if="authorFollowList && authorFollowList.list?.length"
        class="w-full flex justify-start flex-wrap animate-fadeIn30"
      >
        <div
          v-for="(item, index) in authorFollowList.list"
          :key="index"
          class="p-1 text-center cursor-pointer text-gray-500 hover:bg-pink-50 dark:hover:bg-gray-700 hover:text-pink-400"
        >
          <NuxtLink :to="`/author/${item?.toUidRelation?.uuid}/userInfo`">
            <img
              v-lazy="item?.toUidRelation?.avatar"
              class="rounded-full animate-lazyloaded h-20 w-20 border-2 border-transparent hover:border-pink-400 duration-300 object-cover"
              alt=""
            >
            <div class="text-xs mt-1">
              {{ item?.toUidRelation?.nickName }}
            </div>
          </NuxtLink>
        </div>
        <div class="w-full p-1.5 text-gray-500">
          <MoreBtn v-if="hasMore" @click="nextPage" />
        </div>
      </div>
      <NotFound v-else />
    </Spin>
  </div>
</template>
