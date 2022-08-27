<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-27 22:00:31
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:46
 * @FilePath: /pinkmoe_index/components/Postmusic/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Postmusic">
import { usePostMusic } from '/@/hooks/postMusic'
const props = defineProps({
  postId: {
    type: String,
    default: '',
  },
  nickName: {
    type: String,
    default: '',
  },
})
const { postMusic, buyMusic, musicList } = usePostMusic(props)
</script>

<template>
  <aplayer v-if="musicList" autoplay :music="postMusic[0]" :list="musicList" />
</template>
