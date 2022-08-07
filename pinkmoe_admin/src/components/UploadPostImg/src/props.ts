/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-29 15:17:22
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:17:47
 * @FilePath: /pinkmoe_admin/src/components/UploadPostImg/src/props.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import type { PropType } from 'vue';
import { NUpload } from 'naive-ui';

export const basicProps = {
  ...NUpload.props,
  accept: {
    type: String,
    default: '.jpg,.png,.jpeg,.svg,.gif',
  },
  helpText: {
    type: String as PropType<string>,
    default: '',
  },
  maxSize: {
    type: Number as PropType<number>,
    default: 2,
  },
  maxNumber: {
    type: Number as PropType<number>,
    default: Infinity,
  },
  value: {
    type: Array as PropType<string[]>,
    default: () => [],
  },
  cover: {
    type: String as PropType<string>,
    default: () => "",
  },
  width: {
    type: Number as PropType<number>,
    default: 104,
  },
  height: {
    type: Number as PropType<number>,
    default: 104, //建议不小于这个尺寸 太小页面可能显示有异常
  },
};
