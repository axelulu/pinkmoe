<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 10:50:39
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:59
 * @FilePath: /pinkmoe_index/src/components/Vipconfirm/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Publishcomment">
  import Dialog from '/@/components/Dialog/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import BasicInput from '/@/components/Basicinput/index.vue';
  import WhiteBtn from '/@/components/Whitebtn/index.vue';
  import { useVipConfirm } from '/@/hooks/vipConfirm';

  const props = defineProps({
    authorAuthority: {
      type: Object,
      default: null,
    },
  });

  const { showAnimate, dialog, num, submitVip, changeVip, currentVip } = useVipConfirm(props);
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
            <span class="ml-1">确定{{ authorAuthority.authorityName }}规格</span>
          </div>
          <div
            class="text-sm px-4 py-1.5 bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-800 cursor-pointer hover:text-white hover:bg-pink-400 duration-300 rounded-bl-md"
            @click="dialog.hide()"
          >
            <font-awesome-icon :icon="['fas', 'times']" />
          </div>
        </div>
        <div class="pb-4 pt-2 px-4 text-xs">
          <div>
            <BasicInput
              class="w-full text-xs"
              :icon="['fas', 'calendar']"
              :max="100"
              v-model:value="num"
              type="number"
              placeholder="输入开通数量"
            />
            <div class="flex flex-row justify-between items-center">
              <div class="flex flex-row">
                <div v-for="(k, v) in authorAuthority.authorityParams" :key="v">
                  <div
                    v-if="k.authorityParamId.search('Price') !== -1"
                    @click="changeVip(k)"
                    :class="currentVip == k ? 'bg-green-500 text-white' : 'text-gray-500'"
                    class="text-xs px-2 flex justify-center items-center border-2 dark:border-gray-900 dark:text-gray-200 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
                  >
                    {{ k.authorityParamKey }}</div
                  >
                </div>
              </div>
              <div
                >价格：<span class="text-red-500 text-xl"
                  >{{ num * currentVip?.authorityParamValue }}¥</span
                ></div
              >
            </div>
          </div>
        </div>
        <div class="flex flex-row">
          <WhiteBtn
            @click="
              () => {
                dialog.hide();
              }
            "
            classes="w-6/12"
            :icon="['fas', 'xmark']"
            value="关闭"
          />
          <GreenBtn @click="submitVip" classes="w-6/12" :icon="['fas', 'check']" value="开通" />
        </div>
      </form>
    </div>
  </Dialog>
</template>

<style lang="less" scoped></style>
