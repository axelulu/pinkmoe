/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:15:37
 * @FilePath: /pinkmoe_admin/src/components/Form/src/hooks/useFormValues.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { isArray, isFunction, isObject, isString, isNullOrUnDef } from '@/utils/is';
import { unref } from 'vue';
import type { Ref, ComputedRef } from 'vue';
import type { FormSchema } from '../types/form';
import { set } from 'lodash-es';

interface UseFormValuesContext {
  defaultFormModel: Ref<any>;
  getSchema: ComputedRef<FormSchema[]>;
  formModel: Recordable;
}
export function useFormValues({ defaultFormModel, getSchema, formModel }: UseFormValuesContext) {
  // 加工 form values
  function handleFormValues(values: Recordable) {
    if (!isObject(values)) {
      return {};
    }
    const res: Recordable = {};
    for (const item of Object.entries(values)) {
      let [, value] = item;
      const [key] = item;
      if (
        !key ||
        (isArray(value) && value.length === 0) ||
        isFunction(value) ||
        isNullOrUnDef(value)
      ) {
        continue;
      }
      // 删除空格
      if (isString(value)) {
        value = value.trim();
      }
      set(res, key, value);
    }
    return res;
  }

  //初始化默认值
  function initDefault() {
    const schemas = unref(getSchema);
    const obj: Recordable = {};
    schemas.forEach((item) => {
      const { defaultValue } = item;
      if (!isNullOrUnDef(defaultValue)) {
        obj[item.field] = defaultValue;
        formModel[item.field] = defaultValue;
      }
    });
    defaultFormModel.value = obj;
  }

  return { handleFormValues, initDefault };
}
