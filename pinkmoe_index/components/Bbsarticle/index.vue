<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 08:22:26
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 23:43:38
 * @FilePath: /pinkmoe_index/components/Bbsarticle/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Bbsarticle">
import { useBbsArticle } from '/@/hooks/bbs/bbsArticle'
import { useUtil } from '/@/hooks/util'
const props = defineProps({
  post: {
    type: Object,
    default: null,
  },
})
const { formatDate } = useUtil()
const {
  status,
  unFollow,
  follow,
  collect,
  unCollect,
  isCollect,
  auth,
  jump,
  deletePostImg,
  lev,
} = useBbsArticle(props.post)
</script>

<template>
  <article
    class="p-6 flex flex-row dark:bg-gray-700 dark:text-gray-200 text-gray-500 bg-white rounded-md mb-3"
  >
    <NuxtLink :to="`/author/${post?.AuthorRelation?.uuid}/userInfo`" class="w-18">
      <img
        v-lazy="post?.AuthorRelation?.avatar"
        class="w-12 h-12 animate-lazyloaded rounded-full overflow-hidden"
        alt=""
      >
    </NuxtLink>
    <div class="w-full flex flex-col ml-0">
      <div class="flex justify-between items-center">
        <div class="flex flex-col">
          <div class="flex flex-row justify-center items-center">
            <NuxtLink
              :to="`/author/${post?.AuthorRelation?.uuid}/userInfo`"
              class="dark:text-white hover:text-pink-400 duration-300 text-gray-700 text-md"
            >
              {{ post?.AuthorRelation?.nickName }}
            </NuxtLink>
            <span
              :style="`background-color: ${lev?.color};`"
              class="ml-3 select-none cursor-pointer rounded-xl text-xs text-white px-2 py-0.2"
            >{{ lev?.levelName }}</span>
            <span
              :style="`background-color: ${post?.AuthorRelation?.authority?.authorityColor};`"
              class="ml-3 select-none cursor-pointer rounded-xl text-xs text-white px-1 py-0.2"
            >
              {{ post?.AuthorRelation?.authority?.authorityName }}</span>
          </div>
          <div class="text-sm mt-1 text-gray-400">
            {{ formatDate(post?.UpdatedAt) }}更新 - {{ post?.view }}次阅读
          </div>
        </div>
        <div v-if="post?.AuthorRelation?.uuid !== auth.userInfo?.uuid" class="flex flex-row">
          <span
            class="bg-green-500 select-none flex justify-center items-center active:bg-green-500 active:border-green-500 hover:bg-green-600 hover:border-green-600 duration-300 py-1 px-2 cursor-pointer text-xs text-white border-2 border-green-500"
            @click="!status ? unFollow() : follow()"
          >
            <i :class="`mr-1 inline-block ${status ? 'i-material-symbols:check-small' : 'i-material-symbols:add'}`" />
            <span>{{ status ? '已关注' : '加关注' }}</span>
          </span>
          <span
            class="bg-white-500 select-none flex justify-center items-center dark:hover:bg-gray-600 hover:bg-gray-200 active:bg-gray-50 duration-300 py-1 px-2 cursor-pointer text-xs text-gray-500 dark:text-gray-200 border-2 border-gray-200 dark:border-gray-600"
            @click="jump"
          >
            <i class="mr-1 inline-block i-jam:envelope-f" />
            <span>站内信</span>
          </span>
        </div>
      </div>
      <div class="mt-2">
        <div v-if="post?.title" class="pb-2 pt-1">
          <NuxtLink
            :to="`/post/${post?.postId}`"
            class="text-md font-semibold hover:text-gray-700 duration-300 dark:hover:text-gray-50"
          >
            {{ post?.title }}
          </NuxtLink>
        </div>
        <div
          class="text-sm max-h-36 overflow-hidden"
          v-html="post?.type === 'active' ? post?.content : deletePostImg(post?.content)"
        />
        <div class="flex flex-wrap mt-4">
          <photo-provider :default-backdrop-opacity="0.3" :loop="true">
            <photo-consumer
              v-for="src in post?.fileRelation"
              :key="src"
              class="w-1/4"
              :intro="src.name"
              :src="src.url"
            >
              <img
                v-lazy="src.url"
                style="cursor: zoom-in"
                class="h-full w-full mx-2 animate-lazyloaded object-cover overflow-hidden"
                :alt="src.name"
              >
            </photo-consumer>
          </photo-provider>
        </div>
      </div>
      <div class="flex justify-between items-start mt-4">
        <div class="flex flex-wrap w-8/12">
          <NuxtLink
            v-for="(topic, v) in post?.topic"
            :key="v"
            :to="`/topic/${topic.value}`"
            class="border border-pink-400 text-xs mb-2 text-pink-400 p-1 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer mr-2"
          >
            {{ topic.label }}
          </NuxtLink>
        </div>
        <div class="flex flex-row text-xs">
          <div
            class="mx-2 cursor-pointer hover:text-pink-400 duration-300 flex justify-start items-center"
            :class="isCollect ? 'text-pink-400' : ''"
            @click="isCollect ? unCollect() : collect()"
          >
            <i class="mr-1 inline-block i-ic:baseline-star" />
            <span>{{
              isCollect ? '已收藏' : '收藏'
            }}</span>
          </div>
          <NuxtLink
            :to="`/post/${post?.postId}`"
            class="mx-2 cursor-pointer hover:text-pink-400 duration-300 flex justify-start items-center"
          >
            <i class="mr-1 inline-block i-ant-design:comment-outlined" />
            <span>评论</span>
          </NuxtLink>
          <div class="mx-2 cursor-pointer hover:text-pink-400 duration-300 flex justify-start items-center">
            <i class="mr-1 inline-block i-material-symbols:share" />
            <span>分享</span>
          </div>
        </div>
      </div>
    </div>
  </article>
</template>
