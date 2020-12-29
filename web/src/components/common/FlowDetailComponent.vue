<template>
	<div class="form-box">
		<el-form ref="form" :model="form" label-width="80px">
			<el-form-item label="任务key">
				<el-input v-model="form.flow_key"></el-input>
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
				<div id="dags_editor" ></div>
			</el-form-item>
			<el-form-item label="环境变量">
				<div id="env_editor"></div>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
    export default {
        name: "FlowDetailComponent",
        data() {
            return {
                form: {
                    flow_name: "",
                    flow_key: "",
                    flow_desc: "",
                    env: "{}",
                    config: "{}"
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
        props: ['form'],
		methods: {
            reinit() {
                console.log("FlowDetailComponent Reinit.")
                this.form.flow_name = ""
                this.form.flow_key = ""
                this.form.flow_desc = ""
                this.form.env = {}
                this.form.config = {}
                this.initJsonEditor()
			},
            refresh() {
                this.form.config = JSON.stringify(this.dag_editor.get())
                this.form.env = JSON.stringify(this.env_editor.get())
			},
			initJsonEditor() {
                let dag_container = document.getElementById("dags_editor")
				dag_container.innerHTML = "";
                this.dag_editor = new this.$jsoneditor(dag_container, this.options)
                this.dag_editor.set(this.form.config)
                let env_container = document.getElementById("env_editor")
                env_container.innerHTML = "";
                this.env_editor = new this.$jsoneditor(env_container, this.options)
                this.env_editor.set(this.form.env)
			}
		},
        created() {
            this.$nextTick(() => {
                this.reinit()
            })
        }
    };
</script>
