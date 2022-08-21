<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 10:35:44
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 11:20:57
 * @FilePath: /pinkmoe_vitesse/src/components/Goods/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
  <!-- Article -->
  <article
    class="border-2 duration-300 border-transparent hover:border-pink-400 rounded-md overflow-hidden"
  >
    <div
      class="hover:bg-pink-50 dark:bg-gray-700 bg-white dark:hover:bg-gray-800 shadow-sm w-full h-full"
    >
      <div class="relative">
        <router-link :to="`/shop/goods/${goods.goodsId}`">
          <img
            v-lazy="goods.cover"
            class="overflow-hidden animate-lazyloaded object-cover h-40 w-full cursor-pointer object-cover"
            alt=""
            srcset=""
          >
        </router-link>
        <div class="absolute top-0 right-0 text-xs flex justify-center items-center text-white bg-black bg-opacity-30 px-1 py-0.5">
          <i class="inline-block i-ant-design:eye-filled" />
          <span>{{ goods.view }}</span>
        </div>
        <router-link
          v-if="1"
          :to="`/category/${1}`"
          class="absolute bottom-0 right-0 text-xs text-white bg-pink-400 px-1 py-0.5 text-center cursor-pointer"
        >
          {{ goods.CategoryRelation.name }}
        </router-link>
      </div>
      <div class="flex flex-row justify-between">
        <router-link
          :to="`/shop/goods/${goods.goodsId}`"
          class="text-ellipsis flex-1 text-xs p-2 line-clamp-2 h-10 cursor-pointer hover:text-pink-400 dark:hover:text-pink-400 dark:text-gray-200 duration-300"
        >
          {{ goods.title }}
        </router-link>
        <router-link
          :to="`/shop/goods/${goods.goodsId}`"
          class="text-ellipsis w-min-20 text-md p-2 line-clamp-2 h-10 cursor-pointer text-red-500 duration-300"
        >
          ¥{{ goods.sizeValRelation[0].saleAmount }}
        </router-link>
      </div>
      <div class="flex flex-row justify-between pl-2 pr-2 pt-1 pb-2 text-xs">
        <router-link
          :to="`/author/${1}/userInfo`"
          class="cursor-pointer hover:text-pink-400 dark:text-gray-200 duration-300 text-gray-500"
        >
          库存:{{ goods.sizeValRelation[0].stock }}
        </router-link>
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

<style lang="less" scoped></style>
