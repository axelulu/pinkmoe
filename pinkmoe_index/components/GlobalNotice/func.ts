/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-31 12:08:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 10:27:22
 * @FilePath: /pinkmoe_vitesse/src/components/GlobalNotice/func.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
/**
 * 创建通知插件的调用方法
 */
import type { ComponentPublicInstance } from 'vue'
import { createApp } from 'vue'
import Notification from './index.vue'

export interface OptionsType {
  content: string
  duration?: number
  btn: string
}

interface InstanceType {
  id: string
  vm: ComponentPublicInstance<any>
}

const instances: InstanceType[] = []
let seed = 1

const removeInstance = (id: string, removeHeight: number): void => {
  const index = instances.findIndex(item => item.id === id)
  const len = instances.length

  // 删除 instance
  instances.splice(index, 1)

  if (len < 1)
    return

  for (let i = index; i < len - 1; i++) {
    const inst = instances[i].vm
    inst.bottomOffset = inst.bottomOffset - removeHeight - 16
  }
}

const notify = (options: OptionsType): void => {
  const id = `notification_${seed++}`
  const container = document.createElement('div')
  document.body.appendChild(container)

  let verticalOffset = 30
  instances.forEach((item) => {
    verticalOffset += item.vm.$el.offsetHeight + 16
  })

  const instance = createApp({
    data() {
      return {
        bottomOffset: verticalOffset,
      }
    },
    methods: {
      closedFunc(height: number): void {
        removeInstance(id, height)
        document.body.removeChild(container)
        instance.unmount()
      },
    },
    render() {
      return h(
        Notification,
        {
          verticalOffset: this.bottomOffset,
          onClosed: this.closedFunc,
          ...options,
        },
        {},
      )
    },
  })

  instances.push({
    id,
    vm: instance.mount(container),
  })
}

export default notify
