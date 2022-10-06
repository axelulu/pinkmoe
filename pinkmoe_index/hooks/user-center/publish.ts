/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 14:33:05
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 23:31:42
 * @FilePath: /pinkmoe_index/hooks/user-center/publish.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { useUserStore } from '/@/store/modules/user'
import { getCategoryList } from '/@/api/category'
import { deleteFile, fileList, upload } from '/@/api/upload'
import type { ReqFileList } from '/@/api/upload/types'
import { createPost } from '/@/api/post'
import type { ReqPublishPost } from '/@/api/post/types'
import { useUtil } from '../util'

export const useUserCenterPublish = () => {
  const categoryList = ref<any>()
  const wangEditor = ref()
  const message = ref()
  const formPublish = ref()
  const loading = ref(false)
  const publishCategory = ref()
  const { checkForm } = useUtil()
  const auth = useUserStore()
  const router = useRouter()
  const { proxy } = getCurrentInstance()
  const category = ref<any>({
    name: '',
    slug: '',
  })
  const formParams: ReqPublishPost = reactive({
    postId: '',
    title: '',
    exerpt: '',
    content: 'hello',
    category: '',
    author: auth.userInfo?.uuid,
    cover: '',
    postImg: <any>[],
    type: 'post',
    view: 0,
    videoRelation: [
      {
        name: '',
        url: '',
        subtitles: '',
        price: 0,
        priceType: 'credit',
      },
    ],
    musicRelation: [
      {
        name: '',
        url: '',
        price: 0,
        priceType: 'credit',
      },
    ],
    from: 'original',
    downloadRelation: [
      // {
      //   name: '',
      //   url: '',
      //   unpackPwd: '',
      //   extractPwd: '',
      //   price: 0,
      //   priceType: 'credit',
      // },
    ],
    status: 'pending',
    topic: [] as string[],
    commentStatus: 'true',
  })

  const fileFormParams: ReqFileList = reactive({
    page: 1,
    pageSize: 999,
    type: 'post_temporary',
  })

  async function getCategory() {
    categoryList.value = await getCategoryList()
    publishCategory.value.show()
  }

  async function getImgList() {
    loading.value = true
    formParams.postImg = (await fileList(fileFormParams)).list
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  onMounted(() => {
    getImgList()
  })

  function chooseCategory(item: { slug: any; name: any }) {
    category.value.slug = item.slug
    category.value.name = item.name
    formParams.category = item.slug
    publishCategory.value.hide()
  }

  async function uploadPostImg(event: any) {
    const fd = new FormData()
    fd.append('file', event.target.files[0])
    fd.append('uuid', auth.userInfo?.uuid as string)
    fd.append('type', 'post_temporary')
    proxy.$message({
      successMsg: '上传成功',
      failedMsg: '上传失败',
      loadFun: async () => {
        const { code, message, result } = await upload(fd)
        return { code, message, result }
      },
    }).then(async (res) => {
      if (res.status === 200) {
        formParams.postImg?.push(res.result.file!)
        if (formParams.cover?.length === 0)
          formParams.cover = formParams.postImg?.[0].url
      }
    })
    event.target.value = ''
  }

  function insertPostImg(img: any) {
    wangEditor.value.insertCover(img.url)
  }

  async function deletePostImg(index: any) {
    proxy.$message({
      successMsg: '删除成功',
      failedMsg: '删除失败',
      loadFun: async () => {
        const { code, message } = await deleteFile({
          url: formParams.postImg?.[index].url,
        })
        return { code, message }
      },
    }).then(async (res) => {
      if (res.status === 200) {
        if (formParams.postImg?.[index].url === formParams.cover)
          formParams.cover = formParams.postImg?.[0].url
        formParams.postImg?.splice(index, 1)
      }
    })
  }

  async function publishPost() {
    if (
      checkForm(formParams.title, '[\\s\\S]{2,50}$', '请输入正确的文章标题格式')
      && checkForm(formParams.exerpt, '[\\s\\S]{2,120}$', '请输入正确的文章摘要格式')
      && checkForm(formParams.content, '[\\s\\S]{2,120000}$', '请输入正确的文章内容格式')
    ) {
      proxy.$message({
        successMsg: '发布成功',
        failedMsg: '发布失败',
        loadFun: async () => {
          const { code, message } = await createPost(formParams)
          return { code, message }
        },
      }).then(async (res) => {
        if (res.status === 200) {
          setTimeout(() => {
            router.push('/user-center/posts')
          }, 1000)
        }
      })
    }
  }

  return {
    category,
    categoryList,
    publishCategory,
    wangEditor,
    message,
    formParams,
    loading,
    getCategory,
    publishPost,
    chooseCategory,
    uploadPostImg,
    insertPostImg,
    deletePostImg,
    formPublish,
  }
}
