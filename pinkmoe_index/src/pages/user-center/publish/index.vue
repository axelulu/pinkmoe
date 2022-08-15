<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:11:42
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-15 14:55:32
 * @FilePath: /pinkmoe_index/src/pages/user-center/publish/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="UserCenterPublishIndex">
  import UserCenterLayout from '/@/components/Usercenterlayout/index.vue';
  import BasicInput from '/@/components/Basicinput/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import WhiteBtn from '/@/components/Whitebtn/index.vue';
  import TextareaInput from '/@/components/TextareaInput/index.vue';
  import WangEditor from '/@/components/Wangeditor/index.vue';
  import PublishCategory from '/@/components/Publishcategory/index.vue';
  import { useUserCenterPublish } from '/@/hooks/user-center/publish';
  import { useAppStore } from '/@/store/modules/app';
  import { useHead } from '@vueuse/head';
  const {
    category,
    formParams,
    categoryList,
    wangEditor,
    getCategory,
    publishPost,
    publishCategory,
    chooseCategory,
    uploadPostImg,
    insertPostImg,
    deletePostImg,
    formPublish,
  } = useUserCenterPublish();

  const { siteBasic } = useAppStore();
  useHead({
    title: `发布文章 - 用户中心`,
    meta: [
      { name: 'og:type', content: 'publish' },
      {
        name: 'og:title',
        content: `发布文章 - 用户中心`,
      },
      { name: 'og:url', content: siteBasic?.url },
    ],
  });
</script>

