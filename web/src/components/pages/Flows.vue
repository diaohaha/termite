<template>
    <div>
        <div class="container">
			<div class="handle-box">
                <el-input v-model="query.search" placeholder="flow key" class="handle-input mr10"></el-input>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">search</el-button>
				<el-button
					type="primary"
					icon="el-icon-circle-plus"
					class="handle-add mr10"
					style="float: right"
					@click="handleCreate"
				>create</el-button>
            </div>
            <el-table
                :data="tableData"
                border
                class="table"
                ref="multipleTable"
                header-cell-class-name="table-header"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column prop="Key" label="key">
                    <template slot-scope="scope">
                        <a href='javascript:void(0)' @click="jump(scope.row.key)">{{scope.row.key}}</a>
                    </template>
                </el-table-column>
                <el-table-column prop="name" label="name"></el-table-column>
				<el-table-column prop="desc" label="desc"></el-table-column>
				<el-table-column label="operation" width="180" align="center">
                    <template slot-scope="scope">
						<el-button
							type="text"
							icon="el-icon-copy-document"
							@click="handleCopy(scope.$index, scope.row)"
						>copy</el-button>
                        <el-button
                            type="text"
                            icon="el-icon-delete"
                            class="red"
                            @click="handleDelete(scope.$index, scope.row)"
                        >delete</el-button>
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
		<!-- Form -->
		<el-dialog title="创建" :visible.sync="createVisible">
			<FlowDetailComponent ref="flowdetail" v-bind:form="createParams"></FlowDetailComponent>
			<div slot="footer" class="dialog-footer">
				<el-button @click="createVisible = false">取 消</el-button>
				<el-button type="primary" @click="saveCreate">确 定</el-button>
			</div>
		</el-dialog>

		<!-- 复制弹出框 -->
        <el-dialog title="复制" :visible.sync="copyVisible" width="30%">
            <el-form ref="form" :model="form" label-width="70px">
                <el-form-item label="flowKey">
                    <el-input v-model="copyParams.flow_key"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="copyVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveCopy">确 定</el-button>
            </span>
        </el-dialog>
		<!-- 编辑弹出框 -->
		<el-dialog title="编辑" :visible.sync="editVisible" width="30%">
			<el-form ref="form" :model="form" label-width="70px">
				<el-form-item label="flow key">
					<el-input v-model="form.name"></el-input>
				</el-form-item>
			</el-form>
			<span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
		</el-dialog>
	</div>
</template>

<script>
    import FlowDetailComponent from '../common/FlowDetailComponent'
    import { queryFlowConfigs } from '../../api/index';
    import { copyFlowConfig } from '../../api/index';
    import { deleteFlowConfig } from '../../api/index';
    import { createFlowConfig } from '../../api/index';
    export default {
    name: 'iflows',
    data() {
        return {
            query: {
                search: "",
                page_index: 1,
                page_size: 13,
            },
            tableData: [],
            multipleSelection: [],
            delList: [],
            editVisible: false,
            copyVisible: false,
            createVisible: false,
            pageTotal: 0,
            form: {},
            idx: -1,
            id: -1,
			copyParams: {
                src_flow_key: "",
				flow_key: ""
			},
            createParams: {
                flow_name: "",
                flow_key: "",
                flow_desc: "",
                env: "",
                config: ""
            },
			deleteParams: {
                flow_key: ""
			}
        };
    },
    created() {
        this.getData();
    },
	components:{
        FlowDetailComponent
	},
    methods: {
        // 获取 easy-mock 的模拟数据
        getData() {
            queryFlowConfigs(this.query).then(res => {
                this.tableData = res.data.configs;
                this.pageTotal = res.data.count;
            });
        },
		jump(dsturl) {
            console.log("jump", dsturl)
            this.$router.push("iflows/" + dsturl)
		},
        // 触发搜索按钮
        handleSearch() {
            this.$set(this.query, 'page_index', 1);
            this.getData();
        },
        // 删除操作
        handleDelete(index, row) {
            // 二次确认删除
			this.deleteParams.flow_key = row.key
            this.$confirm('确定要删除吗？', '提示', {
                type: 'warning'
            })
                .then(() => {
                    deleteFlowConfig(this.deleteParams).then(res => {
                        if (res.code == "A0001") {
                            this.$message.success('删除成功');
                            this.tableData.splice(index, 1);
						} else {
                            this.$message.error(res.msg)
						}
					})
                })
                .catch(() => {});
        },
        // 多选操作
        handleSelectionChange(val) {
            this.multipleSelection = val;
        },
        delAllSelection() {
            const length = this.multipleSelection.length;
            let str = '';
            this.delList = this.delList.concat(this.multipleSelection);
            for (let i = 0; i < length; i++) {
                str += this.multipleSelection[i].name + ' ';
            }
            this.$message.error(`删除了${str}`);
            this.multipleSelection = [];
        },
        // 编辑操作
        handleCreate() {
            this.createVisible = true;
        },
        // 编辑操作
        handleEdit(index, row) {
            this.idx = index;
            this.form = row;
            this.editVisible = true;
        },
        // 复制操作
        handleCopy(index, row) {
            this.idx = index;
            this.form = row;
            this.copyParams.src_flow_key = this.form.key
			this.copyParams.flow_key = this.form.key
            this.copyVisible = true;
        },
        // 保存复制
        saveCopy() {
            this.copyVisible = false;
            copyFlowConfig(this.copyParams).then(res => {
                this.$message.success(`复制成功`);
                this.getData()
            });
        },
        saveCreate() {
            this.$refs.flowdetail.refresh();
            this.createVisible = false;
            createFlowConfig(this.createParams).then(res => {
                if (res.code == 'A0001') {
                    this.$message.success(`创建成功`);
				} else {
                    this.$message.error(`创建失败`);
				}
                this.getData()
            });
        },
        // 保存编辑
        saveEdit() {
            this.editVisible = false;
            this.$message.success(`修改第 ${this.idx + 1} 行成功`);
            this.$set(this.tableData, this.idx, this.form);
        },
        // 分页导航
        handlePageChange(val) {
            this.$set(this.query, 'page_index', val);
            this.getData();
        }
    }
};
</script>

<style scoped>
.handle-box {
    margin-bottom: 20px;
}

.handle-select {
    width: 120px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}
.table {
    width: 100%;
    font-size: 14px;
}
.red {
    color: #ff0000;
}
.mr10 {
    margin-right: 10px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}
</style>
