<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 09:48:39
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 19:57:15
 * @FilePath: /pinkmoe_index/components/Mobile/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Mobile">
import src from '/@/assets/images/mobile.png'
import { useAppStore } from '/@/store/modules/app'

const res = useAppStore()
</script>

<template>
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