<template>
  <!-- UserCenterPublishIndex -->
  <UserCenterLayout>
    <form ref="formPublish" class="ml-6" onsubmit="return false">
      <div
        class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'paint-brush']" />
          <span class="ml-1 select-none">投稿说明</span>
        </div>
        <div class="p-4">
          <div
            class="px-4 py-3 bg-gray-100 dark:bg-gray-800 dark:text-gray-200 mt-2 text-xs text-gray-500"
          >
            <p
              >欢迎您在粉萌次元进行投稿，在开始前请您简略阅读投稿须知，投稿即代表您同意以下协议啦。</p
            >
          </div>
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'file']" />
          <span class="ml-1 select-none">文章类型</span>
        </div>
        <div class="flex flex-row items-center px-4 pb-4 mt-6">
          <div
            :class="formParams.type === 'post' ? 'bg-green-500 text-white' : 'text-gray-600'"
            @click="formParams.type = 'post'"
            class="text-xs w-1/3 flex justify-center items-center border-y-2 border-gray-100 dark:text-gray-200 dark:border-gray-900 py-1.5 cursor-pointer select-none hover:bg-green-500 duration-300 active:bg-green-600 hover:text-white"
          >
            <font-awesome-icon :icon="['fas', 'file']" />
            <span class="ml-1">文章</span>
          </div>
          <div
            :class="formParams.type === 'video' ? 'bg-green-500 text-white' : 'text-gray-600'"
            @click="formParams.type = 'video'"
            class="text-xs w-1/3 flex justify-center items-center border-y-2 border-gray-100 dark:text-gray-200 dark:border-gray-900 py-1.5 cursor-pointer select-none hover:bg-green-500 duration-300 active:bg-green-600 hover:text-white"
          >
            <font-awesome-icon :icon="['fas', 'video']" />
            <span class="ml-1">视频</span>
          </div>
          <div
            :class="formParams.type === 'music' ? 'bg-green-500 text-white' : 'text-gray-600'"
            @click="formParams.type = 'music'"
            class="text-xs w-1/3 flex justify-center items-center border-y-2 border-gray-100 dark:text-gray-200 dark:border-gray-900 py-1.5 cursor-pointer select-none hover:bg-green-500 duration-300 active:bg-green-600 hover:text-white"
          >
            <font-awesome-icon :icon="['fas', 'music']" />
            <span class="ml-1">音乐</span>
          </div>
        </div>
      </div>
      <div
        v-if="formParams.type === 'video'"
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative animate-fadeIn30"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'video']" />
          <span class="ml-1 select-none">视频</span>
        </div>
        <div class="flex flex-wrap items-center px-4 pb-4 mt-6">
          <div
            v-for="(item, index) in formParams.videoRelation"
            :key="index"
            class="flex flex-row w-full"
          >
            <div class="w-3/12 pr-2">
              <BasicInput
                class="w-full text-xs"
                :icon="['fas', 'video']"
                v-model:value="item.name"
                type="text"
                placeholder="视频名称"
              />
            </div>
            <div class="w-5/12 pr-2">
              <BasicInput
                class="w-full text-xs"
                :icon="['fas', 'link']"
                v-model:value="item.url"
                type="text"
                placeholder="视频链接"
              />
            </div>
            <div class="w-4/12 pr-2">
              <BasicInput
                class="w-full text-xs"
                :icon="['fas', 'link']"
                v-model:value="item.subtitles"
                type="text"
                placeholder="视频字幕地址"
              />
            </div>
            <div class="w-2/12 pr-2 flex flex-row justify-center items-center">
              <div
                :class="item.priceType === 'credit' ? 'bg-green-500 text-white' : 'text-gray-500'"
                @click="item.priceType = 'credit'"
                class="text-xs w-6/12 h-8 flex justify-center items-center dark:border-gray-900 dark:text-gray-200 border-y-2 border-l-2 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
              >
                积分
              </div>
              <div
                :class="item.priceType === 'cash' ? 'bg-green-500 text-white' : 'text-gray-500'"
                @click="item.priceType = 'cash'"
                class="text-xs w-6/12 h-8 flex justify-center items-center dark:border-gray-900 dark:text-gray-200 border-y-2 border-r-2 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
              >
                金币
              </div>
            </div>
            <div class="w-2/12 pr-2">
              <BasicInput
                class="text-xs"
                :icon="['fas', 'cloud-download-alt']"
                v-model:value="item.price"
                type="number"
                placeholder="下载所需"
              />
            </div>
            <div class="w-2/12 flex flex-row items-center">
              <div
                @click="
                  formParams.videoRelation?.length! > 24
                    ? ''
                    : formParams.videoRelation?.push({
                        name: '',
                        url: '',
                        subtitles: '',
                        price: 0,
                        priceType: 'credit',
                      })
                "
                :class="
                  formParams.videoRelation?.length! > 24
                    ? 'opacity-60 cursor-not-allowed'
                    : 'cursor-pointer'
                "
                class="w-full h-8 flex justify-center items-center select-none py-1 hover:bg-green-600 duration-300 active:bg-green-500 bg-green-500 border-2 border-green-600 text-xs text-white text-center"
              >
                <font-awesome-icon :icon="['fas', 'add']" />
              </div>
              <div
                @click="
                  formParams.videoRelation?.length! < 2
                    ? ''
                    : formParams.videoRelation?.splice(index, 1)
                "
                :class="
                  formParams.videoRelation?.length! < 2
                    ? 'opacity-60 cursor-not-allowed'
                    : 'cursor-pointer'
                "
                class="w-full h-8 flex justify-center items-center select-none py-1 cursor-pointer hover:bg-red-600 duration-300 active:bg-red-500 bg-red-500 border-2 border-red-600 text-xs text-white text-center"
              >
                <font-awesome-icon :icon="['fas', 'trash']" />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div
        v-if="formParams.type === 'music'"
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative animate-fadeIn30"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'music']" />
          <span class="ml-1 select-none">音乐</span>
        </div>
        <div class="flex flex-wrap items-center px-4 pb-4 mt-6">
          <div
            v-for="(item, index) in formParams.musicRelation"
            :key="index"
            class="flex flex-row w-full"
          >
            <div class="w-3/12 pr-2">
              <BasicInput
                class="w-full text-xs"
                :icon="['fas', 'music']"
                v-model:value="item.name"
                type="text"
                placeholder="音乐名称"
              />
            </div>
            <div class="w-5/12 pr-2">
              <BasicInput
                class="w-full text-xs"
                :icon="['fas', 'link']"
                v-model:value="item.url"
                type="text"
                placeholder="音乐页面地址（网易云音乐或文件直连）"
              />
            </div>
            <div class="w-2/12 pr-2 flex flex-row justify-center items-center">
              <div
                :class="item.priceType === 'credit' ? 'bg-green-500 text-white' : 'text-gray-500'"
                @click="item.priceType = 'credit'"
                class="text-xs w-6/12 h-8 flex justify-center items-center dark:border-gray-900 dark:text-gray-200 border-y-2 border-l-2 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
              >
                积分
              </div>
              <div
                :class="item.priceType === 'cash' ? 'bg-green-500 text-white' : 'text-gray-500'"
                @click="item.priceType = 'cash'"
                class="text-xs w-6/12 h-8 flex justify-center items-center dark:border-gray-900 dark:text-gray-200 border-y-2 border-r-2 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
              >
                金币
              </div>
            </div>
            <div class="w-2/12 pr-2">
              <BasicInput
                class="text-xs"
                :icon="['fas', 'cloud-download-alt']"
                v-model:value="item.price"
                type="number"
                placeholder="下载所需"
              />
            </div>
            <div class="w-2/12 flex flex-row items-center">
              <div
                @click="
                  formParams.musicRelation?.length! > 24
                    ? ''
                    : formParams.musicRelation?.push({
                        name: '',
                        url: '',
                        price: 0,
                        priceType: 'credit',
                      })
                "
                :class="
                  formParams.musicRelation?.length! > 24
                    ? 'opacity-60 cursor-not-allowed'
                    : 'cursor-pointer'
                "
                class="w-full h-8 flex justify-center items-center select-none py-1 hover:bg-green-600 duration-300 active:bg-green-500 bg-green-500 border-2 border-green-600 text-xs text-white text-center"
              >
                <font-awesome-icon :icon="['fas', 'add']" />
              </div>
              <div
                @click="
                  formParams.musicRelation?.length! < 2
                    ? ''
                    : formParams.musicRelation?.splice(index, 1)
                "
                :class="
                  formParams.musicRelation?.length! < 2
                    ? 'opacity-60 cursor-not-allowed'
                    : 'cursor-pointer'
                "
                class="w-full h-8 flex justify-center items-center select-none py-1 cursor-pointer hover:bg-red-600 duration-300 active:bg-red-500 bg-red-500 border-2 border-red-600 text-xs text-white text-center"
              >
                <font-awesome-icon :icon="['fas', 'trash']" />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'file']" />
          <span class="ml-1 select-none">文章标题</span>
        </div>
        <div class="flex flex-row items-center px-4 pb-4 mt-6">
          <BasicInput
            class="w-full text-xs"
            :icon="['fas', 'file']"
            :required="true"
            pattern="[\u4e00-\u9fa5]{1,50}$|^[\dA-Za-z_]{1,50}$"
            v-model:value="formParams.title"
            type="text"
            placeholder="请输入文章标题"
          />
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'file']" />
          <span class="ml-1 select-none">文章摘要</span>
        </div>
        <div class="flex flex-row items-center px-4 pb-4 mt-6">
          <TextareaInput
            class="w-full text-xs"
            v-model:value="formParams.exerpt"
            :required="true"
            pattern="[\u4e00-\u9fa5]{1,120}$|^[\dA-Za-z_]{1,120}$"
            rows="5"
            maxlen="120"
            placeholder="请输入文章摘要"
          />
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'paint-brush']" />
          <span class="ml-1 select-none">文章内容</span>
        </div>
        <div class="flex flex-row items-center px-4 pb-4 mt-6">
          <WangEditor ref="wangEditor" v-model:value="formParams.content" />
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'images']" />
          <span class="ml-1 select-none">文章图片</span>
        </div>
        <div class="flex flex-col items-center px-4 pb-4 mt-6">
          <label
            for="avatar"
            class="border-2 border-dashed hover:bg-gray-100 dark:hover:bg-gray-800 dark:border-gray-200 duration-300 border-gray-500 cursor-pointer h-12 w-full flex justify-center items-center"
          >
            <font-awesome-icon
              class="text-xs text-gray-500 dark:text-gray-200"
              :icon="['fas', 'hand-paper']"
            />
            <span class="ml-2 text-xs text-gray-500 dark:text-gray-200 select-none">添加图片</span>
            <input
              @change="uploadPostImg"
              accept="image/gif, image/jpg, image/jpeg, image/png, image/webp"
              type="file"
              id="avatar"
              hidden
              value=""
            />
          </label>
          <div class="flex flex-col w-full mt-2">
            <div
              v-for="(item, index) in formParams.postImg"
              :key="index"
              class="flex flex-row w-full"
            >
              <img class="w-16 h-16 mt-2 mr-2 object-cover" :src="item.url" alt="" />
              <div class="flex flex-col flex-1">
                <TextareaInput class="w-full text-xs h-8" :value="item.name" />
                <div class="flex flex-row justify-between">
                  <div class="flex flex-row w-full">
                    <GreenBtn
                      @click="insertPostImg(item)"
                      classes="w-2/12"
                      value="插入到文章"
                      :icon="['fas', 'plus-circle']"
                    />
                    <GreenBtn
                      v-if="formParams.cover === item.url"
                      @click="formParams.cover = item.url"
                      classes="w-2/12"
                      value="设为封面"
                      :icon="['fas', 'star']"
                    />
                    <WhiteBtn
                      v-else
                      @click="formParams.cover = item.url"
                      class="w-2/12"
                      value="设为封面"
                      :icon="['fas', 'star']"
                    />
                  </div>
                  <div
                    @click="deletePostImg(index)"
                    class="w-1/12 h-8 flex justify-center items-center select-none py-1 cursor-pointer hover:bg-red-600 duration-300 active:bg-red-500 bg-red-500 border-2 border-red-600 text-xs text-white text-center"
                  >
                    <font-awesome-icon :icon="['fas', 'trash']" />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'sitemap']" />
          <span class="ml-1 select-none">文章分类</span>
        </div>
        <div
          @click="getCategory"
          class="flex flex-row items-center px-2 py-1 cursor-pointer bg-gray-50 dark:bg-gray-800 dark:hover:bg-gray-600 dark:hover:text-gray-200 dark:text-gray-200 dark:border-gray-900 hover:bg-white hover:text-gray-600 duration-300 mt-6 border-2 border-gray-200 mx-4 mb-4 text-xs text-gray-500"
        >
          <font-awesome-icon class="" :icon="['fas', 'sitemap']" />
          <span class="ml-2 select-none">{{ category.name ? category.name : '选择文章分类' }}</span>
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'tags']" />
          <span class="ml-1 select-none">文章标签</span>
        </div>
        <div class="flex flex-col items-center px-4 pb-4 mt-6">
          <div class="flex flex-wrap items-center justify-start w-full">
            <div v-for="(item, index) in formParams.topic" :key="index" class="w-3/12 pr-4">
              <div class="relative my-1.5 flex flex-row text-xs">
                <input
                  placeholder="请输入标签"
                  @input="(e) => formParams.topic[index] = (e.target as HTMLTextAreaElement).value"
                  class="border-2 w-full border-gray-100 dark:bg-gray-800 dark:border-gray-900 dark:text-gray-200 py-1.5 pl-2 pr-8 focus:border-pink-400 duration-300 focus-visible:outline-0"
                  type="text"
                  :value="item"
                />
                <font-awesome-icon
                  @click="formParams.topic?.length! > 1 ? formParams.topic?.splice(index, 1) : ''"
                  :class="
                    formParams.topic?.length! > 1
                      ? 'text-gray-500 dark:text-gray-300'
                      : 'text-gray-300 dark:text-gray-500 cursor-not-allowed'
                  "
                  class="absolute right-0 top-0 p-2.5 cursor-pointer dark:active:bg-gray-600 active:bg-gray-100 duration-300"
                  :icon="['fas', 'trash']"
                />
              </div>
            </div>
          </div>
          <GreenBtn
            :class="formParams.topic?.length! > 7 ? 'opacity-60 cursor-not-allowed' : ''"
            classes="w-full mt-2"
            @click="formParams.topic?.length! > 7 ? '' : formParams.topic?.push('')"
            value=""
            :icon="['fas', 'plus']"
          />
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'cloud-download-alt']" />
          <span class="ml-1 select-none">下载链接</span>
        </div>
        <div class="flex flex-wrap items-center px-4 pb-4 mt-6">
          <div v-if="formParams?.downloadRelation?.length!">
            <div
              v-for="(item, index) in formParams.downloadRelation"
              :key="index"
              class="flex flex-row w-full"
            >
              <div class="w-2/12 pr-2">
                <BasicInput
                  class="w-full text-xs"
                  :icon="['fas', 'cloud-download-alt']"
                  v-model:value="item.name"
                  type="text"
                  pattern="^[\u4E00-\u9FA5A-Za-z0-9_]+$"
                  placeholder="下载名称"
                />
              </div>
              <div class="w-2/12 pr-2">
                <BasicInput
                  class="w-full text-xs"
                  :icon="['fas', 'link']"
                  v-model:value="item.url"
                  type="text"
                  placeholder="下载链接"
                />
              </div>
              <div class="w-2/12 pr-2">
                <BasicInput
                  class="w-full text-xs"
                  :icon="['fas', 'key']"
                  v-model:value="item.extractPwd"
                  type="text"
                  placeholder="提取密码"
                />
              </div>
              <div class="w-2/12 pr-2">
                <BasicInput
                  class="w-full text-xs"
                  :icon="['fas', 'unlock']"
                  v-model:value="item.unpackPwd"
                  type="text"
                  placeholder="解压密码"
                />
              </div>
              <div class="w-2/12 pr-2 flex flex-row justify-center items-center">
                <div
                  :class="item.priceType === 'credit' ? 'bg-green-500 text-white' : 'text-gray-500'"
                  @click="item.priceType = 'credit'"
                  class="text-xs w-6/12 h-8 flex justify-center items-center dark:border-gray-900 dark:text-gray-200 border-y-2 border-l-2 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
                >
                  积分
                </div>
                <div
                  :class="item.priceType === 'cash' ? 'bg-green-500 text-white' : 'text-gray-500'"
                  @click="item.priceType = 'cash'"
                  class="text-xs w-6/12 h-8 flex justify-center items-center dark:border-gray-900 dark:text-gray-200 border-y-2 border-r-2 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
                >
                  金币
                </div>
              </div>
              <div class="w-2/12 pr-2">
                <BasicInput
                  class="text-xs"
                  :icon="['fas', 'cloud-download-alt']"
                  v-model:value="item.price"
                  type="number"
                  placeholder="下载所需"
                />
              </div>
              <div class="w-2/12 flex flex-row items-center">
                <div
                  @click="
                    formParams.downloadRelation?.length! > 24
                      ? ''
                      : formParams.downloadRelation?.push({
                          name: '',
                          url: '',
                          extractPwd: '',
                          unpackPwd: '',
                          priceType: 'credit',
                          price: 0,
                        })
                  "
                  :class="
                    formParams.downloadRelation.length > 24
                      ? 'opacity-60 cursor-not-allowed'
                      : 'cursor-pointer'
                  "
                  class="w-full h-8 flex justify-center items-center select-none py-1 hover:bg-green-600 duration-300 active:bg-green-500 bg-green-500 border-2 border-green-600 text-xs text-white text-center"
                >
                  <font-awesome-icon :icon="['fas', 'add']" />
                </div>
                <div
                  @click="
                    formParams.downloadRelation?.length! < 1
                      ? ''
                      : formParams.downloadRelation?.splice(index, 1)
                  "
                  :class="
                    formParams.downloadRelation.length < 1
                      ? 'opacity-60 cursor-not-allowed'
                      : 'cursor-pointer'
                  "
                  class="w-full h-8 flex justify-center items-center select-none py-1 cursor-pointer hover:bg-red-600 duration-300 active:bg-red-500 bg-red-500 border-2 border-red-600 text-xs text-white text-center"
                >
                  <font-awesome-icon :icon="['fas', 'trash']" />
                </div>
              </div>
            </div>
          </div>
          <GreenBtn
            v-else
            :class="formParams.downloadRelation?.length! > 7 ? 'opacity-60 cursor-not-allowed' : ''"
            classes="w-full mt-2"
            @click="
              formParams.downloadRelation?.push({
                name: '',
                url: '',
                extractPwd: '',
                unpackPwd: '',
                priceType: 'credit',
                price: 0,
              })
            "
            value=""
            :icon="['fas', 'plus']"
          />
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'truck']" />
          <span class="ml-1 select-none">作品来源</span>
        </div>
        <div class="flex flex-row items-center px-4 pb-4 mt-6">
          <div
            :class="formParams.from == 'original' ? 'bg-green-500 text-white' : 'text-gray-500'"
            @click="formParams.from = 'original'"
            class="text-xs w-1/2 flex justify-center items-center border-2 dark:border-gray-900 dark:text-gray-200 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
          >
            原创
          </div>
          <div
            :class="formParams.from == 'reprinted' ? 'bg-green-500 text-white' : 'text-gray-500'"
            @click="formParams.from = 'reprinted'"
            class="text-xs w-1/2 flex justify-center items-center border-2 dark:border-gray-900 dark:text-gray-200 border-gray-100 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
          >
            转载
          </div>
        </div>
      </div>
      <div
        class="w-full mt-5 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <font-awesome-icon :icon="['fas', 'truck']" />
          <span class="ml-1 select-none">发布状态</span>
        </div>
        <div class="flex flex-row items-center px-4 pb-4 mt-6">
          <div
            :class="formParams.status == 'pending' ? 'bg-green-500 text-white' : 'text-gray-500'"
            @click="formParams.status = 'pending'"
            class="text-xs w-1/2 flex justify-center items-center border-2 dark:text-gray-200 border-gray-100 dark:border-gray-900 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
          >
            提交审核
          </div>
          <div
            :class="formParams.status == 'published' ? 'bg-green-500 text-white' : 'text-gray-500'"
            @click="formParams.status = 'draft'"
            class="text-xs w-1/2 flex justify-center items-center border-2 dark:text-gray-200 border-gray-100 dark:border-gray-900 py-1.5 cursor-pointer select-none hover:bg-green-600 duration-300 active:bg-green-500 hover:text-white"
          >
            保存草稿
          </div>
        </div>
      </div>
      <GreenBtn
        @click="publishPost"
        classes="w-full mt-2"
        value="发布文章"
        :icon="['fas', 'check']"
      />
    </form>
    <PublishCategory
      ref="publishCategory"
      :category-list="categoryList"
      @chooseCategory="chooseCategory"
    />
  </UserCenterLayout>
</template>

<style lang="less" scoped></style>

<route lang="yaml">
meta:
  auth: true
</route>
