
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="ID:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户唯一ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户名" />
       </el-form-item>
        <el-form-item label="显示名称:" prop="displayName">
          <el-input v-model="formData.displayName" :clearable="true"  placeholder="请输入显示名称" />
       </el-form-item>
        <el-form-item label="用户简介:" prop="bio">
          <el-input v-model="formData.bio" :clearable="true"  placeholder="请输入用户简介" />
       </el-form-item>
        <el-form-item label="位置信息:" prop="location">
          <el-input v-model="formData.location" :clearable="true"  placeholder="请输入位置信息" />
       </el-form-item>
        <el-form-item label="网站链接:" prop="website">
          <el-input v-model="formData.website" :clearable="true"  placeholder="请输入网站链接" />
       </el-form-item>
        <el-form-item label="注册日期:" prop="joinDate">
          <el-input v-model="formData.joinDate" :clearable="true"  placeholder="请输入注册日期" />
       </el-form-item>
        <el-form-item label="关注者数量:" prop="followersCount">
          <el-input v-model.number="formData.followersCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关注数量:" prop="followingCount">
          <el-input v-model.number="formData.followingCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总推文数:" prop="tweetCount">
          <el-input v-model.number="formData.tweetCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="点赞数量:" prop="likesCount">
          <el-input v-model.number="formData.likesCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="头像URL:" prop="profileImageUrl">
          <SelectImage v-model="formData.profileImageUrl" file-type="image"/>
       </el-form-item>
        <el-form-item label="封面图片URL:" prop="coverImageUrl">
          <SelectImage v-model="formData.coverImageUrl" file-type="image"/>
       </el-form-item>
        <el-form-item label="最后一次爬取时间:" prop="lastScraped">
          <el-date-picker v-model="formData.lastScraped" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="权重:" prop="weight">
          <el-input-number v-model="formData.weight" :precision="2" :clearable="true"></el-input-number>
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
  createTwitterUsers,
  updateTwitterUsers,
  findTwitterUsers
} from '@/api/Twitter/twitterUsers'

defineOptions({
    name: 'TwitterUsersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            userId: undefined,
            username: '',
            displayName: '',
            bio: '',
            location: '',
            website: '',
            joinDate: '',
            followersCount: undefined,
            followingCount: undefined,
            tweetCount: undefined,
            likesCount: undefined,
            profileImageUrl: "",
            coverImageUrl: "",
            lastScraped: new Date(),
            weight: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTwitterUsers({ ID: route.query.id })
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
               res = await createTwitterUsers(formData.value)
               break
             case 'update':
               res = await updateTwitterUsers(formData.value)
               break
             default:
               res = await createTwitterUsers(formData.value)
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