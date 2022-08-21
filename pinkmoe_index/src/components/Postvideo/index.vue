<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-27 22:00:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 10:10:30
 * @FilePath: /pinkmoe_vitesse/src/components/Postvideo/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Postvideo">
import GreenBtn from '/@/components/Greenbtn/index.vue'
import { usePostVideo } from '/@/hooks/postVideo'
const props = defineProps({
  postId: {
    type: String,
    default: '',
  },
  cover: {
    type: String,
    default: '',
  },
  videoRelation: {
    type: Array,
    default: null,
  },
})
const { postVideo, buyVideo, videoRef, showLogin, currentVideo, auth, options, changeVideo }
    = usePostVideo(props)
</script>

<template>
  <!-- Postvideo -->
  <div v-if="postVideo" class="flex flex-row w-full h-116 relative">
    <div
      v-if="!auth.isLogin"
      class="absolute top-0 flex justify-center items-center w-full h-full z-10"
    >
      <GreenBtn
        classes="w-2/12 h-8 mb-2"
        value="登陆观看"
        icon="material-symbols:check-circle"
        @click="showLogin"
      />
    </div>
    <div class="flex flex-row w-full" :style="auth.isLogin ? '' : 'filter: blur(6px)'">
      <div class="w-19/24 h-112">
        <NPlayer ref="videoRef" :options="options" />
      </div>
      <div
        class="w-5/24 mx-4 py-2 overflow-y-scroll h-112 dark:bg-gray-800 bg-white rounded-md shadow-sm"
      >
        <div v-for="(video, v) in postVideo" :key="v" class="px-2 py-1 w-full">
          <div class="flex flex-row">
            <div
              :class="
                currentVideo === video && options.src === video.url
                  ? 'dark:bg-pink-400 bg-pink-400 text-gray-200'
                  : 'dark:bg-gray-700 bg-gray-200'
              "
              class="cursor-pointer px-2 h-10 w-full rounded-md flex justify-between items-center"
              @click="changeVideo(video)"
            >
              <div class="ml-2 text-sm select-none">
                第{{ v + 1 }}集
              </div>
            </div>
            <div v-if="video.buy === 0" class="h-10 w-24">
              <GreenBtn
                classes="w-full h-10"
                value="购买"
                icon="bi:cart-plus-fill"
                @click="buyVideo(video)"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
