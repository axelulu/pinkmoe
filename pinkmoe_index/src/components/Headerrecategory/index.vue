<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 10:59:57
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:10:31
 * @FilePath: /pinkmoe_index/src/components/Headerrecategory/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Headerrecategory">
  import HeaderReCategory from '/@/components/Headerrecategory/index.vue';
  defineProps({
    item: {
      type: Array,
      default: null,
    },
  });
  const isShow = ref(false);
  const children = ref();
  const headerReCategory = ref();
  const emit = defineEmits(['emitCategoryPost']);
  const mouseenter = (item) => {
    isShow.value = true;
    children.value = item;
    if (headerReCategory.value) {
      headerReCategory.value.hide();
    }
    emit('emitCategoryPost', item);
  };

  function hide() {
    isShow.value = false;
  }

  defineExpose({
    hide,
  });
</script>

<template>
  <!-- Headerrecategory -->
  <ul
    v-if="item"
    class="flex-col dark:bg-gray-600 childrenCategory pt-3 pb-3 bg-white absolute left-32 -top-3 shadow-xl animate-fadeIn30 hover:flex"
  >
    <div class="relative">
      <li
        v-for="(i, index) in item.children"
        @mouseenter="mouseenter(i)"
        :key="index"
        class="flex flex-row text-xs"
      >
        <router-link
          :to="'/category/' + i.slug"
          class="px-4 py-2 w-32 hover:bg-pink-400 hover:text-white cursor-pointer duration-300"
        >
          <font-awesome-icon :icon="[i.iconType, i.icon]" />
          <span class="ml-1">{{ i.name }}</span>
        </router-link>
      </li>
      <HeaderReCategory ref="headerReCategory" v-if="isShow" :item="children" />
    </div>
  </ul>
</template>

<style lang="less" scoped></style>
