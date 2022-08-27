<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 11:39:48
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:37
 * @FilePath: /pinkmoe_index/components/Morebtn/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
