/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 18:13:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:06:42
 * @FilePath: /pinkmoe_index/src/hooks/author/fans.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ResPage } from '/@/api/common/types';
import { getAuthorFansList } from '/@/api/author';
import { ReqAuthorFans, ResFollow } from '/@/api/author/types';

export const useAuthorFans = () => {
  const authorFansList = ref<ResPage<Array<ResFollow>>>();
  const route = useRoute();
  const loading = ref<boolean>(false);
  const hasMore = ref<boolean>(true);
  const formParams: ReqAuthorFans = reactive({
    author: '',
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
  });

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++;
    const res = await getAuthorFansList(formParams);
    if (!res.list || res.list.length <= 0) {
      hasMore.value = false;
    } else {
      authorFansList.value?.list?.push(...res.list!);
    }
    setTimeout(() => {
      loading.value = false;
    }, 300);
  };

  // 获取粉丝列表
  const getAuthorFans = async () => {
    loading.value = true;
    formParams.toUid = route.params.author;
    authorFansList.value = await getAuthorFansList(formParams);
    setTimeout(() => {
      loading.value = false;
    }, 300);
  };

  onMounted(() => {
    getAuthorFans();
  });

  return {
    authorFansList,
    loading,
    hasMore,
    formParams,
    nextPage,
  };
};
