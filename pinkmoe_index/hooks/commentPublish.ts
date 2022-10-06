/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 10:47:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 17:03:14
 * @FilePath: /pinkmoe_index/hooks/commentPublish.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { useUtil } from './util'
import { createComment } from '/@/api/comment'
import { useUserStore } from '/@/store/modules/user'

export const useCommentPublish = (props, emit) => {
  const dialog = ref()
  const font = ref()
  const smile = ref()
  const formComment = ref()
  const { checkForm } = useUtil()
  const commentContent = ref<string>('')
  const commentType = ref<string>('post')
  const commentMetas = ref<any>(null)
  const showAnimate = ref<boolean>(false)
  const auth = useUserStore()
  const { proxy } = getCurrentInstance()

  const show = (commentMeta) => {
    if (commentMeta !== null) {
      commentMetas.value = commentMeta
      commentType.value = 'comment'
    }
    else {
      commentMetas.value = null
      commentType.value = 'post'
    }
    dialog.value.show()
  }

  const showFont = () => {
    font.value.show()
  }

  const showSmile = () => {
    smile.value.show()
  }

  const formParams = reactive({
    postId: props.postId,
    parentId: commentMetas.value ? commentMetas.value.ID : 0,
    formUid: auth.userInfo?.uuid,
    toUid: commentMetas.value ? commentMetas.value.FormUidRelation.uuid : null,
    content: commentContent.value,
    type: commentType.value,
    status: 'published',
  })

  const submitComment = async () => {
    if (checkForm(commentContent.value, '[\\s\\S]{1,100}$', '请输入正确的评论内容')) {
      formParams.content = commentContent.value
      formParams.toUid = commentMetas.value ? commentMetas.value.FormUidRelation.uuid : null
      formParams.parentId = commentMetas.value ? commentMetas.value.ID : 0
      formParams.type = commentType.value
      proxy.$message({
        successMsg: '评论成功',
        failedMsg: '评论失败',
        loadFun: async () => {
          const { code, message } = await createComment(formParams)
          return { code, message }
        },
      }).then((res) => {
        if (res.status === 200) {
          commentContent.value = ''
          setTimeout(() => {
            dialog.value.hide()
            emit('getPostComment')
          }, 1000)
        }
      })
    }
  }

  const insertFont = (item) => {
    commentContent.value = commentContent.value + item
  }

  const insertSmile = (item) => {
    commentContent.value = commentContent.value + item
  }

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
  }
}
