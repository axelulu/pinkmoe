<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:13:05
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-11 10:07:42
 * @FilePath: /pinkmoe_index/src/components/Usercenterlayout/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Usercenterlayout">
  import { useUserStore } from '/@/store';

  const auth = useUserStore();
  const router = useRouter();
  const route = useRoute();
  const { proxy } = getCurrentInstance();
  const setting = ref<any>([
    {
      icon: 'paint-brush',
      iconType: 'fas',
      name: '文章管理',
      children: [
        {
          icon: 'paint-brush',
          iconType: 'fas',
          name: '资源投稿',
          url: '/user-center/publish',
          func: null,
        },
        {
          icon: 'copy',
          iconType: 'fas',
          name: '我的帖子',
          url: '/user-center/posts',
          func: null,
        },
        {
          icon: 'heart',
          iconType: 'fas',
          name: '收藏夹',
          url: '/user-center/stars',
          func: null,
        },
        {
          icon: 'comments',
          iconType: 'fas',
          name: '我的吐槽',
          url: '/user-center/comments',
          func: null,
        },
      ],
    },
    {
      icon: 'bell',
      iconType: 'fas',
      name: '消息管理',
      children: [
        {
          icon: 'history',
          iconType: 'fas',
          name: '消费记录',
          url: '/user-center/record',
          func: null,
        },
        {
          icon: 'bell',
          iconType: 'fas',
          name: '我的提醒',
          url: '/user-center/notice',
          func: null,
        },
        {
          icon: 'envelope',
          iconType: 'fas',
          name: '即时通讯',
          url: '/user-center/im',
          func: null,
        },
      ],
    },
    {
      icon: 'bag-shopping',
      iconType: 'fas',
      name: '商城论坛',
      children: [
        {
          icon: 'address-card',
          iconType: 'fas',
          name: '收货地址',
          url: '/user-center/home',
          func: null,
        },
        {
          icon: 'basket-shopping',
          iconType: 'fas',
          name: '订单记录',
          url: '/user-center/follow',
          func: null,
        },
        {
          icon: 'people-group',
          iconType: 'fas',
          name: '关注圈子',
          url: '/user-center/fans',
          func: null,
        },
      ],
    },
    {
      icon: 'gamepad',
      iconType: 'fas',
      name: '游戏中心',
      children: [
        {
          icon: 'bookmark',
          iconType: 'fas',
          name: '我的荣誉',
          url: '/user-center/medal',
          func: null,
        },
        {
          icon: 'gift',
          iconType: 'fas',
          name: '抽奖',
          url: '/user-center/lottery',
          func: null,
        },
        {
          icon: 'gift',
          iconType: 'fas',
          name: '商城',
          url: '/user-center/shop',
          func: null,
        },
      ],
    },
    {
      icon: 'address-card',
      iconType: 'fas',
      name: '个人设置',
      children: [
        {
          icon: 'address-card',
          iconType: 'fas',
          name: '我的会员',
          url: '/user-center/vip',
          func: null,
        },
        {
          icon: 'cog',
          iconType: 'fas',
          name: '我的设置',
          url: '/user-center/settings',
          func: null,
        },
        {
          icon: 'power-off',
          iconType: 'fas',
          name: '退出登陆',
          url: '/user-center/publish',
          func: () => {
            auth.logout();
            proxy.$message({
              type: 'success',
              msg: '退出登陆成功',
            });
            setTimeout(() => {
              router.push(route.path);
            }, 1000);
          },
        },
      ],
    },
  ]);
</script>

<template>
  <!-- Usercenterlayout -->
  <div class="flex flex-row mb-2 lg:w-3/4 xl:w-5/12 m-auto mt-4">
    <div class="w-2/12 flex flex-col bg-white dark:bg-gray-700 py-4 rounded-md shadow-sm">
      <div v-for="(item, index) in setting" :key="index">
        <div
          class="text-xs text-gray-500 dark:text-gray-200 dark:hover:bg-gray-900 flex justify-start items-center px-3 py-1.5 cursor-pointer hover:bg-pink-50 hover:text-pink-400 duration-300"
        >
          <font-awesome-icon
            class="p-1.5 rounded-full bg-gray-300 dark:bg-gray-800"
            :icon="[item.iconType, item.icon]"
          />
          <span class="pl-2">{{ item.name }}</span>
        </div>
        <div v-for="(childItem, index) in item.children" :key="index">
          <div
            v-if="childItem.func"
            @click="childItem.func()"
            class="text-xs text-gray-500 dark:text-gray-200 dark:hover:bg-gray-900 flex justify-start items-center px-3 py-1 cursor-pointer hover:bg-pink-50 hover:text-pink-400 duration-300"
          >
            <font-awesome-icon class="p-1.5" :icon="[childItem.iconType, childItem.icon]" />
            <span class="pl-2">{{ childItem.name }}</span>
          </div>
          <router-link :to="childItem.url" v-else>
            <div
              :class="
                router.currentRoute.value.fullPath === childItem.url
                  ? 'bg-pink-50 dark:bg-gray-900 text-pink-400 border-l-2 border-pink-400'
                  : ''
              "
              class="text-xs group text-gray-500 dark:text-gray-200 dark:hover:bg-gray-900 flex justify-start items-center px-3 py-1 cursor-pointer hover:bg-pink-50 hover:text-pink-400 duration-300"
            >
              <font-awesome-icon
                class="p-1.5 group-hover:text-pink-400"
                :class="
                  router.currentRoute.value.fullPath === childItem.url
                    ? 'text-pink-400'
                    : 'text-gray-300'
                "
                :icon="[childItem.iconType, childItem.icon]"
              />
              <span class="pl-2">{{ childItem.name }}</span>
            </div>
          </router-link>
        </div>
      </div>
    </div>
    <div class="w-10/12">
      <slot></slot>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
