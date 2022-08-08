<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:10:55
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 22:17:00
 * @FilePath: /pinkmoe_index/src/pages/user-center/im/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="UserCenterImIndex">
  import UserCenterLayout from '/@/components/Usercenterlayout/index.vue';
  import BasicInput from '/@/components/Basicinput/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import { useAppStore } from '/@/store/modules/app';
  import { useHead } from '@vueuse/head';
  import { useUserCenterIm } from '/@/hooks/user-center/im';
  import { useUtil } from '/@/hooks/util';

  const { siteBasic } = useAppStore();
  useHead({
    title: `私信 - 用户中心`,
    meta: [
      { name: 'og:type', content: 'im' },
      {
        name: 'og:title',
        content: `私信 - 用户中心`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
  const {
    sendMsg,
    auth,
    msgRef,
    msg,
    addChat,
    chatList,
    relationChat,
    chatRelationList,
    addChatRelation,
    deleteChatRelation,
    currentSendId,
  } = useUserCenterIm();

  const { formatDate } = useUtil();
</script>

<template>
  <!-- UserCenterImIndex -->
  <UserCenterLayout>
    <div class="ml-6">
      <div
        class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'paint-brush']" />
          <span class="ml-1 select-none">即时通讯</span>
        </div>
        <div class="p-4">
          <div
            class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
          >
            <p>这里会展示即时通讯！</p>
          </div>
        </div>
        <div class="flex flex-row items-start px-4 pb-4">
          <div class="w-5/24">
            <BasicInput
              class="w-full text-xs"
              :icon="['fas', 'user']"
              type="text"
              v-model:value="relationChat"
              @keyup.enter="addChatRelation"
              placeholder="输入用户 UID 以对话"
            />
            <div class="overflow-y-scroll">
              <div
                v-for="(chatRelation, v) in chatRelationList?.list"
                :key="v"
                :class="
                  currentSendId === chatRelation?.sendIdRelation?.uuid
                    ? 'bg-pink-400 dark:bg-pink-400 text-white'
                    : 'bg-pink-50 dark:bg-gray-800'
                "
                class="flex flex-row items-center duration-300 justify-between cursor-pointer"
              >
                <div
                  @click="addChat(chatRelation?.sendIdRelation?.uuid)"
                  class="flex flex-row items-center"
                >
                  <img
                    class="w-6 h-6 rounded-full mr-2 ml-1 my-0.5"
                    v-lazy="`${chatRelation?.sendIdRelation?.avatar}`"
                    alt=""
                  />
                  <span class="text-xs flex-1">{{ chatRelation?.sendIdRelation?.nickName }}</span>
                </div>
                <div
                  @click="deleteChatRelation(chatRelation?.sendIdRelation?.uuid)"
                  class="px-2 hover:bg-pink-400 h-6 hover:text-white duration-300 cursor-pointer"
                >
                  <font-awesome-icon :icon="['fas', 'times']" />
                </div>
              </div>
            </div>
          </div>
          <div class="w-19/24 pl-4">
            <div ref="msgRef" class="h-160 overflow-y-scroll">
              <div v-for="(k, v) in chatList?.list" :key="v">
                <div
                  v-if="k.userIdRelation?.uuid === auth?.userInfo?.uuid"
                  class="flex flex-row w-full items-start w-full pb-2 border-b dark:border-gray-800 border-gray-100"
                >
                  <router-link :to="`/author/${k.userIdRelation?.uuid}/userInfo`">
                    <img
                      class="w-12 h-12 rounded-full m-2"
                      v-lazy="`${k.userIdRelation?.avatar}`"
                      alt=""
                    />
                  </router-link>
                  <div>
                    <div class="flex flex-row items-start text-xs mt-2">
                      <router-link
                        :to="`/author/${k.userIdRelation?.uuid}/userInfo`"
                        class="mr-4 hover:text-pink-400 duration-300"
                        >{{ k.userIdRelation?.nickName }}</router-link
                      >
                      <div
                        :style="`background-color: ${k.userIdRelation?.authority?.authorityColor};`"
                        class="mr-4 px-0.5 rounded-md text-white"
                        >{{ k.userIdRelation?.authority?.authorityName }}</div
                      >
                      <div class="mr-4">{{ formatDate(k?.UpdatedAt) }}</div>
                    </div>
                    <div class="mt-1">{{ k.content }}</div>
                  </div>
                </div>
                <div
                  v-else
                  class="flex flex-row w-full items-start justify-end w-full pb-2 border-b dark:border-gray-800 border-gray-100"
                >
                  <div>
                    <div class="flex flex-row items-start text-right text-xs mt-2">
                      <div class="ml-4">{{ formatDate(k?.UpdatedAt) }}</div>
                      <div
                        :style="`background-color: ${k.userIdRelation?.authority?.authorityColor};`"
                        class="ml-4 px-0.5 rounded-md text-white"
                        >{{ k.userIdRelation?.authority?.authorityName }}</div
                      >
                      <router-link
                        :to="`/author/${k.userIdRelation?.uuid}/userInfo`"
                        class="ml-4 hover:text-pink-400 duration-300"
                        >{{ k.userIdRelation?.nickName }}</router-link
                      >
                    </div>
                    <div class="mt-1 text-right">{{ k.content }}</div>
                  </div>
                  <img
                    class="w-12 h-12 rounded-full m-2"
                    v-lazy="`${k.userIdRelation?.avatar}`"
                    alt=""
                  />
                </div>
              </div>
            </div>
            <div class="flex flex-row justify-between items-center">
              <div class="w-11/12 pr-4">
                <BasicInput
                  class="text-xs"
                  :icon="['fas', 'comment-dots']"
                  type="text"
                  v-model:value="msg"
                  @keyup.enter="sendMsg"
                  placeholder="输入对话消息"
                />
              </div>
              <GreenBtn @click="sendMsg" classes="w-1/12" value="发送" :icon="['fas', 'check']" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </UserCenterLayout>
</template>

<style lang="less" scoped></style>

<route lang="yaml">
meta:
  auth: true
</route>
