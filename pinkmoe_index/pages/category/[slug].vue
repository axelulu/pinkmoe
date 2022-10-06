<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:01:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 23:17:29
 * @FilePath: /pinkmoe_index/pages/category/[slug].vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="CategoryIndex">
import NotFound from '/@/components/NotFound/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import Spin from '/@/components/Spin/index.vue'
import { useCategory } from '/@/hooks/category'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'

const { loading, hasMore, title, route, categoryPostList, sort, sortPost, nextPage, formParams }
    = await useCategory()

const { siteBasic } = useAppStore()
useHead({
  titleTemplate: `${title.value} - 分类页面`,
  meta: [
    { name: 'og:type', content: 'category' },
    {
      name: 'og:title',
      content: `${title.value} - ${siteBasic?.title}`,
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})

definePageMeta({
  layout: 'home',
})
</script>

<template>
  <div class="flex flex-col items-center min-h-86">
    <div class="lg:w-3/4 xl:w-5/12 text-gray-500">
      <div
        class="flex flex-row bg-white dark:bg-gray-700 dark:text-gray-200 mt-3 rounded-md shadow-sm animate-fadeIn30"
      >
        <NuxtLink
          to="/"
          class="text-xs p-3 ml-1 hover:bg-pink-50 dark:hover:bg-gray-800 cursor-pointer hover:text-pink-400 duration-300 flex justify-start items-center"
        >
          <i class="inline-block i-mdi:home-variant" />
        </NuxtLink>
        <NuxtLink
          v-for="(item, index) in categoryPostList?.list?.category"
          :key="index"
          :to="`/category/${item.slug}`"
          class="flex justify-center items-center"
        >
          <i class="text-xs py-3 px-2 inline-block ml-1 i-fluent:caret-right-12-filled" />
          <span
            class="text-xs py-3 px-2 hover:bg-pink-50 cursor-pointer hover:text-pink-400 dark:hover:bg-gray-800 duration-300"
          >{{ item.name }}</span>
        </NuxtLink>
      </div>
      <div
        class="flex flex-row justify-between bg-white dark:bg-gray-700 mt-3 rounded-md shadow-sm animate-fadeIn30 px-4"
      >
        <div class="flex flex-row">
          <div
            v-for="(item, index) in sort"
            :key="index"
            :class="
              formParams.orderKey === item.type ? 'text-pink-400 dark:bg-gray-800 bg-pink-50' : 'text-gray-500 dark:text-gray-200'
            "
            class="text-xs px-3 select-none py-2 flex justify-center items-center cursor-pointer dark:hover:bg-gray-800 hover:text-pink-400 hover:bg-pink-50 duration-300"
            @click="sortPost(item.type, formParams.desc!)"
          >
            {{ $t(item.title) }}
          </div>
        </div>
        <div
          class="text-xs px-3 py-2 select-none flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
          @click="sortPost('updated_at', !formParams.desc)"
        >
          {{ formParams.desc ? $t('reverseOrder') : $t('positiveSequence') }}
        </div>
      </div>
      <Spin :show="loading">
        <div
          v-if="categoryPostList?.list?.post?.length"
          class="w-full flex justify-start flex-wrap mt-4 animate-fadeIn30"
        >
          <div v-for="(post, v) in categoryPostList.list.post" :key="v" class="w-1/6 p-1.5">
            <Article :post="post" img-height="h-60" />
          </div>
          <div class="w-full p-1.5 text-gray-500">
            <MoreBtn v-if="hasMore" @click="nextPage" />
          </div>
        </div>
        <NotFound v-else />
      </Spin>
    </div>
  </div>
</template>
