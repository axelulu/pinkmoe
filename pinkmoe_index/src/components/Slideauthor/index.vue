<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 18:48:20
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 11:38:34
 * @FilePath: /pinkmoe_vitesse/src/components/Slideauthor/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Slideauthor">
import { useSlideAuthor } from '/@/hooks/slideAuthor'

const props = defineProps({
  author: {
    type: Object,
    default: null,
  },
  commentCount: {
    type: Number,
    default: 0,
  },
  postCount: {
    type: Number,
    default: 0,
  },
  fansCount: {
    type: Number,
    default: 0,
  },
  followCount: {
    type: Number,
    default: 0,
  },
  followStatus: {
    type: Boolean,
    default: true,
  },
})
const { unFollow, follow, status, jump, lev } = useSlideAuthor(props)
</script>

<template>
  <!-- Slideauthor -->
  <div class="flex flex-row justify-center relative animate-fadeIn30">
    <div
      class="absolute w-full h-30 bg-cover"
      :style="`background-image: url(${author?.headerImg});`"
    />
    <router-link
      :to="`/author/${author?.uuid}/userInfo`"
      class="w-36 h-36 rounded-full absolute mt-14 border-4 border-white shadow-2xl cursor-pointer overflow-hidden"
    >
      <img v-lazy="author?.avatar" class="w-36 h-36 object-cover animate-lazyloaded" alt="">
    </router-link>
    <div
      class="w-full bg-white dark:bg-gray-700 pt-24 mt-30 text-center rounded-md overflow-hidden"
    >
      <router-link
        :to="`/author/${author?.uuid}/userInfo`"
        class="text-lg text-pink-400 mb-2 hover:text-white hover:bg-pink-400 px-1 py-0.5 duration-300 cursor-pointer"
      >
        {{ author?.nickName }}
      </router-link>
      <div class="p-1 flex flex-row justify-center items-center">
        <span class="text-xs text-white flex justify-center items-center bg-green-500 px-1 py-0.5">
          <i class="mr-1 inline-block i-ph:paw-print-fill" />
          {{ author?.credit }}</span>
        <span
          :style="`background-color: ${author?.authority?.authorityColor};`"
          class="text-xs text-white px-1 py-0.5"
        >{{ author?.authority?.authorityName }}</span>
        <span :style="`background-color: ${lev?.color};`" class="text-xs text-white px-1 py-0.5">{{
          lev?.levelName
        }}</span>
      </div>
      <div class="text-xs text-gray-500 dark:text-gray-200 pt-1">
        {{ author?.desc }}
      </div>
      <div class="pb-4 pt-3">
        <span
          class="bg-green-500 select-none active:bg-green-500 active:border-green-500 hover:bg-green-600 hover:border-green-600 duration-300 py-1 px-2 cursor-pointer text-xs text-white border-2 border-green-500"
          @click="!status ? unFollow() : follow()"
        >
          <i :class="`mr-1 i-${status ? 'material-symbols:check-circle' : 'material-symbols:add'}`" />
          <span>{{ status ? '已关注' : '加关注' }}</span>
        </span>
        <span
          class="bg-white-500 select-none dark:hover:bg-gray-600 hover:bg-gray-200 active:bg-gray-50 duration-300 py-1 px-2 cursor-pointer text-xs text-gray-500 dark:text-gray-200 border-2 border-gray-200 dark:border-gray-600"
          @click="jump"
        >
          <i class="mr-1 i-jam:envelope-f" />
          <span>站内信</span>
        </span>
      </div>
      <div
        class="flex flex-row text-xs text-gray-500 dark:text-gray-200 bg-gray-50 dark:bg-gray-700"
      >
        <router-link
          :to="`/author/${author?.uuid}/follow`"
          class="py-2 flex-1 border-r border-t border-gray-100 dark:border-gray-600 cursor-pointer hover:bg-pink-400 hover:border-pink-400 hover:text-white duration-300"
        >
          <div>{{ followCount }}</div>
          <div>关注</div>
        </router-link>
        <router-link
          :to="`/author/${author?.uuid}/fans`"
          class="py-2 flex-1 border-r border-t border-gray-100 dark:border-gray-600 cursor-pointer hover:bg-pink-400 hover:border-pink-400 hover:text-white duration-300"
        >
          <div>{{ fansCount }}</div>
          <div>粉丝</div>
        </router-link>
        <router-link
          :to="`/author/${author?.uuid}/post`"
          class="py-2 flex-1 border-r border-t border-gray-100 dark:border-gray-600 cursor-pointer hover:bg-pink-400 hover:border-pink-400 hover:text-white duration-300"
        >
          <div>{{ postCount }}</div>
          <div>帖子</div>
        </router-link>
        <router-link
          :to="`/author/${author?.uuid}/comment`"
          class="py-2 flex-1 border-r border-t border-gray-100 dark:border-gray-600 cursor-pointer hover:bg-pink-400 hover:border-pink-400 hover:text-white duration-300"
        >
          <div>{{ commentCount }}</div>
          <div>评论</div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
