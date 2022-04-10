package codegen

import "html/template"

var viewIndexTmpl = template.Must(template.New("index.vue").Parse(`
<template>
    <section class="page-container">
        <!--工具条-->
        <div ref="toolbar" class="toolbar">
            <el-form :inline="true" :model="filters">
                <el-form-item>
                    <el-input v-model="filters.name" placeholder="名称"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" v-on:click="list">查询</el-button>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="handleAdd">新增</el-button>
                </el-form-item>
            </el-form>
        </div>

        <!--列表-->
		<div ref="mainContent" :style="{ height: mainHeight }">
			<el-table
				v-loading="listLoading"
				height="100%"
				:data="results"
				highlight-current-row
				stripe
				border
				@selection-change="handleSelectionChange"
			>
				<el-table-column type="selection" width="55"></el-table-column>
				<el-table-column prop="id" label="编号"></el-table-column>
				{{range .Fields}}
				<el-table-column prop="{{.CamelName}}" label="{{.CamelName}}"></el-table-column>
				{{end}}
				<el-table-column label="操作" width="150">
					<template slot-scope="scope">
						<el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
					</template>
				</el-table-column>
			</el-table>
		</div>

        <!--工具条-->
        <div ref="pagebar" class="pagebar">
            <el-pagination
                           @current-change="handlePageChange"
                           @size-change="handleLimitChange"
                           :page-sizes="[20, 50, 100, 300]"
                           :current-page="page.page"
                           :page-size="page.limit"
                           :total="page.total"
						   layout="total, sizes, prev, pager, next, jumper">
            </el-pagination>
        </div>

        <!--新增界面-->
        <el-dialog title="新增" :visible.sync="addFormVisible" :close-on-click-modal="false">
            <el-form :model="addForm" label-width="80px" ref="addForm">
                {{range .Fields}}
				<el-form-item label="{{.CamelName}}">
					<el-input v-model="addForm.{{.CamelName}}"></el-input>
				</el-form-item>
                {{end}}
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="addFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
            </div>
        </el-dialog>

        <!--编辑界面-->
        <el-dialog title="编辑" :visible.sync="editFormVisible" :close-on-click-modal="false">
            <el-form :model="editForm" label-width="80px" ref="editForm">
                <el-input v-model="editForm.id" type="hidden"></el-input>
                {{range .Fields}}
				<el-form-item label="{{.CamelName}}">
					<el-input v-model="editForm.{{.CamelName}}"></el-input>
				</el-form-item>
                {{end}}
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="editFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="editSubmit" :loading="editLoading">提交</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
  import mainHeight from "@/utils/mainHeight";
  export default {
    data() {
      return {
		mainHeight: "300px",
        results: [],
        listLoading: false,
        page: {},
        filters: {},
        selectedRows: [],

        addForm: {},
        addFormVisible: false,
        addLoading: false,

        editForm: {},
        editFormVisible: false,
        editLoading: false,
      }
    },
    mounted() {
	  mainHeight(this);
      this.list();
    },
    methods: {
		async list() {
		  const params = Object.assign(this.filters, {
			page: this.page.page,
			limit: this.page.limit,
		  })
		  try {
			const data = await this.axios.form('/api/admin/{{.KebabName}}/list', params)
			this.results = data.results
			this.page = data.page
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e.message || e })
		  } finally {
			this.listLoading = false
		  }
		},
		async handlePageChange(val) {
		  this.page.page = val
		  await this.list()
		},
		async handleLimitChange(val) {
		  this.page.limit = val
		  await this.list()
		},
		handleAdd() {
		  this.addForm = {
			name: '',
			description: '',
		  }
		  this.addFormVisible = true
		},
		async addSubmit() {
		  try {
			await this.axios.form('/api/admin/{{.KebabName}}/create', this.addForm)
			this.$message({ message: '提交成功', type: 'success' })
			this.addFormVisible = false
			await this.list()
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e.message || e })
		  }
		},
		async handleEdit(index, row) {
		  try {
			const data = await this.axios.get('/api/admin/{{.KebabName}}/' + row.id)
			this.editForm = Object.assign({}, data)
			this.editFormVisible = true
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e.message || e })
		  }
		},
		async editSubmit() {
		  try {
			await this.axios.form('/api/admin/{{.KebabName}}/update', this.editForm)
			await this.list()
			this.editFormVisible = false
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e.message || e })
		  }
		},
	
		handleSelectionChange(val) {
		  this.selectedRows = val
		},
    }
  }
</script>

<style lang="scss" scoped>

</style>

`))
