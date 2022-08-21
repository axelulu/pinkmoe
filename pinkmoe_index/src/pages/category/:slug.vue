<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:01:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-20 21:39:04
 * @FilePath: /pinkmoe_vitesse/src/pages/category/:slug.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="CategoryIndex">
import NotFound from '/@/components/NotFound/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import Spin from '/@/components/Spin/index.vue'
import { useCategory } from '/@/hooks/category'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'

const { loading, hasMore, title, route, categoryPostList, sort, sortPost, nextPage, formParams }
    = useCategory()

const { siteBasic } = useAppStore()
useHead({
  title: computed(() => `${title.value} - 分类页面`),
  meta: [
    { name: 'og:type', content: 'category' },
    {
      name: 'og:title',
      content: `${route.params.slug} - ${siteBasic?.title}`,
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <!-- Category -->
  <div class="flex flex-col items-center min-h-86">
    <div class="lg:w-3/4 xl:w-5/12 text-gray-500">
      <div
        class="flex flex-row bg-white dark:bg-gray-700 dark:text-gray-200 mt-3 rounded-md shadow-sm animate-fadeIn30"
      >
        <router-link
          to="/"
          class="text-xs p-3 ml-1 hover:bg-pink-50 dark:hover:bg-gray-800 cursor-pointer hover:text-pink-400 duration-300"
        >
          <i class="inline-block i-mdi:home-variant" />
        </router-link>
        <router-link
          v-for="(item, index) in categoryPostList?.list?.category"
          :key="index"
          :to="`/category/${item.slug}`"
          class="flex justify-center items-center"
        >
          <i class="text-xs py-3 px-2 ml-1 i-fluent:caret-right-12-filled" />
          <span
            class="text-xs py-3 px-2 hover:bg-pink-50 cursor-pointer hover:text-pink-400 dark:hover:bg-gray-800 duration-300"
          >{{ item.name }}</span>
        </router-link>
      </div>
      <div
        class="flex flex-row justify-between bg-white dark:bg-gray-700 mt-3 rounded-md shadow-sm animate-fadeIn30 px-4"
      >
        <div class="flex flex-row">
          <div
            v-for="(item, index) in sort"
            :key="index"
            :class="
              formParams.orderKey === item.type ? 'text-pink-400 dark:bg-gray-800 bg-pink-50' : ''
            "
            class="text-xs px-3 select-none py-2 flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
            @click="sortPost(item.type, formParams.desc!)"
          >
            {{ item.title }}
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

<style lang="less" scoped></style>
