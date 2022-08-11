<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 10:35:44
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-10 20:01:32
 * @FilePath: /pinkmoe_index/src/components/Goods/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Article">
  import { useUtil } from '/@/hooks/util';
  defineProps({
    post: {
      type: Object,
      default: null,
    },
    imgHeight: {
      type: String,
      default: '',
    },
  });
  const { formatDate } = useUtil();
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
        <router-link :to="'/post/'">
          <img
            :class="'overflow-hidden animate-lazyloaded object-cover h-40 w-full cursor-pointer object-cover'"
            v-lazy="
              'https://static.7b2.com/wp-content/uploads/2018/06/139e5682282d66.jpg?x-oss-process=image/resize,m_fill,h_600,w_600/sharpen,120/format,webp'
            "
            alt=""
            srcset=""
          />
        </router-link>
        <div class="absolute top-0 right-0 text-xs text-white bg-black bg-opacity-30 px-1 py-0.5">
          <font-awesome-icon :icon="['fas', 'eye']" />
          <span>{{ 1 }}</span>
        </div>
        <router-link
          v-if="1"
          :to="'/category/' + 1"
          class="absolute bottom-0 right-0 text-xs text-white bg-pink-400 px-1 py-0.5 text-center cursor-pointer"
        >
          5折
        </router-link>
      </div>
      <div class="flex flex-row justify-between">
        <router-link
          :to="'/post/' + post.postId"
          class="text-ellipsis flex-1 text-xs p-2 line-clamp-2 h-10 cursor-pointer hover:text-pink-400 dark:hover:text-pink-400 dark:text-gray-200 duration-300"
          >商品一号
        </router-link>
        <router-link
          :to="'/post/' + post.postId"
          class="text-ellipsis w-5/12 text-md p-2 line-clamp-2 h-10 cursor-pointer text-red-500 duration-300"
          >¥100
        </router-link>
      </div>
      <div class="flex flex-row justify-between pl-2 pr-2 pt-1 pb-2 text-xs">
        <router-link
          :to="'/author/' + 1 + '/userInfo'"
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
          >库存:20
        </router-link>
        <div
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
          >售出:1
        </div>
        <div
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
          >人气:100
        </div>
      </div>
    </div>
  </article>
</template>

<style lang="less" scoped></style>
