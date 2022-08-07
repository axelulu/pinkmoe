<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:48
 * @FilePath: /pinkmoe_index/src/components/SvgIcon/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <svg aria-hidden="true" class="svg-icon-spin" :class="calsses">
    <use :xlink:href="symbolId" :fill="color" />
  </svg>
</template>

<script lang="ts" setup>
  const props = defineProps({
    prefix: {
      type: String,
      default: 'icon',
    },
    name: {
      type: String,
      required: true,
    },
    color: {
      type: String,
      default: '#333',
    },
    size: {
      type: String,
      default: 'default',
    },
  });
  const symbolId = computed(() => `#${props.prefix}-${props.name}`);
  const calsses = computed(() => {
    return {
      [`sdms-size-${props.size}`]: props.size,
    };
  });
  const fontSize = reactive({ default: '32px', small: '20px', large: '48px' });
</script>
<style lang="less" scoped>
  .svg-icon-spin {
    width: v-bind('fontSize.default');
    height: v-bind('fontSize.default');
    fill: v-bind(color);
    vertical-align: middle;
    color: v-bind(color);

    &.sdms-size-small {
      font-size: v-bind('fontSize.small');
      height: v-bind('fontSize.small');
    }

    &.sdms-size-large {
      font-size: v-bind('fontSize.large');
      height: v-bind('fontSize.large');
    }
  }
</style>
