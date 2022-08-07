<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:32:24
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:05:00
 * @FilePath: /pinkmoe_index/src/pages/bbs/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="BbsIndex">
  import Bbsarticle from '../../components/Bbsarticle/index.vue';
  import NotFound from '/@/components/NotFound/index.vue';
  import MoreBtn from '/@/components/Morebtn/index.vue';
  import Spin from '/@/components/Spin/index.vue';
  import { useBbs } from '/@/hooks/bbs/bbs';
  import TextareaInput from '/@/components/TextareaInput/index.vue';
  import BasicInput from '/@/components/Basicinput/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import { useHead } from '@vueuse/head';
  import { useAppStore } from '/@/store/modules/app';
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
    showLogin,
    publishPost,
    formPublishRef,
  } = useBbs();

  const { siteBasic } = useAppStore();
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
  });
</script>

<template>
  <!-- Bbs -->
  <div class="flex flex-col items-center">
    <div class="lg:w-3/4 xl:w-5/12 text-gray-500 mt-4 flex flex-row">
      <div class="w-38">
        <ul
          class="rounded-md text-sm shadow-sm text-center text-gray-500 dark:bg-gray-700 dark:text-gray-200 bg-white py-4"
        >
          <li
            v-for="(item, v) in leftBar"
            :key="v"
            @click="sortPost(item.slug)"
            :class="currentSlug === item.slug ? 'text-white bg-pink-400' : ''"
            class="py-2 px-4 cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
            >{{ item.name }}
          </li>
        </ul>
      </div>
      <form ref="formPublishRef" onsubmit="return false" class="w-full">
        <div class="flex flex-col relative justify-between px-4">
          <div
            v-if="!auth.isLogin"
            class="absolute top-0 flex justify-center items-center w-full h-full z-10"
          >
            <GreenBtn
              @click="showLogin"
              classes="w-2/12 h-8 mb-2"
              :value="'登陆'"
              :icon="['fas', 'circle-right']"
          /></div>
          <div
            :style="auth.isLogin ? '' : 'filter: blur(6px)'"
            class="mb-4 pt-2 px-4 bg-white dark:bg-gray-700 rounded-md shadow-sm animate-fadeIn30"
            ><BasicInput
              v-if="isTitle"
              class="w-full text-xs"
              :icon="['fas', 'file']"
              :required="true"
              pattern="[\u4e00-\u9fa5]{1,50}$|^[\dA-Za-z_]{1,50}$"
              v-model:value="postFormParams.title"
              type="text"
              :placeholder="$t('bbs.pleaseEnterTitle')" />
            <TextareaInput
              :required="true"
              pattern="[\u4e00-\u9fa5]{1,50}$|^[\dA-Za-z_]{1,50}$"
              v-model:value="postFormParams.content"
              :placeholder="t('bbs.pleaseEnterContent')"
              class="w-full text-xs h-18" />
            <div class="mb-3 mt-1">
              <span
                @click="isTitle = !isTitle"
                :class="
                  isTitle
                    ? 'bg-pink-400 border-pink-500 active:bg-pink-500 active:border-pink-500 hover:bg-pink-600 hover:border-pink-600 border-pink-500'
                    : 'bg-gray-500 border-gray-500 active:bg-gray-500 active:border-gray-500 hover:bg-gray-600 hover:border-gray-600 border-gray-500'
                "
                class="mr-2 select-none duration-300 py-1 px-2 cursor-pointer text-xs text-white border-2"
              >
                <font-awesome-icon class="mr-1" :icon="['fas', 'file']" />
                <span>{{ $t('bbs.openTitle') }}</span>
              </span>
              <span
                class="bg-gray-500 mr-2 select-none active:bg-gray-500 active:border-gray-500 hover:bg-gray-600 hover:border-gray-600 duration-300 py-1 px-2 cursor-pointer text-xs text-white border-2 border-gray-500"
              >
                <font-awesome-icon class="mr-1" :icon="['fas', 'tags']" />
                <span>{{ $t('bbs.addTopic') }}</span>
              </span>
            </div>
            <GreenBtn
              @click="publishPost"
              classes="w-full mb-2"
              :value="$t('bbs.releaseDynamic')"
              :icon="['fas', 'check']"
          /></div>
        </div>
        <div
          class="flex flex-row mx-4 mb-4 justify-between bg-white dark:bg-gray-700 rounded-md shadow-sm animate-fadeIn30 px-4"
        >
          <div class="flex flex-row">
            <div
              v-for="(item, index) in sort"
              :key="index"
              @click="sortPost(item.type)"
              :class="currentSort === item.type ? 'text-pink-400 dark:bg-gray-800 bg-pink-50' : ''"
              class="text-xs px-3 select-none py-2 flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
            >
              {{ item.title }}
            </div>
          </div>
          <div
            @click="sortPost('desc')"
            class="text-xs px-3 py-2 select-none flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
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
      <div class="w-100">
        <div v-for="(_, v) in 2" :key="v" class="mb-6">
          <div class="flex flex-row justify-between mx-1 animate-fadeIn30 mb-2">
            <div class="text-xs text-gray-500 dark:text-gray-200">
              <font-awesome-icon class="mr-1" :icon="['fas', 'newspaper']" />
              <span>{{ $t('bbs.popularPosts') }}</span>
            </div>
            <router-link
              :to="'/author//post'"
              class="text-xs text-gray-500 dark:text-gray-200 hover:text-pink-400 cursor-pointer duration-300"
            >
              <span class="mr-1">{{ $t('more') }}</span>
              <font-awesome-icon :icon="['fas', 'caret-right']" />
            </router-link>
          </div>
          <div
            class="flex py-10 flex-row dark:bg-gray-700 dark:text-gray-200 text-gray-500 bg-white rounded-md mb-3"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
