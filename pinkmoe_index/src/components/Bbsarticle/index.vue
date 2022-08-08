<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 08:22:26
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 11:16:41
 * @FilePath: /pinkmoe_index/src/components/Bbsarticle/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Bbsarticle">
  import { useBbsArticle } from '/@/hooks/bbs/bbsArticle';
  import { useUtil } from '/@/hooks/util';
  const { formatDate } = useUtil();
  const props = defineProps({
    post: {
      type: Object,
      default: null,
    },
  });
  const {
    status,
    unFollow,
    follow,
    collect,
    unCollect,
    isCollect,
    auth,
    jump,
    deletePostImg,
    lev,
  } = useBbsArticle(props.post);
</script>

<template>
  <!-- Bbsarticle -->
  <article
    class="p-6 flex flex-row dark:bg-gray-700 dark:text-gray-200 text-gray-500 bg-white rounded-md mb-3"
  >
    <router-link :to="'/author/' + post?.AuthorRelation?.uuid + '/userInfo'" class="w-18">
      <img
        class="w-12 h-12 animate-lazyloaded rounded-full overflow-hidden"
        v-lazy="post?.AuthorRelation?.avatar ? '/' + post?.AuthorRelation?.avatar : ''"
        alt=""
      />
    </router-link>
    <div class="w-full flex flex-col ml-0">
      <div class="flex justify-between items-center">
        <div class="flex flex-col">
          <div class="flex flex-row justify-center items-center"
            ><router-link
              :to="'/author/' + post?.AuthorRelation?.uuid + '/userInfo'"
              class="dark:text-white hover:text-pink-400 duration-300 text-gray-700 text-md"
              >{{ post?.AuthorRelation?.nickName }}</router-link
            >
            <span
              :style="`background-color: ${lev?.color};`"
              class="ml-3 select-none cursor-pointer rounded-xl text-xs text-white px-2 py-0.2"
              >{{ lev?.levelName }}</span
            >
            <span
              :style="`background-color: ${post?.AuthorRelation?.authority?.authorityColor};`"
              class="ml-3 select-none cursor-pointer rounded-xl text-xs text-white px-1 py-0.2"
            >
              {{ post?.AuthorRelation?.authority?.authorityName }}</span
            >
          </div>
          <div class="text-sm mt-1 text-gray-400"
            >{{ formatDate(post?.UpdatedAt) }}更新 - {{ post?.view }}次阅读</div
          >
        </div>
        <div v-if="post?.AuthorRelation?.uuid !== auth.userInfo?.uuid" class="flex flex-row">
          <span
            @click="!status ? unFollow() : follow()"
            class="bg-green-500 select-none active:bg-green-500 active:border-green-500 hover:bg-green-600 hover:border-green-600 duration-300 py-1 px-2 cursor-pointer text-xs text-white border-2 border-green-500"
          >
            <font-awesome-icon class="mr-1" :icon="['fas', status ? 'check' : 'add']" />
            <span>{{ status ? '已关注' : '加关注' }}</span>
          </span>
          <span
            @click="jump"
            class="bg-white-500 select-none dark:hover:bg-gray-600 hover:bg-gray-200 active:bg-gray-50 duration-300 py-1 px-2 cursor-pointer text-xs text-gray-500 dark:text-gray-200 border-2 border-gray-200 dark:border-gray-600"
          >
            <font-awesome-icon class="mr-1" :icon="['fas', 'envelope']" />
            <span>站内信</span>
          </span>
        </div>
      </div>
      <div class="mt-2">
        <div v-if="post?.title" class="pb-2 pt-1">
          <router-link
            :to="'/post/' + post?.postId"
            class="text-md font-semibold hover:text-gray-700 duration-300 dark:hover:text-gray-50"
            >{{ post?.title }}</router-link
          >
        </div>
        <div
          class="text-sm max-h-36 overflow-hidden"
          v-html="post?.type === 'active' ? post?.content : deletePostImg(post?.content)"
        ></div>
        <div class="flex flex-wrap mt-4">
          <photo-provider :default-backdrop-opacity="0.3" :loop="true">
            <photo-consumer
              class="w-1/4"
              v-for="src in post?.fileRelation"
              :intro="src.name"
              :key="src"
              :src="'/' + src.url"
            >
              <img
                style="cursor: zoom-in"
                class="h-full w-full mx-2 animate-lazyloaded object-cover overflow-hidden"
                v-lazy="'/' + src.url"
                :alt="src.name"
              />
            </photo-consumer>
          </photo-provider>
        </div>
      </div>
      <div class="flex justify-between items-center mt-4">
        <div>
          <router-link
            :to="'/topic/' + topic.value"
            v-for="(topic, v) in post?.topic"
            :key="v"
            class="border border-pink-400 text-xs text-pink-400 p-1 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer mr-2"
          >
            {{ topic.label }}</router-link
          >
        </div>
        <div class="flex flex-row text-xs">
          <div
            @click="isCollect ? unCollect() : collect()"
            class="mx-2 cursor-pointer hover:text-pink-400 duration-300"
            :class="isCollect ? 'text-pink-400' : ''"
          >
            <font-awesome-icon class="mr-1" :icon="['fas', 'star']" /><span>{{
              isCollect ? '已收藏' : '收藏'
            }}</span>
          </div>
          <router-link
            :to="'/post/' + post?.postId"
            class="mx-2 cursor-pointer hover:text-pink-400 duration-300"
          >
            <font-awesome-icon class="mr-1" :icon="['fas', 'comment']" /><span>评论</span>
          </router-link>
          <div class="mx-2 cursor-pointer hover:text-pink-400 duration-300">
            <font-awesome-icon class="mr-1" :icon="['fas', 'share']" /><span>分享</span>
          </div>
        </div>
      </div>
    </div>
  </article>
</template>

<style lang="less" scoped></style>
