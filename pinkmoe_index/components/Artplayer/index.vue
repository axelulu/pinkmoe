<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-21 20:19:05
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 15:33:30
 * @FilePath: /pinkmoe_index/components/Artplayer/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup>
import Artplayer from 'artplayer'

const props = defineProps({
  option: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['get-instance'])

const instance = ref<any>()

const artRef = ref()

onMounted(() => {
  setTimeout(() => {
    instance.value = new Artplayer({
      ...props.option,
      container: artRef.value,
    })
    nextTick(() => {
      emit('get-instance', instance.value)
    })
  }, 1000)
})

onBeforeUnmount(() => {
  if (instance.value && instance.value?.destroy)
    instance.value?.destroy(false)
})
</script>

<template>
  <div ref="artRef" />
</template>

<style>
.art-video-player {
    z-index: 1;
}
</style>
