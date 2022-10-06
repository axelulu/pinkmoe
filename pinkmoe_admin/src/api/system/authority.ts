/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:28
 * @FilePath: /pinkmoe_admin/src/api/system/authority.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取角色列表
 * @param params
 */
export function getRoleList(params) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Authority/AuthorityList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建角色菜单
 * @param params
 */
export function updateMenuAuthority(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Authority/AuthorityMenuUpdate',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建角色
 * @param params
 */
export function createAuthority(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Authority/AuthorityCreate',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新角色
 * @param params
 */
export function updateAuthority(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Authority/AuthorityUpdate',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除角色
 * @param params
 */
export function deleteAuthority(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Authority/AuthorityDelete',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
