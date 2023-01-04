<template>
  <section class="page-container">
    <div ref="toolbar" class="toolbar">
      <el-form :inline="true" :model="filters">
        <el-form-item>
          <el-input v-model="filters.name" placeholder="名称"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="list">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleEdit($event)">新增</el-button>
        </el-form-item>
      </el-form>
    </div>

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

        <el-table-column prop="name" label="name"></el-table-column>

        <el-table-column prop="eventType" label="eventType"></el-table-column>

        <el-table-column prop="startTime" label="startTime"></el-table-column>

        <el-table-column prop="endTime" label="endTime"></el-table-column>

        <el-table-column prop="coverImages" label="coverImages"></el-table-column>

        <el-table-column prop="content" label="content"></el-table-column>

        <el-table-column prop="organizer" label="organizer"></el-table-column>

        <el-table-column prop="status" label="status"></el-table-column>

        <el-table-column prop="updateTime" label="updateTime"></el-table-column>

        <el-table-column prop="createTime" label="createTime"></el-table-column>

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

    <div ref="pagebar" class="pagebar">
      <el-pagination
        :page-sizes="[20, 50, 100, 300]"
        :current-page="page.page"
        :page-size="page.limit"
        :total="page.total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handlePageChange"
        @size-change="handleLimitChange"
      >
      </el-pagination>
    </div>

    <edit ref="editRef" @success="list" />
  </section>
</template>

<script>
import Edit from "./components/Edit";
import mainHeight from "@/utils/mainHeight";
export default {
  components: { Edit },
  data() {
    return {
      mainHeight: "300px",
      results: [],
      listLoading: false,
      page: {},
      filters: {},
      selectedRows: [],
    };
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
      });
      try {
        const data = await this.axios.form("/api/admin/event/list", params);
        this.results = data.results;
        this.page = data.page;
      } catch (e) {
        this.$notify.error({ title: "错误", message: e.message || e });
      } finally {
        this.listLoading = false;
      }
    },
    async handlePageChange(val) {
      this.page.page = val;
      await this.list();
    },
    async handleLimitChange(val) {
      this.page.limit = val;
      await this.list();
    },
    handleSelectionChange(val) {
      this.selectedRows = val;
    },
    handleEdit(row) {
      if (row.id) {
        this.$refs.editRef.showEdit(row);
      } else {
        this.$refs.editRef.showEdit();
      }
    },
  },
};
</script>

<style lang="scss" scoped></style>
