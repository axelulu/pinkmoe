/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 09:00:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:06:56
 * @FilePath: /pinkmoe_index/src/hooks/bbs/bbsArticle.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { useUtil } from '../util';
import { createFollow, deleteFollow, followStatus } from '/@/api/follow';
import { ReqFollowStatus } from '/@/api/follow/types';
import { collectPost, unCollectPost } from '/@/api/post';
import { useUserStore } from '/@/store/modules/user';

export const useBbsArticle = (post) => {
  const status = ref<boolean>(false);
  const auth = useUserStore();
  const router = useRouter();
  const route = useRoute();
  const isCollect = ref(false);
  const formParams: ReqFollowStatus = reactive({
    toUid: post?.AuthorRelation?.uuid,
  });
  const { proxy } = getCurrentInstance();

  const deletePostImg = (content) => {
    const imgReg = /<img.*?(?:>|\/>)/gi; //定义正则表达式（拿到img标签所有值）
    const deArray = content.match(imgReg); //使用正则表达式，拿到的是数组
    //["<img src="images/01.gif">", "<img src="images/02.gif">"]
    if (deArray != null && deArray != undefined) {
      for (const vm of deArray) {
        content = content.replace(vm, '*'); //放在循环中剔除img标签
      }
    }
  };
  const { getLevel } = useUtil();
  const lev = getLevel(post?.AuthorRelation?.exp);

  const getFollowStatus = async () => {
    if (auth.isLogin) {
      status.value = await followStatus(formParams);
    }
  };

  const follow = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router: router,
        route: route,
      });
    } else {
      const { code, message } = await deleteFollow(formParams);
      if (code === 200) {
        proxy.$message({
          type: 'success',
          msg: '取消关注成功',
        });
        await getFollowStatus();
      } else {
        proxy.$message({
          type: 'warning',
          msg: message || '取消关注失败',
        });
      }
    }
  };

  const unFollow = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router: router,
        route: route,
      });
    } else {
      const { code, message } = await createFollow(formParams);
      if (code === 200) {
        proxy.$message({
          type: 'success',
          msg: '关注成功',
        });
        await getFollowStatus();
      } else {
        proxy.$message({
          type: 'warning',
          msg: message || '关注失败',
        });
      }
    }
  };

  const collect = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router: router,
        route: route,
      });
    } else {
      const { code, message } = await collectPost({ postId: post.postId });
      if (code === 200) {
        isCollect.value = true;
        proxy.$message({
          type: 'success',
          msg: '收藏成功',
        });
        await getFollowStatus();
      } else {
        proxy.$message({
          type: 'warning',
          msg: message || '收藏失败',
        });
      }
    }
  };

  const unCollect = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router: router,
        route: route,
      });
    } else {
      const { code, message } = await unCollectPost({ postId: post.postId });
      if (code === 200) {
        isCollect.value = false;
        proxy.$message({
          type: 'success',
          msg: '取消收藏成功',
        });
        await getFollowStatus();
      } else {
        proxy.$message({
          type: 'warning',
          msg: message || '取消收藏失败',
        });
      }
    }
  };

  const jump = () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router: router,
        route: route,
      });
    } else {
      router.push({
        path: '/user-center/im',
        query: { sendId: post?.AuthorRelation?.uuid },
      });
    }
  };

  onMounted(() => {
    isCollect.value = post?.collectRelation.length;
    getFollowStatus();
  });

  return {
    status,
    auth,
    isCollect,
    jump,
    unFollow,
    follow,
    collect,
    unCollect,
    deletePostImg,
    lev,
  };
};
