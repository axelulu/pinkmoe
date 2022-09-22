<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 17:43:49
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 23:37:07
 * @FilePath: /pinkmoe_index/components/Publishcategory/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Publishcategory">
import PublishReCategory from '/@/components/Publishrecategory/index.vue'

defineProps({
  categoryList: {
    type: Array,
    default: null,
  },
})
const emit = defineEmits(['chooseCategory'])
const dialog = ref()
const searchReCategory = ref()
const isShow = ref<boolean>(false)
const showAnimate = ref<boolean>(false)
const children = ref<any>(null)
const currentCategory = ref<Object>({
  slug: '0',
  type: 'all',
})
const show = () => {
  dialog.value.show()
}
const hide = () => {
  dialog.value.hide()
}
const chooseCategory = (item) => {
  emit('chooseCategory', item)
}

const getCategoryPost = (item) => {
  currentCategory.value = item
}

const getChildCategory = (item) => {
  isShow.value = true
  children.value = item
  currentCategory.value = item
  if (searchReCategory.value)
    searchReCategory.value.hide()
}

defineExpose({
  show,
  hide,
})
</script>

<template>
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
            <i class="inline-block i-codicon:three-bars" />
            <span class="ml-1">选择分类</span>
          </div>
          <div
            class="flex flex-row justify-center items-center text-red-500 dark:text-gray-200 text-xs ml-4 p-0 rounded-bl-md rounded-br-md overflow-hidden"
          >
            <span class="ml-1">单击查看子分类，双击选择分类</span>
          </div>
          <div
            class="text-sm px-4 py-2 flex justify-center items-center bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-800 cursor-pointer hover:text-white hover:bg-pink-400 duration-300 rounded-bl-md"
            @click="dialog.hide()"
          >
            <i class="inline-block i-clarity:times-circle-solid" />
          </div>
        </div>
        <div v-if="categoryList" class="flex flex-row py-6 px-4 overflow-x-auto">
          <div class="bg-gray-200 dark:bg-gray-800">
            <div
              v-for="(item, index) in categoryList.list"
              :key="index"
              :class="
                currentCategory.slug === item.slug && currentCategory.type !== 'all'
                  ? 'text-white bg-pink-400 dark:bg-pink-400'
                  : 'text-gray-500'
              "
              class="text-xs select-none text-center w-26 py-2 dark:bg-gray-800 dark:hover:bg-pink-400 hover:bg-pink-400 dark:text-gray-200 hover:text-white px-2 py-1 duration-300 cursor-pointer"
              @click="getChildCategory(item)"
              @dblclick="chooseCategory(item)"
            >
              {{ item.name }}
              <i v-if="item.children" class="inline-block mr-1 i-fluent:caret-right-12-filled" />
            </div>
          </div>
          <PublishReCategory
            v-if="isShow"
            ref="searchReCategory"
            :current-category="currentCategory"
            :item="children"
            @emitCategoryPost="getCategoryPost"
            @chooseCategory="chooseCategory"
          />
        </div>
      </form>
    </div>
  </Dialog>
</template>
