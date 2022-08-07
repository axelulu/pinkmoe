/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-25 21:03:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:07:17
 * @FilePath: /pinkmoe_index/src/hooks/user-center/stars.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ReqAuthorPost } from '/@/api/author/types';
import { ResPost } from '/@/api/home/types';
import { getCollectPostList } from '/@/api/post';

export const useUserCenterStars = () => {
  const loading = ref(true);
  const hasMore = ref(true);
  const authorCollectList = ref<any>([]);
  const formParams: ReqAuthorPost = reactive({
    page: 1,
    pageSize: 12,
  });

  async function getUserCollectList() {
    loading.value = true;
    authorCollectList.value = await getCollectPostList(formParams);
    setTimeout(() => {
      loading.value = false;
    }, 300);
  }

  async function nextPage() {
    loading.value = true;
    (formParams.page as number)++;
    const res = await getCollectPostList(formParams);
    if (!res.list || res.list?.length <= 0) {
      hasMore.value = false;
    } else {
      authorCollectList.value?.list?.push(...(res?.list as Array<ResPost>));
    }
    setTimeout(() => {
      loading.value = false;
    }, 300);
  }

  onMounted(() => {
    getUserCollectList();
  });

  return {
    authorCollectList,
    hasMore,
    loading,
    nextPage,
  };
};
