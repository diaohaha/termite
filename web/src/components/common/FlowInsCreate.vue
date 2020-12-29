<template>
	<div class="form-box">
		<el-form ref="form" :model="form" label-width="80px" id="upload" enctype="multipart/form-data" method="post">
			<el-row>
				<el-input v-model="form.flow_key" :disabled="true"></el-input>
			</el-row>
			<el-row :gutter="20">
				<el-col :span="12">
					<el-input
						:cols="2"
						type="textarea"
						placeholder="请输出CID"
						rows="20"
						v-model="form.cids_str"
					></el-input>
				</el-col>
				<el-col :span="12">
					<el-upload
						ref="upload"
						drag
						action=""
						:http-request="handlePost"
						:file-list="fileList"
						:auto-upload="false"
						:on-preview="handlePreview"
						>
						<i class="el-icon-upload"></i>
						<div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
					</el-upload>
				</el-col>
			</el-row>
			<el-row>
				<el-button type="primary" @click="onSubmit" style="width: 100%">创建</el-button>
			</el-row>
		</el-form>
	</div>
</template>

<script>
    import { addFlowInstances } from "../../api/index";
    import { addFlowInstancesByFile } from "../../api/index";
    export default {
        name: "flowInsAdd",
        data() {
            return {
                body: {
                    vflow: this.$route.params.flowKey,
                    cids: []
                },
				filePostBody: {
                    vflow: this.$route.params.flowKey,
				},
                form: {
                    flow_key: this.$route.params.flowKey,
                    cids_str: ""
                },
				fileList: [],
            };
        },
        methods: {
            handlePreview(file) {
                console.log(file)
                // this.fileList.push({
                //     name: file.name,
                // });
            },
			handlePost(file) {
                //
				console.log("handle post")
                // var data = document.getElementById('upload');
                // const fd = new window.FormData(data)
                var formElement = document.querySelector("form");
                var formData = new FormData(formElement);
                formData.append('vflow', this.$route.params.flowKey)
				formData.delete('file')
                formData.append('file', file.file)
                addFlowInstancesByFile(formData).then(res => {
                    console.log("imhear")
                    if (res.code == 'A0001') {
                        this.$message.success("上传成功")
                        this.form.cids_str = ""
                        this.body.cids = []
                    } else {
                        this.$message.error("上传失败")
                    }
                    this.$refs.upload.clearFiles()
                    vm.$refs.upload.uploadFiles.length = 0
				})
			},
            onSubmit() {
                this.body.cids = []
				console.log(this.fileList)
                console.log(this.$refs.upload.uploadFiles.length)
                if (this.form.cids_str == "" && this.$refs.upload.uploadFiles.length == 0) {
                    this.$message.error("请输入ID或添加文件");
                    return
				} else {
                    if (this.form.cids_str != "") {
                        this.body.cids = this.form.cids_str.split('\n')
                        addFlowInstances(this.body).then(res => {
                            if (res.code == 'A0001') {
                                this.$message.success("success:" + res.data.success + "  error:" + res.data.fail)
                                this.form.cids_str = ""
                                this.body.cids = []
                            } else {
                                this.$message.error("添加失败")
                            }
                        });
					} else {
                        // 文件上传
                        this.$refs.upload.submit();
                        // this.handlePost(file)
					}
				}
            }
        },
        created() {
        }
    };
</script>

<style>
	.el-row {
		margin-bottom: 20px;
	}
</style>
