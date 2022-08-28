<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:56:27
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 16:59:25
 * @FilePath: /pinkmoe_index/components/Message/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Message">
import Dialog from '/@/components/Dialog/index.vue'
const props = defineProps({
  successMsg: {
    type: String,
    default: '成功',
  },
  failedMsg: {
    type: String,
    default: '失败',
  },
  type: {
    type: String,
    default: '',
  },
  duration: {
    type: Number,
    default: 1500,
  },
  loadFun: {
    type: Function,
    default: null,
  },
  close: {
    type: Function,
    default: null,
  },
})

const dialog = ref()
const msg = ref()
const status = ref(0)
const data = ref()
const showAnimate = ref<boolean>(false)

function show() {
  msg.value = '加载中...'
  status.value = 0
  dialog.value.show()
  const result = props.loadFun()
  result.then(({ code, message, result }) => {
    msg.value = code === 200 ? props.successMsg : message || props.failedMsg
    status.value = code
    data.value = result
    setTimeout(() => {
      if (dialog.value)
        dialog.value.hide()
    }, props.duration)
  })
}

onMounted(() => {
  show()
})
</script>

<template>
  <Dialog
    ref="dialog"
    :hide-fun="() => (showAnimate = false)"
    :show-fun="() => (showAnimate = true)"
    @hide="close({ status, data })"
  >
    <div
      :class="showAnimate ? 'animate-pinkTipZoomInUp30' : 'animate-pinkTipZoomOutDown30'"
      class="fixed top-32 bottom-0 left-0 ring-0 flex justify-center items-start h-full w-full pointer-events-none"
    >
      <div v-if="type">
        <div
          v-if="type === 'success'"
          class="bg-gray-50 text-green-600 dark:bg-gray-700 pointer-events-auto flex flex-row justify-center items-center px-2 py-1 rounded-md overflow-hidden"
        >
          <i class="inline-block i-material-symbols:check-circle" />
          <span class="ml-1">{{ msg }}</span>
        </div>
        <div
          v-if="type === 'info'"
          class="bg-gray-50 text-sky-600 dark:bg-gray-700 pointer-events-auto flex flex-row justify-center items-center px-2 py-1 rounded-md overflow-hidden"
        >
          <i class="inline-block i-bi:exclamation-circle-fill" />
          <span class="ml-1">{{ msg }}</span>
        </div>
        <div
          v-if="type === 'warning'"
          class="bg-gray-50 flex justify-center items-center text-red-600 dark:bg-gray-700 pointer-events-auto flex flex-row justify-center items-center p-2 rounded-md overflow-hidden"
        >
          <i class="inline-block i-clarity:times-circle-solid" />
          <span class="ml-1">{{ msg }}</span>
        </div>
      </div>
      <div v-else>
        <div
          :class="status === 0 ? 'text-yellow-600' : status === 200 ? 'text-green-600' : 'text-red-600'"
          class="bg-gray-50 dark:bg-gray-700 pointer-events-auto flex flex-row justify-center items-center px-2 py-1 rounded-md overflow-hidden"
        >
          <i v-if="status === 0" class="inline-block i-line-md:loading-twotone-loop" />
          <i v-if="status === 200" class="inline-block i-material-symbols:check-circle" />
          <i v-if="status === -1" class="inline-block i-clarity:times-circle-solid" />
          <span class="ml-1">{{ msg }}</span>
        </div>
      </div>
    </div>
  </Dialog>
</template>
