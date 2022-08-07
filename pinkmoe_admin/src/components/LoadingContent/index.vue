<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:15:46
 * @FilePath: /pinkmoe_admin/src/components/LoadingContent/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template></template>
<script lang="ts">
  import { useLoadingBar } from 'naive-ui';

  export default {
    name: 'LoadingContent',
    setup() {
      //挂载在 window 方便与在js中使用
      window['$loading'] = useLoadingBar();
    },
  };
</script>
