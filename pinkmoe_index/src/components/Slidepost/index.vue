<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 18:50:47
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:39
 * @FilePath: /pinkmoe_index/src/components/Slidepost/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Slidepost">
  defineProps({
    posts: {
      type: Array,
      default: null,
    },
  });
</script>

<template>
  <!-- Slidepost -->
  <div class="flex flex-wrap animate-fadeIn30">
    <div v-for="(item, index) in posts" :key="index" class="w-1/2 px-1 my-1">
      <div
        class="border-2 border-transparent hover:border-pink-400 duration-300 hover:bg-pink-50 rounded-md overflow-hidden"
      >
        <router-link :to="'/post/' + item.postId">
          <div class="bg-white dark:bg-gray-700 rounded-md overflow-hidden">
            <img
              class="h-28 w-full animate-lazyloaded object-cover cursor-pointer object-cover"
              v-lazy="item.cover ? '/' + item.cover : ''"
              alt=""
            />
            <div
              class="text-ellipsis text-xs px-2 mb-2 pt-2 line-clamp-2 h-10 cursor-pointer hover:text-pink-400 duration-300"
            >
              {{ item.title }}
            </div>
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
