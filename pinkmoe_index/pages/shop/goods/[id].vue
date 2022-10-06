<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-11 15:03:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 15:38:10
 * @FilePath: /pinkmoe_index/pages/shop/goods/[id].vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="PostIdHtml">
import { useGoodsItem } from '/@/hooks/shop/goods'
import { useHead } from '@vueuse/head'
import { useAppStore } from '/@/store/modules/app'
import { useUtil } from '/@/hooks/util'
import GreenBtn from '/@/components/Greenbtn/index.vue'
import BasicInput from '/@/components/Basicinput/index.vue'
import SlideAuthor from '/@/components/Slideauthor/index.vue'
import avatar from '/@/assets/images/avatar.jpeg'
import PublishComment from '/@/components/Publishcomment/index.vue'
import PostComment from '/@/components/Postcomment/index.vue'
const { formatDate } = useUtil()

const {
  goodsItem,
  hasMore,
  loading,
  nextPage,
  commentList,
  showComment,
  refreshComment,
  user,
  stock,
  saleAmount,
  route,
  comment,
  num,
  cover,
  changeSizeVal,
} = await useGoodsItem()

const { siteBasic } = useAppStore()
useHead({
  titleTemplate: `${goodsItem.value?.goods?.title} - 商品页面`,
  meta: [
    { name: 'og:type', content: 'goods' },
    {
      name: 'og:title',
      content: `${route.params.id} - ${siteBasic?.title}`,
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
    <div class="flex justify-start flex-wrap lg:w-3/4 xl:w-5/12">
      <div
        class="flex flex-row bg-white w-full dark:bg-gray-700 dark:text-gray-200 mt-3 rounded-md shadow-sm animate-fadeIn30"
      >
        <div class="w-4/12 p-4">
          <img
            v-lazy="cover"
            class="w-80 h-80 object-cover animate-lazyloaded rounded-sm overflow-hidden"
            alt=""
          >
          <div class="flex flex-row">
            <img
              v-for="(k, v) in goodsItem?.goods?.goodsImg"
              :key="v"
              v-lazy="k.url"
              :class="cover === k.url ? 'border-pink-400' : ''"
              class="w-12 h-12 cursor-pointer border-2 border-transparent m-1 object-cover animate-lazyloaded rounded-sm overflow-hidden"
              alt=""
              @click="cover = k.url"
            >
          </div>
        </div>
        <div class="w-8/12 p-4">
          <div class="text-xl text-gray-800 dark:text-gray-200 py-2">
            {{
              goodsItem?.goods?.title
            }}
          </div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">
              售价
            </div>
            <div class="flex flex-row w-full justify-between items-center">
              <div class="text-red-500 text-xl min-w-20">
                ¥{{ saleAmount }}
              </div>
              <div class="text-xs mr-4">
                {{ formatDate(goodsItem?.goods?.UpdatedAt) }}
              </div>
            </div>
          </div>
          <div
            v-for="(size, index) in goodsItem?.goods?.sizeRelation"
            :key="index"
            class="flex flex-row my-2 justify-start items-center"
          >
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">
              {{ size.key }}
            </div>
            <div
              v-for="(sizeItem, i) in size.value"
              :key="i"
              :class="
                sizeItem.status
                  ? 'bg-pink-500 text-white dark:border-pink-600 border-pink-600 dark:text-gray-200 border-gray-100 hover:bg-pink-600 duration-300 active:bg-pink-500 hover:text-white'
                  : 'hover:bg-pink-600 duration-300 active:bg-pink-500 hover:text-white dark:hover:border-pink-600 hover:border-pink-600'
              "
              class="text-xs px-2 mr-2 h-8 flex justify-center items-center border-y-2 border-x-2 py-1.5 cursor-pointer select-none"
              @click="changeSizeVal(i, size.value)"
            >
              {{ sizeItem.value }}
            </div>
          </div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">
              库存
            </div>
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">
              {{ stock }}
            </div>
          </div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">
              已售
            </div>
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">
              {{
                goodsItem?.goods?.sizeValRelation?.[0].saleAmount
              }}
            </div>
          </div>
          <div class="flex flex-row my-2 justify-start items-center">
            <div class="text-sm dark:text-gray-200 text-gray-600 min-w-16">
              购买数量
            </div>
            <div
              class="text-sm dark:text-gray-200 text-gray-600 min-w-16 flex flex-row justify-center items-center"
            >
              <GreenBtn
                classes="w-10 h-8"
                value=""
                icon="i-ic:baseline-minus"
                @click="num > 0 ? num-- : null"
              />
              <BasicInput
                v-model:value="num"
                :is-icon="false"
                class="w-20 text-xs"
                icon="i-ic:round-videocam"
                type="number"
                placeholder="请选择数量"
              />
              <GreenBtn classes="w-10 h-8" value="" icon="i-material-symbols:add" @click="num++" />
            </div>
          </div>
          <div class="flex flex-row justify-between mt-6 items-center">
            <div class="flex flex-row items-center">
              <GreenBtn
                classes="w-20 mr-2"
                value="购买"
                icon="i-material-symbols:check-circle"
                @click="publishPost"
              />
              <GreenBtn
                classes="w-20"
                value="加购物车"
                icon="i-ant-design:plus-circle-filled"
                @click="publishPost"
              />
            </div>
            <div
              class="mr-4 w-16 h-6 bg-gray-500 text-white flex justify-center items-center text-xs hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
            >
              <i class="mr-1 inline-block i-ph:flag-fill" />
              <span>报告</span>
            </div>
          </div>
        </div>
      </div>
      <div class="w-full flex flex-row">
        <div class="w-9/12 pr-2 pt-4">
          <div
            class="bg-white w-full p-4 dark:bg-gray-700 dark:text-gray-200 rounded-md shadow-sm animate-fadeIn30"
          >
            <div class="pb-4 text-sm">
              产品简介
            </div>
            <div v-html="goodsItem?.goods?.content" />
          </div>
          <div
            class="bg-pink-400 rounded-md border-2 border-transparent hover:border-pink-500 hover:opacity-80 flex flex-row items-center p-2 cursor-pointer duration-300 mt-4"
            @click="showComment(null)"
          >
            <div class="rounded-full overflow-hidden mr-2">
              <img
                v-lazy="user.isLogin ? user.userInfo?.avatar : avatar"
                class="h-8 w-8 object-cover animate-lazyloaded"
                alt="登陆"
              >
            </div>
            <div class="text-lg text-white">
              {{
                user.isLogin ? '参与讨论聊一聊～' : '登陆后才能评论哦～'
              }}
            </div>
          </div>
          <PostComment
            :has-more="hasMore"
            :next-page="nextPage"
            :loading="loading"
            :post-comment="commentList?.list"
            @showComment="showComment"
          />
        </div>
        <div class="w-3/12 pl-2 pt-4">
          <SlideAuthor
            :author="goodsItem?.goods?.AuthorRelation"
            :comment-count="goodsItem?.commentCount"
            :fans-count="goodsItem?.fansCount"
            :follow-count="goodsItem?.followCount"
            :post-count="goodsItem?.postCount"
            :follow-status="goodsItem?.followStatus"
          />
        </div>
        <PublishComment
          ref="comment"
          :post-id="route.params.id as string"
          @getPostComment="refreshComment"
        />
      </div>
    </div>
  </div>
</template>
