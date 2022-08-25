<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 11:41:28
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 23:18:52
 * @FilePath: /pinkmoe_vitesse/src/components/Footer/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Footer">
import { useFooter } from '/@/hooks/footer'

const { scollTop, appStore } = useFooter()
</script>

<template>
  <!-- Footer -->
  <div class="relative">
    <div class="bg-white dark:bg-gray-700 dark:text-gray-200 w-full mt-10 flex justify-center">
      <div v-if="appStore.siteFooter" class="lg:w-4/5 xl:w-6/12">
        <div class="flex flex-row text-gray-500 dark:text-gray-200">
          <div class="w-1/4 pr-8">
            <div class="text-xs mt-6">
              快速导航
            </div>
            <ul class="text-xs mt-2">
              <li
                v-for="(item, index) in appStore.siteFooter.links"
                :key="index"
                class="pt-1 pb-0.5"
              >
                <a
                  :href="item.url"
                  class="text-xs text-pink-400 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer p-1"
                >{{ item.name }}</a>
              </li>
            </ul>
          </div>
          <div class="w-1/4 pr-8">
            <div class="text-xs mt-6">
              网站申明
            </div>
            <div class="text-xs mt-2">
              {{ appStore.siteFooter.notice }}
            </div>
          </div>
          <div class="w-1/4 pr-8">
            <div class="text-xs mt-6">
              联系我们
            </div>
            <div class="text-xs mt-2">
              {{ appStore.siteFooter.contact }}
            </div>
          </div>
          <div class="w-1/4 pr-8">
            <div class="text-xs mt-6">
              关于小站
            </div>
            <div class="text-xs mt-2">
              {{ appStore.siteFooter.about }}
            </div>
          </div>
        </div>
        <div class="mb-8">
          <div class="text-xs mt-2">
            <i class="inline-block i-material-symbols:link-rounded" />
            <span class="ml-1">友情链接</span>
          </div>
          <ul class="text-xs mt-2 flex flex-row">
            <li v-for="(item, index) in appStore.siteFooter.friendsLinks" :key="index" class="pr-1">
              <a
                :href="item.url"
                class="text-xs text-pink-400 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer p-1"
              >{{ item.name }}
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <div
      class="mb-8 h-14 text-gray-500 dark:text-gray-200 items-center text-xs justify-center flex underline"
    >
      © {{ appStore.siteBasic?.title }} 2022
    </div>
    <div
      style="backdrop-filter: blur(2px)"
      class="fixed bottom-0 w-full h-9 shadow-scollTop bg-opacity-75 dark:bg-gray-700 bg-white flex justify-center"
    >
      <div
        class="w-3/4 flex justify-center items-center text-xs text-gray-500 dark:text-gray-200 cursor-pointer hover:bg-pink-400 hover:text-white duration-300"
        @click="scollTop"
      >
        <i class="inline-block i-ic:outline-arrow-upward" />
        <span class="ml-1">返回顶部</span>
      </div>
    </div>
  </div>
</template>
