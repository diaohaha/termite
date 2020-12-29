<template>
	<div>
		<div class="container">

			<div class="handle-box">
				<el-form :inline="true">
					<el-form-item>
						<el-input v-model="query.cid" placeholder="cid" class="handle-input" style="width: 300px"></el-input>
					</el-form-item>
					<el-form-item>
						<el-checkbox-button :indeterminate="isIndeterminate" v-model="checkAll" @change="handleCheckAllChange">全选</el-checkbox-button>
					</el-form-item>
					<el-form-item>
						<el-checkbox-group v-model="checkedstates" @change="handleCheckedStatesChange">
							<el-checkbox-button v-for="state in states" :label="state" :key="state[0]">{{state[1]}}</el-checkbox-button>
						</el-checkbox-group>
					</el-form-item>
					<el-form-item>
						<el-button type="primary" icon="el-icon-search" @click="handleSearch">search</el-button>
					</el-form-item>
				</el-form>
			</div>
			<el-table
				:data="tableData"
				style="width: 100%"
				:row-class-name="tableRowClassName"
				border
				class="table"
				ref="flowInsTable"
				header-cell-class-name="table-header"
				row-key="flow_id"
				:expand-row-keys="expands"
				@expand-change="expandContext"
			>
				<el-table-column type="expand" @click="expandContext(row)">
					<template slot-scope="props" style="font-size: 10px">
						<el-form label-position="left" inline class="demo-table-expand">
							<el-row type="flex" class="row-bg">
								<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">创建时间:</span></div></el-col>
								<el-col :span="4"><div class="grid-content bg-purple"><span>{{ props.row.create_time }}</span></div></el-col>
								<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">更新时间:</span></div></el-col>
								<el-col :span="4"><div class="grid-content bg-purple"><span>{{ props.row.update_time }}</span></div></el-col>
								<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">对象ID:</span></div></el-col>
								<el-col :span="4"><div class="grid-content bg-purple"><span>{{ props.row.cid }}</span></div></el-col>
							</el-row>
							<el-row type="flex" class="row-bg">
								<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">工作流上下文:</span></div></el-col>
								<el-col :span="22"><div class="grid-content bg-purple"><pre>{{ props.row.context }}</pre></div></el-col>
							</el-row>
							<el-divider content-position="left"><div style="font-size: 12px; font-weight: lighter">任务列表</div></el-divider>
							<el-table
								:data="subTableData"
								style="width: 100%"
								:row-class-name="workTableRowClassName"
								class="table"
								header-cell-class-name="table-header"
							>
								<el-table-column type="expand" >
									<template slot-scope="subprops">
										<el-form label-position="left" inline class="demo-table-expand">
											<el-row type="flex" class="row-bg" justify="space-around">
												<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">创建时间:</span></div></el-col>
												<el-col :span="4"><div class="grid-content bg-purple"><span>{{ subprops.row.create_time }}</span></div></el-col>
												<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">更新时间:</span></div></el-col>
												<el-col :span="4"><div class="grid-content bg-purple"><span>{{ subprops.row.update_time }}</span></div></el-col>
												<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">对象ID:</span></div></el-col>
												<el-col :span="4"><div class="grid-content bg-purple"><span>{{ subprops.row.cid }}</span></div></el-col>
												<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">任务名称:</span></div></el-col>
												<el-col :span="4"><div class="grid-content bg-purple"><span>{{ subprops.row.vwork }}</span></div></el-col>
											</el-row>
											<el-row type="flex" class="row-bg" justify="space-around">
												<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">任务输出:</span></div></el-col>
												<el-col :span="22"><div class="grid-content bg-purple"><pre>{{ subprops.row.output }}</pre></div></el-col>
											</el-row>
											<el-row type="flex" class="row-bg" justify="space-around">
												<el-col :span="2"><div class="grid-content bg-purple"><span style="font-weight: lighter">异常信息:</span></div></el-col>
												<el-col :span="22"><div class="grid-content bg-purple"><pre>{{ subprops.row.error }}</pre></div></el-col>
											</el-row>
										</el-form>
									</template>
								</el-table-column>
								<el-table-column prop="vwork" label="vwork"></el-table-column>
								<el-table-column prop="create_time" label="create_time"></el-table-column>
								<el-table-column prop="update_time" label="update_time"></el-table-column>
								<el-table-column prop="cid" label="cid"></el-table-column>
								<el-table-column prop="state" label="state"></el-table-column>
								<el-table-column label="operation">
									<template slot-scope="scope">
										<el-tooltip class="item" effect="dark" content="recover" placement="top">
											<el-button
												icon="el-icon-refresh"
												circle
												:disabled="(scope.row.state == '初始化' || scope.row.state == '调度延期' || scope.row.state == '成功')"
												@click="recoverWork(scope.row)"
											></el-button>
										</el-tooltip>
									</template>
								</el-table-column>
							</el-table>
						</el-form>
					</template>
				</el-table-column>
				<el-table-column prop="flow_id" label="flow_id"></el-table-column>
				<el-table-column prop="create_time" label="create_time"></el-table-column>
				<el-table-column prop="update_time" label="update_time"></el-table-column>
				<el-table-column prop="cid" label="cid"></el-table-column>
				<el-table-column prop="state" label="state"></el-table-column>
				<el-table-column label="operation">
					<template slot-scope="scope">
						<el-tooltip class="item" effect="dark" content="recover" placement="top">
							<el-button
								icon="el-icon-refresh"
								circle
								:disabled="!(scope.row.state == '失败' || scope.row.state == '异常' || scope.row.state == '超时')"
								@click="recoverFlow(scope.row)"
							></el-button>
						</el-tooltip>
						<el-tooltip class="item" effect="dark" content="delete" placement="top">
							<el-button
								icon="el-icon-delete"
								circle
								@click="deleteFlow(scope.row)"
							></el-button>
						</el-tooltip>
					</template>
				</el-table-column>
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
	/*.el-form-item {*/
	/*	margin-right: 0;*/
	/*	margin-bottom: 0;*/
	/*	width: 50%;*/
	/*}*/
	/*.el-form-item--small.el-form-item {*/
	/*	margin-bottom: 5px;*/
	/*}*/
	/*.el-form-item--small .el-form-item__content, .el-form-item--small .el-form-item__label {*/
	/*	line-height: 22px;*/
	/*}*/
