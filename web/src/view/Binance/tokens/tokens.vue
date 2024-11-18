
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
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
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
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
        
          <el-table-column align="left" label="ID" prop="id" width="120" />
          <el-table-column align="left" label="地址" prop="address" width="360" />
          <el-table-column align="left" label="标识" prop="chainName" width="120" />
          <el-table-column align="left" label="市值" prop="marketCap" width="120" />
          <el-table-column align="left" label="isMeme" prop="isMeme" width="120" />
          <el-table-column align="left" label="ok上线" prop="okx" width="120" />
          <el-table-column align="left" label="gate上线" prop="gate" width="120" />
          <el-table-column align="left" label="bybit上线" prop="bybit" width="120" />
          <el-table-column align="left" label="kucoin上线" prop="kucoin" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateTokensFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="地址:"  prop="address" >
              <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
            </el-form-item>
            <el-form-item label="标识:"  prop="chainName" >
              <el-input v-model="formData.chainName" :clearable="true"  placeholder="请输入标识" />
            </el-form-item>
            <el-form-item label="市值:"  prop="marketCap" >
              <el-input v-model.number="formData.marketCap" :clearable="true" placeholder="请输入市值" />
            </el-form-item>
            <el-form-item label="isMeme:"  prop="isMeme" >
              <el-input v-model.number="formData.isMeme" :clearable="true" placeholder="请输入isMeme" />
            </el-form-item>
            <el-form-item label="ok上线:"  prop="okx" >
              <el-input v-model.number="formData.okx" :clearable="true" placeholder="请输入ok上线" />
            </el-form-item>
            <el-form-item label="gate上线:"  prop="gate" >
              <el-input v-model.number="formData.gate" :clearable="true" placeholder="请输入gate上线" />
            </el-form-item>
            <el-form-item label="bybit上线:"  prop="bybit" >
              <el-input v-model.number="formData.bybit" :clearable="true" placeholder="请输入bybit上线" />
            </el-form-item>
            <el-form-item label="kucoin上线:"  prop="kucoin" >
              <el-input v-model.number="formData.kucoin" :clearable="true" placeholder="请输入kucoin上线" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="ID">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="地址">
                        {{ detailFrom.address }}
                    </el-descriptions-item>
                    <el-descriptions-item label="标识">
                        {{ detailFrom.chainName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="市值">
                        {{ detailFrom.marketCap }}
                    </el-descriptions-item>
                    <el-descriptions-item label="isMeme">
                        {{ detailFrom.isMeme }}
                    </el-descriptions-item>
                    <el-descriptions-item label="ok上线">
                        {{ detailFrom.okx }}
                    </el-descriptions-item>
                    <el-descriptions-item label="gate上线">
                        {{ detailFrom.gate }}
                    </el-descriptions-item>
                    <el-descriptions-item label="bybit上线">
                        {{ detailFrom.bybit }}
                    </el-descriptions-item>
                    <el-descriptions-item label="kucoin上线">
                        {{ detailFrom.kucoin }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createTokens,
  deleteTokens,
  deleteTokensByIds,
  updateTokens,
  findTokens,
  getTokensList
} from '@/api/Binance/tokens'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'Tokens'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
  const table = await getTokensList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteTokensFunc(row)
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
      const res = await deleteTokensByIds({ ids })
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
const updateTokensFunc = async(row) => {
    const res = await findTokens({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTokensFunc = async (row) => {
    const res = await deleteTokens({ id: row.id })
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
        address: '',
        chainName: '',
        marketCap: undefined,
        isMeme: undefined,
        okx: undefined,
        gate: undefined,
        bybit: undefined,
        kucoin: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
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
  const res = await findTokens({ id: row.id })
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