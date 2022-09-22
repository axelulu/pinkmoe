<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:11:34
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 16:51:47
 * @FilePath: /pinkmoe_index/pages/user-center/posts/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="UserCenterPostsIndex">
import UserCenterLayout from '/@/components/Usercenterlayout/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import Spin from '/@/components/Spin/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import { useUserCenterPosts } from '/@/hooks/user-center/posts'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'
import { useUtil } from '/@/hooks/util'
const { formatDate, getThumbnail } = useUtil()
const { userPostList, nextPage, hasMore, loading } = useUserCenterPosts()

definePageMeta({
  middleware: ['user-auth'],
  layout: 'user-center',
})

const { siteBasic } = useAppStore()
useHead({
  title: '文章 - 用户中心',
  meta: [
    { name: 'og:type', content: 'posts' },
    {
      name: 'og:title',
      content: '文章 - 用户中心',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <div class="ml-6">
    <div
      class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
    >
      <div class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
        <i class="inline-block i-fluent:copy-24-filled" />
        <span class="ml-1 select-none">我的帖子</span>
      </div>
      <div class="p-4">
        <div
          class="px-4 py-3 bg-gray-100 dark:bg-gray-800 mt-2 text-xs text-gray-500 dark:text-gray-200"
        >
          <p>这里会展示您发布的帖子！</p>
        </div>
      </div>
      <div class="flex flex-row items-center px-4 pb-4">
        <Spin :show="loading">
          <table v-if="userPostList?.list && userPostList?.list?.length" class="w-full">
            <thead class="text-xs text-gray-500 bg-pink-200">
              <tr>
                <th class="px-2 py-3 w-2/24">
                  封面
                </th>
                <th class="px-2 py-3 w-6/24">
                  标题
                </th>
                <th class="px-2 py-3 w-2/24">
                  操作
                </th>
                <th class="px-2 py-3 w-2/24">
                  状态
                </th>
                <th class="px-2 py-3 w-2/24">
                  类型
                </th>
                <th class="px-2 py-3 w-2/24">
                  日期
                </th>
                <th class="px-2 py-3 w-2/24">
                  评论
                </th>
                <th class="px-2 py-3 w-2/24">
                  下载
                </th>
                <th class="px-2 py-3 w-2/24">
                  查看数
                </th>
                <th class="px-2 py-3 w-2/24">
                  帖子点赞
                </th>
              </tr>
            </thead>
            <tbody class="text-xs text-gray-500">
              <tr v-for="(item, index) in userPostList?.list" :key="index">
                <th class="px-2 py-3 flex justify-center items-center">
                  <img
                    v-lazy="getThumbnail(item.cover)"
                    class="w-10 h-10 animate-lazyloaded object-cover"
                    alt=""
                  >
                </th>
                <th class="px-2 py-3">
                  <div
                    class="cursor-pointer hover:bg-pink-400 hover:text-white duration-300 py-1.5 font-normal"
                  >
                    <span>{{ item.type === 'active' ? item.content : item.title }}</span>
                  </div>
                </th>
                <th class="px-2 py-3">
                  <div
                    class="cursor-pointer hover:bg-pink-400 hover:text-white duration-300 py-1.5"
                  >
                    <div class="flex justify-center items-center">
                      <i class="inline-block i-material-symbols:edit" />
                      <span class="ml-1 font-normal">编辑</span>
                    </div>
                  </div>
                </th>
                <th class="px-2 py-3 font-normal">
                  {{
                    item.status === 'published'
                      ? '已发布'
                      : item.status === 'draft'
                        ? '草稿'
                        : '等待审核'
                  }}
                </th>
                <th class="px-2 py-3 font-normal">
                  {{
                    item.type === 'post'
                      ? '文章'
                      : item.type === 'music'
                        ? '音乐'
                        : item.type === 'active'
                          ? '动态'
                          : '视频'
                  }}
                </th>
                <th class="px-2 py-3 font-normal">
                  {{ formatDate(item?.UpdatedAt) }}
                </th>
                <th class="px-2 py-3">
                  <div class="flex justify-center items-center">
                    <i class="inline-block i-fa6-solid:comment" />
                    <span class="ml-1 font-normal">{{ item.commentRelation.length }}</span>
                  </div>
                </th>
                <th class="px-2 py-3">
                  <div class="flex justify-center items-center">
                    <i class="inline-block i-material-symbols:cloud-download" />
                    <span class="ml-1 font-normal">{{ item.downloadRelation.length }}</span>
                  </div>
                </th>
                <th class="px-2 py-3">
                  <div class="flex justify-center items-center">
                    <i class="inline-block i-ant-design:eye-filled" />
                    <span class="ml-1 font-normal">{{ item.view }}</span>
                  </div>
                </th>
                <th class="px-2 py-3">
                  <div class="flex justify-center items-center">
                    <i class="inline-block i-ion:ios-thumbs-up" />
                    <span class="ml-1 font-normal">{{ item.starRelation.length }}</span>
                  </div>
                </th>
              </tr>
            </tbody>
            <tfoot class="text-xs text-gray-500 bg-pink-200">
              <tr>
                <th class="px-2 py-3 w-2/24">
                  封面
                </th>
                <th class="px-2 py-3 w-6/24">
                  标题
                </th>
                <th class="px-2 py-3 w-2/24">
                  操作
                </th>
                <th class="px-2 py-3 w-2/24">
                  状态
                </th>
                <th class="px-2 py-3 w-2/24">
                  类型
                </th>
                <th class="px-2 py-3 w-2/24">
                  日期
                </th>
                <th class="px-2 py-3 w-2/24">
                  评论
                </th>
                <th class="px-2 py-3 w-2/24">
                  下载
                </th>
                <th class="px-2 py-3 w-2/24">
                  查看数
                </th>
                <th class="px-2 py-3 w-2/24">
                  帖子点赞
                </th>
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
</template>
