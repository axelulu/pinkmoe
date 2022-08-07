<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 17:31:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:04:55
 * @FilePath: /pinkmoe_index/src/pages/author/:author/userinfo.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="AuthorAuthorUserinfo">
  import { useHead } from '@vueuse/head';
  import AuthorLayout from '/@/components/Authorlayout/index.vue';
  import { useAppStore } from '/@/store/modules/app';
  const detail = ref();
  function userInfos(res) {
    detail.value = [
      {
        key: '昵称',
        value: res.value?.nickName,
        icon: [],
      },
      {
        key: 'UID',
        value: res.value?.uuid,
        icon: [],
      },
      {
        key: '描述',
        value: res.value?.desc,
        icon: [],
      },
      {
        key: '积分',
        value: res.value?.credit,
        icon: [],
      },
      {
        key: '经验',
        value: res.value?.exp,
        icon: [],
      },
      {
        key: '性别',
        value: res.value?.sex === 1 ? '保密' : res.value?.sex === 2 ? '女' : '男',
        icon: [],
      },
    ];
  }

  const { siteBasic } = useAppStore();
  useHead({
    title: `用户信息 - 用户主页`,
    meta: [
      { name: 'og:type', content: 'userInfo' },
      {
        name: 'og:title',
        content: `用户信息 - 用户主页`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
</script>

<template>
  <!-- AuthorAuthorUserinfo -->
  <AuthorLayout @userInfo="userInfos">
    <div class="mt-4 bg-white dark:bg-gray-700 py-4 animate-fadeIn30 min-h-68">
      <div
        v-for="(item, index) in detail"
        :key="index"
        class="px-4 py-3 flex flex-row hover:bg-gray-100 dark:hover:bg-gray-800 duration-300"
      >
        <div class="text-xs text-gray-500 dark:text-gray-200 w-28 font-semibold">{{
          item.key
        }}</div>
        <div class="text-xs text-gray-500 dark:text-gray-200">
          <font-awesome-icon
            v-if="item.icon.length > 0"
            class="text-gray-700 dark:text-gray-200 text-xs"
            :icon="item.icon"
          />
          <span class="ml-1">{{ item.value }}</span>
        </div>
      </div>
    </div></AuthorLayout
  >
</template>

<style lang="less" scoped></style>
