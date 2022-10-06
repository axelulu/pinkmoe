/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:16:15
 * @FilePath: /pinkmoe_admin/src/components/Modal/src/hooks/useModal.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { ref, onUnmounted, unref, getCurrentInstance, watch, nextTick } from 'vue';
import { isProdMode } from '@/utils/env';
import { ModalMethods, UseModalReturnType } from '../type';
import { getDynamicProps } from '@/utils';
import { tryOnUnmounted } from '@vueuse/core';
export function useModal(props): UseModalReturnType {
  const modalRef = ref<Nullable<ModalMethods>>(null);
  const currentInstance = getCurrentInstance();

  const getInstance = () => {
    const instance = unref(modalRef.value);
    if (!instance) {
      console.error('useModal instance is undefined!');
    }
    return instance;
  };

  const register = (modalInstance: ModalMethods) => {
    isProdMode() &&
      tryOnUnmounted(() => {
        modalRef.value = null;
      });
    modalRef.value = modalInstance;
    currentInstance?.emit('register', modalInstance);

    watch(
      () => props,
      () => {
        props && modalInstance.setProps(getDynamicProps(props));
      },
      {
        immediate: true,
        deep: true,
      }
    );
  };

  const methods: ModalMethods = {
    setProps: (props): void => {
      getInstance()?.setProps(props);
    },
    openModal: () => {
      getInstance()?.openModal();
    },
    closeModal: () => {
      getInstance()?.closeModal();
    },
    setSubLoading: (status) => {
      getInstance()?.setSubLoading(status);
    },
  };

  return [register, methods];
}
