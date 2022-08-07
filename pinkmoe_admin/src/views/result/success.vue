<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:58
 * @FilePath: /pinkmoe_admin/src/views/result/success.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <n-card :bordered="false" class="proCard">
    <div class="result-box">
      <n-result
        status="success"
        title="操作成功"
        description="提交结果页用于反馈一系列操作任务的处理结果，如果仅是简单操作，灰色区域可以显示一些补充的信息。"
      >
        <div class="result-box-extra">
          <p>已提交申请，等待财务部门审核。</p>
        </div>
        <template #footer>
          <div class="flex justify-center mb-4">
            <n-space align="center">
              <n-button type="info" @click="goHome">回到首页</n-button>
              <n-button>查看详情</n-button>
              <n-button>打印</n-button>
            </n-space>
          </div>
        </template>
      </n-result>
    </div>
  </n-card>
</template>
<script lang="ts" setup>
  import { computed } from 'vue';
  import { useRouter } from 'vue-router';
  import { useThemeVars } from 'naive-ui';

  const router = useRouter();
  const themeVars = useThemeVars();

  const getTableHeaderColor = computed(() => {
    return themeVars.value.tableHeaderColor;
  });

  function goHome() {
    router.push('/');
  }
</script>
<style lang="less" scoped>
  .result-box {
    width: 72%;
    margin: 0 auto;
    text-align: center;
    padding-top: 5px;

    &-extra {
      padding: 24px 40px;
      text-align: left;
      background: v-bind(getTableHeaderColor);
      border-radius: 4px;
    }
  }
</style>
