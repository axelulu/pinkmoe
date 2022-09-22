<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:12:27
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 16:50:39
 * @FilePath: /pinkmoe_index/pages/user-center/vip/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="UserCenterVipIndex">
import UserCenterLayout from '/@/components/Usercenterlayout/index.vue'
import GreenBtn from '/@/components/Greenbtn/index.vue'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'
import { useUserCenterVip } from '/@/hooks/user-center/vip'
import NotFound from '/@/components/NotFound/index.vue'

definePageMeta({
  middleware: ['user-auth'],
  layout: 'user-center',
})

const { siteBasic } = useAppStore()
const { authorAuthorityList, auth, currentIndex, loading, buyVip } = useUserCenterVip()
useHead({
  title: '会员 - 用户中心',
  meta: [
    { name: 'og:type', content: 'vip' },
    {
      name: 'og:title',
      content: '会员 - 用户中心',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <div class="ml-6">
    <div
      class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
    >
      <div class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
        <i class="inline-block i-icon-park-solid:vip-one" />
        <span class="ml-1 select-none">我的会员</span>
      </div>
      <div class="p-4">
        <div
          class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
        >
          <p>这里会展示我的会员信息！</p>
        </div>
      </div>
      <div class="flex flex-col items-center px-4 pb-4">
        <div
          class="px-4 py-3 w-full bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
        >
          <p class="text-xs py-1 flex justify-start items-center">
            <span class="text-xs">我的积分</span><span class="ml-4 bg-green-400 rounded-xl text-white px-1 py-0.5 flex justify-center items-center">
              <i class="inline-block mr-1 i-ph:paw-print-fill" /> 积分{{ auth.userInfo?.credit }}</span>
          </p>
          <p class="text-xs py-1 flex justify-start items-center">
            <span class="text-xs">我的现金</span><span class="ml-4 bg-red-400 rounded-xl text-white px-1 py-0.5 flex justify-center items-center">
              <i class="inline-block mr-1 i-lucide:gem" /> 现金{{ auth.userInfo?.cash }}</span>
          </p>
          <p class="text-xs py-1 flex justify-start items-center">
            <span class="text-xs">我的角色</span><span class="ml-4 bg-pink-400 rounded-xl text-white px-1 py-0.5 flex justify-center items-center">
              <i class="inline-block mr-1 i-bx:bxs-user-check" />
              {{ auth.userInfo?.authority?.authorityName }}</span>
          </p>
        </div>
        <Spin :show="loading" class="flex flex-wrap">
          <div class="w-full py-4 flex flex-wrap justify-center min-h-79">
            <div
              v-for="(authority, v) in authorAuthorityList?.list"
              v-if="authorAuthorityList?.list?.length > 0"
              :key="v"
              class="w-1/4 px-4 py-4"
            >
              <div
                :class="
                  currentIndex === v
                    ? 'dark:bg-gray-900 bg-gray-100 border-pink-400 hover:bg-gray-200 dark:hover:bg-gray-900'
                    : ' dark:bg-gray-800 bg-gray-50 border-gray-200 dark:border-gray-900 hover:bg-gray-100 dark:hover:bg-gray-900'
                "
                class="flex flex-row w-full duration-300 border-2 cursor-pointer rounded-md"
                @click="currentIndex = v"
              >
                <div class="text-center w-full p-2">
                  <div class="text-base pb-2 pt-2 text-blue-400">
                    {{
                      authority.authorityName
                    }}
                  </div>
                  <div v-for="(k, v) in authority.authorityParams" :key="v">
                    <div
                      v-if="k.authorityParamId.search('Price') == -1"
                      class="text-xs px-2 py-1 text-gray-800 dark:text-gray-200"
                    >
                      <i class="mr-2 inline-block text-green-500 i-material-symbols:check-circle" />
                      {{
                        k.authorityParamKey
                      }}{{ k.authorityParamValue }}%
                    </div>
                    <div v-else class="text-base py-0.5 text-gray-800 dark:text-gray-200">
                      ￥{{ k.authorityParamValue }}/
                      <span class="text-gray-400">{{ k.authorityParamKey }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <NotFound v-else />
          </div>
        </Spin>
        <GreenBtn
          classes="w-full mt-2"
          value="立即购买"
          icon="i-bi:cart-plus-fill"
          @click="buyVip"
        />
      </div>
    </div>
  </div>
</template>
