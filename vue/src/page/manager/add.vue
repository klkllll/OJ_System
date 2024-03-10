<template>
  <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="120px" class="demo-ruleForm" :size="formSize">
    <el-form-item label="问题名称" prop="title">
      <el-input v-model="ruleForm.title" />
    </el-form-item>
    <el-form-item label="内容" prop="content ">
      <el-input v-model="ruleForm.content" type="textarea" />
    </el-form-item>
    <el-form-item label="分类" prop="problem_categories">
      <el-select v-model="ruleForm.problem_categories" placeholder="请选择" multiple style="width: 100%">
        <el-option :label="mi.name" :value="mi.id" v-for="mi in sortList" :key="mi.id" />
      </el-select>
    </el-form-item>
    <el-form-item label="最大运行时间" prop="max_runtime">
      <el-input v-model.number="ruleForm.max_runtime" />
    </el-form-item>
    <el-form-item label="最大占用内存" prop="max_mem">
      <el-input v-model.number="ruleForm.max_mem" />
    </el-form-item>

    <el-form-item label="测试案例" prop="test_cases">
      <el-form-item label-width="0" v-for="(mi, i) in ruleForm.test_cases" :key="i">
        <el-row :gutter="20">
          <el-col :span="4" style="text-align: right">输入：</el-col>
          <el-col :span="6">
            <el-input type="textarea" v-model="mi.input" rows="6" cols="40" />
          </el-col>
          <el-col :span="4" style="text-align: right">输出：</el-col>
          <el-col :span="6">
            <el-input type="textarea" v-model="mi.output" rows="6" cols="40" />
          </el-col>
          <el-col :span="4">
            <el-icon @click="addCase">
              <circle-plus />
            </el-icon>
            <el-icon @click="removeCase(i)">
              <remove />
            </el-icon>
          </el-col>
        </el-row>
      </el-form-item>
    </el-form-item>
    <el-form-item label="起始代码" prop="origin_code">
      <div id="codeEditBox" style="height: 300px;"></div>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="submitForm(ruleFormRef)">创建
      </el-button>
      <el-button @click="closeBox">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker'
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker'
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker'
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker'
import EditorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'
import * as monaco from 'monaco-editor';

import { nextTick, onBeforeUnmount, reactive, ref } from "vue";
import type { FormInstance, FormRules } from "element-plus";
import { ElMessage } from "element-plus";
import { CirclePlus, Remove } from "@element-plus/icons-vue";
//import api from "api";
import api from "../../api/api";



const formSize = ref("default");
const ruleFormRef = ref<FormInstance>();
// interface formBase{
//    title :string,
//   content : '',
//   max_runtime : '',
//   max_mem : '',
//   test_cases : [],
//   category_ids: [],
// }
const ruleForm = ref({
  title: "",
  content: "",
  max_runtime: 0,
  max_mem: 0,
  test_cases: [{ input: "", output: "" }],
  problem_categories: [],
  origin_code: "",
});
const emits = defineEmits(["cancel", "submit"]);
const props = defineProps(["sortList", "selectQues"]);
const text = ref('')

if (props.selectQues) {
  console.log("1", props.selectQues);

  let se = props.selectQues;
  // se.problem_categories 数组中的每个元素的 category_id 属性提取出来，然后重新赋值给 se.problem_categories
  if (se.problem_categories instanceof Array) {
    se.problem_categories = se.problem_categories.map((it: any) => {
      return it.category_id;
    });
  }

  if (!Array.isArray(se.test_cases) || se.test_cases.length == 0) {
    se.test_cases = [{ input: "", output: "" }];
    console.log("12", se.test_cases);
  }
  ruleForm.value = se;
  console.log(ruleForm.value);
  // 识别\n \t
  // replacedText.value = ruleForm.value.origin_code.replace(/\n/g, '<br>')
  // replacedText.value = replacedText.value.replace(/\t/g, '&nbsp;&nbsp;&nbsp;&nbsp;')
  text.value = se.origin_code
}
const closeBox = () => {
  emits("cancel");
};

const rules = reactive<FormRules>({// 用于提交表单数据
  title: [{ required: true, message: "请输入", trigger: "blur" }],
  content: [{ required: true, message: "请输入", trigger: "blur" }],
  max_runtime: [
    { required: true, message: "请输入", trigger: "blur" },
    { type: "number", message: "请输入数字", trigger: "blur" },
  ],
  max_mem: [
    { required: true, message: "请输入", trigger: "blur" },
    { type: "number", message: "请输入数字", trigger: "blur" },
  ],
  test_cases: [{ required: true, message: "请输入", trigger: "blur" }],
  problem_categories: [
    {
      type: "array",
      required: true,
      message: "Please select at least one activity type",
      trigger: "change",
    },
  ],
  origin_code: [{required: true }]
});
const addCase = () => {
  ruleForm.value.test_cases.push({
    input: "",
    output: "",
  });
  // console.log(props.selectQues.test_cases);

};
const removeCase = (i: number) => {
  if (ruleForm.value.test_cases.length > 1) {
    ruleForm.value.test_cases.splice(i, 1);
  }
};
const submitForm = async (formEl: FormInstance | undefined) => {// 提交表单
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      if (ruleForm.value.identity) {
        console.log(ruleForm.value);
        
        api.editProblem(ruleForm.value).then((res) => {
          if (res.data.code == 200) {
            ElMessage.success("成功");
            emits("submit");
          } else {
            ElMessage(res.data.msg);
          }
        });
      } else {
        api.addProblem(ruleForm.value).then((res: any) => {
          if (res.data.code == 200) {
            ElMessage.success("成功");
            emits("submit");
          } else {
            ElMessage(res.data.msg);
          }
        });
      }
    } else {
      console.log("error submit!", fields);
    }
  });
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};

// 代码编辑器

onBeforeUnmount(() => {
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
  nextTick(() => {

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
      value: text.value, // 编辑器初始显示文字
      language: 'go', // 语言支持自行查阅demo
      automaticLayout: true, // 自适应布局  
      theme: 'vs-dark', // 官方自带三种主题vs, hc-black, or vs-dark
      foldingStrategy: 'indentation',
      renderLineHighlight: 'all', // 行亮
      selectOnLineNumbers: true, // 显示行号
      minimap: {
        enabled: false,
      },
      readOnly: false, // 只读
      fontSize: 16, // 字体大小
      scrollBeyondLastLine: false, // 取消代码后面一大段空白 
      overviewRulerBorder: false, // 不要滚动条的边框  
    }) : editor.setValue(text.value);
    // console.log(editor)
    // 监听值的变化
    editor.onDidChangeModelContent((val: any) => {
      text.value = editor.getValue();
      ruleForm.value.origin_code = editor.getValue();
    })
  })
}
editorInit()
</script>

<style scoped lang="scss">
// .el-input {
//   white-space: pre-line;
// }
#codeEditBox {
  width: 100%;
  height: 50%;
}
</style>