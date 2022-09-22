<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 17:31:32
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 15:39:10
 * @FilePath: /pinkmoe_index/pages/author/:author/fans.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="AuthorAuthorFans">
import AuthorLayout from '/@/components/Authorlayout/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import Spin from '/@/components/Spin/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import { useAuthorFans } from '/@/hooks/author/fans'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'

const { authorFansList, loading, hasMore, nextPage } = useAuthorFans()

const { siteBasic } = useAppStore()
useHead({
  title: '用户粉丝 - 用户主页',
  meta: [
    { name: 'og:type', content: 'fans' },
    {
      name: 'og:title',
      content: '用户粉丝 - 用户主页',
    },
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
        v-if="authorFansList && authorFansList.list?.length"
        class="w-full flex justify-start flex-wrap animate-fadeIn30"
      >
        <div
          v-for="(item, index) in authorFansList.list"
          :key="index"
          class="p-1 text-center cursor-pointer text-gray-500 hover:bg-pink-50 dark:hover:bg-gray-700 hover:text-pink-400"
        >
          <NuxtLink :to="`/author/${item.formUidRelation?.uuid}/userInfo`">
            <img
              v-lazy="item.formUidRelation?.avatar"
              class="rounded-full animate-lazyloaded h-20 w-20 border-2 border-transparent hover:border-pink-400 duration-300 object-cover"
              alt=""
            >
            <div class="text-xs mt-1">
              {{ item.formUidRelation?.nickName }}
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
