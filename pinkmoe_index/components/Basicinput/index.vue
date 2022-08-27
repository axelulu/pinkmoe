<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 08:18:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:02
 * @FilePath: /pinkmoe_index/components/Basicinput/index.vue
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
  isIcon: {
    type: Boolean,
    default: true,
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
    type: String,
    default: '',
  },
  placeholder: {
    type: String,
    default: '请输入内容',
  },
})
</script>

<template>
  <div class="relative my-1.5 flex flex-row justify-center items-center">
    <i
      v-if="isIcon" :class="`border-gray-100 text-center absolute left-0 top-0 m-2 p-2 text-gray-500 inline-block dark:text-gray-200 ${icon}`"
    />
    <input
      :max="max"
      :min="min"
      :placeholder="placeholder"
      :pattern="pattern"
      :required="required"
      :class="isIcon ? `pl-7 ${width}` : `${width}`"
      class="border-2 dark:text-gray-200 dark:bg-gray-800 border-gray-100 dark:border-gray-800 py-1.5 px-2 focus:border-pink-400 duration-300 focus-visible:outline-0"
      :type="type"
      :value="value"
      @input="
        $emit(
          'update:value',
          type === 'number'
            ? parseInt(($event.target as HTMLTextAreaElement).value)
            : ($event.target as HTMLTextAreaElement).value,
        )
      "
    >
    <slot />
  </div>
</template>
