<template>
    <div>
        <div class="container">
            <div class="handle-box">
                <el-input v-model="query.search" placeholder="work key" class="handle-input mr10"></el-input>
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
                ref="worksTable"
                header-cell-class-name="table-header"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" align="center"></el-table-column>
				<el-table-column prop="Key" label="key">
					<template slot-scope="scope">
						<a href='javascript:void(0)' @click="jump(scope.row.key)">{{scope.row.key}}</a>
<!--						<a :href=" '/#/iworks/' + scope.row.key">{{scope.row.key}}</a>-->
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
                    :total="itemTotal"
                    @current-change="handlePageChange"
                ></el-pagination>
            </div>
        </div>

		<!-- Form -->
		<el-dialog title="创建" :visible.sync="createVisible">
			<WorkDetailComponent ref="workdetail" v-bind:form="createParams"></WorkDetailComponent>
			<div slot="footer" class="dialog-footer">
				<el-button @click="createVisible = false">取 消</el-button>
				<el-button type="primary" @click="saveCreate">确 定</el-button>
			</div>
		</el-dialog>
		<!-- 复制弹出框 -->
		<el-dialog title="复制" :visible.sync="copyVisible" width="30%">
			<el-form ref="copyParams" :model="copyParams" label-width="70px">
				<el-form-item label="key">
					<el-input v-model="copyParams.work_key"></el-input>
				</el-form-item>
			</el-form>
			<span slot="footer" class="dialog-footer">
                <el-button @click="copyVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveCopy">确 定</el-button>
            </span>
		</el-dialog>
	</div>
</template>

<script>
    import { fetchWorksConfig } from '../../api/index';
    import { queryWorkConfigs } from '../../api/index';
    import { createWorkConfig } from '../../api/index';
    import { copyWorkConfig } from '../../api/index';
    import { deleteWorkConfig } from '../../api/index';
    import WorkDetailComponent from '../common/WorkDetailComponent'
export default {
    name: 'iflows',
    components: {WorkDetailComponent},
    data() {
        return {
            query: {
                search: "",
				page_index: 1,
				page_size: 13,
			},
            tableData: [],
            itemTotal: 0,
            multipleSelection: [],
            delList: [],
            editVisible: false,
            copyVisible: false,
            createVisible: false,
			copyParams: {
                src_work_key: "",
				work_key: ""
			},
            deleteParams: {
                work_key: ""
            },
            createParams: {
                work_name: "",
                work_key: "",
                work_desc: "",
                config: ""
            },
            form: {},
            idx: -1,
        };
    },
    created() {
        this.getData();
    },
    methods: {
        getData() {
            queryWorkConfigs(this.query).then(res => {
                this.tableData = res.data.configs;
                this.itemTotal = res.data.count
            });
        },
        jump(dsturl) {
            this.$router.push("iworks/" + dsturl)
        },
        // 创建
        handleCreate() {
            this.createVisible = true;
        },
        // 删除操作
        handleDelete(index, row) {
            // 二次确认删除
			this.deleteParams.work_key = row.key
            this.$confirm('确定要删除吗？', '提示', {
                type: 'warning'
            })
                .then(() => {
                    deleteWorkConfig(this.deleteParams).then(res => {
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
        handleEdit(index, row) {
            console.log("edit", index, row)
            this.idx = index;
            this.form = row;
            this.editVisible = true;
        },
        // 复制操作
        handleCopy(index, row) {
            this.copyParams.work_key = row.key
            this.copyParams.src_work_key = row.key
            this.copyVisible = true;
        },
		saveCreate() {
            this.$refs.workdetail.refresh();
            this.createVisible = false;
            createWorkConfig(this.createParams).then(res => {
                if (res.code == 'A0001') {
                    this.$message.success(`创建成功`);
                } else {
                    this.$message.error(`创建失败`);
                }
                this.getData()
            });
		},
        // 保存复制
        saveCopy() {
            copyWorkConfig(this.copyParams).then(res => {
                if (res.code == "A0001"){
                    this.$message.success(`复制成功`);
                    this.getData()
				} else {
                    this.$message.error(res.msg);
				}
			})
            this.copyVisible = false;
        },
        // 分页导航
        handlePageChange(val) {
            this.$set(this.query, 'page_index', val);
            this.getData();
        },
        // 触发搜索按钮
        handleSearch() {
            this.$set(this.query, 'page_index', 1);
            this.getData();
        },
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
