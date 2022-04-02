<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Twitter名">
          <el-input v-model="searchInfo.userName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
<!--        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>-->
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="推特id" prop="userId" width="120" />
        <el-table-column align="left" label="Twitter名称" prop="userName" width="120" />
        <el-table-column align="left" label="类别" prop="userTypeName" width="120" />
        <el-table-column align="left" label="头像url" :show-overflow-tooltip="true" prop="avatarUrl" width="120" />
        <el-table-column align="left" label="背景url" :show-overflow-tooltip="true" prop="bannerUrl" width="120" />
        <el-table-column align="left" label="个人介绍" :show-overflow-tooltip="true" prop="biography" width="120" />
        <el-table-column align="left" label="推特地址" :show-overflow-tooltip="true" prop="websiteUrl" width="120" />
        <el-table-column align="left" label="位置" prop="location" width="120" />
        <el-table-column align="left" label="出生日期" prop="birthday" width="120" />
        <el-table-column align="left" label="喜欢数" prop="likesCount" width="120" />
        <el-table-column align="left" label="粉丝数" prop="followersCount" width="120" />
        <el-table-column align="left" label="追随数" prop="followingCount" width="120" />
        <el-table-column align="left" label="好友数" prop="friendsCount" width="120" />
        <el-table-column align="left" label="推文数" prop="tweetsCount" width="120" />
        <el-table-column align="left" label="是否私有" prop="isPrivate" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isPrivate) }}</template>
        </el-table-column>
        <el-table-column align="left" label="是否校验" prop="isVerified" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isVerified) }}</template>
        </el-table-column>
        <el-table-column align="left" label="加入日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.joinedTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="创建日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
<!--        <el-table-column align="left" label="listedCount字段" prop="listedCount" width="120" />-->
        <el-table-column align="left" label="操作" width="150">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateTTwitterUserFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户头像url:">
          <el-input v-model="formData.avatarUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="背景url:">
          <el-input v-model="formData.bannerUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="个人介绍:">
          <el-input v-model="formData.biography" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出生日期:">
          <el-input v-model="formData.birthday" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="粉丝数:">
          <el-input v-model.number="formData.followersCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="追随数:">
          <el-input v-model.number="formData.followingCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="好友数:">
          <el-input v-model.number="formData.friendsCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否私有:">
          <el-switch v-model="formData.isPrivate" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="是否校验:">
          <el-switch v-model="formData.isVerified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="加入日期:">
          <el-date-picker v-model="formData.joinedTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="喜欢数:">
          <el-input v-model.number="formData.likesCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="listedCount:">
          <el-input v-model.number="formData.listedCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="location:">
          <el-input v-model="formData.location" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="推文数:">
          <el-input v-model.number="formData.tweetsCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:">
          <el-input v-model="formData.userId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名称:">
          <el-input v-model="formData.userName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="外键id:">
          <el-input v-model="formData.userTypeId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="个人推特地址:">
          <el-input v-model="formData.websiteUrl" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'TTwitterUser'
}
</script>

<script setup>
import {
  createTTwitterUser,
  deleteTTwitterUser,
  deleteTTwitterUserByIds,
  updateTTwitterUser,
  findTTwitterUser,
  getTTwitterUserList
} from '@/api/tTwitterUser'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { getTSystemAll } from '@/api/tSystem'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  avatarUrl: '',
  bannerUrl: '',
  biography: '',
  birthday: '',
  followersCount: 0,
  followingCount: 0,
  friendsCount: 0,
  isPrivate: false,
  isVerified: false,
  joinedTime: new Date(),
  likesCount: 0,
  listedCount: 0,
  location: '',
  tweetsCount: 0,
  userId: '',
  userName: '',
  userTypeId: '',
  websiteUrl: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const systemList = ref([])

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.isPrivate === '') {
    searchInfo.value.isPrivate = null
  }
  if (searchInfo.value.isVerified === '') {
    searchInfo.value.isVerified = null
  }
  getTableData()
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
  const table = await getTTwitterUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    await getSystemData()
    table.data.list.forEach((item, i) => {
      systemList.value.forEach((tsystem) => {
        if (item.userTypeId === tsystem.key) {
          item.userTypeName = tsystem.value
          return
        }
      })
    })
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
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
    deleteTTwitterUserFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
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
          ids.push(item.ID)
        })
  const res = await deleteTTwitterUserByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTTwitterUserFunc = async(row) => {
  const res = await findTTwitterUser({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.retTwitterUser
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteTTwitterUserFunc = async(row) => {
  const res = await deleteTTwitterUser({ ID: row.ID })
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
    avatarUrl: '',
    bannerUrl: '',
    biography: '',
    birthday: '',
    followersCount: 0,
    followingCount: 0,
    friendsCount: 0,
    isPrivate: false,
    isVerified: false,
    joinedTime: new Date(),
    likesCount: 0,
    listedCount: 0,
    location: '',
    tweetsCount: 0,
    userId: '',
    userName: '',
    userTypeId: '',
    websiteUrl: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createTTwitterUser(formData.value)
      break
    case 'update':
      res = await updateTTwitterUser(formData.value)
      break
    default:
      res = await createTTwitterUser(formData.value)
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
}

const getSystemData = async() => {
  const rs = await getTSystemAll()
  if (rs.code === 0) {
    systemList.value = rs.data
  }
}
</script>

<style>
</style>
