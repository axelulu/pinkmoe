<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-04 07:40:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-20 20:55:43
 * @FilePath: /pinkmoe_vitesse/src/components/Shopconfirm/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Publishcomment">
import Dialog from '/@/components/Dialog/index.vue'
import GreenBtn from '/@/components/Greenbtn/index.vue'
import BasicInput from '/@/components/Basicinput/index.vue'
import WhiteBtn from '/@/components/Whitebtn/index.vue'
import { useShopConfirm } from '/@/hooks/shopConfirm'

const props = defineProps({
  authorShop: {
    type: Object,
    default: null,
  },
})

const { showAnimate, dialog, num, submitShop, key, currentVip } = useShopConfirm(props)
</script>

<template>
  <!-- Publishcomment -->
  <Dialog
    ref="dialog"
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
            <i class="inline-block i-material-symbols:check-circle" />
            <span class="ml-1">确定{{ authorShop.shopName }}规格</span>
          </div>
          <div
            class="text-sm px-4 py-1.5 bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-800 cursor-pointer hover:text-white hover:bg-pink-400 duration-300 rounded-bl-md"
            @click="dialog.hide()"
          >
            <i class="inline-block i-clarity:times-circle-solid" />
          </div>
        </div>
        <div class="pb-4 pt-2 px-4 text-xs">
          <div v-if="authorShop.priceType === 'credit' || authorShop.priceType === 'cash'">
            <BasicInput
              v-model:value="num"
              class="w-full text-xs"
              icon="material-symbols:calendar-month"
              :max="100"
              type="number"
              placeholder="输入开通数量"
            />
            <div class="flex flex-row justify-between items-center">
              <div>
                价格：<span class="text-red-500 text-xl">{{ num * currentVip?.shopCredit }}¥</span>
              </div>
            </div>
          </div>
          <div v-else-if="authorShop.priceType === 'key'">
            <BasicInput
              v-model:value="key"
              class="w-full text-xs"
              icon="material-symbols:key"
              :max="100"
              type="text"
              placeholder="输入充值卡密"
            />
          </div>
        </div>
        <div class="flex flex-row">
          <WhiteBtn
            classes="w-6/12"
            icon="fa6-solid:circle-xmark"
            value="关闭"
            @click="
              () => {
                dialog.hide();
              }
            "
          />
          <GreenBtn classes="w-6/12" icon="material-symbols:check-circle" value="开通" @click="submitShop" />
        </div>
      </form>
    </div>
  </Dialog>
</template>

<style lang="less" scoped></style>
