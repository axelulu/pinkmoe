<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:28:28
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 15:37:40
 * @FilePath: /pinkmoe_index/pages/post/[id].vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="PostIdHtml">
import { usePostItem } from '/@/hooks/post'
import PostDownload from '/@/components/Postdownload/index.vue'
import PostVideo from '/@/components/Postvideo/index.vue'
import PostMusic from '/@/components/Postmusic/index.vue'
import PublishComment from '/@/components/Publishcomment/index.vue'
import PostComment from '/@/components/Postcomment/index.vue'
import SlideAuthor from '/@/components/Slideauthor/index.vue'
import SlidePost from '/@/components/Slidepost/index.vue'
import SlideUser from '/@/components/Slideuser/index.vue'
import SlideComment from '/@/components/Slidecomment/index.vue'
import avatar from '/@/assets/images/avatar.jpeg'
import { useHead } from '@vueuse/head'
import { useAppStore } from '/@/store/modules/app'
import { useUtil } from '/@/hooks/util'

import { useI18n } from 'vue-i18n'
const { formatDate } = useUtil()

const {
  postItem,
  recommendPost,
  hasMore,
  loading,
  nextPage,
  commentList,
  showComment,
  share,
  refreshComment,
  user,
  route,
  comment,
  loadingPost,
} = await usePostItem()

