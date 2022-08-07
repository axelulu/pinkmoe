/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:30
 * @FilePath: /pinkmoe_admin/src/api/system/menu.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取当前用户菜单
 */
export function adminMenus() {
  return http.request({
    url: '/Admin/Menu/Menu',
    method: 'GET',
  });
}

/**
 * @description: 获取tree菜单
 * @param params
 */
export function getMenuList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Menu/MenuList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 获取所有菜单
 * @param params
 */
export function getAllMenu(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Menu/AllMenu',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 获取角色菜单
 * @param params
 */
export function getMenuAuthority(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Authority/AuthorityMenu',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新菜单
 * @param params
 */
export function updateMenu(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Menu/MenuUpdate',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除菜单
 * @param params
 */
export function deleteMenu(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Menu/MenuDelete',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建菜单
 * @param params
 */
export function createMenu(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Menu/MenuCreate',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
