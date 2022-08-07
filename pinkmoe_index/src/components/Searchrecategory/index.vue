<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 11:42:51
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:23
 * @FilePath: /pinkmoe_index/src/components/Searchrecategory/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Searchrecategory">
  import SearchReCategory from '/@/components/Searchrecategory/index.vue';
  defineProps({
    item: {
      type: Object,
      default: null,
    },
    currentCategory: {
      type: Object,
      default: null,
    },
  });
  const show = ref(false);
  const searchReCategory = ref();
  const childrenCategory = ref<any>(null);
  const emit = defineEmits(['emitCategoryPost']);
  function getChildCategory(item) {
    show.value = true;
    childrenCategory.value = item;
    if (searchReCategory.value) {
      searchReCategory.value.hide();
    }
    emit('emitCategoryPost', item);
  }

  function emitCategoryPost(item) {
    emit('emitCategoryPost', item);
  }

  function hide() {
    show.value = false;
  }

  defineExpose({
    hide,
  });
</script>

<template>
  <!-- Searchrecategory -->
  <div v-if="item.children">
    <div class="flex flex-row">
      <div
        @click="
          getChildCategory({
            slug: item.slug,
            type: 'all',
          })
        "
        :class="
          currentCategory.slug === item.slug && currentCategory.type === 'all'
            ? 'text-white bg-pink-400'
            : 'text-gray-500'
        "
        class="text-xs hover:bg-pink-400 dark:text-gray-200 hover:text-white px-2 py-1 duration-300 cursor-pointer"
      >
        全部
      </div>
      <div
        v-for="(item, index) in item.children"
        :key="index"
        @click="getChildCategory(item)"
        :class="
          currentCategory.slug === item.slug && currentCategory.type !== 'all'
            ? 'text-white bg-pink-400'
            : 'text-gray-500'
        "
        class="text-xs hover:bg-pink-400 dark:text-gray-200 hover:text-white px-2 py-1 duration-300 cursor-pointer"
      >
        {{ item.name }}
      </div>
    </div>
    <SearchReCategory
      ref="searchReCategory"
      @emitCategoryPost="emitCategoryPost"
      v-if="show"
      :current-category="currentCategory"
      :item="childrenCategory"
    />
  </div>
</template>

<style lang="less" scoped></style>
