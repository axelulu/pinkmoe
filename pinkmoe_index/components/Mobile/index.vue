<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 09:48:39
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-26 12:02:05
 * @FilePath: /pinkmoe/pinkmoe_index/components/Mobile/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Mobile">
import { useAppStore } from '/@/store'
import src from '/@/assets/images/mobile.png'

const res = useAppStore()
</script>

<template>
  <!-- Mobile -->
  <div class="h-full">
    <div class="flex items-center justify-between border border-gray-200">
      <div class="flex justify-center w-full ml-12">
        <img class="h-14 pt-2" :src="res.siteBasic?.logo[0]">
      </div>
      <i class="inline-block text-xl mr-6 pt-2 text-gray-500 i-codicon:three-bars" />
    </div>
    <div>
      <img class="relative flex" :src="src">
      <div class="flex flex-col justify-center w-full absolute top-28">
        <div class="w-full flex justify-center">
          <img class="h-24 w-40 object-cover" :src="res.siteBasic?.logo[0]">
        </div>
        <div class="w-full flex justify-center">
          <button class="rounded-xl bg-pink-400 px-8 py-3 text-white mt-8">
            下载APP(尽情期待)
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
