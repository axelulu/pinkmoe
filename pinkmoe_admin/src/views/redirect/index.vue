<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:51
 * @FilePath: /pinkmoe_admin/src/views/redirect/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="tsx">
  import { defineComponent, onBeforeMount } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { NEmpty } from 'naive-ui';

  export default defineComponent({
    name: 'Redirect',
    setup() {
      const route = useRoute();
      const router = useRouter();
      onBeforeMount(() => {
        const { params, query } = route;
        const { path } = params;
        router.replace({
          path: '/' + (Array.isArray(path) ? path.join('/') : path),
          query,
        });
      });
      return () => <NEmpty />;
    },
  });
</script>
