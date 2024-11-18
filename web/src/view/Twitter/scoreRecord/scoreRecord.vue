
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="条件">
            <el-input
              v-model="searchInfo.query"
              placeholder="条件"
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
            <el-button  type="primary" icon="plus" @click="openDialog">热度查询</el-button>
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
          <el-table-column align="left" label="条件" prop="query" width="540" />
          <el-table-column align="left" label="总分" prop="scoreAdded" width="220" />
         <el-table-column align="left" label="日期" prop="timestamp" width="180">
            <template #default="scope">{{ formatDate(scope.row.timestamp) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
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
                <span class="text-lg">推特搜索</span>
              </div>
            </template>
            <el-form :model="formData" label-width="100px">
            <div v-for="(item, index) in formFields" :key="index" class="field-group">
              <el-form-item :label="'条件' + (index + 1)">
                <el-input v-model="item.value" placeholder="请输入条件" style="width: 300px;margin-right: 20px;"/> <el-button type="danger" @click="removeField(index)">移除</el-button>
              </el-form-item>
             
            </div>
            <el-button type="primary" @click="addField">添加字段</el-button>
            <el-button type="success" @click="submitForm">提交</el-button>
          </el-form>

    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="ID">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="条件">
                        {{ detailFrom.query }}
                    </el-descriptions-item>
                    <el-descriptions-item label="总分">
                        {{ detailFrom.scoreAdded }}
                    </el-descriptions-item>
                    <el-descriptions-item label="日期">
                        {{ detailFrom.timestamp }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createScoreRecord,
  deleteScoreRecord,
  deleteScoreRecordByIds,
  updateScoreRecord,
  findScoreRecord,
  getScoreRecordList
} from '@/api/Twitter/scoreRecord'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox ,ElLoading} from 'element-plus'
import { ref, reactive } from 'vue'
import axios from 'axios';


const formFields = reactive([{ value: "" }]);

const addField = () => {
  formFields.push({ value: "" });
};

const removeField = (index) => {
  formFields.splice(index, 1);
};

const submitForm = () => {
  // 校验是否所有字段都填写了值
  const hasEmptyField = formFields.some((field) => field.value.trim() === "");
  if (hasEmptyField) {
    ElMessage({
          type: 'warning',
          message: '请移除无用的条件'
        })
    return;
  }
  if(formFields.length>=1){

      let loadingInstance = null; // 提前定义变量
      const queryString = formFields.map((field) => `querys=${encodeURIComponent(field.value)}`)
        .join('&');
      // 启动加载动画
      loadingInstance = ElLoading.service({
          lock: true,
          text: '加载中...', // 提示文本
          background: 'rgba(0, 0, 0, 0.7)', // 半透明背景
      });
      axios.get(`http://167.71.204.201:5000/TwitterViews?${queryString}`
      ).then((response) => {
      // 确保在请求成功后执行后续逻辑
        ElMessage({
          type: 'success',
          message: response.data,
        });
        onReset()
      }) .catch((error) => {
      ElMessage({
        type: 'error',
        message: `接口调用失败: ${error.message}`,
      });
    })
    .finally(() => {
      // 确保在无论成功还是失败时关闭加载动画
      loadingInstance.close();
    });

  }

};

defineOptions({
    name: 'ScoreRecord'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(true)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            query: '',
            scoreAdded: 0,
            timestamp: new Date(),
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
  const table = await getScoreRecordList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteScoreRecordFunc(row)
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
      const res = await deleteScoreRecordByIds({ ids })
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




// 删除行
const deleteScoreRecordFunc = async (row) => {
    const res = await deleteScoreRecord({ id: row.id })
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
        query: '',
        scoreAdded: 0,
        timestamp: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createScoreRecord(formData.value)
                  break
                case 'update':
                  res = await updateScoreRecord(formData.value)
                  break
                default:
                  res = await createScoreRecord(formData.value)
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
  const res = await findScoreRecord({ id: row.id })
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