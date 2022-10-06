<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 10:19:02
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:40
 * @FilePath: /pinkmoe_index/components/NotFound/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Notfound">
import src from '/@/assets/images/404.png'
defineProps({
  classes: {
    default: '',
  },
})
</script>

<template>
  <div class="mt-4 w-full flex justify-center items-start">
    <div class="relative w-125 h-48">
      <img :class="classes" :src="src" alt=""><div
        class="absolute text-sm dark:text-white select-none left-60 top-24"
      >
        这里空空如也哦 (Ｔ▽Ｔ)
      </div>
    </div>
  </div>
</template>
