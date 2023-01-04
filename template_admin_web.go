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
          <el-button type="primary" @click="handleEdit($event)">新增</el-button>
        </el-form-item>
      </el-form>
    </div>

      <!--列表-->
    <div ref="mainContent" :style="{ height: mainHeight }">
      <el-table
        v-loading="loading"
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
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
        <template #empty>
          <el-empty />
        </template>
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

    <edit ref="editRef" @success="list" />

  </section>
</template>

<script>
  import Edit from "./components/Edit";
  import mainHeight from "@/utils/mainHeight";
  export default {
    name: "{{.Name}}",
    components: { Edit },
    data() {
      return {
        mainHeight: "300px",
        results: [],
        loading: true,
        page: {},
        filters: {},
        selectedRows: [],
      }
    },
    mounted() {
      mainHeight(this);
      this.list();
    },
    methods: {
      async list() {
		    this.loading = true;
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
          this.loading = false
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
      handleSelectionChange(val) {
        this.selectedRows = val
      },
      handleEdit(row) {
        if (row.id) {
          this.$refs.editRef.showEdit(row);
        } else {
          this.$refs.editRef.showEdit();
        }
      },
    }
  }
</script>

<style lang="scss" scoped></style>

`))

var viewEditTmpl = template.Must(template.New("edit.vue").Parse(`
<template>
  <el-dialog
    :visible.sync="dialogFormVisible"
    :close-on-click-modal="false"
    :title="title"
    @close="close"
  >
    <el-form ref="formRef" :model="form" label-width="80px">
      {{range .Fields}}
      <el-form-item label="{{.CamelName}}" prop="{{.CamelName}}">
        <el-input v-model="form.{{.CamelName}}"></el-input>
      </el-form-item>
      {{end}}
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="close">取消</el-button>
      <el-button type="primary" :loading="loading" @click="save">提交</el-button>
    </div>
  </el-dialog>
</template>

<script>
export default {
  data() {
    return {
      title: "",
      form: {},
      rules: {
        // type: [{ required: true, trigger: "blur", message: "请选择类型" }],
        // name: [{ required: true, trigger: "blur", message: "请输入名称" }],
      },
      dialogFormVisible: false,
      loading: false,
    };
  },
  methods: {
    async showEdit(row) {
      if (!row) {
        this.title = "添加";
      } else {
        this.title = "编辑";
        this.form = await this.axios.get("/api/admin/{{.KebabName}}/" + row.id);
      }
      this.dialogFormVisible = true;
    },
    close() {
      this.$refs.formRef.resetFields();
      this.form = {};
      this.dialogFormVisible = false;
    },
    save() {
      this.$refs.formRef.validate(async (valid) => {
        if (valid) {
          let url;
          if (this.title === "添加") {
            url = "/api/admin/{{.KebabName}}/create";
          } else if (this.title === "编辑") {
            url = "/api/admin/{{.KebabName}}/update";
          }

          const params = {
            ...this.form,
          };
          try {
            await this.axios.form(url, params);
            this.$message.success("保存成功");
            this.$emit("success");
            this.close();
          } catch (e) {
            this.$notify.error({ title: "错误", message: e.message || e });
          }
        }
      });
    },
  },
};
</script>

<style lang="scss" scoped></style>

`))
