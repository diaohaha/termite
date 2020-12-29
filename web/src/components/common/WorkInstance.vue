<template>
	<div>
		<div class="container">

			<div class="handle-box">
				<el-input v-model="query.cid" placeholder="cid" class="handle-input" style="width: 300px"></el-input>
				<el-button type="primary" icon="el-icon-search" @click="handleSearch">search</el-button>
			</div>
			<br>
			<el-table
				:data="tableData"
				style="width: 100%"
				border
				:row-class-name="tableRowClassName"
				class="table"
				ref="multipleTable"
				header-cell-class-name="table-header"
				@row-click="expandContext"
				row-key="work_id"
				:expand-row-keys="expands"
			>
				<el-table-column type="expand" >
					<template slot-scope="props">
						<el-row type="flex" class="row-bg" justify="space-around">
							<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">创建时间:</span></div></el-col>
							<el-col :span="4"><div class="grid-content bg-purple"><span>{{ props.row.create_time }}</span></div></el-col>
							<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">更新时间:</span></div></el-col>
							<el-col :span="4"><div class="grid-content bg-purple"><span>{{ props.row.update_time }}</span></div></el-col>
							<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">对象ID:</span></div></el-col>
							<el-col :span="4"><div class="grid-content bg-purple"><span>{{ props.row.cid }}</span></div></el-col>
							<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">任务名称:</span></div></el-col>
							<el-col :span="4"><div class="grid-content bg-purple"><span>{{ props.row.vwork }}</span></div></el-col>
						</el-row>
						<el-row type="flex" class="row-bg" justify="space-around">
							<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">任务输出:</span></div></el-col>
							<el-col :span="22"><div class="grid-content bg-purple"><pre>{{ props.row.output }}</pre></div></el-col>
						</el-row>
						<el-row type="flex" class="row-bg" justify="space-around">
							<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">异常信息:</span></div></el-col>
							<el-col :span="22"><div class="grid-content bg-purple"><pre>{{ props.row.error }}</pre></div></el-col>
						</el-row>
<!--						<el-form label-position="left" inline class="demo-table-expand">-->
<!--							<el-form-item label="对象ID">-->
<!--								<span>{{ props.row.cid }}</span>-->
<!--							</el-form-item>-->
<!--							<el-form-item label="任务名称">-->
<!--								<span>{{ props.row.vwork }}</span>-->
<!--							</el-form-item>-->
<!--							<el-form-item label="所属工作流">-->
<!--								<span>{{ props.row.vflow }}</span>-->
<!--							</el-form-item>-->
<!--							<el-form-item label="创建时间">-->
<!--								<span>{{ props.row.create_time }}</span>-->
<!--							</el-form-item>-->
<!--							<el-form-item label="更新时间">-->
<!--								<span>{{ props.row.update_time }}</span>-->
<!--							</el-form-item>-->
<!--							<el-form-item label="任务输出">-->
<!--								<pre>{{ props.row.output }}</pre>-->
<!--							</el-form-item>-->
<!--							<el-form-item label="异常信息">-->
<!--								<pre>{{ props.row.error }}</pre>-->
<!--							</el-form-item>-->
<!--						</el-form>-->
					</template>
				</el-table-column>
				<el-table-column prop="work_id" label="work_id"></el-table-column>
				<el-table-column prop="create_time" label="create_time"></el-table-column>
				<el-table-column prop="update_time" label="update_time"></el-table-column>
				<el-table-column prop="cid" label="cid"></el-table-column>
				<el-table-column prop="state" label="state"></el-table-column>
			</el-table>
			<div class="pagination">
				<el-pagination
					background
					layout="total, prev, pager, next"
					:current-page="query.page_index"
					:page-size="query.page_size"
					:total="pageTotal"
					@current-change="handlePageChange"
				></el-pagination>
			</div>
		</div>
	</div>
</template>

<style>
	.el-table .warning-row {
		background: oldlace;
	}
	.el-table .failed-row {
		background: #ff8888;
	}

	.el-table .success-row {
		background: #f0f9eb;
	}
</style>

<script>
    import { fetchWorkInstances } from "../../api/index";
    export default {
        name: "workInstance",
        data() {
            return {
                pageTotal: 0,
                tableData: [],
                query: {
                    cid: "",
                    vwork: this.$route.params.workKey,
                    page_index: 1,
                    page_size: 15,
                },
                expands: [],
            };
        },
        methods: {
            tableRowClassName({row, rowIndex}) {
                console.log("row is:", row)
                if (row.state === "超时") {
                    return 'warning-row';
                } else if (row.state === "异常") {
                    return 'failed-row';
                } else if (row.state === "失败") {
                    return 'warning-row';
                } else if (row.state === "初始化") {
                    return 'success-row';
                } else if (row.state === "已下发") {
                    return 'success-row';
                } else if (row.state === "执行中") {
                    return 'success-row';
                } else if (row.state === "成功") {
                    return 'success-row';
                } else if (row.state === "延期调度") {
                    return 'success-row';
                }
                return '';
            },
            expandContext(row, column, event) {
                Array.prototype.remove = function (val) {
                    let index = this.indexOf(val);
                    if (index > -1) {
                        this.splice(index, 1);
                    }
                };
                this.expands = []
                if (this.expands.indexOf(row.flow_id) < 0) {
                    this.expands.push(row.flow_id);
                } else {
                    this.expands.remove(row.flow_id);
                }

                if (this.expands.indexOf(row.flow_id) < 0) {
                    this.expands.push(row.flow_id);
                } else {
                    this.expands.remove(row.flow_id);
                }
            },
            getWorkInstanceData() {
                console.log(this.$route.params.workkey)
                fetchWorkInstances(this.query).then(res => {
                    this.tableData = res.data.works;
                    this.pageTotal = res.data.count;
                });
            },
            // 触发搜索按钮
            handleSearch() {
                this.$set(this.query, 'page_index', 1);
                this.getWorkInstanceData();
            },
            // 分页导航
            handlePageChange(val) {
                this.$set(this.query, 'page_index', val);
                this.getWorkInstanceData();
            }
        },
        created() {
            this.getWorkInstanceData();
        }
    };
</script>
