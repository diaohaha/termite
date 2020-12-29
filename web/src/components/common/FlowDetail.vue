<template>
	<div class="form-box">
		<el-form ref="form" :model="form" label-width="80px">
			<el-form-item label="任务key">
				<el-input v-model="form.flow_key" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="任务名称">
				<el-input v-model="form.flow_name"></el-input>
			</el-form-item>
			<el-form-item label="描述">
				<el-input
					type="textarea"
					rows="5"
					v-model="form.flow_desc"
				></el-input>
			</el-form-item>
			<el-form-item label="依赖配置">
				<div id="dags_editor" style="height: 500px"></div>
			</el-form-item>
			<el-form-item label="环境变量">
				<div id="env_editor"></div>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="onSubmit">更新</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
    import { fetchFlows } from "../../api/index";
    import { queryFlowConfigs } from "../../api/index";
    import { updateFlowConfig } from "../../api/index";
    import { lowerJSONKey } from "../../utils/utils";
export default {
	name: "flowdetail",
	data() {
		return {
			flowData: {},
			query: {
				vflow: this.$route.params.flowKey
			},
			form: {
				flow_name: "",
				flow_key: "",
				flow_desc: "",
                env: "",
                config: ""
			},
            options: {
                mode: 'code',
                indentation: 4,
                search: true,
            },
			dag_editor: null,
			env_editor: null
		};
	},
    methods: {
        getFlowDetailData() {
            console.log(this.$route.params.flowKey)
            queryFlowConfigs(this.query).then(res => {
                this.flowData = res.data.configs[0];
                if (this.flowData) {
                    this.form.flow_key = this.flowData.key;
                    this.form.flow_name = this.flowData.name;
                    this.form.flow_desc = this.flowData.desc;
                    let dag_container = document.getElementById("dags_editor")
					this.dag_editor = new this.$jsoneditor(dag_container, this.options)
                    this.dag_editor.set(this.flowData.config)
                    let env_container = document.getElementById("env_editor")
                    this.env_editor = new this.$jsoneditor(env_container, this.options)
                    this.env_editor.set(this.flowData.env)
                }
            });
        },
        onSubmit() {
            this.form.config = JSON.stringify(this.dag_editor.get())
            this.form.env = JSON.stringify(this.env_editor.get())
            updateFlowConfig(this.form).then( res => {
			    if (res.code == "A0001") {
                    this.$message.success("更新成功！");
				} else {
                    this.$message.error(res.msg);
				}
			})
        }
    },
	created() {
		this.getFlowDetailData();
	}
};
</script>
