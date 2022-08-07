<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 17:37:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:12:02
 * @FilePath: /pinkmoe_index/src/components/Wangeditor/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Wangeditor">
  import '@wangeditor/editor/dist/css/style.css'; // 引入 css
  import { Editor, Toolbar } from '@wangeditor/editor-for-vue';

  const props = defineProps({
    value: {
      type: String,
      default: '<p></p>',
    },
  });
  const emit = defineEmits(['update:value']);

  // 编辑器实例，必须用 shallowRef，重要！
  const editorRef = shallowRef();

  const toolbarConfig = {};
  const editorConfig = { placeholder: '请输入内容...' };
  // 内容 HTML
  const valueHtml = ref('');
  // 模拟 ajax 异步获取内容
  onMounted(() => {
    setTimeout(() => {
      valueHtml.value = props.value;
    }, 1000);
  });

  // 组件销毁时，也及时销毁编辑器，重要！
  onBeforeUnmount(() => {
    const editor = editorRef.value;
    if (editor == null) return;

    editor.destroy();
  });

  // 编辑器回调函数
  const handleCreated = (editor) => {
    editorRef.value = editor; // 记录 editor 实例，重要！
  };
  const handleChange = (editor) => {
    emit('update:value', editor.getHtml());
  };
  const handleDestroyed = (editor) => {
    console.log('destroyed', editor);
  };
  const handleFocus = (editor) => {
    console.log('focus', editor);
  };
  const handleBlur = (editor) => {
    console.log('blur', editor);
  };
  const customAlert = (info, type) => {
    alert(`【自定义提示】${type} - ${info}`);
  };
  const customPaste = (editor, event, callback) => {
    // const html = event.clipboardData.getData("text/html"); // 获取粘贴的 html
    const text = event.clipboardData.getData('text/plain'); // 获取粘贴的纯文本
    // const rtf = event.clipboardData.getData('text/rtf') // 获取 rtf 数据（如从 word wsp 复制粘贴）
    // 自定义插入内容
    editor.insertText(text);

    // 返回值（注意，vue 事件的返回值，不能用 return）
    callback(false); // 返回 false ，阻止默认粘贴行为
    // callback(true) // 返回 true ，继续默认的粘贴行为
  };

  const insertText = () => {
    const editor = editorRef.value;
    if (editor == null) return;
    editor.insertText('hello world');
  };

  const printHtml = () => {
    const editor = editorRef.value;
    if (editor == null) return;
  };

  const disable = () => {
    const editor = editorRef.value;
    if (editor == null) return;
    editor.disable();
  };

  function insertCover(cover) {
    const editor = editorRef.value;
    if (editor == null) return;
    editor.dangerouslyInsertHtml('<img src="' + cover + '">');
  }

  defineExpose({
    insertCover,
  });
</script>

<template>
  <!-- Wangeditor -->
  <div style="border: 1px solid #ccc; margin-top: 10px">
    <Toolbar
      :editor="editorRef"
      :defaultConfig="toolbarConfig"
      mode="default"
      style="border-bottom: 1px solid #ccc"
    />
    <Editor
      :defaultConfig="editorConfig"
      mode="default"
      v-model="valueHtml"
      style="height: 400px; overflow-y: hidden"
      @onCreated="handleCreated"
      @onChange="handleChange"
      @onDestroyed="handleDestroyed"
      @onFocus="handleFocus"
      @onBlur="handleBlur"
      @customAlert="customAlert"
      @customPaste="customPaste"
    />
  </div>
</template>

<style lang="less">
  /* 暗色主题 */
  html.dark {
    --w-e-textarea-bg-color: #333;
    --w-e-textarea-color: #fff;
    --w-e-toolbar-color: #fff;
    --w-e-toolbar-bg-color: #333;
    --w-e-toolbar-active-bg-color: #333;
    --w-e-toolbar-active-color: #fff;
    --w-e-modal-button-bg-color: #333;
    /* ...其他... */
  }
</style>
