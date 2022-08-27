<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 21:22:44
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:24:14
 * @FilePath: /pinkmoe_index/components/Commentfont/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="Commentfont">
import Dialog from '/@/components/Dialog/index.vue'
defineProps({
  postId: {
    type: String,
    default: 'login',
  },
})
const emit = defineEmits(['insertFont'])
const dialog = ref()
const font = ref([
  '(⊙⊙！) ',
  'ƪ(‾ε‾“)ʃƪ(',
  'Σ(°Д°;',
  '눈_눈',
  '(๑>◡<๑) ',
  '(❁´▽`❁)',
  '(,,Ծ▽Ծ,,)',
  '（⺻▽⺻ ）',
  '乁( ◔ ౪◔)「',
  'ლ(^o^ლ)',
  '(◕ܫ◕)',
  '❀.(*´▽`*)❀.',
  '(´･ω･`)',
  'ヽ(ﾟ∀ﾟ)ﾉ',
  '(ฅ>ω<*ฅ)',
  'ヽ(*≧ω≦)ﾉ',
  '( ≧Д≦)',
  '(Ｔ▽Ｔ)',
  '(ノへ￣、) ',
  '・ﾟ・(ノД`)・ﾟ・',
  '(,, ･∀･)ﾉ゛',
  '(o´ω`o)ﾉ',
  '(￣^￣)ゞ',
  '⁄(⁄ ⁄•⁄ω⁄•⁄ ⁄)⁄',
  '（＊／ω＼＊ )',
  '(´∇ﾉ｀*)ノ',
  '(〃ﾉωﾉ)',
  '(╯°□°）╯︵ ┻━┻',
  '(=ﾟДﾟ=) ▄︻┻┳━ ·.`.`.`. ',
  '╰(￣ω￣ｏ) [摸摸头]',
  '( ｰ̀дｰ́ )ง ',
  '(๑•̀ㅂ•́)و✧',
  '_(:з」∠)_',
  '╮(╯▽╰)╭',
  '（ﾉ´д｀）',
])
const showAnimate = ref<boolean>(false)

function show() {
  dialog.value.show()
}

function insertFont(item) {
  emit('insertFont', item)
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
            v-for="(item, index) in font"
            :key="index"
            class="px-2 py-1 bg-gray-200 dark:bg-gray-600 dark:hover:bg-gray-700 hover:bg-gray-100 duration-300 select-none text-xs m-0.5 flex-grow cursor-pointer"
            @click="insertFont(item)"
          >
            {{ item }}
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
