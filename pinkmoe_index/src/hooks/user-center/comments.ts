/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 14:35:57
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:07:00
 * @FilePath: /pinkmoe_index/src/hooks/user-center/comments.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { getUserComment } from '/@/api/comment';
import { ReqAuthorPost } from '/@/api/author/types';
import { useUserStore } from '/@/store';
import { ResPost } from '/@/api/home/types';

export const useUserCenterComments = () => {
  const loading = ref(true);
  const hasMore = ref(true);
  const authorCommentList = ref<any>([]);
  const auth = useUserStore();
  const formParams: ReqAuthorPost = reactive({
    page: 1,
    pageSize: 12,
    formUid: auth.userInfo?.uuid,
  });

  async function getUserCommentList() {
    loading.value = true;
    authorCommentList.value = await getUserComment(formParams);
    setTimeout(() => {
      loading.value = false;
    }, 300);
  }

  async function nextPage() {
    loading.value = true;
    (formParams.page as number)++;
    const res = await getUserComment(formParams);
    if (!res.list || res.list?.length <= 0) {
      hasMore.value = false;
    } else {
      authorCommentList.value?.list?.push(...(res?.list as Array<ResPost>));
    }
    setTimeout(() => {
      loading.value = false;
    }, 300);
  }

  onMounted(() => {
    getUserCommentList();
  });

  return {
    authorCommentList,
    hasMore,
    loading,
    nextPage,
  };
};
