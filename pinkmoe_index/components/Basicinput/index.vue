<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 08:18:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 16:33:42
 * @FilePath: /pinkmoe_index/components/Basicinput/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
          'update:value', ($event.target as HTMLTextAreaElement).value,
        )
      "
    >
    <slot />
  </div>
</template>
