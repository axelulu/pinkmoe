<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 20:37:02
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:10:04
 * @FilePath: /pinkmoe_index/src/components/Dialog/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
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
  });
  const isShow = ref<boolean>(false);
  const showAnimate = ref<boolean>(false);
  const emit = defineEmits(['hide']);

  function show() {
    isShow.value = true;
    showAnimate.value = true;
    props.showFun ? props.showFun() : '';
  }

  function hide() {
    showAnimate.value = false;
    props.hideFun ? props.hideFun() : '';
    setTimeout(() => {
      isShow.value = false;
      emit('hide');
    }, 300);
  }

  defineExpose({
    hide,
    show,
  });
</script>

<template>
  <!-- Dialog -->
  <div
    v-show="isShow"
    :class="showAnimate ? 'animate-fadeIn30' : 'animate-fadeOut30'"
    style="backdrop-filter: blur(12px)"
    class="z-20 fixed top-0 bottom-0 left-0 ring-0 w-full h-full bg-black bg-opacity-50"
  >
    <div @click="hide" class="fixed top-0 bottom-0 left-0 ring-0 w-full h-full"></div>
    <slot></slot>
  </div>
</template>

<style lang="less" scoped></style>
