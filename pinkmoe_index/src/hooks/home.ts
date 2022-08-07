/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 09:10:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:07:38
 * @FilePath: /pinkmoe_index/src/hooks/home.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { getHomeList } from '/@/api/home';
import { ResContent, ResPost } from '/@/api/home/types';

export const useHomeList = () => {
  const content = ref<Array<ResContent>>();
  const popular = ref<Array<ResPost>>();
  const recommend = ref<Array<ResPost>>();

  // 获取首页
  const getHome = async () => {
    const homeList = await getHomeList();
    content.value = homeList.content;
    popular.value = homeList.popular;
    recommend.value = homeList.recommend;
  };

  onMounted(() => {
    getHome();
  });

  return {
    content,
    popular,
    recommend,
  };
};
