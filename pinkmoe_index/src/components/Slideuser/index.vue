<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 18:50:59
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:42
 * @FilePath: /pinkmoe_index/src/components/Slideuser/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Slideuser">
  defineProps({
    users: {
      type: Array,
      default: null,
    },
  });
</script>

<template>
  <!-- Slideuser -->
  <div class="mt-4 animate-fadeIn30">
    <div class="flex flex-row justify-between mx-1">
      <div class="text-xs text-gray-500 dark:text-gray-200">
        <font-awesome-icon class="mr-1" :icon="['fas', 'users']" />
        <span>喵爪怪们</span>
      </div>
    </div>
    <div class="flex flex-wrap">
      <div
        v-for="(item, index) in users"
        :key="index"
        class="w-1/3 p-1 text-center cursor-pointer text-gray-500 dark:text-gray-200 hover:bg-pink-50 dark:hover:bg-gray-600 hover:text-pink-400"
      >
        <router-link :to="'/author/' + item.uuid + '/userInfo'">
          <img
            class="rounded-full animate-lazyloaded object-cover h-20 w-20 border-2 border-transparent hover:border-pink-400 duration-300"
            v-lazy="item.avatar ? '/' + item.avatar : ''"
            alt=""
          />
          <div class="text-xs mt-1">{{ item.nickName }}</div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
