<template>
	<div>
		<el-row :gutter="20">
			<el-col :span="4">
				<el-card class="mycard">work config<div class="num">{{ workConfigNum }}</div></el-card>
			</el-col>
			<el-col :span="4">
				<el-card class="mycard">active work config<div class="num">{{ workActiveNum }}</div></el-card>
			</el-col>
			<el-col :span="4">
				<el-card class="mycard">work instance<div class="num">{{ workInsNum }}</div></el-card>
			</el-col>
			<el-col :span="4">
				<el-card class="mycard">flow config<div class="num">{{ flowConfigNum }}</div></el-card>
			</el-col>
			<el-col :span="4">
				<el-card class="mycard">active flow config<div class="num">{{ flowActiveNum }}</div></el-card>
			</el-col>
			<el-col :span="4">
				<el-card class="mycard">flow instance<div class="num">{{ flowInsNum }}</div></el-card>
			</el-col>
		</el-row>
		<el-row :gutter="20">
			<el-col :span="12">
				<el-card shadow="hover">
					<schart
						ref="bar"
						class="schart"
						canvasId="bar1"
						:data="workData"
						type="bar"
						:options="optionsWork"
					></schart>
				</el-card>
			</el-col>
			<el-col :span="12">
				<el-card shadow="hover">
					<schart
						ref="bar"
						class="schart"
						canvasId="bar2"
						:data="flowData"
						type="bar"
						:options="optionsFlow"
					></schart>
				</el-card>
			</el-col>
		</el-row>
		<el-row :gutter="20">
			<el-col :span="24">
				<el-card shadow="hover">
					<schart
						ref="bar"
						class="schart"
						canvasId="bar3"
						:data="flowDataByFlow"
						type="bar"
						:options="optionsFlowByFlow"
					></schart>
				</el-card>
			</el-col>
		</el-row>
	</div>
</template>

<script>
import Schart from "vue-schart";
import bus from "../common/bus";
import { getFlowCountInfo } from "../../api/index";
import { getWorkCountInfo } from "../../api/index";

export default {
	name: "dashboard",
	data() {
		return {
			name: localStorage.getItem("ms_username"),
			workData: [],
			flowData: [],
            flowDataByFlow: [],
			workConfigNum: 0,
            workActiveNum: 0,
            workInsNum: 0,
            flowConfigNum: 0,
            flowActiveNum: 0,
            flowInsNum: 0,
			optionsWork: {
				title: "工作状态分布",
				showValue: false,
				fillColor: "rgb(45, 140, 240)",
				bottomPadding: 30,
				topPadding: 30
			},
            optionsFlow: {
                title: "工作流状态分布",
                showValue: false,
                fillColor: "rgb(45, 140, 240)",
                bottomPadding: 30,
                topPadding: 30
            },
            optionsFlowByFlow: {
                title: "工作流调度分布",
                showValue: false,
                fillColor: "rgb(45, 140, 240)",
                bottomPadding: 30,
                topPadding: 30
            },
		};
	},
	components: {
		Schart
	},
	computed: {
		role() {
			return this.name === "admin" ? "超级管理员" : "普通用户";
		}
	},
	created() {
		this.handleListener();
		this.getDataFlowCount();
		this.getDataWorkCount();
	},
	activated() {
		this.handleListener();
	},
	deactivated() {
		window.removeEventListener("resize", this.renderChart);
		bus.$off("collapse", this.handleBus);
	},
	methods: {
        getDataWorkCount() {
            getWorkCountInfo(this.query).then(res => {
                this.workConfigNum = res.data.work_config_num
                this.workActiveNum = res.data.active_work_num
                this.workInsNum = res.data.work_instance_num
                let length = res.data.work_num_detail.length
                for (let i = 0; i < length; i++) {
                    this.workData.push({
                        "name": res.data.work_num_detail[i].state,
                        "value": res.data.work_num_detail[i].num,
                    })
                }
            });
        },
        getDataFlowCount() {
            getFlowCountInfo(this.query).then(res => {
                this.flowConfigNum = res.data.flow_config_num
                this.flowActiveNum = res.data.active_flow_num
                this.flowInsNum = res.data.flow_instance_num
                let length = res.data.flow_num_detail_by_state.length
                for (let i = 0; i < length; i++) {
                    this.flowData.push({
						"name": res.data.flow_num_detail_by_state[i].state,
                        "value": res.data.flow_num_detail_by_state[i].num,
					})
                }
                let length1 = res.data.flow_num_detail_by_flow.length
                for (let i = 0; i < length1; i++) {
                    this.flowDataByFlow.push({
                        "name": res.data.flow_num_detail_by_flow[i].vflow,
                        "value": res.data.flow_num_detail_by_flow[i].num,
                    })
                }
            });
		},
		handleListener() {
			bus.$on("collapse", this.handleBus);
			// 调用renderChart方法对图表进行重新渲染
			window.addEventListener("resize", this.renderChart);
		},
		handleBus(msg) {
			setTimeout(() => {
				this.renderChart();
			}, 300);
		},
		renderChart() {
			this.$refs.bar.renderChart();
			this.$refs.line.renderChart();
		}
	}
};
</script>

<style scoped>
.el-row {
	margin-bottom: 20px;
}

.grid-content {
	display: flex;
	align-items: center;
	height: 100px;
}

.grid-cont-right {
	flex: 1;
	text-align: center;
	font-size: 14px;
	color: #999;
}

.grid-num {
	font-size: 30px;
	font-weight: bold;
}

.grid-con-icon {
	font-size: 50px;
	width: 100px;
	height: 100px;
	text-align: center;
	line-height: 100px;
	color: #fff;
}

.grid-con-1 .grid-con-icon {
	background: rgb(45, 140, 240);
}

.grid-con-1 .grid-num {
	color: rgb(45, 140, 240);
}

.grid-con-2 .grid-con-icon {
	background: rgb(100, 213, 114);
}

.grid-con-2 .grid-num {
	color: rgb(45, 140, 240);
}

.grid-con-3 .grid-con-icon {
	background: rgb(242, 94, 67);
}

.grid-con-3 .grid-num {
	color: rgb(242, 94, 67);
}

.user-info {
	display: flex;
	align-items: center;
	padding-bottom: 20px;
	border-bottom: 2px solid #ccc;
	margin-bottom: 20px;
}

.user-avator {
	width: 120px;
	height: 120px;
	border-radius: 50%;
}

.user-info-cont {
	padding-left: 50px;
	flex: 1;
	font-size: 14px;
	color: #999;
}

.user-info-cont div:first-child {
	font-size: 30px;
	color: #222;
}

.user-info-list {
	font-size: 14px;
	color: #999;
	line-height: 25px;
}

.user-info-list span {
	margin-left: 70px;
}

.mgb20 {
	margin-bottom: 20px;
}

.todo-item {
	font-size: 14px;
}

.todo-item-del {
	text-decoration: line-through;
	color: #999;
}

.schart {
	width: 100%;
	height: 300px;
}

.num {
	font-size: 12px;
	opacity: .69;
	line-height: 24px;
}
.mycard {
	background: #a0cfff;
}
</style>
