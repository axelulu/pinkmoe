<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 21:22:55
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 23:18:17
 * @FilePath: /pinkmoe_vitesse/src/components/Commentsmile/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Commentsmile">
import Dialog from '/@/components/Dialog/index.vue'
defineProps({
  postId: {
    type: String,
    default: 'login',
  },
})
const emit = defineEmits(['insertSmile'])
const dialog = ref()
const smile = ref([
  {
    name: 'bcr_no',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/no.webp',
  },
  {
    name: 'bcr_ok',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/ok.webp',
  },
  {
    name: 'bcr_一切都会好起来的',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/一切都会好起来的.webp',
  },
  {
    name: 'bcr_了解',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/了解.webp',
  },
  {
    name: 'bcr_交给我吧',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/交给我吧.webp',
  },
  {
    name: 'bcr_做得好',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/做得好.webp',
  },
  {
    name: 'bcr_加油',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/加油.webp',
  },
  {
    name: 'bcr_参战',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/参战.webp',
  },
  {
    name: 'bcr_吸溜',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/吸溜.webp',
  },
  {
    name: 'bcr_呜呼呼',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/呜呼呼.webp',
  },
  {
    name: 'bcr_哇哇',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/哇哇.webp',
  },
  {
    name: 'bcr_哈罗',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/哈罗.webp',
  },
  {
    name: 'bcr_喵喵喵',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/喵喵喵.webp',
  },
  {
    name: 'bcr_在那里',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/在那里.webp',
  },
  {
    name: 'bcr_已经举报了',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/已经举报了.webp',
  },
  {
    name: 'bcr_干杯',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/干杯.webp',
  },
  {
    name: 'bcr_快回复吧',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/快回复吧.webp',
  },
  {
    name: 'bcr_抱歉啦',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/抱歉啦.webp',
  },
  {
    name: 'bcr_接下来拜托了',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/接下来拜托了.webp',
  },
  {
    name: 'bcr_早上好',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/早上好.webp',
  },
  {
    name: 'bcr_晚上好',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/晚上好.webp',
  },
  {
    name: 'bcr_晚安zzz',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/晚安zzz.webp',
  },
  {
    name: 'bcr_来吧',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/来吧.webp',
  },
  {
    name: 'bcr_气呼呼',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/气呼呼.webp',
  },
  {
    name: 'bcr_真是难办啊',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/真是难办啊.webp',
  },
  {
    name: 'bcr_糟糕',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/糟糕.webp',
  },
  {
    name: 'bcr_给我等一下',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/给我等一下.webp',
  },
  {
    name: 'bcr_给我记住了',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/给我记住了.webp',
  },
  {
    name: 'bcr_要上了哦',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/要上了哦.webp',
  },
  {
    name: 'bcr_谢谢',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/谢谢.webp',
  },
  {
    name: 'bcr_跟随妾身而来吧',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/跟随妾身而来吧.webp',
  },
  {
    name: 'bcr_颤颤抖抖',
    url: 'https://cdn.inn-studio.com/themes/zero/images/bcr/颤颤抖抖.webp',
  },
  {
    name: '滑稽',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1f0wwjnedhoj200u00ugld.jpg',
  },
  {
    name: '感谢分享',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5ir4rznj202s02eglh.jpg',
  },
  {
    name: '脸红',
    url: 'https://sinaimg.inn-studio.com/large/686ee05djw1eu8ijxc3p7g201c01c3yd.gif',
  },
  {
    name: '杯具',
    url: 'https://sinaimg.inn-studio.com/large/686ee05djw1eu8ikpw34jg201e01emx1.gif',
  },
  {
    name: '亚历山大',
    url: 'https://sinaimg.inn-studio.com/large/686ee05djw1eu8iliwosmg201e01e74h.gif',
  },
  {
    name: '想要',
    url: 'https://sinaimg.inn-studio.com/large/686ee05djw1eu8ilzci2jg202s02sglo.gif',
  },
  {
    name: '吃惊',
    url: 'https://sinaimg.inn-studio.com/large/686ee05djw1eu8j1vay4ej204h049jrb.jpg',
  },
  {
    name: '好样的',
    url: 'https://sinaimg.inn-studio.com/large/686ee05djw1eu8iomh5cbg203g03cdgx.gif',
  },
  {
    name: '感谢分享2',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1f44hq6g2ftj202s02swef.jpg',
  },
  {
    name: '生气',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fzbvagqj202s02sa9z.jpg',
  },
  {
    name: '卖萌',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8g1m0y49j202s02sjrb.jpg',
  },
  {
    name: 'OK！',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fza5ep2j202s02s3yg.jpg',
  },
  {
    name: '不行',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fz69u1qj202s02st8n.jpg',
  },
  {
    name: '叹气',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5jxgin9j202s02e743.jpg',
  },
  {
    name: '棒！',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fz1nm1rj202s02smx3.jpg',
  },
  {
    name: '偷笑',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1f124864q4qj202s02e3ye.jpg',
  },
  {
    name: '噫',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1f0z83hc4ndj202s02ea9x.jpg',
  },
  {
    name: 'good',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1f0z83fr4l7j202s02emx2.jpg',
  },
  {
    name: '不明真相',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5fo2790j202s02eweb.jpg',
  },
  {
    name: '太棒了',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5lerczej202s02edfn.jpg',
  },
  {
    name: '不知所措',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5qzpg9aj202s02eq2u.jpg',
  },
  {
    name: '盯..',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5pi74hgj202s02et8i.jpg',
  },
  {
    name: '所以呢',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1f0z838j5d4j202s02emx1.jpg',
  },
  {
    name: '残念',
    url: 'https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5hkglbsj202s02e3yf.jpg',
  },
])
const showAnimate = ref<boolean>(false)

function show() {
  dialog.value.show()
}

function insertSmile(item) {
  emit('insertSmile', `[${item.name}]`)
  dialog.value.hide()
}

defineExpose({
  show,
})
</script>

<template>
  <Dialog
    ref="dialog"
    :hide-fun="() => (showAnimate = false)"
    :show-fun="() => (showAnimate = true)"
  >
    <div
      :class="showAnimate ? 'animate-pinkZoomInUp30' : 'animate-pinkZoomOutDown30'"
      class="fixed top-0 bottom-0 left-0 ring-0 flex justify-center items-center h-full w-full pointer-events-none"
    >
      <div
        class="lg:w-5/12 xl:w-5/12 bg-gray-50 dark:bg-gray-700 pointer-events-auto flex flex-col rounded-md overflow-hidden"
      >
        <div class="flex flex-wrap space-between items-center text-center px-3 pt-3 pb-1">
          <span
            v-for="(item, index) in smile"
            :key="index"
            class="px-2 py-1 animate-lazyloaded bg-gray-200 dark:bg-gray-600 dark:hover:bg-gray-700 hover:bg-gray-100 duration-300 select-none text-xs m-0.5 flex-grow cursor-pointer"
            @click="insertSmile(item)"
          >
            <img v-lazy="item.url" class="w-12 h-12 object-cover" alt="登陆">
          </span>
        </div>
        <div
          class="bg-gray-200 flex justify-center items-center py-2 dark:bg-gray-600 dark:hover:bg-gray-700 text-center py-1 cursor-pointer hover:bg-gray-100 duration-300"
          @click="dialog.hide()"
        >
          <i class="inline-block i-carbon:close-filled" />
        </div>
      </div>
    </div>
  </Dialog>
</template>
