/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 10:47:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 21:29:01
 * @FilePath: /xanaduCms/pinkmoe_index/src/hooks/commentPublish.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useUtil } from './util';
import { createComment } from '/@/api/comment';
import { useUserStore } from '/@/store';

export const useCommentPublish = (props, emit) => {
  const dialog = ref();
  const font = ref();
  const smile = ref();
  const formComment = ref();
  const { checkForm } = useUtil();
  const commentContent = ref<string>('');
  const commentType = ref<string>('post');
  const commentMetas = ref<any>(null);
  const showAnimate = ref<boolean>(false);
  const auth = useUserStore();
  const { proxy } = getCurrentInstance();

  const show = (commentMeta) => {
    if (commentMeta !== null) {
      commentMetas.value = commentMeta;
      commentType.value = 'comment';
    } else {
      commentMetas.value = null;
      commentType.value = 'post';
    }
    dialog.value.show();
  };

  const showFont = () => {
    font.value.show();
  };

  const showSmile = () => {
    smile.value.show();
  };

  const formParams = reactive({
    postId: props.postId,
    parentId: commentMetas.value ? commentMetas.value.ID : 0,
    formUid: auth.userInfo?.uuid,
    toUid: commentMetas.value ? commentMetas.value.FormUidRelation.uuid : null,
    content: commentContent.value,
    type: commentType.value,
    status: 'published',
  });

  const submitComment = async () => {
    if (checkForm(commentContent.value, '[\\s\\S]{1,100}$', '请输入正确的评论内容')) {
      formParams.content = commentContent.value;
      formParams.toUid = commentMetas.value ? commentMetas.value.FormUidRelation.uuid : null;
      formParams.parentId = commentMetas.value ? commentMetas.value.ID : 0;
      formParams.type = commentType.value;
      const { code, message } = await createComment(formParams);
      if (code === 200) {
        commentContent.value = '';
        proxy.$message({
          type: 'success',
          msg: '评论成功',
        });
        setTimeout(() => {
          dialog.value.hide();
          emit('getPostComment');
        }, 1000);
      } else {
        proxy.$message({
          type: 'success',
          msg: message || '评论失败',
        });
      }
    }
  };

  const insertFont = (item) => {
    commentContent.value = commentContent.value + item;
  };

  const insertSmile = (item) => {
    commentContent.value = commentContent.value + item;
  };

  return {
    insertSmile,
    insertFont,
    showSmile,
    showFont,
    show,
    submitComment,
    showAnimate,
    commentMetas,
    formComment,
    commentType,
    commentContent,
    smile,
    font,
    dialog,
  };
};
