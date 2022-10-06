<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 11:39:48
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:37
 * @FilePath: /pinkmoe_index/components/Morebtn/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Morebtn">
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
</script>

<template>
  <div
    class="w-full select-none shadow-sm p-1.5 dark:hover:bg-gray-800 bg-white border-2 border-transparent rounded-md cursor-pointer text-center text-xs hover:bg-pink-50 dark:bg-gray-700 dark:text-gray-200 hover:border-pink-400 duration-300"
  >
    {{ $t('more') }}
  </div>
</template>
