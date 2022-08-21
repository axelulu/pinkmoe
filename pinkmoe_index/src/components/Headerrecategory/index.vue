<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 10:59:57
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 11:07:34
 * @FilePath: /pinkmoe_vitesse/src/components/Headerrecategory/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Headerrecategory">
import HeaderReCategory from '/@/components/Headerrecategory/index.vue'
defineProps({
  item: {
    type: Array,
    default: null,
  },
  index: {
    type: Number,
    default: 0,
  },
})
const emit = defineEmits(['emitCategoryPost'])
const isShow = ref(false)
const children = ref()
const childIndex = ref()
const headerReCategory = ref()
const mouseenter = (item, index) => {
  isShow.value = true
  children.value = item
  childIndex.value = index
  if (headerReCategory.value)
    headerReCategory.value.hide()

  emit('emitCategoryPost', item)
}

function hide() {
  isShow.value = false
}

defineExpose({
  hide,
})
</script>

<template>
  <!-- Headerrecategory -->
  <ul
    v-if="item"
    :style="index === 0 ? `top: -12px` : `top: ${32 * index - 12}px`"
    class="flex-col dark:bg-gray-600 bg-opacity-90 -top-3 childrenCategory pt-3 pb-3 bg-white absolute left-32 shadow-xl animate-fadeIn30 hover:flex"
  >
    <div class="relative">
      <li
        v-for="(i, index) in item.children"
        :key="index"
        class="flex flex-row text-xs"
        @mouseenter="mouseenter(i, index)"
      >
        <router-link
          :to="`/category/${i.slug}`"
          class="pl-4 py-2 w-32 hover:bg-pink-400 hover:text-white cursor-pointer duration-300"
        >
          <i :class="`${i.icon}`" />
          <span class="ml-1 mr-4">{{ i.name }}</span>
          <i v-if="i.children" class="i-fluent:caret-right-12-filled" />
        </router-link>
      </li>
      <HeaderReCategory v-if="isShow" ref="headerReCategory" :item="children" :index="childIndex" />
    </div>
  </ul>
</template>

<style lang="less" scoped></style>
