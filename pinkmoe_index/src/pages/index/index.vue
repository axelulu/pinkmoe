<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:05:07
 * @FilePath: /pinkmoe_index/src/pages/index/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script setup lang="ts">
  import { useHomeList } from '/@/hooks/home';
  import NotFound from '/@/components/NotFound/index.vue';
  import MoreBtn from '/@/components/Morebtn/index.vue';
  import Article from '/@/components/Article/index.vue';
  import Popular from './components/Popular/index.vue';
  import Recommend from './components/Recommend/index.vue';
  import { useAppStore } from '/@/store/modules/app';
  import { useHead } from '@vueuse/head';

  const contentRef = ref();
  const { content, recommend, popular } = useHomeList();
  const data = reactive<any>({
    showBar: false,
    showIndex: 0,
  });
  function jump(i: any) {
    data.showIndex = i;
    document.documentElement.scrollTop = contentRef.value.children[i].offsetTop;
  }

  function scrollHandler() {
    // 当前滚动高度
    let scrollY = document.documentElement.scrollTop;
    if (scrollY > 700) {
      data.showBar = true;
      data.showIndex = -1;
    } else {
      data.showBar = false;
    }
    for (var i = 0; i < contentRef.value.children.length; i++) {
      if (scrollY + 50 > contentRef.value.children[i].offsetTop) {
        data.showIndex = i;
      }
    }
  }

  watchEffect(() => {
    nextTick(() => {
      if (content.value) {
        if (typeof window !== 'undefined') {
          window.removeEventListener('scroll', scrollHandler);
          window.addEventListener('scroll', scrollHandler);
        }
      }
    });
  });
  onUnmounted(() => {
    if (typeof window !== 'undefined') {
      window.removeEventListener('scroll', scrollHandler);
    }
  });
  onMounted(() => {
    scrollHandler();
  });

  const { siteBasic } = useAppStore();
  useHead({
    title: `${siteBasic?.title} - ${siteBasic?.desc}`,
    meta: [
      { name: 'og:type', content: 'index' },
      {
        name: 'og:title',
        content: `${siteBasic?.title} - ${siteBasic?.desc}`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
</script>
<template>
  <div class="flex flex-col items-center relative">
    <Popular :popular="popular" />
    <Recommend :recommend="recommend" />
    <div
      v-if="data?.showBar"
      class="fixed top-20 right-siderBar z-10 bg-white dark:bg-gray-700 dark:text-gray-200 border border-pink-400"
    >
      <div v-for="(item, index) in content" :key="index">
        <div
          @click="jump(index)"
          :class="data.showIndex === index ? 'bg-pink-400 text-white' : 'text-gray-500'"
          class="text-xs p-2 cursor-pointer hover:bg-pink-400 hover:text-white dark:text-gray-200 duration-300 flex flex-col"
        >
          <font-awesome-icon
            class="text-gray-700 dark:text-gray-200"
            :icon="[item.iconType, item.icon]"
          />
          <span :class="data.showIndex === index ? 'block' : 'hidden'" class="col-text mt-1">{{
            item.name
          }}</span>
        </div>
      </div>
    </div>
    <div ref="contentRef" class="w-full animate-fadeIn30">
      <div
        v-for="(item, index) in content"
        :key="index"
        class="lg:w-3/4 xl:w-5/12 mt-4 flex flex-col m-auto"
      >
        <div class="ml-2 flex flex-row">
          <div
            class="w-7 h-7 bg-sky-400 rounded-full flex justify-center items-center text-white mr-1"
          >
            <font-awesome-icon :icon="[item.iconType, item.icon]" />
          </div>
          <div class="text-lg">{{ item.name }}</div>
          <div class="flex flex-row ml-6 flex-1 justify-start items-center">
            <router-link
              v-for="(topic, v) in item.topic"
              :key="v"
              :to="'/topic/' + topic.value"
              class="h-6 border-2 border-gray-200 dark:border-gray-600 dark:text-gray-200 text-xs px-2 py-0.5 mx-0.5 cursor-pointer text-gray-500 hover:text-black hover:border-gray-300 duration-300"
            >
              {{ topic.label }}
            </router-link>
          </div>
          <router-link
            :to="'/category/' + item.slug"
            class="text-xs text-gray-500 mr-2 flex items-center hover:text-pink-400 dark:text-gray-200 cursor-pointer duration-300"
          >
            更多
            <font-awesome-icon class="ml-1" :icon="['fas', 'caret-right']" />
          </router-link>
        </div>
        <div class="flex justify-start flex-wrap min-h-86">
          <div
            v-if="item.post?.length"
            class="w-full flex justify-start flex-wrap mt-4 animate-fadeIn30"
          >
            <div v-for="(post, v) in item.post" :key="v" class="w-1/6 p-1.5">
              <Article :post="post" :imgHeight="item.style === 1 ? 'h-60' : 'h-32'" />
            </div>
          </div>
          <NotFound v-else />
          <div class="w-full p-1.5 text-gray-500">
            <router-link :to="'/category/' + item.slug">
              <MoreBtn />
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped></style>
