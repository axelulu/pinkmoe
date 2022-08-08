<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 17:34:13
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 22:43:55
 * @FilePath: /pinkmoe_index/src/components/Authorlayout/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Authorlayout">
  import { useAuthorLayout } from '/@/hooks/author/layout';
  const emit = defineEmits(['userInfo']);

  const { unFollow, follow, lev, menu, userInfo, currentMenu, status, jump } =
    useAuthorLayout(emit);
</script>

<template>
  <!-- Authorlayout -->
  <div class="flex flex-col lg:w-3/4 xl:w-5/12 m-auto">
    <div
      class="bg-cover min-h-43 w-full relative"
      :style="`background-image: url(${userInfo?.headerImg});`"
    >
      <div :style="`backdrop-filter: blur(6px);`" class="pb-4 pt-4 h-full h-full">
        <div class="flex justify-center items-center">
          <img
            class="w-24 h-24 animate-lazyloaded rounded-full border-4 border-white dark:border-gray-700 shadow-lg object-cover"
            v-lazy="userInfo?.avatar"
            alt=""
          />
          <span class="text-xl text-gray-500 dark:text-gray-200 mx-2">{{
            userInfo?.nickName
          }}</span>
          <span
            :style="`background-color: ${userInfo?.authority?.authorityColor};`"
            class="text-xs px-1.5 py-1 text-white"
            >{{ userInfo?.authority?.authorityName }}</span
          >
          <span
            :style="`background-color: ${lev?.color};`"
            class="ml-3 select-none cursor-pointer rounded-xl text-xs text-white px-2 py-0.2"
            >{{ lev?.levelName }}</span
          >
        </div>
        <div class="flex justify-center items-center mt-4">
          <span
            @click="!status ? unFollow() : follow()"
            class="bg-green-500 select-none active:bg-green-500 active:border-green-500 hover:bg-green-600 hover:border-green-600 duration-300 py-1 px-4 cursor-pointer text-xs text-white border-2 border-green-500"
          >
            <font-awesome-icon class="mr-1" :icon="['fas', status ? 'check' : 'add']" />
            <span>{{ status ? '已关注' : '加关注' }}</span>
          </span>
          <a
            @click="jump"
            class="bg-white-500 select-none dark:hover:bg-gray-600 hover:bg-gray-200 active:bg-gray-50 duration-300 dark:border-gray-700 dark:text-gray-200 py-1 px-4 cursor-pointer text-xs text-gray-500 border-2 border-gray-200"
          >
            <font-awesome-icon class="mr-1" :icon="['fas', 'envelope']" />
            <span>站内信</span>
          </a>
        </div></div
      >
    </div>
    <div>
      <div
        class="bg-white text-gray-500 dark:bg-gray-700 dark:text-gray-200 opacity-95 px-3 flex flex-row rounded-md justify-around shadow-sm"
      >
        <router-link
          :to="item.url"
          v-for="(item, index) in menu"
          :key="index"
          :class="
            currentMenu === item.value
              ? 'border-pink-400 dark:border-pink-400'
              : 'dark:border-gray-600'
          "
          class="px-16 border-b-2 w-full flex justify-center items-center py-2.5 cursor-pointer hover:bg-pink-400 hover:text-white duration-300"
        >
          <font-awesome-icon class="text-xs" :icon="item.icon" />
          <span class="ml-1 text-xs">{{ item.name }}</span>
        </router-link>
      </div>
      <slot></slot>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
