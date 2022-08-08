<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 17:59:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 10:12:09
 * @FilePath: /pinkmoe_index/src/components/Publishrecategory/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Publishrecategory">
  import PublishReCategory from '/@/components/Publishrecategory/index.vue';
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
  const emit = defineEmits(['emitCategoryPost', 'chooseCategory']);
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

  function chooseCategory(item) {
    emit('chooseCategory', item);
  }

  function hide() {
    show.value = false;
  }

  defineExpose({
    hide,
  });
</script>

<template>
  <!-- Publishrecategory -->
  <div class="flex flex-row bg-gray-200 dark:bg-gray-800" v-if="item.children">
    <div>
      <div
        v-for="(item, index) in item.children"
        :key="index"
        @click="getChildCategory(item)"
        @dblclick="chooseCategory(item)"
        :class="
          currentCategory.slug === item.slug && currentCategory.type !== 'all'
            ? 'text-white bg-pink-400 dark:bg-pink-400'
            : 'text-gray-500'
        "
        class="text-xs select-none text-center w-26 py-2 dark:bg-gray-800 dark:hover:bg-pink-400 hover:bg-pink-400 dark:text-gray-200 hover:text-white px-2 py-1 duration-300 cursor-pointer"
      >
        {{ item.name }}
        <font-awesome-icon v-if="item.children" class="mr-1" :icon="['fas', 'caret-right']" />
      </div>
    </div>
    <PublishReCategory
      ref="searchReCategory"
      @emitCategoryPost="emitCategoryPost"
      @chooseCategory="chooseCategory"
      v-if="show"
      :current-category="currentCategory"
      :item="childrenCategory"
    />
  </div>
</template>

<style lang="less" scoped></style>
