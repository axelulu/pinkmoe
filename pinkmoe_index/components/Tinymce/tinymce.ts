/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-19 19:24:32
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 19:24:38
 * @FilePath: /pinkmoe_vitesse/src/components/Tinymce/tinymce.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
// tinymce.ts

// imagetools
export const plugins = [
  'advlist anchor autolink code codesample  directionality  fullscreen hr insertdatetime link lists nonbreaking noneditable pagebreak paste preview print save searchreplace tabfocus  template  textpattern visualblocks visualchars wordcount table image toc',
]

export const toolbar = 'undo redo | bold italic underline strikethrough | fontselect fontsizeselect formatselect | toc alignleft aligncenter alignright alignjustify lineheight | outdent indent | numlist bullist | forecolor backcolor | pagebreak | charmap emoticons | fullscreen preview save print | hr link image | anchor pagebreak | insertdatetime | blockquote removeformat subscript superscript code codesample | searchreplace'

export const fontFormats = '微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif,Andale Mono=andale mono,times;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;Comic Sans MS=comic sans ms,sans-serif;Courier New=courier new,courier;Georgia=georgia,palatino;Helvetica=helvetica;Impact=impact,chicago;Symbol=symbol;Tahoma=tahoma,arial,helvetica,sans-serif;Terminal=terminal,monaco;Times New Roman=times new roman,times;Trebuchet MS=trebuchet ms,geneva;Verdana=verdana,geneva;Webdings=webdings;Wingdings=wingdings,zapf dingbats'

