<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 17:59:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:53
 * @FilePath: /pinkmoe_index/components/Publishrecategory/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Publishrecategory">
import PublishReCategory from '/@/components/Publishrecategory/index.vue'
defineProps({
  item: {
    type: Object,
    default: null,
  },
  currentCategory: {
    type: Object,
    default: null,
  },
})
const emit = defineEmits(['emitCategoryPost', 'chooseCategory'])
const show = ref(false)
const searchReCategory = ref()
const childrenCategory = ref<any>(null)
function getChildCategory(item) {
  show.value = true
  childrenCategory.value = item
  if (searchReCategory.value)
    searchReCategory.value.hide()

  emit('emitCategoryPost', item)
}

function emitCategoryPost(item) {
  emit('emitCategoryPost', item)
}

function chooseCategory(item) {
  emit('chooseCategory', item)
}

function hide() {
  show.value = false
}

defineExpose({
  hide,
})
</script>

<template>
  <div v-if="item.children" class="flex flex-row bg-gray-200 dark:bg-gray-800">
    <div>
      <div
        v-for="(item, index) in item.children"
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
      v-if="show"
      ref="searchReCategory"
      :current-category="currentCategory"
      :item="childrenCategory"
      @emitCategoryPost="emitCategoryPost"
      @chooseCategory="chooseCategory"
    />
  </div>
</template>
