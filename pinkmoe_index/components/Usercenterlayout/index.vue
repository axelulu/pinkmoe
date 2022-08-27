<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:13:05
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 14:58:10
 * @FilePath: /pinkmoe_index/components/Usercenterlayout/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Usercenterlayout">
import { useUserStore } from '/@/store/modules/user';

const auth = useUserStore()
const router = useRouter()
const route = useRoute()
const { proxy } = getCurrentInstance()
// @unocss-include
const setting = ref<any>([
  {
    icon: 'i-ph:paint-brush-broad-fill',
    name: '文章管理',
    children: [
      {
        icon: 'i-ph:paint-brush-broad-fill',
        name: '资源投稿',
        url: '/user-center/publish',
        func: null,
      },
      {
        icon: 'i-fluent:copy-24-filled',
        name: '我的帖子',
        url: '/user-center/posts',
        func: null,
      },
      {
        icon: 'i-mdi:cards-heart',
        name: '收藏夹',
        url: '/user-center/stars',
        func: null,
      },
      {
        icon: 'i-fa-solid:comments',
        name: '我的吐槽',
        url: '/user-center/comments',
        func: null,
      },
    ],
  },
  {
    icon: 'i-ph:bell-fill',
    name: '消息管理',
    children: [
      {
        icon: 'i-ri:chat-history-fill',
        name: '消费记录',
        url: '/user-center/record',
        func: null,
      },
      {
        icon: 'i-ph:bell-fill',
        name: '我的提醒',
        url: '/user-center/notice',
        func: null,
      },
      {
        icon: 'i-jam:envelope-f',
        name: '即时通讯',
        url: '/user-center/im',
        func: null,
      },
    ],
  },
  {
    icon: 'i-fa6-solid:bag-shopping',
    name: '商城论坛',
    children: [
      {
        icon: 'i-fa-solid:address-card',
        name: '收货地址',
        url: '/user-center/home',
        func: null,
      },
      {
        icon: 'i-fa6-solid:basket-shopping',
        name: '订单记录',
        url: '/user-center/follow',
        func: null,
      },
      {
        icon: 'i-fa6-solid:people-group',
        name: '关注圈子',
        url: '/user-center/fans',
        func: null,
      },
    ],
  },
  {
    icon: 'i-mdi:gamepad-square',
    name: '游戏中心',
    children: [
      {
        icon: 'i-material-symbols:bookmark',
        name: '我的荣誉',
        url: '/user-center/medal',
        func: null,
      },
      {
        icon: 'i-mdi:gift',
        name: '抽奖',
        url: '/user-center/lottery',
        func: null,
      },
      {
        icon: 'i-material-symbols:shopping-bag',
        name: '商城',
        url: '/user-center/shop',
        func: null,
      },
    ],
  },
  {
    icon: 'i-material-symbols:settings-applications-rounded',
    name: '个人设置',
    children: [
      {
        icon: 'i-icon-park-solid:vip-one',
        name: '我的会员',
        url: '/user-center/vip',
        func: null,
      },
      {
        icon: 'i-mdi:cog',
        name: '我的设置',
        url: '/user-center/settings',
        func: null,
      },
      {
        icon: 'i-fa6-solid:power-off',
        name: '退出登陆',
        url: '/user-center/publish',
        func: () => {
          auth.logout()
          proxy.$message({
            type: 'success',
            msg: '退出登陆成功',
          })
          setTimeout(() => {
            router.push(route.path)
          }, 1000)
        },
      },
    ],
  },
])
</script>

<template>
  <div class="flex flex-row mb-2 lg:w-3/4 xl:w-5/12 m-auto mt-4">
    <div class="w-2/12 flex flex-col bg-white dark:bg-gray-700 py-4 rounded-md shadow-sm">
      <div v-for="(item, index) in setting" :key="index">
        <div
          class="text-xs text-gray-500 dark:text-gray-200 dark:hover:bg-gray-900 flex justify-start items-center px-3 py-1.5 cursor-pointer hover:bg-pink-50 hover:text-pink-400 duration-300"
        >
          <div class="p-1 flex justify-center items-center rounded-full bg-gray-300 dark:bg-gray-800">
            <i :class="`inline-block ${item.icon}`" />
          </div>
          <span class="pl-2">{{ item.name }}</span>
        </div>
        <div v-for="(childItem, index) in item.children" :key="index">
          <div
            v-if="childItem.func"
            class="text-xs text-gray-500 dark:text-gray-200 dark:hover:bg-gray-900 flex justify-start items-center px-3 py-1 cursor-pointer hover:bg-pink-50 hover:text-pink-400 duration-300"
            @click="childItem.func()"
          >
            <i :class="`p-1.5 inline-block ${childItem.icon}`" />
            <span class="pl-2">{{ childItem.name }}</span>
          </div>
          <NuxtLink v-else :to="childItem.url">
            <div
              :class="
                router.currentRoute.value.fullPath === childItem.url
                  ? 'bg-pink-50 dark:bg-gray-900 text-pink-400 border-l-2 border-pink-400'
                  : ''
              "
              class="text-xs group py-2 text-gray-500 dark:text-gray-200 dark:hover:bg-gray-900 flex justify-start items-center px-3 py-1 cursor-pointer hover:bg-pink-50 hover:text-pink-400 duration-300"
            >
              <i
                class="p-1.5 group-hover:text-pink-400"
                :class="
                  router.currentRoute.value.fullPath === childItem.url
                    ? `text-pink-400 inline-block ${childItem.icon}`
                    : `text-gray-300 inline-block ${childItem.icon}`
                "
              />
              <span class="pl-2">{{ childItem.name }}</span>
            </div>
          </NuxtLink>
        </div>
      </div>
    </div>
    <div class="w-10/12">
      <slot />
    </div>
  </div>
</template>
