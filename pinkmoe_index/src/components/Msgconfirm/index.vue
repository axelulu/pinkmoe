<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-28 13:10:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:10:52
 * @FilePath: /pinkmoe_index/src/components/Msgconfirm/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Publishcomment">
  import Dialog from '/@/components/Dialog/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import WhiteBtn from '/@/components/Whitebtn/index.vue';
  import { useMsgConfirm } from '/@/hooks/msgConfirm';

  const props = defineProps({
    msg: {
      type: String,
      default: 'login',
    },
    close: {
      type: Function,
      default: null,
    },
    ok: {
      type: Function,
      default: null,
    },
  });
  const emit = defineEmits(['getPostComment']);

  const { showAnimate, dialog } = useMsgConfirm(props, emit);
</script>

<template>
  <!-- Publishcomment -->
  <Dialog
    ref="dialog"
    @hide="close"
    :hide-fun="
      () => {
        showAnimate = false;
      }
    "
    :show-fun="
      () => {
        showAnimate = true;
      }
    "
  >
    <div
      :class="showAnimate ? 'animate-pinkZoomInUp30' : 'animate-pinkZoomOutDown30'"
      class="fixed top-0 bottom-0 left-0 ring-0 flex justify-center items-center h-full w-full pointer-events-none"
    >
      <form
        onsubmit="return false"
        class="lg:w-1/6 xl:w-2/12 bg-gray-50 dark:bg-gray-700 pointer-events-auto flex flex-col rounded-md overflow-hidden"
      >
        <div class="flex flex-row justify-between items-center">
          <div
            class="flex flex-row justify-center items-center text-gray-500 dark:text-gray-200 text-xs ml-4 p-0 rounded-bl-md rounded-br-md overflow-hidden"
          >
            <font-awesome-icon :icon="['fas', 'circle-check']" />
            <span class="ml-1">确定选择</span>
          </div>
          <div
            class="text-sm px-4 py-1.5 bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-800 cursor-pointer hover:text-white hover:bg-pink-400 duration-300 rounded-bl-md"
            @click="dialog.hide()"
          >
            <font-awesome-icon :icon="['fas', 'times']" />
          </div>
        </div>
        <div class="pb-4 pt-2 px-4 text-xs h-14"> {{ msg }} </div>
        <div class="flex flex-row">
          <WhiteBtn
            @click="
              () => {
                dialog.hide();
                ok(false);
              }
            "
            classes="w-6/12"
            :icon="['fas', 'xmark']"
            value="关闭"
          />
          <GreenBtn
            @click="
              () => {
                dialog.hide();
                ok(true);
              }
            "
            classes="w-6/12"
            :icon="['fas', 'check']"
            value="确定"
          />
        </div>
      </form>
    </div>
  </Dialog>
</template>

<style lang="less" scoped></style>
