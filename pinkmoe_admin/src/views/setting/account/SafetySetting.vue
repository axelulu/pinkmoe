<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:25:07
 * @FilePath: /pinkmoe_admin/src/views/setting/account/SafetySetting.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <n-grid cols="1" responsive="screen" class="-mt-5">
    <n-grid-item>
      <n-list>
        <n-list-item>
          <template #suffix>
            <n-button type="primary" text>修改</n-button>
          </template>
          <n-thing title="账户密码">
            <template #description
              ><span class="text-gray-400">绑定手机和邮箱，并设置密码，帐号更安全</span></template
            >
          </n-thing>
        </n-list-item>
        <n-list-item>
          <template #suffix>
            <n-button type="primary" text>修改</n-button>
          </template>
          <n-thing title="绑定手机">
            <template #description
              ><span class="text-gray-400">已绑定手机号：+86189****4877</span></template
            >
          </n-thing>
        </n-list-item>
        <n-list-item>
          <template #suffix>
            <n-button type="primary" text>设置</n-button>
          </template>
          <n-thing title="密保问题">
            <template #description
              ><span class="text-gray-400"
                >未设置密保问题，密保问题可有效保护账户安全</span
              ></template
            >
          </n-thing>
        </n-list-item>
        <n-list-item>
          <template #suffix>
            <n-button type="primary" text>修改</n-button>
          </template>
          <n-thing title="个性域名">
            <template #description
              ><span class="text-gray-400">已绑定域名：https://www.naiveui.com</span></template
            >
          </n-thing>
        </n-list-item>
      </n-list>
    </n-grid-item>
  </n-grid>
</template>

<script lang="ts" setup></script>
