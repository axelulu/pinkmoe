/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 08:40:11
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-11 07:21:23
 * @FilePath: /pinkmoe_index/src/hooks/shop/shop.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { ResPage } from '/@/api/common/types';
import { ResPost } from '/@/api/home/types';
import { ReqBbsPost } from '/@/api/post/types';
import { useUserStore } from '/@/store';

export const useShop = () => {
  const postList = ref<ResPage<Array<ResPost>>>();
  const { proxy } = getCurrentInstance();
  const shopMenu = ref();
  const sort = ref<any>([
    {
      title: '按最新',
      type: 'updated_at',
    },
    {
      title: '按标题',
      type: 'title',
    },
    {
      title: '按作者',
      type: 'author',
    },
    {
      title: '按查看',
      type: 'view',
    },
  ]);
  const auth = useUserStore();
  const formParams: ReqBbsPost = reactive({
    userId: auth?.userInfo?.uuid,
    category: '',
    type: '',
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
  });

  const loading = ref<boolean>(false);
  const hasMore = ref<boolean>(true);

  const scrollMenu = (isRight) => {
    shopMenu.value.scrollTo({
      // y方向坐标800px（代码中不需要写明单位）
      left: isRight ? shopMenu.value.scrollLeft + 140 : shopMenu.value.scrollLeft - 140,
      // 滚动方式是平滑滚动 默认是auto 即instant 直接跳到目标位置  不常用
      behavior: 'smooth',
    });
  };

  return {
    postList,
    sort,
    formParams,
    hasMore,
    loading,
    auth,
    scrollMenu,
    shopMenu,
  };
};
