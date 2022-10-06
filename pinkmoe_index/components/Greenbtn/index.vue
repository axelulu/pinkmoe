<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 18:24:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 11:01:13
 * @FilePath: /pinkmoe_vitesse/src/components/Greenbtn/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Greenbtn">
defineProps({
  classes: {
    type: String,
    default: null,
  },
  value: {
    type: String,
    default: '登陆',
  },
  icon: {
    type: String,
    default: '',
  },
  class: {
    type: String,
    default: '',
  },
})
</script>

<template>
  <div
    :class="classes"
    class="w-full flex justify-center items-center select-none py-1.5 cursor-pointer hover:bg-green-600 duration-300 active:bg-green-500 bg-green-500 border-0 border-green-600 text-xs text-white text-center"
  >
    <i :class="`inline-block ${icon}`" />
    <span :class="value ? 'ml-1' : ''">{{ value }}</span>
  </div>
</template>
