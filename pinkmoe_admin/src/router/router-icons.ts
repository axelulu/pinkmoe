/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-01 22:37:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:20:25
 * @FilePath: /pinkmoe_admin/src/router/router-icons.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { renderIcon } from '@/utils/index';
import { DashboardOutlined } from "@vicons/antd";
import {
  OptionsSharp,
  ColorPaletteOutline,
  ReaderOutline,
  MegaphoneOutline,
  LogoAndroid,
  BagHandleOutline,
  ChatbubblesOutline,
  LogInOutline,
  ExtensionPuzzleOutline,
  CogSharp,
  DesktopOutline,
  People,
} from "@vicons/ionicons5";

//前端路由图标映射表
export const constantRouterIcon = {
  DashboardOutlined: renderIcon(DashboardOutlined),
  ReaderOutline: renderIcon(ReaderOutline),
  OptionsSharp: renderIcon(OptionsSharp),
  ColorPaletteOutline: renderIcon(ColorPaletteOutline),
  MegaphoneOutline: renderIcon(MegaphoneOutline),
  BagHandleOutline: renderIcon(BagHandleOutline),
  ChatbubblesOutline: renderIcon(ChatbubblesOutline),
  LogInOutline: renderIcon(LogInOutline),
  ExtensionPuzzleOutline: renderIcon(ExtensionPuzzleOutline),
  CogSharp: renderIcon(CogSharp),
  DesktopOutline: renderIcon(DesktopOutline),
  People: renderIcon(People),
  LogoAndroid: renderIcon(LogoAndroid),
};
