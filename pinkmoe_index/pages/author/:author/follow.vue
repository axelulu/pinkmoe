<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 17:31:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-23 21:54:16
 * @FilePath: /pinkmoe_index/src/pages/author/:author/follow.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
  meta: [
    { name: 'og:type', content: 'follow' },
    {
      name: 'og:title',
      content: '用户关注 - 用户主页',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <!-- AuthorAuthorFollow -->
  <AuthorLayout>
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
            <router-link :to="`/author/${item?.toUidRelation?.uuid}/userInfo`">
              <img
                v-lazy="item?.toUidRelation?.avatar"
                class="rounded-full animate-lazyloaded h-20 w-20 border-2 border-transparent hover:border-pink-400 duration-300 object-cover"
                alt=""
              >
              <div class="text-xs mt-1">
                {{ item?.toUidRelation?.nickName }}
              </div>
            </router-link>
          </div>
          <div class="w-full p-1.5 text-gray-500">
            <MoreBtn v-if="hasMore" @click="nextPage" />
          </div>
        </div>
        <NotFound v-else />
      </Spin>
    </div>
  </AuthorLayout>
</template>
