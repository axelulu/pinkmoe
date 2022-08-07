/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:03:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:07:48
 * @FilePath: /pinkmoe_index/src/hooks/postDownload.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { useUserStore } from '/@/store';
import { buyPostDownload, getPostDownload } from '/@/api/post';
import { useClipboard } from '@vueuse/core';
import { createDialogModal } from '../utils/createModal';
import MsgConfirm from '/@/components/Msgconfirm/index.vue';

export const usePostDownload = (props) => {
  const auth = useUserStore();
  const postDownload = ref();
  const route = useRoute();
  const router = useRouter();
  const { proxy } = getCurrentInstance();
  const getDownload = async () => {
    postDownload.value = await getPostDownload({ postId: props.postId });
  };

  const showLogin = () => {
    proxy.$login({
      type: 'login',
      router: router,
      route: route,
    });
  };

  const copyText = async (text: string) => {
    const { copy } = useClipboard({ source: text });
    await copy();
    proxy.$message({
      type: 'success',
      msg: '复制成功',
    });
  };

  const buyDownload = async (item) => {
    const res = await createDialogModal(MsgConfirm, {
      msg: '确认购买此商品？',
    });
    if (res) {
      const { code, message } = await buyPostDownload({
        postId: props.postId,
        downloadId: item.ID,
      });
      if (code === 200) {
        proxy.$message({
          type: 'success',
          msg: '购买成功',
        });
        getDownload();
      } else {
        proxy.$message({
          type: 'warning',
          msg: message || '购买失败',
        });
      }
    }
  };

  onMounted(() => {
    getDownload();
  });

  return {
    auth,
    postDownload,
    showLogin,
    copyText,
    buyDownload,
  };
};
