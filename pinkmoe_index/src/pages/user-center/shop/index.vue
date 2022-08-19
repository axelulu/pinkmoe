<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:12:12
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 08:11:37
 * @FilePath: /pinkmoe_index/src/pages/user-center/shop/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="UserCenterShopIndex">
  import UserCenterLayout from '/@/components/Usercenterlayout/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import { useAppStore } from '/@/store/modules/app';
  import { useHead } from '@vueuse/head';
  import { useUserCenterShop } from '/@/hooks/user-center/shop';

  const { siteBasic } = useAppStore();
  const { auth, currentIndex, userShop, loading, buyShop } = useUserCenterShop();
  useHead({
    title: `商城 - 用户中心`,
    meta: [
      { name: 'og:type', content: 'shop' },
      {
        name: 'og:title',
        content: `商城 - 用户中心`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
</script>

<template>
  <!-- UserCenterShopIndex -->
  <UserCenterLayout>
    <div class="ml-6">
      <div
        class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'paint-brush']" />
          <span class="ml-1 select-none">商城</span>
        </div>
        <div class="p-4">
          <div
            class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
          >
            <p>这里会展示我的商城！</p>
          </div>
        </div>
        <div class="flex flex-col items-center px-4 pb-4">
          <div
            class="px-4 py-3 w-full bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
          >
            <p class="text-xs py-1"
              ><span class="text-xs">我的积分</span
              ><span class="ml-4 bg-green-400 rounded-xl text-white px-1 py-0.5"
                ><font-awesome-icon :icon="['fas', 'gem']" /> 积分{{ auth.userInfo?.credit }}</span
              ></p
            >
            <p class="text-xs py-1"
              ><span class="text-xs">我的现金</span
              ><span class="ml-4 bg-red-400 rounded-xl text-white px-1 py-0.5"
                ><font-awesome-icon :icon="['fas', 'gem']" /> 现金{{ auth.userInfo?.cash }}</span
              ></p
            >
            <p class="text-xs py-1">
              <span class="text-xs">我的角色</span>
              <span class="ml-4 bg-pink-400 rounded-xl text-white px-1 py-0.5">
                <font-awesome-icon :icon="['fas', 'gem']" />
                {{ auth.userInfo?.authority?.authorityName }}</span
              >
            </p>
          </div>
          <Spin :show="loading" class="flex flex-wrap">
            <div class="w-full py-4 flex flex-wrap">
              <div v-for="(shop, v) in userShop" :key="v" class="w-1/2 p-1">
                <div
                  @click="currentIndex = v"
                  :class="
                    currentIndex === v
                      ? 'dark:bg-gray-900 bg-gray-100 border-pink-400 hover:bg-gray-200 dark:hover:bg-gray-900'
                      : ' dark:bg-gray-800 bg-gray-50 border-gray-200 dark:border-gray-900 hover:bg-gray-100 dark:hover:bg-gray-900'
                  "
                  class="flex flex-row items-center border-2 cursor-pointer duration-300 rounded-md"
                >
                  <font-awesome-icon class="text-4xl p-5" :icon="['fas', 'gift']" />
                  <div class="p-2">
                    <div class="text-base pb-1">{{ shop.shopName }}</div>
                    <div class="text-xs">{{ shop.shopDesc }}</div>
                  </div>
                </div>
              </div>
            </div>
          </Spin>
          <GreenBtn
            @click="buyShop"
            classes="w-full mt-2"
            value="立即购买"
            :icon="['fas', 'cart-plus']"
          />
        </div>
      </div>
    </div>
  </UserCenterLayout>
</template>

<style lang="less" scoped></style>

<route lang="yaml">
meta:
  auth: true
</route>