const { t } = useI18n()
const { siteBasic } = useAppStore()
useHead({
  titleTemplate: `${postItem.value?.post?.title} - 文章页面`,
  meta: [
    { name: 'og:type', content: 'post' },
    {
      name: 'og:title',
      content: `${postItem.value?.post?.title} - ${siteBasic?.title}`,
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
    <PostVideo
      v-if="
        postItem?.post?.type === 'video'
          && postItem?.post?.videoRelation
          && postItem?.post?.videoRelation.length > 0
      "
      :post-id="postItem.post.postId"
      :cover="postItem.post.cover as string"
    />
    <div class="flex justify-start flex-wrap lg:w-3/4 xl:w-5/12 mt-4">
      <div class="w-9/12 pr-2" style="min-height: 800px">
        <Spin :show="loadingPost" class="flex flex-wrap">
          <div
            class="bg-white w-full min-h-68 dark:bg-gray-700 text-gray-500 dark:text-gray-200 rounded-md overflow-hidden animate-fadeIn30"
          >
            <div v-if="postItem?.post?.title" class="text-center text-2xl pt-4 pb-4 px-4">
              {{ postItem?.post?.title }}
            </div>
            <div
              class="flex flex-row justify-center pt-3 pb-3 border-y border-gray-100 bg-gray-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-200"
            >
              <NuxtLink
                :to="`/category/${postItem?.post?.CategoryRelation?.slug}`"
                class="text-xs mr-4 cursor-pointer flex justify-center items-center hover:bg-pink-50 hover:text-pink-400 dark:hover:bg-gray-800 dark:hover:text-pink-400 duration-300"
              >
                <i class="mr-1 inline-block i-material-symbols:folder-open-rounded" />
                <span class="mr-1">{{ postItem?.post?.CategoryRelation?.name }}</span>
              </NuxtLink>
              <div class="text-xs mr-4 flex justify-center items-center">
                <i class="mr-1 inline-block i-material-symbols:nest-clock-farsight-analog-outline-rounded" />
                <span>{{ formatDate(postItem?.post?.UpdatedAt) }}</span>
              </div>
              <NuxtLink
                :to="`/author/${postItem?.post?.author}/userInfo`"
                class="text-xs mr-4 cursor-pointer flex justify-center items-center hover:bg-pink-50 hover:text-pink-400 dark:hover:bg-gray-800 dark:hover:text-pink-400 duration-300"
              >
                <i class="mr-1 inline-block i-bx:bxs-user-circle" />
                <span>{{ postItem?.post?.AuthorRelation?.nickName }}</span>
              </NuxtLink>
              <div class="text-xs mr-4 flex justify-center items-center">
                <i class="mr-1 inline-block i-ic:outline-play-circle-filled" />
                <span>{{ postItem?.post?.view }}</span>
              </div>
            </div>
            <PostMusic
              v-if="
                postItem?.post?.type === 'music'
                  && postItem?.post?.musicRelation
                  && postItem?.post?.musicRelation.length > 0
              "
              :nick-name="postItem.post.AuthorRelation?.nickName"
              :post-id="postItem.post.postId"
            />
            <div
              class="p-4 break-words text-gray-500 dark:text-gray-200 text-sm"
              v-html="postItem?.post?.content"
            />
            <div class="flex flex-wrap pb-2">
              <div class="pl-4 pb-2 flex flex-wrap">
                <NuxtLink
                  v-for="(topicItem, v) in postItem?.post?.topic"
                  :key="v"
                  :to="`/topic/${topicItem.value}`"
                  class="border border-pink-400 text-xs mb-2 text-pink-400 p-1 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer mr-2"
                >
                  {{ topicItem.value }}
                </NuxtLink>
              </div>
              <div class="pl-4 pb-2 flex flex-row flex-1">
                <div
                  class="w-6 h-6 bg-gray-500 text-white flex justify-center items-center text-xs hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
                  @click="share('weibo')"
                >
                  <i class="inline-block i-ant-design:weibo-circle-filled" />
                </div>
                <div
                  class="w-6 h-6 bg-gray-500 text-white flex justify-center items-center text-xs hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
                  @click="share('qq')"
                >
                  <i class="inline-block i-ri:qq-fill" />
                </div>
                <div
                  class="w-6 h-6 bg-gray-500 text-white flex justify-center items-center text-xs hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
                  @click="share('weixin')"
                >
                  <i class="inline-block i-uiw:weixin" />
                </div>
                <div
                  class="w-6 h-6 bg-gray-500 text-white flex justify-center items-center text-xs hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
                  @click="share('bold')"
                >
                  <i class="inline-block i-ph:aperture-bold" />
                </div>
              </div>
              <div
                class="mr-4 w-16 h-6 mx-4 bg-gray-500 text-white flex justify-center items-center text-xs hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
              >
                <i class="inline-block i-ph:flag-fill" />
                <span>报告</span>
              </div>
            </div>
            <div
              class="px-4 py-3 border-t border-gray-100 dark:border-gray-600 text-gray-500 dark:bg-gray-700 dark:text-gray-200 bg-gray-50"
            >
              <ul class="text-xs list-disc ml-4">
                <li>
                  本作品是由 粉萌次元 会员
                  <NuxtLink
                    class="text-pink-400"
                    :to="`/author/${postItem?.post?.AuthorRelation?.uuid}/userInfo`"
                  >
                    {{ postItem?.post?.AuthorRelation?.nickName }}
                  </NuxtLink>
                  的投递作品。
                </li>
                <li>转载请务请署名并注明出处</li>
                <li>
                  禁止再次修改后发布；任何商业用途均须联系作者。如未经授权用作他处，作者将保留追究侵权者法律责任的权利。
                </li>
              </ul>
            </div>
          </div>
        </Spin>
        <PostDownload
          v-if="postItem?.post?.downloadRelation && postItem?.post?.downloadRelation.length > 0"
          :post-id="postItem.post.postId"
        />
        <div class="flex flex-col mt-4 animate-fadeIn30">
          <div class="text-xs text-gray-500 dark:text-gray-200 mr-4 flex justify-start items-center">
            <i class="mr-1 inline-block i-mdi:cards-heart" />
            <span class="mr-1">推荐内容</span>
          </div>
          <div v-if="recommendPost" class="flex overflow-x-auto">
            <div v-for="(post, v) in recommendPost.list" :key="v" class="w-44 p-1.5 flex-shrink-0">
              <Article :post="post" img-height="h-60" />
            </div>
          </div>
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
      <div v-if="postItem" class="w-3/12 pl-2 relative">
        <SlideAuthor
          :author="postItem?.post?.AuthorRelation"
          :comment-count="postItem.commentCount"
          :fans-count="postItem.fansCount"
          :follow-count="postItem.followCount"
          :post-count="postItem.postCount"
          :follow-status="postItem.followStatus"
        />
        <div>
          <div class="flex flex-row justify-between mx-1 mt-4 animate-fadeIn30">
            <div class="text-xs text-gray-500 dark:text-gray-200 flex justify-center items-center">
              <i class="mr-1 inline-block i-ph:user-fill" />
              <span>Ta的帖子</span>
            </div>
            <NuxtLink
              :to="`/author/${postItem?.post?.AuthorRelation?.uuid}/post`"
              class="text-xs text-gray-500 flex justify-center items-center dark:text-gray-200 hover:text-pink-400 cursor-pointer duration-300"
            >
              <span class="mr-1">{{ $t('more') }}</span>
              <i class="mr-1 inline-block i-fluent:caret-right-12-filled" />
            </NuxtLink>
          </div>
          <SlidePost :posts="postItem.authorPosts" />
          <SlideUser :users="postItem.users" />
          <SlideComment :comments="postItem.comments" />
        </div>
      </div>
    </div>
    <PublishComment
      ref="comment"
      :post-id="route.params.id as string"
      @getPostComment="refreshComment"
    />
  </div>
</template>
