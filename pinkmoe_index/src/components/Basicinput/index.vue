<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 08:18:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:49
 * @FilePath: /pinkmoe_index/src/components/Basicinput/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Basicinput">
  defineProps({
    value: {
      type: String || Number,
      default: '',
    },
    required: {
      type: Boolean,
      default: false,
    },
    pattern: {
      type: String,
      default: '.*',
    },
    type: {
      type: String,
      default: 'text',
    },
    width: {
      type: String,
      default: 'w-full',
    },
    max: {
      type: Number,
      default: 10000000,
    },
    min: {
      type: Number,
      default: 0,
    },
    icon: {
      type: Array,
      default: null,
    },
    placeholder: {
      type: String,
      default: '请输入内容',
    },
  });
</script>

<template>
  <!-- Basicinput -->
  <div class="relative my-1.5 flex flex-row">
    <font-awesome-icon
      class="border-gray-100 text-center absolute left-0 top-0 p-2.5 text-gray-500 dark:text-gray-200 dark:bg-gray-800"
      :icon="icon"
    />
    <input
      :class="width"
      :max="max"
      :min="min"
      @input="
        $emit(
          'update:value',
          type === 'number'
            ? parseInt(($event.target as HTMLTextAreaElement).value)
            : ($event.target as HTMLTextAreaElement).value,
        )
      "
      :placeholder="placeholder"
      :pattern="pattern"
      :required="required"
      class="border-2 dark:text-gray-200 dark:bg-gray-800 border-gray-100 dark:border-gray-800 py-1.5 pl-9 px-2 focus:border-pink-400 duration-300 focus-visible:outline-0"
      :type="type"
      :value="value"
    />
    <slot></slot>
  </div>
</template>

<style lang="less" scoped></style>
