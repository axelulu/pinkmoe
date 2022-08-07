<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 09:38:14
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:20
 * @FilePath: /pinkmoe_index/src/components/Search/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Search">
  import Dialog from '/@/components/Dialog/index.vue';
  import { useSearchDia } from '/@/hooks/searchDia';
  const props = defineProps({
    close: {
      type: Function,
      default: null,
    },
    router: {
      type: Object,
      default: null,
    },
  });
  const { jumpTo, keyword, dialog } = useSearchDia(props);
</script>

<template>
  <!-- Search -->
  <Dialog ref="dialog" @hide="close">
    <div class="fixed top-12 left-0 ring-0 flex justify-center w-full pointer-events-none">
      <div class="w-3/6 pointer-events-auto flex flex-col rounded-md overflow-hidden">
        <div class="text-center text-white text-xl p-2">搜索</div>
        <input
          @keyup.enter="jumpTo(keyword)"
          v-model="keyword"
          type="text"
          placeholder="请输入关键词"
          class="border-b-4 p-2 placeholder:text-lg focus:border-pink-400 dark:border-gray-700 dark:focus:border-pink-400 duration-300 focus-visible:outline-0 placeholder:text-white text-white text-center text-lg border-gray-200 bg-transparent"
        />
        <div class="flex flex-wrap my-2 justify-center">
          <span
            v-for="(item, index) in 13"
            :key="index"
            @click="jumpTo('测试')"
            class="py-1.5 px-4 cursor-pointer hover:bg-pink-400 hover:text-white dark:hover:bg-pink-400 dark:hover:text-white dark:bg-gray-700 dark:text-gray-200 duration-300 bg-gray-200 text-xs m-1.5"
            >关键词一</span
          >
        </div>
      </div>
    </div>
  </Dialog>
</template>

<style lang="less" scoped></style>
