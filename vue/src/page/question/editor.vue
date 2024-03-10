<template>
  <div class="e-box">
    <div class="select">
      <el-select v-model="language" @change="changeLanguage">
        <el-option value="go">go</el-option>
      </el-select>
    </div>
    <!-- 代码编辑器 -->
    <div id="codeEditBox"></div>
    <div class="submit">
      <el-button type="primary" @click="submitCode" :loading="loading">提交</el-button>
    </div>
    <div class="sub-box">
     
     编译结果： <p> {{msg}}</p>
    </div>
  </div>
</template>
<script lang="ts" setup>
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker'
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker'
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker'
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker'
import EditorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'
import * as monaco from 'monaco-editor';
import { nextTick,ref,onBeforeUnmount, onMounted, watch } from 'vue'
import {useRoute} from 'vue-router'
import api from "../../api/api";
 
import {ElMessage} from 'element-plus'
const text=ref('') // 怎么让text一进去就去获取problem的数据
const route=useRoute()
const language=ref('go')
const msg=ref()
const loading=ref(false)

const props = defineProps({
  Detail: Object
});
console.log("editor:", props.Detail);

watch(() => props.Detail, (newVal, oldVal) => {
  // 在这里处理父组件传递过来的 Detail 数据的变化
  console.log('New value of Detail:', newVal.origin_code);
  //text.value = newVal.origin_code
  editor.setValue(newVal.origin_code)
});

// @ts-ignore
//text.value = props.detail.origin_code
// 
// MonacoEditor start
//

/*
onMounted(() => {
  text.value = props.Detail.origin_code
  console.log("text:",text.value);
  
})*/

onBeforeUnmount(()=>{
   editor.dispose() 
 })
// @ts-ignore
self.MonacoEnvironment = {
  getWorker(_: string, label: string) {
    if (label === 'json') {
      return new jsonWorker()
    }
    if (label === 'css' || label === 'scss' || label === 'less') {
      return new cssWorker()
    }
    if (label === 'html' || label === 'handlebars' || label === 'razor') {
      return new htmlWorker()
    }
    if (['typescript', 'javascript'].includes(label)) {
      return new tsWorker()
    }
    return new EditorWorker()
  },
}
let editor: monaco.editor.IStandaloneCodeEditor;

const editorInit = () => {
    nextTick(()=>{
        
        monaco.languages.typescript.javascriptDefaults.setDiagnosticsOptions({
            noSemanticValidation: true,
            noSyntaxValidation: false
        });
        monaco.languages.typescript.javascriptDefaults.setCompilerOptions({
            target: monaco.languages.typescript.ScriptTarget.ES2016,
            allowNonTsExtensions: true
        });
        // 设置代码编辑器的样式，这是一个条件表达式，检查 editor 是否已经存在，如果不存在则创建一个新的编辑器实例，否则将编辑器的内容设置为空字符串。
        !editor ? editor = monaco.editor.create(document.getElementById('codeEditBox') as HTMLElement, {
            value:text.value, // 编辑器初始显示文字
            language: 'go', // 语言支持自行查阅demo
            automaticLayout: true, // 自适应布局  
            theme: 'vs-dark', // 官方自带三种主题vs, hc-black, or vs-dark
            foldingStrategy: 'indentation',
            renderLineHighlight: 'all', // 行亮
            selectOnLineNumbers: true, // 显示行号
            minimap:{
                enabled: false,
            },
            readOnly: false, // 只读
            fontSize: 16, // 字体大小
            scrollBeyondLastLine: false, // 取消代码后面一大段空白 
            overviewRulerBorder: false, // 不要滚动条的边框  
        }) : editor.setValue(text.value);
        // console.log(editor)
        // 监听值的变化
        editor.onDidChangeModelContent((val:any) => {
            text.value = editor.getValue();
             
        })
    })
}
editorInit()
// @ts-ignore
const changeLanguage=()=>{
   monaco.editor.setModelLanguage(editor.getModel(), language.value) // 不用管的，这里用ts-ignore给忽略掉了
        

  //  editor.updateOptions({
  //           language: "objective-c"
  //       });
}
const submitCode=()=>{
  loading.value=true
  api.submitCode(text.value,route.query.identity).then(res=>{
    loading.value=false
      if(res.data.code==200){
        msg.value=res.data.data.msg
      
        if(res.data.data.status==1){
            //ElMessage.success(res.data.data.msg)
            ElMessage.success("答案正确")

        }else{
             //ElMessage.warning(res.data.data.msg)
             ElMessage.warning("答案错误")
        }
       

      }else{
        ElMessage.error(res.data.msg)
      }
  })
}
/***
 * editor.setValue(newValue)

editor.getValue()

editor.onDidChangeModelContent((val)=>{ //监听值的变化  })

editor.getAction('editor.action.formatDocument').run()    //格式化代码

editor.dispose()    //销毁实例

editor.onDidChangeOptions　　//配置改变事件

editor.onDidChangeLanguage　　//语言改变事件
 */
</script>
<style scoped lang="scss">
#codeEditBox {
  //  width: 100%;
  height: 50%;
}
.select{
  padding: 10px 0;
}
.submit{
  text-align: center;
  padding: 10px 0;
  
}
.e-box{
  width: 100%;
  padding: 10px ;
  box-sizing: border-box;
  height: 100%;
  display: flex;
  justify-content: space-between;
  box-sizing: border-box;
  flex-direction: column;
}
.sub-box{
  border-radius: 10px;
  background-color: rgb(85, 83, 83);
  padding: 10px;
  box-sizing: border-box;
  flex: 1;
  color: white;
  overflow: scroll;
  
  p {
        white-space: pre-wrap;
        word-wrap: break-word;
    }
}
</style>