</style>

<script>
    import { fetchFlowInstances } from "../../api/index";
    import { fetchWorkInstances } from "../../api/index";
    import { recoverFlowInstances } from "../../api/index";
    import { deleteFlowInstances } from "../../api/index";
    import { recoverWorkInstances } from "../../api/index";
    Array.prototype.remove = function (val) {
        let index = this.indexOf(val);
        if (index > -1) {
            this.splice(index, 1);
        }
    };
    const stateOptions = [
		[0, "待调度"],
		[1, "调度中"],
		[2, "完成"],
		[3, "超时"],
		[4, "异常"],
		[5, "失败"],
		[6, "调度延期"]
	]
    export default {
        name: "flowinstance",
        data() {
            return {
                checkAll: true,
                checkedstates: [],
                states: stateOptions,
                isIndeterminate: true,
                pageTotal: 0,
                tableData: [],
				subTableData: [],
                query: {
                    cid: "",
					vflow: this.$route.params.flowKey,
					vstates:  [],
					page_index: 1,
					page_size: 15,
                },
				workInstancesQuery: {
                    cid: "",
					vflow: this.$route.params.flowKey,
				},
                recoverFlowBody: {
                    flow_ids: [],
				},
                deleteFlowBody: {
                    flow_ids: [],
                },
                recoverWorkBody: {
                    work_ids: [],
                },
                expands: [],
            };
        },
        methods: {
            handleCheckAllChange(val) {
                this.checkedstates = val ? stateOptions : [];
                this.isIndeterminate = false;
            },
            handleCheckedStatesChange(value) {
                console.log(this.states)
                console.log(this.checkedstates)
                let checkedCount = value.length;
                this.checkAll = checkedCount === this.states.length;
                this.isIndeterminate = checkedCount > 0 && checkedCount < this.states.length;
            },
            workTableRowClassName({row, rowIndex}) {
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
            tableRowClassName({row, rowIndex}) {
                if (row.state === "超时") {
                    return 'warning-row';
                } else if (row.state === "异常") {
					return 'failed-row';
                } else if (row.state === "失败") {
                    return 'warning-row';
                } else if (row.state === "待调度") {
                    return 'success-row';
                } else if (row.state === "调度中") {
                    return 'success-row';
                } else if (row.state === "完成") {
                    return 'success-row';
                } else if (row.state === "调度延期") {
                    return 'success-row';
                }
                return '';
            },
            expandContext(row) {
                console.log("expand row")
				this.subTableData = []
                this.workInstancesQuery.cid = row.cid
                this.getWorkInstanceData()
                // this.$refs.flowInsTable.toggleRowExpansion(row)
                // this.expands = []
                if (this.expands.indexOf(row.flow_id) < 0) {
                    this.expands = []
                    this.expands.push(row.flow_id);
                    console.log("if:", this.expands)
                } else {
                    this.expands.remove(row.flow_id);
                    console.log("else:", this.expands)
                }
			},
            getFlowInstanceData() {
                console.log(this.$route.params.flowKey)
                fetchFlowInstances(this.query).then(res => {
					this.tableData = res.data.flows;
					this.pageTotal = res.data.count;
                });
            },
            getWorkInstanceData() {
                fetchWorkInstances(this.workInstancesQuery).then(res => {
                    this.subTableData = res.data.works;
                });
            },
            recoverWork(row) {
                console.log("recover work")
                this.recoverWorkBody.work_ids = [row.work_id]
                recoverWorkInstances(this.recoverWorkBody).then(res => {
                    if (res.code == "A0001" && res.data.success > 0) {
                        this.$message.success("任务重置成功")
                    }
                    this.getWorkInstanceData()
                })
            },
            deleteFlow(row) {
                this.deleteFlowBody.flow_ids = [row.flow_id]
                deleteFlowInstances(this.deleteFlowBody).then(res => {
                    if (res.code == "A0001" && res.data.success > 0) {
                        this.$message.success("任务流删除成功")
                    }
                    this.getFlowInstanceData()
                })
            },
            recoverFlow(row) {
                this.recoverFlowBody.flow_ids = [row.flow_id]
                recoverFlowInstances(this.recoverFlowBody).then(res => {
                    if (res.code == "A0001" && res.data.success > 0) {
                        this.$message.success("任务流重置成功")
					}
                    this.getFlowInstanceData()
				})
			},
            // 触发搜索按钮
            handleSearch() {
                this.$set(this.query, 'page_index', 1);
                let vstates = []
                for (let i = 0; i < this.checkedstates.length; i++) {
                    vstates.push(this.checkedstates[i][0])
				}
                this.$set(this.query, 'vstates', vstates);
                this.getFlowInstanceData();
            },
            // 分页导航
            handlePageChange(val) {
                this.$set(this.query, 'page_index', val);
                this.getFlowInstanceData();
            }
        },
        created() {
            console.log("im create")
            this.getFlowInstanceData();
        }
    };
</script>
