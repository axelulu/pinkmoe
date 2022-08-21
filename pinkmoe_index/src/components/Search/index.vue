<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 09:38:14
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-20 19:38:08
 * @FilePath: /pinkmoe_vitesse/src/components/Search/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Search">
import Dialog from '/@/components/Dialog/index.vue'
import { useSearchDia } from '/@/hooks/searchDia'
const props = defineProps({
  close: {
    type: Function,
    default: null,
  },
  router: {
    type: Object,
    default: null,
  },
})
const { jumpTo, scrollMenu, shopMenu, keyword, dialog, userSearch } = useSearchDia(props)
</script>

<template>
  <!-- Search -->
  <Dialog ref="dialog" @hide="close">
    <div class="fixed top-12 left-0 ring-0 flex justify-center w-full pointer-events-none">
      <div class="w-3/6 pointer-events-auto flex flex-col rounded-md overflow-hidden">
        <div class="text-center text-white text-xl p-2">
          搜索
        </div>
        <input
          v-model="keyword"
          type="text"
          placeholder="请输入关键词"
          class="border-b-4 p-2 placeholder:text-lg focus:border-pink-400 dark:border-gray-700 dark:focus:border-pink-400 duration-300 focus-visible:outline-0 placeholder:text-white text-white text-center text-lg border-gray-200 bg-transparent"
          @keyup.enter="jumpTo(keyword)"
        >
        <div class="flex flex-wrap my-2 justify-center">
          <span
            v-for="(item, index) in userSearch"
            :key="index"
            class="py-1.5 px-4 cursor-pointer hover:bg-pink-400 hover:text-white dark:hover:bg-pink-400 dark:hover:text-white dark:bg-gray-700 dark:text-gray-200 duration-300 bg-gray-200 text-xs m-1.5"
            @click="jumpTo(item.searchWord)"
          >{{ item.searchWord }}</span>
        </div>
        <div class="mt-4 flex flex-nowrap justify-center items-center">
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
            <div
              v-for="(k, v) in 11"
              :key="v"
              class="m-2 cursor-pointer"
              style="background-image: url('/uploads/file/default/background.jpeg')"
            >
              <div
                class="w-30 h-15 select-none flex-shrink-0 flex justify-center items-center text-white text-shadow-bg-white"
                style="backdrop-filter: blur(6px)"
              >
                分类一
              </div>
            </div>
          </div>
          <div
            class="hover:text-pink-400 dark:hover:text-pink-400 duration-300 ml-3 cursor-pointer text-xl"
            @click="scrollMenu(true)"
          >
            <i class="inline-block i-pepicons:triangle-right-filled" />
          </div>
        </div>
      </div>
    </div>
  </Dialog>
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
