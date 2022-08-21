<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 10:19:02
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:45:29
 * @FilePath: /pinkmoe_vitesse/src/components/NotFound/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
  <!-- NotFound -->
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

<style lang="less" scoped></style>
