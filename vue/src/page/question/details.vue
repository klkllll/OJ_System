<template>
    <div class="ques-cont">
        <div class="left">
            <h3>{{ detail.title }}</h3>
            <div class="msg">
                分类：<span v-for="mi in detail.problem_categories" :key="mi.id">
                    {{ mi.category_basic.name }}
                </span>
            </div>
            <div class="msg">
                <span>创建时间{{ detail.created_at }}</span>
                <span>提交次数：{{ detail.submit_num }}</span>
                <span>通过次数：{{ detail.pass_num }}</span>
                <span>最大耗时：{{ detail.max_runtime }}ms</span>

            </div>
            <div v-html="detail.content"></div>

        </div>
        <div class="right" v-if="dataLoaded">
            <Editor :Detail="detail"></Editor>
            
        </div>
        <div v-else>loading</div>
    </div>
</template>
<script lang="ts" setup>
import { reactive, ref } from '@vue/reactivity'
import { useRoute } from 'vue-router'

//import api from 'api'
import api from "../../api/api";
import Editor from './editor.vue'
import {
    Edit
} from '@element-plus/icons-vue'
import { onMounted } from 'vue';

const dataLoaded = ref(false);
const route = useRoute()
const getDetail = () => {
    api.getProblemDetail({
        identity: route.query.identity
    }).then(res => {
        if (res.data.data) {
            detail.value = res.data.data
            console.log("details:",detail.value);
            dataLoaded.value = true; // 设置数据已加载标志为true
        }
    })
}

onMounted(() => {
    getDetail() 
    dataLoaded.value = true;
})

const detail = ref(<any>{})
</script>
<style scoped lang="scss">
.ques-cont {
    display: flex;
    height: 100%;

    .left {
        width: 50%;
        line-height: 2;
        border-right: 10px solid #eee;
        padding: 10px;

        .msg {
            font-size: 12px;

            span {
                margin-right: 20px;
                color: #999;
            }
        }
    }

    .right {
        width: 50%;
    }
}
</style>