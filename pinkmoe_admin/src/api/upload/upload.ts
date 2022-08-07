/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-29 21:47:11
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:35
 * @FilePath: /pinkmoe_admin/src/api/upload/upload.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from "@/utils/http/axios";

/**
 * @description: 删除api
 * @param params
 */
export function deleteImgFile(params) {
  return http.request(
    {
      url: `/Admin/Upload/ImgFileDelete`,
      method: "POST",
      params
    },
    {
      isTransformResponse: false
    }
  );
}
