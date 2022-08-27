<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:24
 * @FilePath: /pinkmoe_index/components/Header/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script setup lang="ts">
// @unocss-include
import HeaderReCategory from '/@/components/Headerrecategory/index.vue'
import { useHeader } from '/@/hooks/header'
const {
  data,
  auth,
  categoryList,
  toggleTheme,
  setting,
  i18n,
  theme,
  headerBarLeft,
  headerBarRight,
  appStore,
  isShow,
  children,
  headerReCategory,
  categoryIndex,
  showSearchDialog,
  checkInUser,
  showLogin,
  mouseenter,
  seleLanguage,
} = useHeader()
</script>

<template>
  <div
    class="h-48 bg-blue-200 w-full flex justify-end items-center flex-col bg-cover"
    :style="`background-image: url(${appStore.siteBasic?.background});`"
  >
    <div class="w-full z-10 h-full">
      <div class="mb-2 lg:w-3/4 xl:w-5/12 h-32 m-auto animate-fadeIn30">
        <div
          class="flex justify-between items-center pt-2"
          style="text-shadow: 0 0px 2px #fff, 0 0px 5px #fff, 0 0px 10px #fff !important"
        >
          <NuxtLink
            v-for="(k, v) in headerBarLeft"
            :key="v"
            :to="`/${k.url}`"
            :class="v === 0 ? 'rounded-bl-2xl rounded-tl-2xl' : 'rounded-br-2xl rounded-tr-2xl'"
            class="bg-white dark:hover:bg-opacity-50 flex justify-center items-center dark:bg-gray-700 dark:bg-opacity-30 bg-opacity-30 text-xs px-2 py-0.5 hover:bg-opacity-50 duration-300"
          >
            <i :class="`text-gray-700 dark:text-gray-200 inline-block ${k.icon}`" />
            <span class="ml-1">{{ i18n.t(k.name) }}</span>
          </NuxtLink>
          <div class="flex-1" />
          <div
            v-for="(k, v) in headerBarRight"
            :key="v"
            :class="v === 0 ? 'rounded-bl-2xl rounded-tl-2xl' : 'rounded-br-2xl rounded-tr-2xl'"
            class="bg-white cursor-pointer select-none flex justify-center items-center dark:hover:bg-opacity-50 dark:bg-gray-700 dark:bg-opacity-30 bg-opacity-30 text-xs px-2 py-0.5 hover:bg-opacity-50 duration-300"
            @click="k.type === 'theme' ? toggleTheme() : seleLanguage()"
          >
            <i v-if="k.type === 'theme'" :class="`text-gray-700 dark:text-gray-200 inline-block ${theme === 'light' ? 'i-material-symbols:wb-sunny' : 'i-ion:ios-moon'}`" />
            <i v-else :class="`text-gray-700 dark:text-gray-200 inline-block ${k.icon}`" />
            <span class="ml-1">{{
              k.type === 'theme'
                ? theme === 'light'
                  ? i18n.t('head.day')
                  : i18n.t('head.night')
                : i18n.locale.value === 'cn'
                  ? i18n.t('head.chinese')
                  : i18n.t('head.japanese')
            }}</span>
          </div>
        </div>
        <div class="w-52 h-18 mt-4">
          <NuxtLink to="/">
            <img
              v-lazy="appStore.siteBasic?.logo"
              class="hover:animate-swing30 animate-lazyloaded w-full h-full object-cover"
              alt=""
            >
          </NuxtLink>
        </div>
      </div>
      <div
        :style="data.fixed === true ? 'backdrop-filter: blur(6px)' : ''"
        :class="
          data.fixed === true
            ? 'fixed z-10 top-0 w-full left-0 bg-white shadow-menuFixed bg-opacity-30 rounded-t-md'
            : ''
        "
      >
        <div
          :style="data.fixed === false ? 'backdrop-filter: blur(6px)' : ''"
          :class="
            data.fixed === true
              ? ''
              : 'bg-white dark:bg-gray-700 dark:bg-opacity-50 bg-opacity-30 rounded-t-md shadow-category'
          "
          class="flex justify-between lg:w-3/4 xl:w-5/12 m-auto"
        >
          <ul class="flex flex-row justify-start animate-fadeIn30">
            <NuxtLink
              to="/"
              class="h-14 pr-3 pl-3 flex justify-center items-center cursor-pointer text-sm hover:bg-black hover:bg-opacity-10 duration-300"
            >
              <div class="flex flex-col">
                <div
                  class="bg-white bg-opacity-40 rounded-full w-6 h-6 flex justify-center items-center text-shadow-bg-white"
                >
                  <i class="text-gray-700 dark:text-gray-200 i-mdi:home-variant" />
                </div>
              </div>
            </NuxtLink>
            <li
              v-for="(category, v) in categoryList"
              :key="v"
              class="relative group"
              @mouseleave="isShow = false"
            >
              <NuxtLink
                :to="`/category/${category.slug}`"
                class="flex-col h-14 pr-3 pl-3 flex justify-center items-center cursor-pointer text-xs hover:bg-black hover:bg-opacity-10 duration-300"
              >
                <div
                  class="bg-white bg-opacity-40 rounded-full w-6 h-6 flex justify-center items-center text-shadow-bg-white"
                >
                  <span class="iconify text-gray-700 dark:text-gray-200 inline-block " :data-icon="`${category.icon}`" />
                </div>
                <div class="text-shadow-bg-white">
                  {{ category.name }}
                </div>
              </NuxtLink>
              <ul
                class="group-hover:flex dark:bg-gray-600 bg-opacity-90 flex-col pt-3 pb-3 bg-white absolute top-full left-0 hidden shadow-xl animate-fadeIn30 hover:flex"
              >
                <div class="relative">
                  <li
                    v-for="(item, index) in category.children"
                    :key="index"
                    class="flex flex-row childCategory text-xs"
                    @mouseenter="mouseenter(item, index)"
                  >
                    <NuxtLink
                      :to="`/category/${item.slug}`"
                      class="pl-4 py-2 w-32 flex justify-center items-center hover:bg-pink-400 hover:text-white cursor-pointer duration-300"
                    >
                      <span class="iconify inline-block " :data-icon="`${item.icon}`" />
                      <span class="ml-1 mr-4">{{ item.name }}</span>
                      <i v-if="item.children" class="inline-block i-fluent:caret-right-12-filled" />
                    </NuxtLink>
                  </li>
                  <HeaderReCategory
                    v-if="isShow"
                    ref="headerReCategory"
                    :item="children"
                    :index="categoryIndex"
                  />
                </div>
              </ul>
            </li>
          </ul>
          <div class="flex flex-row">
            <div
              v-show="!auth.isLogin"
              class="hover:animate-jello100 relative flex shadow-md shadow-gray-300 hover:shadow-gray-400 select-none bg-pink-400 rounded-2xl justify-center items-center px-3 cursor-pointer text-white text-sm duration-300"
              @click="showLogin"
            >
              <span class="loginBtn bg-pink-400">登陆</span>
            </div>
            <div
              v-show="auth.isLogin" class="flex flex-row"
            >
              <div
                :class="auth.checkInStatus ? 'opacity-30 cursor-not-allowed' : ''"
                class="text-xs hover:animate-jello100 flex justify-center items-center w-16 px-3 bg-pink-400 text-white cursor-pointer select-none"
                @click="checkInUser"
              >
                <i class="inline-block i-mdi:map-marker" />
                <span class="ml-1 font-normal">{{ auth.checkInStatus ? '已签' : '签到' }}</span>
              </div>
              <div class="relative userMenu h-full w-14">
                <div class="flex justify-center items-center px-2 h-full cursor-pointer">
                  <img
                    v-lazy="auth.userInfo?.avatar"
                    class="hover:animate-fadeIn30 animate-lazyloaded h-10 w-10 rounded-full z-10 object-cover"
                    alt=""
                  >
                </div>
                <div class="userMenuList duration-300 flex-col dark:bg-gray-700">
                  <div class="h-14 flex justify-center items-center bg-pink-400 text-xs text-white">
                    <div
                      class="hover:bg-pink-300 cursor-pointer flex justify-center items-center h-full px-4 duration-300"
                    >
                      <i class="inline-block i-ph:paw-print-fill" />
                      <span class="ml-1 font-normal">积分: {{ auth.userInfo?.credit || 0 }}</span>
                    </div>
                  </div>
                  <div class="flex flex-row-reverse">
                    <div
                      v-for="(item, index) in setting"
                      :key="index"
                      class="flex flex-col border-r border-gray-100 dark:border-gray-800 dark:text-gray-200 my-3 group hover:text-pink-400 text-gray-500"
                    >
                      <div class="flex flex-col duration-300 justify-center items-center">
                        <div class="p-1 w-5 h-5 rounded-full flex justify-center items-center group-hover:bg-pink-400 duration-300 bg-gray-500 text-white dark:bg-gray-800">
                          <i :class="`text-xs inline-block ${item.icon}`" />
                        </div>
                        <div class="py-2 text-xs">
                          {{ item.name }}
                        </div>
                      </div>
                      <div v-for="(childitem, index) in item.children" :key="index">
                        <div
                          v-if="childitem.func"
                          class="text-xs text-gray-500 flex justify-center items-center dark:text-gray-200 px-2 py-2 cursor-pointer hover:bg-pink-400 w-24 text-center hover:text-white duration-300"
                          @click="childitem.func()"
                        >
                          <i :class="`inline-block ${childitem.icon}`" />
                          <span class="ml-1">{{ childitem.name }}</span>
                        </div>
                        <NuxtLink v-else :to="childitem.url">
                          <div
                            class="text-xs text-gray-500 flex justify-center items-center dark:text-gray-200 px-2 py-2 cursor-pointer hover:bg-pink-400 w-24 text-center hover:text-white duration-300"
                          >
                            <i :class="`inline-block ${childitem.icon}`" />
                            <span class="ml-1">{{ childitem.name }}</span>
                          </div>
                        </NuxtLink>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div
              class="flex justify-center items-center px-2.5 cursor-pointer text-shadow-bg-white hover:bg-black hover:bg-opacity-10 duration-300"
              @click="showSearchDialog"
            >
              <i class="text-gray-700 dark:text-gray-200 text-2xl i-fluent:search-12-filled" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
  .text-shadow-bg-white {
    text-shadow: 0 0 2px #fff, 0 0 5px #fff, 0 0 10px #fff;
  }

  .childCategory:hover .childrenCategory {
    display: flex;
  }

  .userMenu:hover .userMenuList {
    -webkit-transform: scaleX(1);
    transform: scaleX(1);
    opacity: 1;
    visibility: visible;
    pointer-events: auto;
  }

  .userMenuList {
    display: flex;
    background: hsla(0, 0%, 100%, 0.95);
    -webkit-backdrop-filter: blur(5px);
    backdrop-filter: blur(5px);
    position: absolute;
    top: 0;
    right: 0;
    -webkit-box-shadow: 0 15px 30px 5px rgb(0 0 0 / 15%);
    box-shadow: 0 15px 30px 5px rgb(0 0 0 / 15%);
    opacity: 0;
    visibility: hidden;
    pointer-events: none;
    -webkit-animation-delay: 2s;
    animation-delay: 2s;
    -webkit-transform: scale3d(0.7, 0.5, 1);
    transform: scale3d(0.7, 0.5, 1);
    -webkit-transform-origin: top right;
    -ms-transform-origin: top right;
    transform-origin: top right;
    z-index: 1;
  }

  .loginBtn:after {
    top: 65%;
    bottom: auto;
    border-radius: 0 0 10rem 10rem;
    position: absolute;
    content: '';
    left: 0;
    width: 100%;
    background: rgb(244 114 181);
    height: 25px;
  }

  .loginBtn:before {
    position: absolute;
    content: '';
    bottom: 65%;
    left: 0;
    width: 100%;
    height: 25px;
    background: rgb(244 114 181);
    border-radius: 10rem 10rem 0 0;
  }
</style>
