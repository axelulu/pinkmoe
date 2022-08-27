<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:32:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:25:54
 * @FilePath: /pinkmoe_index/pages/shop/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="ShopIndex">
// Import Swiper Vue.js components
import { Swiper, SwiperSlide } from 'swiper/vue'
// import required modules
import { FreeMode, Pagination } from 'swiper'
// Import Swiper styles
import 'swiper/css'
import 'swiper/css/free-mode'
import 'swiper/css/pagination'

import { useShop } from '/@/hooks/shop/shop'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '/@/store/modules/app'

const { content, popular, shopCategory, scrollMenu, shopMenu, loading } = useShop()
const { t } = useI18n()
const { siteBasic } = useAppStore()
useHead({
  title: `商城 - ${siteBasic?.title}`,
  meta: [
    { name: 'og:type', content: 'bbs' },
    {
      name: 'og:title',
      content: `商城 - ${siteBasic?.title}`,
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <div class="flex flex-col items-center min-h-86">
    <Swiper
      :slides-per-view="5"
      :space-between="0"
      :free-mode="true"
      :pagination="{
        clickable: true,
      }"
      :modules="[FreeMode, Pagination]"
      class="mySwiper"
    >
      <SwiperSlide v-for="(k, v) in popular" :key="v" class="cursor-pointer">
        <NuxtLink :to="`/shop/category/${k.category}`">
          <img :src="k.cover" alt="">
        </NuxtLink>
      </SwiperSlide>
    </Swiper>
    <div class="text-gray-500 w-full">
      <div class="lg:w-3/4 xl:w-5/12 m-auto mt-4 flex flex-nowrap justify-center items-center">
        <div
          class="hover:text-pink-400 dark:hover:text-pink-400 duration-300 mr-3 cursor-pointer text-xl"
          @click="scrollMenu(false)"
        >
          <i class="inline-block i-pepicons:triangle-left-filled" />
        </div>
        <div
          ref="shopMenu"
          class="scrollCat flex flex-nowrap overflow-x-auto justify-start items-center"
        >
          <NuxtLink
            v-for="(k, v) in shopCategory"
            :key="v"
            :to="`/shop/category/${k.slug}`"
            class="m-2 cursor-pointer"
            style="background-image: url('/uploads/file/default/background.jpeg')"
          >
            <div
              class="w-30 h-15 select-none flex-shrink-0 flex justify-center items-center text-white text-shadow-bg-white"
              style="backdrop-filter: blur(6px)"
            >
              {{ k.name }}
            </div>
          </NuxtLink>
        </div>
        <div
          class="hover:text-pink-400 dark:hover:text-pink-400 duration-300 ml-3 cursor-pointer text-xl"
          @click="scrollMenu(true)"
        >
          <i class="inline-block i-pepicons:triangle-right-filled" />
        </div>
      </div>
      <div v-for="(k, v) in content" :key="v" class="mt-4">
        <div v-if="v !== 2" class="lg:w-3/4 xl:w-5/12 m-auto">
          <div class="ml-2 flex flex-row">
            <div
              class="w-7 h-7 bg-sky-400 rounded-full flex justify-center items-center text-white mr-1"
            >
              <i :class="`i-${k.icon}`" />
            </div>
            <div class="text-lg">
              {{ k.name }}
            </div>
            <div class="flex-1" />
            <NuxtLink
              :to="`/shop/category/${k.slug}`"
              class="text-xs text-gray-500 mr-2 flex items-center hover:text-pink-400 dark:text-gray-200 cursor-pointer duration-300"
            >
              {{ $t('more') }}
              <i class="ml-1 i-fluent:caret-right-12-filled" />
            </NuxtLink>
          </div>
          <Spin :show="loading">
            <div
              v-if="k.goods?.length"
              class="w-full flex justify-start flex-wrap mt-4 animate-fadeIn30"
            >
              <div v-for="(post, v) in k.goods" :key="v" class="w-1/5 p-1.5">
                <Goods :goods="post" img-height="h-60" />
              </div>
            </div>
            <NotFound v-else />
          </Spin>
        </div>
        <div v-else style="background-image: url('/uploads/file/default/background.jpeg')">
          <div style="backdrop-filter: blur(12px)">
            <div class="lg:w-3/4 xl:w-5/12 m-auto flex flex-wrap py-2">
              <div v-for="(k, v) in 10" :key="v" class="w-1/5 h-28 p-1">
                <div class="flex flex-row bg-yellow-500 w-full h-full rounded-md text-white">
                  <div class="w-7/12 p-2">
                    <div class="text-xs text-gray-100">
                      限制商品
                    </div>
                    <div class="text-xs text-white">
                      <span class="text-xl text-red-500 font-600">¥30</span>
                      优惠券
                    </div>
                    <div class="text-xs text-gray-100 mt-1">
                      使用时效
                    </div>
                    <div class="text-xs text-white">
                      一天内使用有效
                    </div>
                  </div>
                  <div class="w-5/12 p-2 text-center border-l-2 border-gray-200 border-dashed">
                    <div class="text-xs text-gray-100">
                      领取时间
                    </div>
                    <div class="text-xs text-white">
                      无法领取
                    </div>
                    <div
                      class="my-1 py-1 cursor-pointer flex justify-center items-center rounded-md bg-white text-gray-700 text-xs"
                    >
                      立刻获取
                    </div>
                    <div class="text-xs text-gray-100">
                      查看详情
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  .scrollCat::-webkit-scrollbar {
    display: none; /* Chrome Safari */
    -ms-overflow-style: none;
    scrollbar-width: none;
  }
  .swiper {
    width: 100%;
    height: 240px;
  }

  .swiper-slide {
    text-align: center;
    font-size: 18px;
    background: #fff;

    /* Center slide text vertically */
    display: -webkit-box;
    display: -ms-flexbox;
    display: -webkit-flex;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    -webkit-justify-content: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    -webkit-align-items: center;
    align-items: center;
  }

  .swiper-slide img {
    display: block;
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
</style>
