<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 18:50:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 11:39:16
 * @FilePath: /pinkmoe_vitesse/src/components/Slidecomment/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Slidecomment">
import { useUtil } from '/@/hooks/util'
defineProps({
  comments: {
    type: Array,
    default: null,
  },
})
const { formatDate } = useUtil()
</script>

<template>
  <!-- Slidecomment -->
  <div class="mt-4 animate-fadeIn30">
    <div class="flex flex-row justify-between mx-1">
      <div class="text-xs text-gray-500 flex justify-center items-center dark:text-gray-200">
        <i class="mr-1 inline-block i-ph:user-fill" />
        <span>最新吐槽</span>
      </div>
    </div>
    <div
      v-for="(item, index) in comments"
      :key="index"
      class="flex flex-row h-full w-full p-2 bg-white dark:bg-gray-700 dark:text-gray-200 rounded-md duration-300 my-1 border-2 border-transparent hover:border-pink-400 text-center cursor-pointer text-gray-500 hover:bg-pink-50"
    >
      <router-link
        :to="`/author/${item.FormUidRelation.uuid}/userInfo`"
        class="mr-2 w-18 h-full"
      >
        <img
          v-lazy="item.FormUidRelation.avatar"
          class="rounded-full animate-lazyloaded border-2 border-transparent hover:border-pink-400 duration-300 w-12 h-12 object-cover"
          alt=""
        >
      </router-link>
      <div class="text-xs w-full">
        <div class="flex flex-row justify-between w-full">
          <router-link
            :to="`/author/${item.FormUidRelation.uuid}/userInfo`"
            class="text-sm text-pink-400"
          >
            {{ item.FormUidRelation.nickName }}
          </router-link>
          <div>{{ formatDate(item?.UpdatedAt) }}</div>
        </div>
        <div class="text-left mt-1 text-ellipsis text-xs line-clamp-3">
          {{ item.content }}
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
