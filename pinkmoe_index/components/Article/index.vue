<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 10:35:44
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 14:01:57
 * @FilePath: /pinkmoe_index/components/Article/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Article">
import { useUtil } from '/@/hooks/util'
defineProps({
  post: {
    type: Object,
    default: null,
  },
  imgHeight: {
    type: String,
    default: '',
  },
})
const { formatDate, getThumbnail } = useUtil()
</script>

<template>
  <!-- Article -->
  <article
    class="border-2 duration-300 border-transparent hover:border-pink-400 rounded-md overflow-hidden"
  >
    <div
      class="hover:bg-pink-50 dark:bg-gray-700 bg-white dark:hover:bg-gray-800 shadow-sm w-full h-full"
    >
      <div class="relative">
        <router-link :to="`/post/${post.postId}`">
          <img
            v-lazy="getThumbnail(post.cover)"
            :class="
              `overflow-hidden animate-lazyloaded object-cover ${
                imgHeight
              } w-full cursor-pointer object-cover`
            "
            alt=""
            srcset=""
          >
        </router-link>
        <div class="absolute top-0 right-0 text-xs flex justify-center items-center text-white bg-black bg-opacity-30 px-1 py-0.5">
          <i class="inline-block i-ant-design:eye-filled" />
          <span>{{ post.view }}</span>
        </div>
        <router-link
          v-if="post.CategoryRelation.slug"
          :to="`/category/${post.CategoryRelation.slug}`"
          class="absolute bottom-0 text-xs text-white bg-sky-400 bg-opacity-30 px-1 py-0.5 text-center cursor-pointer"
        >
          {{ post.CategoryRelation.name }}
        </router-link>
      </div>
      <router-link
        :to="`/post/${post.postId}`"
        class="text-ellipsis text-xs p-2 line-clamp-2 h-10 cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300"
      >
        {{ post.type === 'active' ? post.content : post.title }}
      </router-link>
      <div class="flex flex-row justify-between pl-2 pr-2 pt-1 pb-2 text-xs">
        <router-link
          :to="`/author/${post.AuthorRelation.uuid}/userInfo`"
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
        >
          {{ post.AuthorRelation.nickName }}
        </router-link>
        <div
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
        >
          {{ formatDate(post?.UpdatedAt) }}
        </div>
      </div>
    </div>
  </article>
</template>
