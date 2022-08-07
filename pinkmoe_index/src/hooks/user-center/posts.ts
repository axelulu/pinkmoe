/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 14:43:08
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:07:05
 * @FilePath: /pinkmoe_index/src/hooks/user-center/posts.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { getUserPost } from '/@/api/post';
import { ReqPage } from '/@/api/common/types';
import { ResPost } from '/@/api/home/types';

export const useUserCenterPosts = () => {
  const hasMore = ref(true);
  const loading = ref(false);
  const userPostList = ref<any>();

  const formParams: ReqPage = reactive({
    page: 1,
    pageSize: 12,
  });

  async function getUserPostList() {
    loading.value = true;
    userPostList.value = await getUserPost(formParams);
    setTimeout(() => {
      loading.value = false;
    }, 300);
  }

  async function nextPage() {
    loading.value = true;
    (formParams.page as number)++;
    const res = await getUserPost(formParams);
    if (!res?.list || res?.list?.length <= 0) {
      hasMore.value = false;
    } else {
      userPostList.value?.list?.push(...(res?.list as Array<ResPost>));
    }
    setTimeout(() => {
      loading.value = false;
    }, 300);
  }

  onMounted(() => {
    getUserPostList();
  });

  return {
    nextPage,
    hasMore,
    loading,
    userPostList,
  };
};
