<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 09:47:23
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 07:56:54
 * @FilePath: /pinkmoe_index/src/pages/index/components/Recommend/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup>
  import NotFound from '/@/components/NotFound/index.vue';
  import Article from '/@/components/Article/index.vue';

  defineProps({
    recommend: {
      type: Object,
      default: null,
    },
    loading: {
      type: Boolean,
      default: false,
    },
  });
</script>

<template>
  <!-- Recommend -->
  <div class="lg:w-2/3 xl:w-5/12 mt-4 justify-center items-center animate-fadeIn30">
    <div class="flex flex-row justify-center mb-2">
      <div class="w-7 h-7 bg-sky-400 rounded-full flex justify-center items-center text-white mr-1">
        <font-awesome-icon :icon="['fas', 'star']" />
      </div>
      <div class="text-lg">{{ $t('home.recommendedBar') }}</div>
    </div>
    <Spin :show="loading" class="flex flex-wrap justify-center">
      <div class="flex justify-center flex-wrap w-full">
        <div v-if="recommend" class="w-full flex justify-center flex-wrap mt-4 animate-fadeIn30">
          <div v-for="(post, v) in recommend" :key="v" class="w-1/6 h-80 p-1.5 overflow-hidden">
            <Article :post="post" imgHeight="h-60" />
          </div>
        </div>
        <NotFound v-else />
      </div>
    </Spin>
  </div>
</template>

<style lang="less" scoped></style>
