<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:59:17
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 19:58:01
 * @FilePath: /pinkmoe_index/components/Spin/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Spin">
defineProps({
  show: {
    type: Boolean,
    default: false,
  },
})
</script>

<template>
  <div class="relative flex flex-wrap w-full">
    <div
      v-if="show"
      :class="show ? 'animate-spinlazyloaded' : ''"
      class="absolute top-0 left-0 w-full h-full justify-center items-center flex z-1 duration-300"
    >
      <i class="animate-rotating inline-block i-fa6-solid:circle-notch" />
    </div>
    <slot />
  </div>
</template>
