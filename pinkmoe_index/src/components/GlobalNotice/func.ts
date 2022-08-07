/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-31 12:08:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:10:09
 * @FilePath: /pinkmoe_index/src/components/GlobalNotice/func.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
/**
 * 创建通知插件的调用方法
 */
import { createApp, ComponentPublicInstance } from 'vue';
import Notification from './index.vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

export type OptionsType = {
  content: string;
  duration?: number;
  btn: string;
};

type InstanceType = {
  id: string;
  vm: ComponentPublicInstance<any>;
};

const instances: InstanceType[] = [];
let seed = 1;

const removeInstance = (id: string, removeHeight: number): void => {
  const index = instances.findIndex((item) => item.id === id);
  const len = instances.length;

  // 删除 instance
  instances.splice(index, 1);

  if (len < 1) return;

  for (let i = index; i < len - 1; i++) {
    const inst = instances[i].vm;
    inst.bottomOffset = inst.bottomOffset - removeHeight - 16;
  }
};

const notify = (options: OptionsType): void => {
  const id = `notification_${seed++}`;
  const container = document.createElement('div');
  document.body.appendChild(container);

  let verticalOffset = 30;
  instances.forEach((item) => {
    verticalOffset += item.vm.$el.offsetHeight + 16;
  });

  const instance = createApp({
    data() {
      return {
        bottomOffset: verticalOffset,
      };
    },
    methods: {
      closedFunc(height: number): void {
        removeInstance(id, height);
        document.body.removeChild(container);
        instance.unmount();
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
      );
    },
  });

  instance.component('FontAwesomeIcon', FontAwesomeIcon); // 渲染到创建的div节点上
  instances.push({
    id,
    vm: instance.mount(container),
  });
};

export default notify;
