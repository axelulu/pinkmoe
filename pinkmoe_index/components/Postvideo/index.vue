<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-27 22:00:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 14:47:39
 * @FilePath: /pinkmoe_index/components/Postvideo/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
const { postVideo, buyVideo, showLogin, getInstance, style, currentVideo, auth, options, changeVideo }
    = await usePostVideo(props)
</script>

<template>
  <div v-if="postVideo" class="flex flex-row w-full h-116 relative">
    <div :style="`background: url('${cover}');background-repeat: no-repeat;background-attachment: local;background-size: cover;filter: blur(8px);`" class="absolute top-0 left-0 w-full h-full" />
    <div class="flex flex-row lg:w-3/4 xl:w-5/12 h-full relative m-auto mt-2">
      <div
        v-if="!auth.isLogin"
        class="absolute top-0 flex justify-center items-center w-full h-full z-10"
      >
        <div
          class="w-2/12 h-8 mb-2"
        >
          <GreenBtn
            value="登陆观看"
            icon="i-material-symbols:check-circle"
            @click="showLogin"
          />
        </div>
      </div>
      <div class="flex flex-row w-full">
        <div class="w-19/24 h-112">
          <Artplayer :option="options" :style="style" @get-instance="getInstance" />
        </div>
        <div
          v-if="auth.isLogin"
          class="w-5/24 mx-4 py-2 overflow-y-scroll h-112 dark:bg-gray-800 bg-white rounded-md shadow-sm"
        >
          <div v-for="(video, v) in postVideo" :key="v" class="px-2 py-1 w-full">
            <div class="flex flex-row">
              <div
                :class="
                  currentVideo === video && options.url === video.url
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
                  icon="i-bi:cart-plus-fill"
                  @click="buyVideo(video)"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
