<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-31 06:52:39
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 19:56:32
 * @FilePath: /pinkmoe_index/components/GlobalNotice/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts">
import {
  computed,
  defineComponent,
  onBeforeUnmount,
  onMounted,
  reactive,
  ref,
  toRefs,
} from 'vue'

export default defineComponent({
  name: 'Notification',
  props: {
    content: {
      type: Object,
      required: true,
    },
    duration: {
      type: Number,
      default: 5000,
    },
    btn: {
      type: String,
      default: '关闭',
    },
    type: {
      type: String,
      default: 'chat',
    },
    router: {
      type: Object,
      default: null,
    },
    verticalOffset: {
      type: Number,
      default: 0,
    },
    onClosed: {
      type: Function,
    },
  },
  setup(props) {
    const root = ref(null!)

    const state = reactive({
      height: 30,
      visible: false,
    })

    const styleObj = computed(() => ({
      position: 'fixed',
      right: '20px',
      zIndex: '10',
      top: `${props.verticalOffset}px`,
    }))

    const timer = ref(0)

    const handleClose = (e: MouseEvent): void => {
      e.preventDefault()
      state.visible = false
    }

    const afterLeave = () => {
      (props as any).onClosed(state.height)
    }

    const afterEnter = () => {
      state.height = (root as any).value.offsetHeight
    }

    const createTimer = () => {
      if (props.duration) {
        timer.value = setTimeout(() => {
          state.visible = false
        }, props.duration) as unknown as number
      }
    }

    const clearTimer = () => {
      if (timer.value) {
        clearTimeout(timer.value)
        timer.value = 0
      }
    }

    const jump = () => {
      props.router.push({
        path: '/user-center/im',
        params: { sendId: props.content.sendIdRelation?.uuid },
      })
    }

    onMounted(() => {
      state.visible = true
      createTimer()
    })

    onBeforeUnmount(() => {
      clearTimer()
    })

    // toRefs 把reactive创建出的数据变更为响应式
    return {
      ...toRefs(state),
      root,
      styleObj,
      jump,
      handleClose,
      afterLeave,
      afterEnter,
      clearTimer,
      createTimer,
    }
  },
})
</script>

<template>
  <transition name="fade" @after-leave="afterLeave" @after-enter="afterEnter">
    <div
      v-show="visible"
      ref="root"
      :style="styleObj"
      @mouseenter="clearTimer"
      @mouseleave="createTimer"
    >
      <div
        v-if="type === 'chat'"
        to="/user-center/im"
        class="select-none relative cursor-pointer mt-2 text-xs z-10 bg-pink-400 duration-300 hover:bg-pink-300 px-8 py-4 rounded-md text-white shadow-md"
        @click="jump"
      >
        <span class="content">来自用户
          <span class="bg-yellow-500 rounded-md text-white px-1 py-0.5 rounded-md">{{
            content?.userIdRelation?.nickName
          }}</span>
          的消息</span>
        <i
          class="i-fa6-solid:xmark inline-block absolute -top-1 -right-1 dark:bg-gray-700 dark:hover:bg-gray-900 dark:hover:text-gray-100 dark:text-gray-300 bg-gray-300 hover:bg-gray-100 hover:text-gray-800 text-gray-700 duration-300 rounded-full w-2 h-2 p-1"
          @click="handleClose"
        />
      </div>
    </div>
  </transition>
</template>

<style scoped></style>
