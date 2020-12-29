<template>
	<div class="form-box">
		<el-form ref="form" :model="form" label-width="80px">
			<el-form-item label="工作key">
				<el-input v-model="form.work_key" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="工作名称">
				<el-input v-model="form.work_name"></el-input>
			</el-form-item>
			<el-form-item label="描述">
				<el-input
					type="textarea"
					rows="5"
					v-model="form.work_desc"
				></el-input>
			</el-form-item>
			<el-form-item label="工作配置">
				<div id="config_editor" style="height: 500px"></div>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="onSubmit">更新</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
    import { fetchWorksConfig } from "../../api/index";
    import { updateWorkConfig } from "../../api/index";
export default {
	name: "workDetail",
	data() {
		return {
			workData: {},
			query: {
				work_key: this.$route.params.workKey
			},
			form: {
				work_name: "",
				work_key: "",
				work_desc: "",
                config: ""
			},
            options: {
                mode: 'code',
                indentation: 4,
                search: true,
            },
            config_editor: null,
		};
	},
    methods: {
        getWorkDetailData() {
            fetchWorksConfig(this.query).then(res => {
                this.workData = res.data[0];
                console.log("work Data: ", this.workData)
                if (this.workData) {
                    this.form.work_key = this.workData.key;
                    this.form.work_name = this.workData.name;
                    this.form.work_desc = this.workData.desc;
                    let config_container = document.getElementById("config_editor")
					this.config_editor = new this.$jsoneditor(config_container, this.options)
                    this.config_editor.set(this.workData.config)
                }
            });
        },
        onSubmit() {
            this.form.config = JSON.stringify(this.config_editor.get())
            updateWorkConfig(this.form).then( res => {
			    if (res.code == "A0001") {
                    this.$message.success("更新成功！");
				} else {
                    this.$message.error(res.msg);
				}
			})
        }
    },
	created() {
		this.getWorkDetailData();
	}
};
</script>
