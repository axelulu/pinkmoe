<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 17:31:23
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 15:39:00
 * @FilePath: /pinkmoe_index/pages/author/:author/comment.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="AuthorAuthorComment">
import AuthorLayout from '/@/components/Authorlayout/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import Spin from '/@/components/Spin/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import { useAuthorComment } from '/@/hooks/author/comment'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'
import { useUtil } from '/@/hooks/util'
const { formatDate } = useUtil()

const { authorCommentList, loading, hasMore, nextPage } = useAuthorComment()

const { siteBasic } = useAppStore()
useHead({
  title: '用户评论 - 用户主页',
  meta: [
    { name: 'og:type', content: 'comment' },
    {
      name: 'og:title',
      content: '用户评论 - 用户主页',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})

definePageMeta({
  layout: 'author',
})
</script>

<template>
  <Spin :show="loading" class="mt-4">
    <div
      v-if="authorCommentList && authorCommentList?.list?.length"
      class="w-full animate-fadeIn30"
    >
      <div
        v-for="(item, index) in authorCommentList.list"
        :key="index"
        class="flex flex-row justify-between items-center dark:hover:bg-gray-800 bg-white dark:bg-gray-700 px-4 py-2 hover:bg-gray-100 duration-300"
      >
        <div class="flex flex-col justify-start">
          <NuxtLink
            :to="`/post/${item?.postRelation?.postId}`"
            class="text-sm text-gray-500 hover:text-pink-400 dark:text-gray-200 duration-300"
          >
            {{ item?.postRelation?.title }}
          </NuxtLink>
          <div class="mt-2">
            <span
              class="text-xs text-gray-500 dark:text-gray-200 dark:bg-gray-800 bg-gray-100 px-2 py-1"
              v-html="item.content"
            />
          </div>
        </div>
        <div class="text-xs">
          {{ formatDate(item?.UpdatedAt) }}
        </div>
      </div>
      <div class="w-full py-3 text-gray-500 dark:text-gray-200">
        <MoreBtn v-if="hasMore" @click="nextPage" />
      </div>
    </div>
    <NotFound v-else />
  </Spin>
</template>
