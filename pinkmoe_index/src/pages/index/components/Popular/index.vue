<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 09:47:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 22:16:36
 * @FilePath: /pinkmoe_index/src/pages/index/components/Popular/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup>
  defineProps({
    popular: {
      type: Object,
      default: null,
    },
  });
</script>

<template>
  <!-- Popular -->
  <div
    class="border-y-4 border-pink-400 w-full animate-fadeIn30 min-h-62"
    style="max-height: 480px"
  >
    <div
      class="popular flex flex-wrap overflow-hidden w-full duration-300"
      style="max-height: 472px"
    >
      <router-link
        v-for="(item, index) in popular"
        :key="index"
        :to="'/post/' + item.postId"
        class="popular-item after:duration-500 flex-grow flex-initial relative cursor-pointer after:opacity-30"
      >
        <img
          class="h-60 max-w-full min-w-full object-cover animate-lazyloaded"
          v-lazy="item.cover"
          alt=""
        />
        <div
          class="absolute top-1/2 left-0 bottom-0 ring-0 w-full"
          style="background: linear-gradient(rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.65))"
        ></div>
        <div class="absolute bottom-1 px-8 w-full text-center text-white text-md"
          >{{ item.title }}
        </div>
      </router-link>
    </div>
  </div>
</template>

<style lang="less" scoped>
  .popular:hover .popular-item:after {
    opacity: 0.5;
  }

  .popular-item:after {
    content: '';
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    background: #000;
    opacity: 0;
  }
  .popular-item:hover:after {
    opacity: 0 !important;
  }
</style>
