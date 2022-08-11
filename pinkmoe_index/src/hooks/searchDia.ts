import { useAppStore } from '../store';

/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 13:24:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-11 07:18:32
 * @FilePath: /pinkmoe_index/src/hooks/searchDia.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
export const useSearchDia = (props) => {
  const keyword = ref<string | string[]>('');
  const dialog = ref();
  const shopMenu = ref();
  const { userSearch } = useAppStore();
  function jumpTo(key: string | string[]) {
    keyword.value = key;
    props.router.push({
      path: '/search',
      query: {
        keyword: key,
        t: Date.parse(new Date().toString()),
      },
    });
    dialog.value.hide();
  }

  const scrollMenu = (isRight) => {
    shopMenu.value.scrollTo({
      // y方向坐标800px（代码中不需要写明单位）
      left: isRight ? shopMenu.value.scrollLeft + 140 : shopMenu.value.scrollLeft - 140,
      // 滚动方式是平滑滚动 默认是auto 即instant 直接跳到目标位置  不常用
      behavior: 'smooth',
    });
  };

  onMounted(() => {
    dialog.value.show();
  });
  return {
    jumpTo,
    scrollMenu,
    shopMenu,
    userSearch,
    dialog,
    keyword,
  };
};
