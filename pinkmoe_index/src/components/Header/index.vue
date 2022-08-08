<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 22:43:59
 * @FilePath: /pinkmoe_index/src/components/Header/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script setup lang="ts">
  import HeaderReCategory from '/@/components/Headerrecategory/index.vue';
  import { useHeader } from '/@/hooks/header';
  import i18n from '/@/locales';
  const {
    data,
    status,
    auth,
    categoryList,
    toggleTheme,
    setting,
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
  } = useHeader();
</script>

<template>
  <!--	Header-->
  <div
    class="h-48 bg-blue-200 w-full flex justify-end items-center flex-col bg-cover"
    :style="'background-image: url(' + appStore.siteBasic?.background + ');'"
  >
    <div class="w-full z-10 h-full">
      <div class="mb-2 lg:w-3/4 xl:w-5/12 h-32 m-auto animate-fadeIn30">
        <div
          class="flex justify-between items-center pt-2"
          style="text-shadow: 0 0px 2px #fff, 0 0px 5px #fff, 0 0px 10px #fff !important"
        >
          <router-link
            :to="`/${k.url}`"
            v-for="(k, v) in headerBarLeft"
            :key="v"
            :class="v === 0 ? 'rounded-bl-2xl rounded-tl-2xl' : 'rounded-br-2xl rounded-tr-2xl'"
            class="bg-white dark:hover:bg-opacity-50 dark:bg-gray-700 dark:bg-opacity-30 bg-opacity-30 text-xs px-2 py-0.5 hover:bg-opacity-50 duration-300"
          >
            <font-awesome-icon class="text-gray-700 dark:text-gray-200" :icon="['fas', k.icon]" />
            <span class="ml-1">{{ k.name }}</span>
          </router-link>
          <div class="flex-1"></div>
          <div
            v-for="(k, v) in headerBarRight"
            @click="k.type === 'theme' ? toggleTheme() : seleLanguage()"
            :key="v"
            :class="v === 0 ? 'rounded-bl-2xl rounded-tl-2xl' : 'rounded-br-2xl rounded-tr-2xl'"
            class="bg-white cursor-pointer select-none dark:hover:bg-opacity-50 dark:bg-gray-700 dark:bg-opacity-30 bg-opacity-30 text-xs px-2 py-0.5 hover:bg-opacity-50 duration-300"
          >
            <font-awesome-icon
              class="text-gray-700 dark:text-gray-200"
              :icon="['fas', k.type === 'theme' ? (theme === 'light' ? 'sun' : 'moon') : k.icon]"
            />
            <span class="ml-1">{{
              k.type === 'theme'
                ? theme === 'light'
                  ? '白天'
                  : '黑夜'
                : i18n.global.locale === 'zh'
                ? '中文'
                : '日文'
            }}</span>
          </div>
        </div>
        <div class="w-52 h-18 mt-4">
          <router-link to="/">
            <img
              class="hover:animate-swing30 animate-lazyloaded w-full h-full object-cover"
              v-lazy="appStore.siteBasic?.logo"
              alt=""
            />
          </router-link>
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
            <router-link
              to="/"
              class="h-14 pr-3 pl-3 flex justify-center items-center cursor-pointer text-sm hover:bg-black hover:bg-opacity-10 duration-300"
            >
              <div class="flex flex-col">
                <div
                  class="bg-white bg-opacity-40 rounded-full w-6 h-6 flex justify-center items-center text-shadow-bg-white"
                >
                  <font-awesome-icon class="text-gray-700 dark:text-gray-200" icon="home" />
                </div>
              </div>
            </router-link>
            <li
              v-for="(category, v) in categoryList"
              :key="v"
              @mouseleave="isShow = false"
              class="relative group"
            >
              <router-link
                :to="'/category/' + category.slug"
                class="flex-col h-14 pr-3 pl-3 flex justify-center items-center cursor-pointer text-xs hover:bg-black hover:bg-opacity-10 duration-300"
              >
                <div
                  class="bg-white bg-opacity-40 rounded-full w-6 h-6 flex justify-center items-center text-shadow-bg-white"
                >
                  <font-awesome-icon
                    class="text-gray-700 dark:text-gray-200"
                    :icon="[category.iconType, category.icon]"
                  />
                </div>
                <div class="text-shadow-bg-white">{{ category.name }}</div>
              </router-link>
              <ul
                class="group-hover:flex dark:bg-gray-600 bg-opacity-90 flex-col pt-3 pb-3 bg-white absolute top-full left-0 hidden shadow-xl animate-fadeIn30 hover:flex"
              >
                <div class="relative">
                  <li
                    v-for="(item, index) in category.children"
                    :key="index"
                    @mouseenter="mouseenter(item, index)"
                    class="flex flex-row childCategory text-xs"
                  >
                    <router-link
                      :to="'/category/' + item.slug"
                      class="pl-4 py-2 w-32 hover:bg-pink-400 hover:text-white cursor-pointer duration-300"
                    >
                      <font-awesome-icon :icon="[item.iconType, item.icon]" />
                      <span class="ml-1 mr-4">{{ item.name }}</span>
                      <font-awesome-icon v-if="item.children" :icon="['fas', 'caret-right']" />
                    </router-link>
                  </li>
                  <HeaderReCategory
                    ref="headerReCategory"
                    v-if="isShow"
                    :item="children"
                    :index="categoryIndex"
                  />
                </div>
              </ul>
            </li>
          </ul>
          <div class="flex flex-row">
            <div
              v-if="!auth.isLogin"
              @click="showLogin"
              class="hover:animate-jello100 relative flex shadow-md shadow-gray-300 hover:shadow-gray-400 select-none bg-pink-400 rounded-2xl justify-center items-center px-3 cursor-pointer text-white text-sm duration-300"
            >
              <span class="loginBtn bg-pink-400">登陆</span>
            </div>
            <div v-else class="flex flex-row">
              <div
                @click="checkInUser"
                :class="status ? 'opacity-30 cursor-not-allowed' : ''"
                class="text-xs hover:animate-jello100 flex justify-center items-center w-16 px-3 bg-pink-400 text-white cursor-pointer select-none"
              >
                <font-awesome-icon :icon="['fas', 'map-marker']" />
                <span class="ml-1 font-normal">{{ status ? '已签' : '签到' }}</span>
              </div>
              <div class="relative userMenu h-full w-14">
                <div class="flex justify-center items-center px-2 h-full cursor-pointer">
                  <img
                    class="hover:animate-fadeIn30 animate-lazyloaded h-10 w-10 rounded-full z-10 object-cover"
                    v-lazy="auth.userInfo?.avatar"
                    alt=""
                  />
                </div>
                <div class="userMenuList duration-300 flex-col dark:bg-gray-700">
                  <div class="h-14 flex justify-center items-center bg-pink-400 text-xs text-white">
                    <div
                      class="hover:bg-pink-300 cursor-pointer flex justify-center items-center h-full px-4 duration-300"
                    >
                      <font-awesome-icon :icon="['fas', 'paw']" />
                      <span class="ml-1 font-normal">积分: {{ auth.userInfo?.credit || 0 }}</span>
                    </div>
                  </div>
                  <div class="flex flex-row-reverse">
                    <div
                      v-for="(item, index) in setting"
                      :key="index"
                      class="flex flex-col border-r border-gray-100 dark:border-gray-800 dark:text-gray-200 my-3 group hover:text-pink-400 text-gray-500"
                    >
                      <div class="flex flex-col justify-center items-center duration-300">
                        <font-awesome-icon
                          class="text-xs p-1.5 rounded-full group-hover:bg-pink-400 duration-300 bg-gray-500 text-white dark:bg-gray-800"
                          :icon="[item.iconType, item.icon]"
                        />
                        <div class="py-2 text-xs">{{ item.name }}</div>
                      </div>
                      <div v-for="(childitem, index) in item.children" :key="index">
                        <div
                          v-if="childitem.func"
                          @click="childitem.func()"
                          class="text-xs text-gray-500 dark:text-gray-200 px-2 py-2 cursor-pointer hover:bg-pink-400 w-24 text-center hover:text-white duration-300"
                        >
                          <font-awesome-icon :icon="[childitem.iconType, childitem.icon]" />
                          <span class="ml-1">{{ childitem.name }}</span>
                        </div>
                        <router-link :to="childitem.url" v-else>
                          <div
                            class="text-xs text-gray-500 dark:text-gray-200 px-2 py-2 cursor-pointer hover:bg-pink-400 w-24 text-center hover:text-white duration-300"
                          >
                            <font-awesome-icon :icon="[childitem.iconType, childitem.icon]" />
                            <span class="ml-1">{{ childitem.name }}</span>
                          </div>
                        </router-link>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div
              @click="showSearchDialog"
              class="flex justify-center items-center px-2.5 cursor-pointer text-shadow-bg-white hover:bg-black hover:bg-opacity-10 duration-300"
            >
              <font-awesome-icon
                class="text-gray-700 dark:text-gray-200 text-2xl"
                :icon="['fas', 'search']"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less">
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
