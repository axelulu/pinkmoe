<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-11 15:03:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 13:34:08
 * @FilePath: /pinkmoe_index/src/pages/shop/goods/:id.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="PostIdHtml">
  import { useGoodsItem } from '/@/hooks/shop/goods';
  import { useHead } from '@vueuse/head';
  import { useAppStore } from '/@/store/modules/app';
  import { useUtil } from '/@/hooks/util';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import BasicInput from '/@/components/Basicinput/index.vue';
  const { formatDate } = useUtil();

  const {
    goodsItem,
    hasMore,
    loading,
    nextPage,
    commentList,
    showComment,
    share,
    refreshComment,
    user,
    stock,
    saleAmount,
    route,
    comment,
    num,
    cover,
    changeSizeVal,
  } = useGoodsItem();

  const { siteBasic } = useAppStore();
  useHead({
    title: `${route.params.id} - 商品页面`,
    meta: [
      { name: 'og:type', content: 'goods' },
      {
        name: 'og:title',
        content: `${route.params.id} - ${siteBasic?.title}`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
</script>

<template>
  <!-- Post -->
  <div class="flex flex-col items-center">
    <div class="flex justify-start flex-wrap lg:w-3/4 xl:w-5/12">
      <div
        class="flex flex-row bg-white w-full dark:bg-gray-700 dark:text-gray-200 mt-3 rounded-md shadow-sm animate-fadeIn30"
      >
        <div class="w-4/12 p-4">
          <img
            class="w-80 h-80 object-cover animate-lazyloaded rounded-sm overflow-hidden"
            v-lazy="cover"
            alt=""
          />
          <div class="flex flex-row">
            <img
              v-for="(k, v) in goodsItem?.goods?.goodsImg"
              :key="v"
              @click="cover = k.url"
              :class="cover === k.url ? 'border-pink-400' : ''"
              class="w-12 h-12 cursor-pointer border-2 border-transparent m-1 object-cover animate-lazyloaded rounded-sm overflow-hidden"
              v-lazy="k.url"
              alt=""
            />
          </div>
        </div>
        <div class="w-8/12 p-4">
          <div class="text-xl text-gray-800 dark:text-gray-200 py-2">{{
            goodsItem?.goods?.title
          }}</div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">售价</div>
            <div class="flex flex-row w-full justify-between items-center">
              <div class="text-red-500 text-xl min-w-20">¥{{ saleAmount }}</div>
              <div class="text-xs mr-4">
                <font-awesome-icon class="mr-1" :icon="['fas', 'eye']" />
                <span>{{ goodsItem?.goods?.view }}</span>
              </div>
            </div>
          </div>
          <div
            v-for="(size, index) in goodsItem?.goods?.sizeRelation"
            :key="index"
            class="flex flex-row my-2 justify-start items-center"
          >
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">{{ size.key }}</div>
            <div
              v-for="(sizeItem, i) in size.value"
              :key="i"
              @click="changeSizeVal(i, size.value)"
              :class="
                sizeItem.status
                  ? 'bg-pink-500 text-white dark:border-pink-600 border-pink-600 dark:text-gray-200 border-gray-100 hover:bg-pink-600 duration-300 active:bg-pink-500 hover:text-white'
                  : 'hover:bg-pink-600 duration-300 active:bg-pink-500 hover:text-white dark:hover:border-pink-600 hover:border-pink-600'
              "
              class="text-xs px-2 mr-2 h-8 flex justify-center items-center border-y-2 border-x-2 py-1.5 cursor-pointer select-none"
              >{{ sizeItem.value }}</div
            >
          </div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">库存</div>
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">{{ stock }}</div>
          </div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">已售</div>
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">{{
              goodsItem?.goods?.sizeValRelation?.[0].saleAmount
            }}</div>
          </div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">购买数量</div>
            <div
              class="text-sm dark:text-gray-200 text-gray-600 min-w-16 flex flex-row justify-center items-center"
            >
              <GreenBtn
                @click="num > 0 ? num-- : null"
                classes="w-10 h-8"
                value=""
                :icon="['fas', 'minus']"
              />
              <BasicInput
                :is-icon="false"
                class="w-20 text-xs"
                :icon="['fas', 'video']"
                v-model:value="num"
                type="number"
                placeholder="请选择数量"
              />
              <GreenBtn @click="num++" classes="w-10 h-8" value="" :icon="['fas', 'add']" />
            </div>
          </div>
          <div class="flex flex-row justify-between mt-6">
            <div class="flex flex-row">
              <GreenBtn
                @click="publishPost"
                classes="w-20 mr-2"
                value="购买"
                :icon="['fas', 'check']"
              />
              <GreenBtn
                @click="publishPost"
                classes="w-20"
                value="加购物车"
                :icon="['fas', 'cart-plus']"
              />
            </div>
            <div
              class="mr-4 w-16 h-6 bg-gray-500 text-white flex justify-center items-center text-xs hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
            >
              <font-awesome-icon class="mr-1" :icon="['fas', 'flag']" />
              <span>报告</span>
            </div>
          </div>
        </div>
      </div>
      <div class="w-full flex flex-row">
        <div class="w-9/12 pr-1 pt-2">
          <div
            class="bg-white w-full p-4 dark:bg-gray-700 dark:text-gray-200 rounded-md shadow-sm animate-fadeIn30"
          >
            <div class="pb-4 text-sm">产品简介</div>
            <div v-html="goodsItem?.goods?.content"></div>
          </div>
        </div>
        <div class="w-3/12 pl-1 pt-2">
          <div
            class="bg-white w-full dark:bg-gray-700 dark:text-gray-200 rounded-md shadow-sm animate-fadeIn30"
            >12</div
          >
        </div>
      </div>
    </div>
  </div>
</template>
