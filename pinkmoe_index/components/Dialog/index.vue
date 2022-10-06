<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 20:37:02
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:17
 * @FilePath: /pinkmoe_index/components/Dialog/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Dialog">
const props = defineProps({
  showFun: {
    type: Function,
    default: null,
  },
  hideFun: {
    type: Function,
    default: null,
  },
})
const emit = defineEmits(['hide'])
const isShow = ref<boolean>(false)
const showAnimate = ref<boolean>(false)
function show() {
  isShow.value = true
  showAnimate.value = true
  props.showFun ? props.showFun() : ''
}

function hide() {
  showAnimate.value = false
  props.hideFun ? props.hideFun() : ''
  setTimeout(() => {
    isShow.value = false
    emit('hide')
  }, 300)
}

defineExpose({
  hide,
  show,
})
</script>

<template>
  <div
    v-show="isShow"
    :class="showAnimate ? 'animate-fadeIn30' : 'animate-fadeOut30'"
    style="backdrop-filter: blur(12px)"
    class="z-20 fixed top-0 bottom-0 left-0 ring-0 w-full h-full bg-black bg-opacity-50"
  >
    <div class="fixed top-0 bottom-0 left-0 ring-0 w-full h-full" @click="hide" />
    <slot />
  </div>
</template>
