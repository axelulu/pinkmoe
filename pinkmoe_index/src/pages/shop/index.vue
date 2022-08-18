<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:32:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 09:47:10
 * @FilePath: /pinkmoe_index/src/pages/shop/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="ShopIndex">
  // Import Swiper Vue.js components
  import { Swiper, SwiperSlide } from 'swiper/vue';
  // import required modules
  import { FreeMode, Pagination } from 'swiper';

  // Import Swiper styles
  import 'swiper/css';

  import 'swiper/css/free-mode';
  import 'swiper/css/pagination';

  import { useShop } from '/@/hooks/shop/shop';

  const { content, popular, shopCategory, scrollMenu, shopMenu, loading } = useShop();
</script>

<template>
  <!-- shop -->
  <div class="flex flex-col items-center min-h-86">
    <swiper
      :slidesPerView="5"
      :spaceBetween="0"
      :freeMode="true"
      :pagination="{
        clickable: true,
      }"
      :modules="[FreeMode, Pagination]"
      class="mySwiper"
    >
      <swiper-slide v-for="(k, v) in popular" :key="v" class="cursor-pointer">
        <router-link :to="'/shop/category/' + k.category">
          <img :src="k.cover" alt="" />
        </router-link>
      </swiper-slide>
    </swiper>
    <div class="text-gray-500 w-full">
      <div class="lg:w-3/4 xl:w-5/12 m-auto mt-4 flex flex-nowrap justify-center items-center">
        <div
          @click="scrollMenu(false)"
          class="hover:text-pink-400 dark:hover:text-pink-400 duration-300 mr-3 cursor-pointer text-xl"
        >
          <font-awesome-icon :icon="['fas', 'angle-left']" />
        </div>
        <div
          v-drag-scroll.options="{ speed: 80, direction: 'x' }"
          ref="shopMenu"
          class="scrollCat flex flex-nowrap overflow-x-auto justify-start items-center"
        >
          <router-link
            :to="'/shop/category/' + k.slug"
            class="m-2 cursor-pointer"
            style="background-image: url('/uploads/file/default/background.jpeg')"
            v-for="(k, v) in shopCategory"
            :key="v"
          >
            <div
              class="w-30 h-15 select-none flex-shrink-0 flex justify-center items-center text-white text-shadow-bg-white"
              style="backdrop-filter: blur(6px)"
            >
              {{ k.name }}
            </div>
          </router-link>
        </div>
        <div
          @click="scrollMenu(true)"
          class="hover:text-pink-400 dark:hover:text-pink-400 duration-300 ml-3 cursor-pointer text-xl"
        >
          <font-awesome-icon :icon="['fas', 'angle-right']"
        /></div>
      </div>
      <div class="mt-4" v-for="(k, v) in content" :key="v">
        <div v-if="v !== 2" class="lg:w-3/4 xl:w-5/12 m-auto">
          <div class="ml-2 flex flex-row">
            <div
              class="w-7 h-7 bg-sky-400 rounded-full flex justify-center items-center text-white mr-1"
            >
              <font-awesome-icon :icon="[k.iconType, k.icon]" />
            </div>
            <div class="text-lg">{{ k.name }}</div>
            <div class="flex-1"></div>
            <router-link
              :to="'/shop/category/' + k.slug"
              class="text-xs text-gray-500 mr-2 flex items-center hover:text-pink-400 dark:text-gray-200 cursor-pointer duration-300"
            >
              更多
              <font-awesome-icon class="ml-1" :icon="['fas', 'caret-right']" />
            </router-link>
          </div>
          <Spin :show="loading">
            <div
              v-if="k.goods?.length"
              class="w-full flex justify-start flex-wrap mt-4 animate-fadeIn30"
            >
              <div v-for="(post, v) in k.goods" :key="v" class="w-1/5 p-1.5">
                <Goods :goods="post" imgHeight="h-60" />
              </div>
            </div>
            <NotFound v-else />
          </Spin>
        </div>
        <div style="background-image: url('/uploads/file/default/background.jpeg')" v-else>
          <div style="backdrop-filter: blur(12px)"
            ><div class="lg:w-3/4 xl:w-5/12 m-auto flex flex-wrap py-2">
              <div v-for="(k, v) in 10" :key="v" class="w-1/5 h-28 p-1">
                <div class="flex flex-row bg-yellow-500 w-full h-full rounded-md text-white">
                  <div class="w-7/12 p-2">
                    <div class="text-xs text-gray-100">限制商品</div>
                    <div class="text-xs text-white">
                      <span class="text-xl text-red-500 font-600">¥30</span>
                      优惠券
                    </div>
                    <div class="text-xs text-gray-100 mt-1">使用时效</div>
                    <div class="text-xs text-white">一天内使用有效</div>
                  </div>
                  <div class="w-5/12 p-2 text-center border-l-2 border-gray-200 border-dashed">
                    <div class="text-xs text-gray-100">领取时间</div>
                    <div class="text-xs text-white">无法领取</div>
                    <div
                      class="my-1 py-1 cursor-pointer flex justify-center items-center rounded-md bg-white text-gray-700 text-xs"
                      >立刻获取</div
                    >
                    <div class="text-xs text-gray-100">查看详情</div>
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

<style lang="less" scoped>
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
