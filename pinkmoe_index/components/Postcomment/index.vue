<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 18:45:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 15:54:33
 * @FilePath: /pinkmoe_index/src/components/Postcomment/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Postcomment">
import MoreBtn from '/@/components/Morebtn/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import CommentContent from '/@/components/Commentcontent/index.vue'
import Spin from '/@/components/Spin/index.vue'
import { useUtil } from '/@/hooks/util'
defineProps({
  postComment: {
    type: Array,
    default: null,
  },
  loading: {
    type: Boolean,
    default: false,
  },
  hasMore: {
    type: Boolean,
    default: false,
  },
  nextPage: {
    type: Function,
    default: null,
  },
})

const emit = defineEmits(['showComment'])

const { formatDate } = useUtil()

function showComment(item) {
  emit('showComment', item)
}
</script>

<template>
  <div v-if="postComment" class="flex flex-col mt-4 animate-fadeIn30">
    <div class="text-xs text-gray-500 dark:text-gray-200 mr-4 mb-2">
      <i class="mr-1 inline-block i-fa6-solid:comment" />
      <span class="mr-1">评论</span>
    </div>
    <Spin :show="loading">
      <div v-if="postComment" class="w-full">
        <div
          v-for="(comment, v) in postComment"
          :key="v"
          class="flex flex-col w-full bg-white dark:bg-gray-700 p-2 rounded-md border-l-4 border-transparent hover:border-pink-400 duration-300"
        >
          <div class="flex flex-row w-full">
            <div class="w-14">
              <img
                v-lazy="comment.FormUidRelation.avatar"
                class="w-10 h-10 animate-lazyloaded rounded-full object-cover"
                alt="用户"
              >
            </div>
            <div class="flex flex-col w-full">
              <div class="flex flex-row justify-between">
                <div class="flex flex-row justify-center items-center">
                  <NuxtLink
                    :to="`/author/${comment.FormUidRelation.uuid}/userInfo`"
                    class="text-xs h-5 text-pink-400 dark:bg-pink-400 dark:text-gray-200 px-1 flex items-center bg-gray-200 rounded-sm mr-2"
                  >
                    {{ comment.FormUidRelation.nickName }}
                  </NuxtLink>
                  <div
                    class="text-xs h-5 text-white dark:bg-yellow-500 dark:text-gray-200 bg-yellow-500 px-1 flex items-center rounded-sm mr-2"
                  >
                    {{ comment.FormUidRelation.authority.authorityName }}
                  </div>
                  <div
                    v-if="true"
                    class="text-xs h-5 text-pink-400 dark:bg-gray-700 dark:text-gray-200 dark:hover:bg-pink-400 dark:hover:text-white bg-white border px-2 flex items-center border-pink-400 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer rounded-sm mr-2"
                  >
                    我
                  </div>
                  <div class="text-xs text-gray-500 dark:text-gray-200 px-1 rounded-sm">
                    {{ formatDate(comment?.UpdatedAt) }}
                  </div>
                </div>
                <div class="flex flex-row">
                  <div
                    class="text-xs text-pink-400 bg-white dark:bg-gray-700 dark:text-gray-200 dark:hover:bg-pink-400 dark:hover:text-white border-l border-t border-b w-10 py-1 flex justify-center items-center border-pink-400 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
                  >
                    <i class="inline-block i-carbon:thumbs-up-filled" />
                  </div>
                  <div
                    class="text-xs text-pink-400 bg-white dark:bg-gray-700 dark:text-gray-200 dark:hover:bg-pink-400 dark:hover:text-white border w-10 py-1 flex justify-center items-center border-pink-400 hover:bg-pink-400 hover:text-white duration-300 cursor-pointer mr-2"
                    @click="emit('showComment', comment)"
                  >
                    回复
                  </div>
                </div>
              </div>
              <div
                class="text-xs text-gray-500 dark:text-gray-200 py-2 flex flex-row items-center"
                v-html="comment.content"
              />
              <div class="mb-2">
                <div v-for="(item, index) in comment.children" :key="index">
                  <div
                    class="flex flex-row justify-between items-center bg-gray-100 dark:bg-gray-800 dark:border-gray-500 pl-2 border-l-2 border-gray-300"
                  >
                    <div class="text-xs text-gray-500 dark:text-gray-200 py-2 flex flex-row">
                      <NuxtLink
                        :to="`/author/${item.FormUidRelation.uuid}/userInfo`"
                        class="text-pink-400"
                      >
                        {{ item.FormUidRelation.nickName }}
                      </NuxtLink>
                      :
                      <div class="flex flex-row" v-html="item.content" />
                    </div>
                    <div
                      class="text-xs text-pink-400 dark:hover:bg-pink-400 dark:hover:text-white w-10 py-2 flex justify-center items-center hover:bg-pink-400 hover:text-white duration-300 cursor-pointer"
                      @click="emit('showComment', item)"
                    >
                      回复
                    </div>
                  </div>
                  <CommentContent :item="item" @showComment="showComment" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="mt-4">
          <MoreBtn v-if="hasMore" @click="nextPage" />
        </div>
      </div>
      <NotFound v-else classes="h-36" />
    </Spin>
  </div>
</template>
