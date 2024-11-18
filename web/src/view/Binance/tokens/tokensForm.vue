
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="ID:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="标识:" prop="chainName">
          <el-input v-model="formData.chainName" :clearable="true"  placeholder="请输入标识" />
       </el-form-item>
        <el-form-item label="市值:" prop="marketCap">
          <el-input v-model.number="formData.marketCap" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="isMeme:" prop="isMeme">
          <el-input v-model.number="formData.isMeme" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ok上线:" prop="okx">
          <el-input v-model.number="formData.okx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="gate上线:" prop="gate">
          <el-input v-model.number="formData.gate" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="bybit上线:" prop="bybit">
          <el-input v-model.number="formData.bybit" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="kucoin上线:" prop="kucoin">
          <el-input v-model.number="formData.kucoin" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTokens,
  updateTokens,
  findTokens
} from '@/api/Binance/tokens'

defineOptions({
    name: 'TokensForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            address: '',
            chainName: '',
            marketCap: undefined,
            isMeme: undefined,
            okx: undefined,
            gate: undefined,
            bybit: undefined,
            kucoin: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTokens({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTokens(formData.value)
               break
             case 'update':
               res = await updateTokens(formData.value)
               break
             default:
               res = await createTokens(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>