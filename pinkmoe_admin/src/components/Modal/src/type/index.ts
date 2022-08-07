/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:16:12
 * @FilePath: /pinkmoe_admin/src/components/Modal/src/type/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import type { DialogOptions } from 'naive-ui/lib/dialog';
/**
 * @description: 弹窗对外暴露的方法
 */
export interface ModalMethods {
  setProps: (props) => void;
  openModal: () => void;
  closeModal: () => void;
  setSubLoading: (status) => void;
}

/**
 * 支持修改，DialogOptions 參數
 */
export type ModalProps = DialogOptions;

export type RegisterFn = (ModalInstance: ModalMethods) => void;

export type UseModalReturnType = [RegisterFn, ModalMethods];
