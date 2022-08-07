/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:20:27
 * @FilePath: /pinkmoe_admin/src/router/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import type { RouteRecordRaw, RouteMeta } from 'vue-router';
import { defineComponent } from 'vue';

export type Component<T extends any = any> =
  | ReturnType<typeof defineComponent>
  | (() => Promise<typeof import('*.vue')>)
  | (() => Promise<T>);

export interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta'> {
  name: string;
  meta: RouteMeta;
  component?: Component | string;
  components?: Component;
  children?: AppRouteRecordRaw[];
  props?: Recordable;
  fullPath?: string;
}

export interface Meta {
  // 名称
  title: string;
  // 是否忽略权限
  ignoreAuth?: boolean;
  permissions?: string[];
  // 是否不缓存
  noKeepAlive?: boolean;
  // 是否固定在tab上
  affix?: boolean;
  // tab上的图标
  icon?: string;
  // 跳转地址
  frameSrc?: string;
  // 外链跳转地址
  externalLink?: string;
  //隐藏
  hidden?: boolean;
}

export interface Menu {
  title: string;
  label: string;
  key: string;
  meta: RouteMeta;
  name: string;
  component?: Component | string;
  components?: Component;
  children?: AppRouteRecordRaw[];
  props?: Recordable;
  fullPath?: string;
  icon?: any;
  path: string;
  permissions?: string[];
  redirect?: string;
  sort?: number;
}
