<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:10:18
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 18:14:33
 * @FilePath: /pinkmoe_index/pages/user-center/comments/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="UserCenterCommentsIndex">
import UserCenterLayout from '/@/components/Usercenterlayout/index.vue'
import Spin from '/@/components/Spin/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import { useUserCenterComments } from '/@/hooks/user-center/comments'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'
import { useUtil } from '/@/hooks/util'
const { formatDate } = useUtil()
const { authorCommentList, hasMore, loading, nextPage } = useUserCenterComments()

definePageMeta({
  middleware: ['user-auth'],
})

const { siteBasic } = useAppStore()
useHead({
  title: '评论 - 用户中心',
  meta: [
    { name: 'og:type', content: 'comment' },
    {
      name: 'og:title',
      content: '评论 - 用户中心',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <!-- UserCenterCommentsIndex -->
  <UserCenterLayout>
    <div class="ml-6">
      <div
        class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <i class="inline-block i-fa-solid:comments" />
          <span class="ml-1 select-none">我的吐槽</span>
        </div>
        <div class="p-4">
          <div
            class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
          >
            <p>这里会展示您的一些评论！</p>
          </div>
        </div>
        <div class="flex flex-wrap items-center px-4 pb-4">
          <Spin :show="loading">
            <div v-if="authorCommentList?.list && authorCommentList?.list?.length" class="w-full">
              <div
                v-for="(item, index) in authorCommentList?.list"
                :key="index"
                class="flex justify-between flex-row items-center text-xs w-full text-gray-500 dark:text-gray-200 dark:hover:bg-gray-800 py-2 px-2 duration-300 cursor-pointer hover:bg-pink-100"
              >
                <div class="flex flex-row justify-center items-center">
                  <div>
                    您对
                    <router-link class="text-pink-400" :to="`/post/${item.postRelation.postId}`">
                      《{{ item.postRelation.title }}》
                    </router-link>
                    进行了一次吐槽：
                    <span v-html="item.content" />
                  </div>
                </div>
                <div class="ml-4 min-w-16">
                  {{ formatDate(item.UpdatedAt) }}
                </div>
              </div>
            </div>
            <NotFound v-else />
          </Spin>
        </div>
      </div>
      <div class="w-full text-gray-500">
        <MoreBtn v-if="hasMore" @click="nextPage" />
      </div>
    </div>
  </UserCenterLayout>
</template>
