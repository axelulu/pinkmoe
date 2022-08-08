<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:12:20
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 22:17:24
 * @FilePath: /pinkmoe_index/src/pages/user-center/stars/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="UserCenterStarsIndex">
  import UserCenterLayout from '/@/components/Usercenterlayout/index.vue';
  import MoreBtn from '/@/components/Morebtn/index.vue';
  import Spin from '/@/components/Spin/index.vue';
  import NotFound from '/@/components/NotFound/index.vue';
  import { useUserCenterStars } from '/@/hooks/user-center/stars';
  import { useAppStore } from '/@/store/modules/app';
  import { useHead } from '@vueuse/head';
  import { useUtil } from '/@/hooks/util';
  const { formatDate } = useUtil();
  const { authorCollectList, nextPage, hasMore, loading } = useUserCenterStars();

  const { siteBasic } = useAppStore();
  useHead({
    title: `收藏 - 用户中心`,
    meta: [
      { name: 'og:type', content: 'stars' },
      {
        name: 'og:title',
        content: `收藏 - 用户中心`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
</script>

<template>
  <!-- UserCenterStarsIndex -->
  <UserCenterLayout>
    <div class="ml-6">
      <div
        class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'paint-brush']" />
          <span class="ml-1 select-none">我的收藏</span>
        </div>
        <div class="p-4">
          <div
            class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
          >
            <p>这里会展示您收藏的帖子！</p>
          </div>
        </div>

        <div class="flex flex-row items-center px-4 pb-4">
          <Spin :show="loading">
            <table v-if="authorCollectList?.list && authorCollectList?.list?.length" class="w-full">
              <thead class="text-xs text-gray-500 bg-pink-200">
                <tr>
                  <th class="px-2 py-3">封面</th>
                  <th class="px-2 py-3">标题</th>
                  <th class="px-2 py-3">状态</th>
                  <th class="px-2 py-3">类型</th>
                  <th class="px-2 py-3">日期</th>
                  <th class="px-2 py-3">下载</th>
                  <th class="px-2 py-3">查看数</th>
                  <th class="px-2 py-3">帖子点赞</th>
                </tr>
              </thead>
              <tbody class="text-xs text-gray-500">
                <tr v-for="(item, index) in authorCollectList?.list" :key="index">
                  <th class="px-2 py-3 flex justify-center items-center">
                    <img
                      v-lazy="item?.PostIdRelation?.cover"
                      class="w-10 h-10 animate-lazyloaded object-cover"
                      alt=""
                    />
                  </th>
                  <th class="px-2 py-3">
                    <router-link
                      :to="'/post/' + item?.PostIdRelation?.postId"
                      class="cursor-pointer hover:bg-pink-400 hover:text-white duration-300 py-1.5 font-normal"
                    >
                      <span>{{
                        item?.PostIdRelation?.type === 'active'
                          ? item?.PostIdRelation?.content
                          : item?.PostIdRelation?.title
                      }}</span>
                    </router-link>
                  </th>
                  <th class="px-2 py-3 font-normal"
                    >{{
                      item?.PostIdRelation?.status === 'published'
                        ? '已发布'
                        : item?.PostIdRelation?.status === 'draft'
                        ? '草稿'
                        : '等待审核'
                    }}
                  </th>
                  <th class="px-2 py-3 font-normal"
                    >{{
                      item?.PostIdRelation?.type === 'post'
                        ? '文章'
                        : item?.PostIdRelation?.type === 'music'
                        ? '音乐'
                        : item?.PostIdRelation?.type === 'video'
                        ? '视频'
                        : '动态'
                    }}
                  </th>
                  <th class="px-2 py-3 font-normal">{{
                    formatDate(item?.PostIdRelation?.UpdatedAt)
                  }}</th>
                  <th class="px-2 py-3">
                    <font-awesome-icon :icon="['fas', 'cloud-download-alt']" />
                    <span class="ml-1 font-normal">{{
                      item?.PostIdRelation?.downloadRelation?.length
                    }}</span>
                  </th>
                  <th class="px-2 py-3">
                    <font-awesome-icon :icon="['fas', 'eye']" />
                    <span class="ml-1 font-normal">{{ item?.PostIdRelation?.view }}</span>
                  </th>
                  <th class="px-2 py-3">
                    <font-awesome-icon :icon="['fas', 'thumbs-up']" />
                    <span class="ml-1 font-normal">{{
                      item?.PostIdRelation?.starRelation?.length
                    }}</span>
                  </th>
                </tr>
              </tbody>
              <tfoot class="text-xs text-gray-500 bg-pink-200">
                <tr>
                  <th class="px-2 py-3">封面</th>
                  <th class="px-2 py-3">标题</th>
                  <th class="px-2 py-3">状态</th>
                  <th class="px-2 py-3">类型</th>
                  <th class="px-2 py-3">日期</th>
                  <th class="px-2 py-3">下载</th>
                  <th class="px-2 py-3">查看数</th>
                  <th class="px-2 py-3">帖子点赞</th>
                </tr>
              </tfoot>
            </table>
            <NotFound v-else />
          </Spin>
        </div>
      </div>
      <div class="w-full text-gray-500">
        <MoreBtn v-if="hasMore" @click="nextPage" />
      </div>
    </div>
  </UserCenterLayout>
</template>

<style lang="less" scoped></style>

<route lang="yaml">
meta:
  auth: true
</route>
