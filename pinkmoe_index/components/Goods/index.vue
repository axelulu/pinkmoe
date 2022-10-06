<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 10:35:44
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:21
 * @FilePath: /pinkmoe_index/components/Goods/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Article">
import { useUtil } from '/@/hooks/util'
defineProps({
  goods: {
    type: Object,
    default: null,
  },
  imgHeight: {
    type: String,
    default: '',
  },
})
const { formatDate } = useUtil()
</script>

<template>
  <article
    class="border-2 duration-300 border-transparent hover:border-pink-400 rounded-md overflow-hidden"
  >
    <div
      class="hover:bg-pink-50 dark:bg-gray-700 bg-white dark:hover:bg-gray-800 shadow-sm w-full h-full"
    >
      <div class="relative">
        <NuxtLink :to="`/shop/goods/${goods.goodsId}`">
          <img
            v-lazy="goods.cover"
            class="overflow-hidden animate-lazyloaded object-cover h-40 w-full cursor-pointer object-cover"
            alt=""
            srcset=""
          >
        </NuxtLink>
        <div class="absolute top-0 right-0 text-xs flex justify-center items-center text-white bg-black bg-opacity-30 px-1 py-0.5">
          <i class="inline-block i-ant-design:eye-filled" />
          <span>{{ goods.view }}</span>
        </div>
        <NuxtLink
          v-if="1"
          :to="`/category/${1}`"
          class="absolute bottom-0 right-0 text-xs text-white bg-pink-400 px-1 py-0.5 text-center cursor-pointer"
        >
          {{ goods.CategoryRelation.name }}
        </NuxtLink>
      </div>
      <div class="flex flex-row justify-between">
        <NuxtLink
          :to="`/shop/goods/${goods.goodsId}`"
          class="text-ellipsis flex-1 text-xs p-2 line-clamp-2 h-10 cursor-pointer hover:text-pink-400 dark:hover:text-pink-400 dark:text-gray-200 duration-300"
        >
          {{ goods.title }}
        </NuxtLink>
        <NuxtLink
          :to="`/shop/goods/${goods.goodsId}`"
          class="text-ellipsis w-min-20 text-md p-2 line-clamp-2 h-10 cursor-pointer text-red-500 duration-300"
        >
          ¥{{ goods.sizeValRelation[0].saleAmount }}
        </NuxtLink>
      </div>
      <div class="flex flex-row justify-between pl-2 pr-2 pt-1 pb-2 text-xs">
        <NuxtLink
          :to="`/author/${1}/userInfo`"
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
        >
          库存:{{ goods.sizeValRelation[0].stock }}
        </NuxtLink>
        <div
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
        >
          售出:1
        </div>
        <div
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
        >
          人气:{{ goods.view }}
        </div>
      </div>
    </div>
  </article>
</template>
