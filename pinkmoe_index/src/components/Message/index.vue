<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:56:27
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:10:43
 * @FilePath: /pinkmoe_index/src/components/Message/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Message">
  import Dialog from '/@/components/Dialog/index.vue';
  const props = defineProps({
    msg: {
      type: String,
      default: '提示',
    },
    type: {
      type: String,
      default: 'success',
    },
    duration: {
      type: Number,
      default: 1500,
    },
    close: {
      type: Function,
      default: null,
    },
  });

  const dialog = ref();
  const showAnimate = ref<boolean>(false);

  function show() {
    dialog.value.show();
    setTimeout(() => {
      if (dialog.value) {
        dialog.value.hide();
      }
    }, props.duration);
  }

  onMounted(() => {
    show();
  });
</script>

<template>
  <!-- Message -->
  <Dialog
    ref="dialog"
    @hide="close"
    :hide-fun="() => (showAnimate = false)"
    :show-fun="() => (showAnimate = true)"
  >
    <div
      :class="showAnimate ? 'animate-pinkTipZoomInUp30' : 'animate-pinkTipZoomOutDown30'"
      class="fixed top-32 bottom-0 left-0 ring-0 flex justify-center items-start h-full w-full pointer-events-none"
    >
      <div
        v-if="type === 'success'"
        class="bg-gray-50 text-green-600 dark:bg-gray-700 pointer-events-auto flex flex-row justify-center items-center px-2 py-1 rounded-md overflow-hidden"
      >
        <font-awesome-icon :icon="['fas', 'check-circle']" />
        <span class="ml-1">{{ msg }}</span>
      </div>
      <div
        v-if="type === 'info'"
        class="bg-gray-50 text-sky-600 dark:bg-gray-700 pointer-events-auto flex flex-row justify-center items-center px-2 py-1 rounded-md overflow-hidden"
      >
        <font-awesome-icon :icon="['fas', 'exclamation-circle']" />
        <span class="ml-1">{{ msg }}</span>
      </div>
      <div
        v-if="type === 'warning'"
        class="bg-gray-50 text-red-600 dark:bg-gray-700 pointer-events-auto flex flex-row justify-center items-center px-2 py-1 rounded-md overflow-hidden"
      >
        <font-awesome-icon :icon="['fas', 'times-circle']" />
        <span class="ml-1">{{ msg }}</span>
      </div>
    </div>
  </Dialog>
</template>

<style lang="less" scoped></style>
