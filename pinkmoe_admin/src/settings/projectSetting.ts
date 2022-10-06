/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-29 20:07:43
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:20:40
 * @FilePath: /pinkmoe_admin/src/settings/projectSetting.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
const setting = {
  //导航模式 vertical 左侧菜单模式 horizontal 顶部菜单模式
  navMode: 'vertical',
  //导航风格 dark 暗色侧边栏 light 白色侧边栏 header-dark 暗色顶栏
  navTheme: 'dark',
  // 是否处于移动端模式
  isMobile: false,
  //顶部
  headerSetting: {
    //背景色
    bgColor: '#fff',
    //固定顶部
    fixed: true,
    //显示重载按钮
    isReload: true,
  },
  //页脚
  showFooter: true,
  //多标签
  multiTabsSetting: {
    //背景色
    bgColor: '#fff',
    //是否显示
    show: true,
    //固定多标签
    fixed: true,
  },
  //菜单
  menuSetting: {
    //最小宽度
    minMenuWidth: 64,
    //菜单宽度
    menuWidth: 200,
    //固定菜单
    fixed: true,
    //分割菜单
    mixMenu: false,
    //触发移动端侧边栏的宽度
    mobileWidth: 800,
  },
  //面包屑
  crumbsSetting: {
    //是否显示
    show: true,
    //显示图标
    showIcon: false,
  },
  //菜单权限模式 FIXED 前端固定路由  BACK 动态获取
  permissionMode: 'BACK',
  //是否开启路由动画
  isPageAnimate: true,
  //路由动画类型
  pageAnimateType: 'zoom-fade',
};
export default setting;
