<template>
	<div class="form-box">
		<el-form ref="form" :model="form" label-width="80px">
			<el-form-item label="工作key">
				<el-input v-model="form.work_key"></el-input>
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
				<div id="config_editor"></div>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
    export default {
        name: "WorkDetailComponent",
        data() {
            return {
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
		props: ["form"],
        methods: {
            reinit() {
                console.log("WorkDetailComponent Reinit.")
                this.form.work_name = ""
                this.form.work_desc = ""
                this.form.work_key = ""
                this.form.config = {}
                this.initJsonEditor()
            },
            refresh() {
                this.form.config = JSON.stringify(this.config_editor.get())
            },
            initJsonEditor() {
                let config_container = document.getElementById("config_editor")
                config_container.innerHTML = "";
                this.config_editor = new this.$jsoneditor(config_container, this.options)
                this.config_editor.set(this.form.config)
            }
        },
        created() {
            this.$nextTick(() => {
                this.reinit()
            })
        }
    };
</script>
