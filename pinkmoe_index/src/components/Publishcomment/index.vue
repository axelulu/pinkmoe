<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 20:55:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:12
 * @FilePath: /pinkmoe_index/src/components/Publishcomment/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Publishcomment">
  import Dialog from '/@/components/Dialog/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import CommentSmile from '/@/components/Commentsmile/index.vue';
  import CommentFont from '/@/components/Commentfont/index.vue';
  import TextareaInput from '/@/components/TextareaInput/index.vue';
  import { useCommentPublish } from '/@/hooks/commentPublish';

  const props = defineProps({
    postId: {
      type: String,
      default: 'login',
    },
  });
  const emit = defineEmits(['getPostComment']);

  const {
    insertSmile,
    insertFont,
    showSmile,
    showFont,
    show,
    submitComment,
    showAnimate,
    commentMetas,
    formComment,
    commentContent,
    smile,
    font,
    dialog,
  } = useCommentPublish(props, emit);

  defineExpose({
    show,
  });
</script>

<template>
  <!-- Publishcomment -->
  <Dialog
    ref="dialog"
    :hide-fun="
      () => {
        showAnimate = false;
      }
    "
    :show-fun="
      () => {
        showAnimate = true;
      }
    "
  >
    <div
      :class="showAnimate ? 'animate-pinkZoomInUp30' : 'animate-pinkZoomOutDown30'"
      class="fixed top-0 bottom-0 left-0 ring-0 flex justify-center items-center h-full w-full pointer-events-none"
    >
      <form
        ref="formComment"
        onsubmit="return false"
        class="lg:w-2/6 xl:w-3/12 bg-gray-50 dark:bg-gray-700 pointer-events-auto flex flex-col rounded-md overflow-hidden"
      >
        <div class="flex flex-row justify-between items-center">
          <div
            class="flex flex-row justify-center items-center text-gray-500 dark:text-gray-200 text-xs ml-4 p-0 rounded-bl-md rounded-br-md overflow-hidden"
          >
            <font-awesome-icon :icon="['fas', 'comments']" />
            <router-link
              :to="'/author/' + commentMetas.FormUidRelation.uuid + '/userInfo'"
              v-if="commentMetas !== null"
              class="ml-1"
              >回复
              <span class="text-pink-400">{{ commentMetas.FormUidRelation.nickName }}</span> 的评论
            </router-link>
            <span v-else class="ml-1">发表评论</span>
          </div>
          <div
            class="text-sm px-4 py-1.5 bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-800 cursor-pointer hover:text-white hover:bg-pink-400 duration-300 rounded-bl-md"
            @click="dialog.hide()"
          >
            <font-awesome-icon :icon="['fas', 'times']" />
          </div>
        </div>
        <div class="pb-4 pt-2 px-4 text-xs h-24">
          <TextareaInput
            :required="true"
            pattern="[\u4e00-\u9fa5]{1,100}$|^[\dA-Za-z_]{1,100}$"
            class="h-full"
            v-model:value="commentContent"
            type="email"
            placeholder="请输入评论内容"
          />
        </div>
        <div class="flex flex-row">
          <div
            @click="showFont"
            class="py-3 w-2/12 flex justify-center dark:bg-gray-700 dark:hover:bg-gray-800 dark:text-gray-200 items-center text-xs bg-gray-100 cursor-pointer hover:bg-gray-200 duration-300"
          >
            <font-awesome-icon :icon="['fas', 'font']" />
          </div>
          <div
            @click="showSmile"
            class="py-3 w-2/12 flex justify-center dark:bg-gray-700 dark:hover:bg-gray-800 dark:text-gray-200 items-center text-xs bg-gray-100 cursor-pointer hover:bg-gray-200 duration-300"
          >
            <font-awesome-icon :icon="['fas', 'smile']" />
          </div>
          <GreenBtn
            @click="submitComment"
            classes="w-8/12"
            :icon="['fas', 'check']"
            value="我说完了"
          />
        </div>
      </form>
    </div>
    <CommentFont @insertFont="insertFont" ref="font" />
    <CommentSmile @insertSmile="insertSmile" ref="smile" />
  </Dialog>
</template>

<style lang="less" scoped></style>
