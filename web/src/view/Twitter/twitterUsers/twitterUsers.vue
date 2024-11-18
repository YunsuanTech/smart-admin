
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          
          <el-form-item label="用户名">
            <el-input
              v-model="searchInfo.username"
              placeholder="用户名"
            />
          </el-form-item>
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" @click="fetchThirdPartyData">重制用户</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column label="头像" prop="profileImageUrl" width="180">
              <template #default="scope">
                <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.profileImageUrl)" fit="cover"/>
              </template>
          </el-table-column>
          <!-- <el-table-column align="left" label="ID" prop="id" width="120" />
          <el-table-column align="left" label="用户唯一ID" prop="userId" width="120" /> -->
          <el-table-column align="left" label="用户名" prop="username" width="120" />
          <el-table-column align="left" label="权重" prop="weight" width="120" />
          <!-- <el-table-column align="left" label="显示名称" prop="displayName" width="120" />
          <el-table-column align="left" label="用户简介" prop="bio" width="120" />
          <el-table-column align="left" label="位置信息" prop="location" width="120" />
          <el-table-column align="left" label="网站链接" prop="website" width="120" />
          <el-table-column align="left" label="注册日期" prop="joinDate" width="120" /> -->
          <el-table-column align="left" label="关注者数量" prop="followersCount" width="120" />
          <el-table-column align="left" label="关注数量" prop="followingCount" width="120" />
          <el-table-column align="left" label="总推文数" prop="tweetCount" width="120" />
          <el-table-column align="left" label="点赞数量" prop="likesCount" width="120" />
<!-- 
          <el-table-column label="封面图片URL" prop="coverImageUrl" width="200">
              <template #default="scope">
                <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.coverImageUrl)" fit="cover"/>
              </template> 
          </el-table-column>-->
         <el-table-column align="left" label="最后一次爬取时间" prop="lastScraped" width="180">
            <template #default="scope">{{ formatDate(scope.row.lastScraped) }}</template>
         </el-table-column>

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateTwitterUsersFunc(scope.row)">编辑</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="ID:"  prop="id" >
              <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入ID" />
            </el-form-item>
            <el-form-item label="用户唯一ID:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户唯一ID" />
            </el-form-item>
            <el-form-item label="用户名:"  prop="username" >
              <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="显示名称:"  prop="displayName" >
              <el-input v-model="formData.displayName" :clearable="true"  placeholder="请输入显示名称" />
            </el-form-item>
            <el-form-item label="用户简介:"  prop="bio" >
              <el-input v-model="formData.bio" :clearable="true"  placeholder="请输入用户简介" />
            </el-form-item>
            <el-form-item label="位置信息:"  prop="location" >
              <el-input v-model="formData.location" :clearable="true"  placeholder="请输入位置信息" />
            </el-form-item>
            <el-form-item label="网站链接:"  prop="website" >
              <el-input v-model="formData.website" :clearable="true"  placeholder="请输入网站链接" />
            </el-form-item>
            <el-form-item label="注册日期:"  prop="joinDate" >
              <el-input v-model="formData.joinDate" :clearable="true"  placeholder="请输入注册日期" />
            </el-form-item>
            <el-form-item label="关注者数量:"  prop="followersCount" >
              <el-input v-model.number="formData.followersCount" :clearable="true" placeholder="请输入关注者数量" />
            </el-form-item>
            <el-form-item label="关注数量:"  prop="followingCount" >
              <el-input v-model.number="formData.followingCount" :clearable="true" placeholder="请输入关注数量" />
            </el-form-item>
            <el-form-item label="总推文数:"  prop="tweetCount" >
              <el-input v-model.number="formData.tweetCount" :clearable="true" placeholder="请输入总推文数" />
            </el-form-item>
            <el-form-item label="点赞数量:"  prop="likesCount" >
              <el-input v-model.number="formData.likesCount" :clearable="true" placeholder="请输入点赞数量" />
            </el-form-item>
            <el-form-item label="头像URL:"  prop="profileImageUrl" >
                <SelectImage
                 v-model="formData.profileImageUrl"
                 file-type="image"
                />
            </el-form-item>
            <el-form-item label="封面图片URL:"  prop="coverImageUrl" >
                <SelectImage
                 v-model="formData.coverImageUrl"
                 file-type="image"
                />
            </el-form-item>
            <el-form-item label="最后一次爬取时间:"  prop="lastScraped" >
              <el-date-picker v-model="formData.lastScraped" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="权重:"  prop="weight" >
              <el-input-number v-model="formData.weight"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="ID">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户唯一ID">
                        {{ detailFrom.userId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户名">
                        {{ detailFrom.username }}
                    </el-descriptions-item>
                    <el-descriptions-item label="显示名称">
                        {{ detailFrom.displayName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户简介">
                        {{ detailFrom.bio }}
                    </el-descriptions-item>
                    <el-descriptions-item label="位置信息">
                        {{ detailFrom.location }}
                    </el-descriptions-item>
                    <el-descriptions-item label="网站链接">
                        {{ detailFrom.website }}
                    </el-descriptions-item>
                    <el-descriptions-item label="注册日期">
                        {{ detailFrom.joinDate }}
                    </el-descriptions-item>
                    <el-descriptions-item label="关注者数量">
                        {{ detailFrom.followersCount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="关注数量">
                        {{ detailFrom.followingCount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="总推文数">
                        {{ detailFrom.tweetCount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="点赞数量">
                        {{ detailFrom.likesCount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="头像URL">
                            <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailFrom.profileImageUrl)" :src="getUrl(detailFrom.profileImageUrl)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="封面图片URL">
                            <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailFrom.coverImageUrl)" :src="getUrl(detailFrom.coverImageUrl)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="最后一次爬取时间">
                        {{ detailFrom.lastScraped }}
                    </el-descriptions-item>
                    <el-descriptions-item label="权重">
                        {{ detailFrom.weight }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createTwitterUsers,
  deleteTwitterUsers,
  deleteTwitterUsersByIds,
  updateTwitterUsers,
  findTwitterUsers,
  getTwitterUsersList
} from '@/api/Twitter/twitterUsers'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox,ElLoading } from 'element-plus'
import { ref, reactive } from 'vue'
import axios from 'axios';

// 调用第三方接口的函数
const fetchThirdPartyData = async () => {
  ElMessageBox.confirm('确定要重置用户吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
      let loadingInstance = null; // 提前定义变量
            try {
              // 启动加载动画
              loadingInstance = ElLoading.service({
                lock: true,
                text: '加载中...', // 提示文本
                background: 'rgba(0, 0, 0, 0.7)', // 半透明背景
              });
              axios.get('http://167.71.204.201:5000/UserVieWs', {}).then((response) => {
                  // 确保在请求成功后执行后续逻辑
                  ElMessage({
                    type: 'success',
                    message: response.data,
                  });
                  onReset()
                  loadingInstance.close();
                })

            } catch (error) {
                // 显示错误消息
                ElMessage({
                  type: 'error',
                  message: `接口调用失败: ${error.message}`,
                });
                loadingInstance.close();
            }
    
      })

};

defineOptions({
    name: 'TwitterUsers'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(true)

// 自动化生成的字典（可能为空）以及字段
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

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getTwitterUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTwitterUsersFunc(row)
        })
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteTwitterUsersByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTwitterUsersFunc = async(row) => {
    const res = await findTwitterUsers({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTwitterUsersFunc = async (row) => {
    const res = await deleteTwitterUsers({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTwitterUsers({ id: row.id })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>