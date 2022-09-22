<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:32:24
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 15:37:11
 * @FilePath: /pinkmoe_index/pages/bbs/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="BbsIndex">
import Bbsarticle from '../../components/Bbsarticle/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import Spin from '/@/components/Spin/index.vue'
import { useBbs } from '/@/hooks/bbs/bbs'
import SlideUser from '/@/components/Slideuser/index.vue'
import SlideComment from '/@/components/Slidecomment/index.vue'
import TextareaInput from '/@/components/TextareaInput/index.vue'
import BasicInput from '/@/components/Basicinput/index.vue'
import GreenBtn from '/@/components/Greenbtn/index.vue'
import { useHead } from '@vueuse/head'
import { useAppStore } from '/@/store/modules/app'
const {
  postList,
  leftBar,
  currentSlug,
  sortPost,
  sort,
  t,
  currentSort,
  formParams,
  nextPage,
  hasMore,
  postFormParams,
  loading,
  isTitle,
  auth,
  bbsSilder,
  showLogin,
  publishPost,
  formPublishRef,
} = useBbs()

const { siteBasic } = useAppStore()
useHead({
  title: `论坛 - ${siteBasic?.title}`,
  meta: [
    { name: 'og:type', content: 'bbs' },
    {
      name: 'og:title',
      content: `论坛 - ${siteBasic?.title}`,
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})

definePageMeta({
  layout: 'home',
})
</script>

<template>
  <div class="flex flex-col items-center">
    <div class="lg:w-3/4 xl:w-5/12 text-gray-500 mt-4 flex flex-row">
      <div class="w-3/24">
        <ul
          class="rounded-md text-sm shadow-sm text-center text-gray-500 dark:bg-gray-700 dark:text-gray-200 bg-white py-4"
        >
          <li
            v-for="(item, v) in leftBar"
            :key="v"
            :class="currentSlug === item.slug ? 'text-white bg-pink-400' : ''"
            class="py-2 px-4 cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
            @click="sortPost(item.slug)"
          >
            {{ $t(item.name) }}
          </li>
        </ul>
      </div>
      <form ref="formPublishRef" onsubmit="return false" class="w-16/24">
        <div class="flex flex-col relative justify-between px-4">
          <div
            v-if="!auth.isLogin"
            class="absolute top-0 flex justify-center items-center w-full h-full z-10"
          >
            <div class="w-30">
              <GreenBtn
                class="w-full h-8 mb-2"
                value="登陆"
                icon="i-ic:sharp-arrow-circle-right"
                @click="showLogin"
              />
            </div>
          </div>
          <div
            :style="auth.isLogin ? '' : 'filter: blur(6px)'"
            class="mb-4 pt-2 px-4 bg-white dark:bg-gray-700 rounded-md shadow-sm animate-fadeIn30"
          >
            <BasicInput
              v-if="isTitle"
              v-model:value="postFormParams.title"
              class="w-full text-xs"
              icon="i-ic:baseline-insert-drive-file"
              :required="true"
              pattern="[\u4e00-\u9fa5]{1,50}$|^[\dA-Za-z_]{1,50}$"
              type="text"
              :placeholder="$t('bbs.pleaseEnterTitle')"
            />
            <TextareaInput
              v-model:value="postFormParams.content"
              :required="true"
              pattern="[\u4e00-\u9fa5]{1,50}$|^[\dA-Za-z_]{1,50}$"
              :placeholder="t('bbs.pleaseEnterContent')"
              class="w-full text-xs h-18"
            />
            <div class="mb-3 mt-1 flex justify-start items-center">
              <span
                :class="
                  isTitle
                    ? 'bg-pink-400 border-pink-500 active:bg-pink-500 active:border-pink-500 hover:bg-pink-600 hover:border-pink-600 border-pink-500'
                    : 'bg-gray-500 border-gray-500 active:bg-gray-500 active:border-gray-500 hover:bg-gray-600 hover:border-gray-600 border-gray-500'
                "
                class="mr-2 select-none flex justify-center items-center duration-300 py-1 px-2 cursor-pointer text-xs text-white border-2"
                @click="isTitle = !isTitle"
              >
                <i class="mr-1 inline-block i-ic:baseline-insert-drive-file" />
                <span>{{ $t('bbs.openTitle') }}</span>
              </span>
              <span
                class="bg-gray-500 flex justify-center items-center mr-2 select-none active:bg-gray-500 active:border-gray-500 hover:bg-gray-600 hover:border-gray-600 duration-300 py-1 px-2 cursor-pointer text-xs text-white border-2 border-gray-500"
              >
                <i class="mr-1 inline-block i-uiw:tags" />
                <span>{{ $t('bbs.addTopic') }}</span>
              </span>
            </div>
            <GreenBtn
              classes="w-full mb-2"
              :value="$t('bbs.releaseDynamic')"
              icon="i-material-symbols:check-circle"
              @click="publishPost"
            />
          </div>
        </div>
        <div
          class="flex flex-row mx-4 mb-4 justify-between bg-white dark:bg-gray-700 rounded-md shadow-sm animate-fadeIn30 px-4"
        >
          <div class="flex flex-row">
            <div
              v-for="(item, index) in sort"
              :key="index"
              :class="currentSort === item.type ? 'text-pink-400 dark:bg-gray-800 bg-pink-50' : 'text-gray-500 dark:text-gray-200'"
              class="text-xs px-3 select-none py-2 flex justify-center items-center cursor-pointer dark:hover:bg-gray-800 hover:text-pink-400 hover:bg-pink-50 duration-300"
              @click="sortPost(item.type)"
            >
              {{ $t(item.title) }}
            </div>
          </div>
          <div
            class="text-xs px-3 py-2 select-none flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
            @click="sortPost('desc')"
          >
            {{ formParams.desc ? $t('reverseOrder') : $t('positiveSequence') }}
          </div>
        </div>
        <Spin :show="loading">
          <div v-if="postList?.list?.length" class="mx-4 w-full">
            <Bbsarticle v-for="(post, v) in postList.list" :key="v" :post="post" />
            <div class="w-full text-gray-500">
              <MoreBtn v-if="hasMore" @click="nextPage" />
            </div>
          </div>
          <NotFound v-else />
        </Spin>
      </form>
      <div class="w-5/24">
        <SlideUser :users="bbsSilder?.users" />
        <SlideComment :comments="bbsSilder?.comments" />
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
