<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:10
 * @FilePath: /pinkmoe_admin/src/views/form/detail/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
-->
<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="表单详情">
        表单除了提交数据，有时也用于显示只读信息。
      </n-card>
    </div>
    <n-card
      :bordered="false"
      title="基本信息"
      class="mt-4 proCard"
      size="small"
      :segmented="{ content: 'hard' }"
    >
      <n-descriptions label-placement="left" class="py-2">
        <n-descriptions-item>
          <template #label>收款人姓名</template>
          啊俊
        </n-descriptions-item>
        <n-descriptions-item label="收款账户">NaiveUiAdmin@qq.com</n-descriptions-item>
        <n-descriptions-item label="付款类型">支付宝</n-descriptions-item>
        <n-descriptions-item label="付款账户">NaiveUiAdmin@163.com</n-descriptions-item>
        <n-descriptions-item label="转账金额">￥1980.00</n-descriptions-item>
        <n-descriptions-item label="状态">
          <n-tag type="success"> 已到账</n-tag>
        </n-descriptions-item>
      </n-descriptions>
    </n-card>
    <n-card
      :bordered="false"
      title="其它信息"
      class="mt-4 proCard"
      size="small"
      :segmented="{ content: 'hard' }"
    >
      <n-descriptions label-placement="left" class="py-2">
        <n-descriptions-item>
          <template #label>城市</template>
          深圳
        </n-descriptions-item>
        <n-descriptions-item label="性别">男</n-descriptions-item>
        <n-descriptions-item label="邮箱">NaiveUiAdmin@qq.com</n-descriptions-item>
        <n-descriptions-item label="地址">广东省深圳市南山区</n-descriptions-item>
        <n-descriptions-item label="生日">1991-06-04</n-descriptions-item>
        <n-descriptions-item label="认证">
          <n-tag type="success"> 已认证</n-tag>
        </n-descriptions-item>
      </n-descriptions>
    </n-card>
    <n-card
      :bordered="false"
      title="表格信息"
      class="mt-4 proCard"
      size="small"
      :segmented="{ content: 'hard' }"
    >
      <n-table :bordered="false" :single-line="false">
        <thead>
          <tr>
            <th>姓名</th>
            <th>性别</th>
            <th>城市</th>
            <th>生日</th>
            <th width="150">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Ah jung</td>
            <td>男</td>
            <td>深圳</td>
            <td>1993-11-09</td>
            <td>
              <n-space>
                <n-button size="small" type="error">删除</n-button>
                <n-button size="small" type="info">查看</n-button>
              </n-space>
            </td>
          </tr>
          <tr>
            <td>西门飞雪</td>
            <td>男</td>
            <td>广州</td>
            <td>1991-09-11</td>
            <td>
              <n-space>
                <n-button size="small" type="error">删除</n-button>
                <n-button size="small" type="info">查看</n-button>
              </n-space>
            </td>
          </tr>
          <tr>
            <td>泰坦巨人</td>
            <td>男</td>
            <td>北京</td>
            <td>1990-11-03</td>
            <td>
              <n-space>
                <n-button size="small" type="error">删除</n-button>
                <n-button size="small" type="info">查看</n-button>
              </n-space>
            </td>
          </tr>
          <tr>
            <td>猎魔人</td>
            <td>女</td>
            <td>上海</td>
            <td>1992-03-11</td>
            <td>
              <n-space>
                <n-button size="small" type="error">删除</n-button>
                <n-button size="small" type="info">查看</n-button>
              </n-space>
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-card>
  </div>
</template>

<script setup></script>

<style lang="less" scoped></style>
