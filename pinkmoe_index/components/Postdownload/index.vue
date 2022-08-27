<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 18:22:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:44
 * @FilePath: /pinkmoe_index/components/Postdownload/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="PostDownload">
import GreenBtn from '/@/components/Greenbtn/index.vue'
import { usePostDownload } from '/@/hooks/postDownload'

const props = defineProps({
  postId: {
    type: String,
    default: '',
  },
})
const { postDownload, auth, showLogin, copyText, buyDownload } = usePostDownload(props)
</script>

<template>
  <div>
    <div class="text-xs text-gray-500 dark:text-gray-200 mr-4 mt-4">
      <i class="text-gray-700 dark:text-gray-200 mr-1 inline-block i-material-symbols:cloud-download" />
      <span class="mr-1">下载链接</span>
    </div>
    <div
      class="bg-white mt-2 dark:bg-gray-700 dark:text-gray-200 rounded-md overflow-hidden animate-fadeIn30 p-4"
    >
      <div v-if="!auth.isLogin">
        <GreenBtn
          classes="w-full"
          value="登陆"
          icon="i-ic:sharp-arrow-circle-right"
          @click="showLogin"
        />
      </div>
      <div
        v-for="(item, index) in postDownload"
        v-else
        :key="index"
        class="bg-gray-100 dark:bg-gray-800 pb-2 pt-6 px-4 mt-6 relative"
      >
        <div class="text-xs bg-green-500 text-white p-1 absolute -top-2">
          {{ item.name }}
        </div>
        <div v-if="item.buy" class="flex flex-col">
          <div class="flex flex-row justify-center items-center mb-4">
            <div class="text-xs text-gray-500 dark:text-gray-200 w-2/12 flex justify-start items-center">
              <i class="inline-block i-ic:baseline-vpn-key" />
              <span class="ml-1">提取密码</span>
            </div>
            <div
              class="text-xs py-1 w-9/12 h-7 border-2 dark:bg-gray-700 dark:hover:border-pink-400 dark:border-gray-800 border-gray-200 text-center bg-white cursor-text duration-300 hover:border-pink-400"
            >
              {{ item.extractPwd ? item.extractPwd : '无' }}
            </div>
            <div
              class="w-1/12"
            >
              <GreenBtn
                value=""
                icon="i-fluent:copy-24-filled"
                @click="copyText(item.extractPwd)"
              />
            </div>
          </div>
          <div class="flex flex-row justify-center items-center mb-4">
            <div class="text-xs text-gray-500 dark:text-gray-200 w-2/12 flex justify-start items-center">
              <i class="inline-block i-nimbus:box-unpacked" />
              <span class="ml-1">解压密码</span>
            </div>
            <div
              class="text-xs py-1 h-7 w-9/12 border-2 dark:bg-gray-700 dark:hover:border-pink-400 dark:border-gray-800 border-gray-200 text-center bg-white cursor-text duration-300 hover:border-pink-400"
            >
              {{ item.unpackPwd ? item.unpackPwd : '无' }}
            </div>
            <div
              class="w-1/12"
            >
              <GreenBtn
                value=""
                icon="i-fluent:copy-24-filled"
                @click="copyText(item.unpackPwd)"
              />
            </div>
          </div>
          <div class="flex flex-row">
            <a class="w-11/12" :href="item.url" target="_blank">
              <GreenBtn value="下载" icon="i-material-symbols:cloud-download" />
            </a>
            <div
              class="w-1/12"
            >
              <GreenBtn
                classes="h-7"
                value=""
                icon="i-fluent:copy-24-filled"
                @click="copyText(item.url)"
              />
            </div>
          </div>
        </div>
        <div v-else-if="item.buy === 0">
          <GreenBtn
            classes="w-full"
            :value="`使用${item.price}积分购买`"
            icon="i-fa6-solid:sack-dollar"
            @click="buyDownload(item)"
          />
        </div>
      </div>
    </div>
  </div>
</template>
