<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 09:47:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 16:40:56
 * @FilePath: /pinkmoe_index/pages/index/components/Popular/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup>
import { useUtil } from '/@/hooks/util'

defineProps({
  popular: {
    type: Object,
    default: null,
  },
  loading: {
    type: Boolean,
    default: false,
  },
})
const { getThumbnail } = useUtil()
</script>

<template>
  <div
    class="border-y-4 border-pink-400 w-full animate-fadeIn30 min-h-62"
    style="max-height: 480px"
  >
    <Spin :show="loading" class="flex flex-wrap" style="min-height:248px;">
      <div
        class="popular flex flex-wrap overflow-hidden w-full duration-300"
        style="max-height: 472px"
      >
        <NuxtLink
          v-for="(item, index) in popular"
          :key="index"
          :to="`/post/${item.postId}`"
          class="popular-item after:duration-500 flex-grow flex-initial relative cursor-pointer after:opacity-30"
        >
          <img
            v-lazy="getThumbnail(item.cover)"
            class="h-60 max-w-full min-w-full object-cover animate-lazyloaded"
            alt=""
          >
          <div
            class="absolute top-1/2 left-0 bottom-0 ring-0 w-full"
            style="background: linear-gradient(rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.65))"
          />
          <div class="absolute mb-1 text-sm bottom-1 px-8 w-full text-center text-white text-md">
            {{ item.title }}
          </div>
        </NuxtLink>
      </div>
    </Spin>
  </div>
</template>

<style scoped>
  .popular:hover .popular-item:after {
    opacity: 0.5;
  }

  .popular-item:after {
    content: '';
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    background: #000;
    opacity: 0;
  }
  .popular-item:hover:after {
    opacity: 0 !important;
  }
</style>
