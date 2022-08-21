<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 09:46:34
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-20 19:22:03
 * @FilePath: /pinkmoe_vitesse/src/pages/search/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="SearchIndex">
import MoreBtn from '/@/components/Morebtn/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import Spin from '/@/components/Spin/index.vue'
import SearchReCategory from '/@/components/Searchrecategory/index.vue'
import { useSearch } from '/@/hooks/search'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'

const {
  jumpTo,
  sortPost,
  nextPage,
  getChildCategory,
  getCategoryPost,
  searchReCategory,
  currentCategory,
  categoryList,
  formParams,
  sort,
  searchPostList,
  loading,
  hasMore,
  isShow,
  children,
  route,
} = useSearch()

const { siteBasic } = useAppStore()
useHead({
  title: `${route.query.keyword} - 搜索页面`,
  meta: [
    { name: 'og:type', content: 'search' },
    {
      name: 'og:title',
      content: `${route.query.keyword} - ${siteBasic?.title}`,
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <!-- SearchIndex -->
  <div class="flex flex-col items-center min-h-86">
    <div class="lg:w-3/4 xl:w-5/12">
      <div class="p-4 bg-white dark:bg-gray-700 rounded-md mt-4">
        <input
          v-model="formParams.title"
          type="text"
          placeholder="请输入关键词"
          class="border-b-4 w-full p-2 placeholder:text-lg focus:border-pink-400 dark:text-gray-200 dark:focus:border-pink-400 dark:border-gray-800 duration-300 focus-visible:outline-0 placeholder:text-gray-500 text-gray-500 text-center text-lg border-gray-200 bg-transparent"
          @keyup.enter="jumpTo(formParams.title as string)"
        >
        <div class="flex flex-row pt-4">
          <div class="text-xs text-gray-500 dark:text-gray-200 font-semibold pr-6 py-1">
            分类
          </div>
          <div class="flex flex-col">
            <div v-if="categoryList">
              <div class="flex flex-row">
                <div
                  :class="
                    currentCategory.slug === '0' && currentCategory.type === 'all'
                      ? 'text-white bg-pink-400'
                      : 'text-gray-500'
                  "
                  class="text-xs hover:bg-pink-400 dark:text-gray-200 hover:text-white px-2 py-1 duration-300 cursor-pointer"
                  @click="
                    getChildCategory({
                      slug: '0',
                      type: 'all',
                    })
                  "
                >
                  全部
                </div>
                <div
                  v-for="(item, index) in categoryList"
                  :key="index"
                  :class="
                    currentCategory.slug === item.slug && currentCategory.type !== 'all'
                      ? 'text-white bg-pink-400'
                      : 'text-gray-500'
                  "
                  class="text-xs hover:bg-pink-400 dark:text-gray-200 hover:text-white px-2 py-1 duration-300 cursor-pointer"
                  @click="getChildCategory(item)"
                >
                  {{ item.name }}
                  <i v-if="item.children" class="inline-block i-ion:caret-down-outline" />
                </div>
              </div>
              <SearchReCategory
                v-if="isShow"
                ref="searchReCategory"
                :current-category="currentCategory"
                :item="children"
                @emitCategoryPost="getCategoryPost"
              />
            </div>
          </div>
        </div>
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
            @click="sortPost(item.type, formParams.desc as boolean)"
          >
            {{ item.title }}
          </div>
        </div>
        <div
          class="text-xs px-3 py-2 select-none flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
          @click="sortPost('updated_at', !formParams.desc)"
        >
          {{ formParams.desc ? '倒序' : '正序' }}
        </div>
      </div>
      <div class="flex justify-start flex-wrap mt-4 animate-fadeIn30">
        <Spin :show="loading">
          <div
            v-if="searchPostList && searchPostList.list?.length"
            class="w-full flex justify-start flex-wrap animate-fadeIn30"
          >
            <div v-for="(post, v) in searchPostList.list" :key="v" class="w-1/6 p-1.5">
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
  </div>
</template>

<style lang="less" scoped></style>
