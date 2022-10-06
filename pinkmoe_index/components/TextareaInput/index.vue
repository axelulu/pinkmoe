<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 21:23:36
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:25:13
 * @FilePath: /pinkmoe_index/components/TextareaInput/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
