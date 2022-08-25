<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 21:23:36
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:46:17
 * @FilePath: /pinkmoe_vitesse/src/components/TextareaInput/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="TextareaInput">
defineProps({
  value: {
    type: String,
    default: '',
  },
  required: {
    type: Boolean,
    default: false,
  },
  pattern: {
    type: String,
    default: '',
  },
  maxlen: {
    type: String,
    default: '120',
  },
  rows: {
    type: String,
    default: '5',
  },
  width: {
    type: String,
    default: 'w-full',
  },
  placeholder: {
    type: String,
    default: '请输入内容',
  },
})
</script>

<template>
  <!-- TextareaInput -->
  <div class="relative my-1.5 flex flex-row">
    <textarea
      :class="width"
      :placeholder="placeholder"
      :pattern="pattern"
      :required="required"
      class="border-2 border-gray-100 py-1.5 pl-2 pr-2 dark:bg-gray-800 dark:border-gray-800 focus:border-pink-400 duration-300 focus-visible:outline-0"
      :maxlength="maxlen"
      :rows="rows"
      :value="value"
      @input="$emit('update:value', ($event.target as HTMLTextAreaElement).value)"
    />
    <slot />
  </div>
</template>
