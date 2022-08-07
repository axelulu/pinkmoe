<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:28:53
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:05:29
 * @FilePath: /pinkmoe_index/src/pages/topic/:slug.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="TopicSlug">
  import NotFound from '/@/components/NotFound/index.vue';
  import MoreBtn from '/@/components/Morebtn/index.vue';
  import Spin from '/@/components/Spin/index.vue';
  import { useTopic } from '/@/hooks/topic';
  import { useHead } from '@vueuse/head';
  import { useAppStore } from '/@/store/modules/app';
  const { loading, hasMore, route, topicPostList, sort, sortPost, nextPage, formParams } =
    useTopic();

  const { siteBasic } = useAppStore();
  useHead({
    title: `${route.params.slug} - ${siteBasic?.title}`,
    meta: [
      { name: 'og:type', content: 'topic' },
      {
        name: 'og:title',
        content: `${route.params.slug} - ${siteBasic?.title}`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
</script>

<template>
  <div>
    <!-- Your content -->
    <div class="flex flex-col items-center min-h-86">
      <div v-if="topicPostList" class="lg:w-3/4 xl:w-5/12 text-gray-500">
        <div
          class="flex flex-row bg-white dark:bg-gray-700 dark:text-gray-200 mt-3 rounded-md shadow-sm animate-fadeIn30"
        >
          <router-link
            to="/"
            class="text-xs p-3 ml-1 hover:bg-pink-50 cursor-pointer hover:text-pink-400 dark:hover:bg-gray-800 duration-300"
          >
            <font-awesome-icon :icon="['fas', 'home']" />
          </router-link>
          <router-link
            :to="'/topic/' + topicPostList.list?.value"
            class="flex justify-center items-center"
          >
            <font-awesome-icon class="text-xs py-3 px-2 ml-1" :icon="['fas', 'caret-right']" />
            <span
              class="text-xs py-3 px-2 hover:bg-pink-50 cursor-pointer hover:text-pink-400 dark:hover:bg-gray-800 duration-300"
              >{{ topicPostList.list?.label }}
            </span>
          </router-link>
        </div>
        <div
          class="flex flex-row justify-between bg-white dark:bg-gray-700 mt-3 rounded-md shadow-sm animate-fadeIn30 px-4"
        >
          <div class="flex flex-row">
            <div
              v-for="(item, index) in sort"
              :key="index"
              @click="sortPost(item.type, formParams.desc as boolean)"
              :class="
                formParams.orderKey === item.type ? 'text-pink-400 dark:bg-gray-800 bg-pink-50' : ''
              "
              class="text-xs px-3 select-none py-2 flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
            >
              {{ item.title }}
            </div>
          </div>
          <div
            @click="sortPost('updated_at', !formParams.desc)"
            class="text-xs px-3 py-2 select-none flex justify-center items-center cursor-pointer text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 hover:text-pink-400 hover:bg-pink-50 duration-300"
          >
            {{ formParams.desc ? '倒序' : '正序' }}
          </div>
        </div>
        <Spin :show="loading">
          <div
            v-if="topicPostList.list?.post && topicPostList.list?.post.length"
            class="w-full flex justify-start flex-wrap mt-4 animate-fadeIn30"
          >
            <div v-for="(post, v) in topicPostList.list.post" :key="v" class="w-1/6 p-1.5">
              <Article :post="post" imgHeight="h-60" />
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
