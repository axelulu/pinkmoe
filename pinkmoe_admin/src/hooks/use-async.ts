/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:18:44
 * @FilePath: /pinkmoe_admin/src/hooks/use-async.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { isReactive, isRef } from 'vue';

function setLoading(loading, val) {
  if (loading != undefined && isRef(loading)) {
    loading.value = val;
  } else if (loading != undefined && isReactive(loading)) {
    loading.loading = val;
  }
}

export const useAsync = async (func: Promise<any>, loading: any): Promise<any> => {
  setLoading(loading, true);

  return await func.finally(() => setLoading(loading, false));
};
