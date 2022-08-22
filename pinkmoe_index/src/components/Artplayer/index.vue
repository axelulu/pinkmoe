<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-21 20:19:05
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 21:25:35
 * @FilePath: /pinkmoe_index/src/components/Artplayer/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
  instance.value = new Artplayer({
    ...props.option,
    container: artRef.value,
  })

  nextTick(() => {
    emit('get-instance', instance.value)
  })
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
