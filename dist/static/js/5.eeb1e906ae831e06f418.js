webpackJsonp([5],{"3jL2":function(e,r,s){"use strict";Object.defineProperty(r,"__esModule",{value:!0});var t={data:function(){return{param:{username:"admin",password:"123123"},rules:{username:[{required:!0,message:"请输入用户名",trigger:"blur"}],password:[{required:!0,message:"请输入密码",trigger:"blur"}]}}},methods:{submitForm:function(){var e=this;this.$refs.login.validate(function(r){if(!r)return e.$message.error("请输入账号和密码"),console.log("error submit!!"),!1;e.$message.success("登录成功"),localStorage.setItem("ms_username",e.param.username),e.$router.push("/")})}}},a={render:function(){var e=this,r=e.$createElement,s=e._self._c||r;return s("div",{staticClass:"login-wrap"},[s("div",{staticClass:"ms-login"},[s("div",{staticClass:"ms-title"},[e._v("Termite")]),e._v(" "),s("el-form",{ref:"login",staticClass:"ms-content",attrs:{model:e.param,rules:e.rules,"label-width":"0px"}},[s("el-form-item",{attrs:{prop:"username"}},[s("el-input",{attrs:{placeholder:"username"},model:{value:e.param.username,callback:function(r){e.$set(e.param,"username",r)},expression:"param.username"}},[s("el-button",{attrs:{slot:"prepend",icon:"el-icon-lx-people"},slot:"prepend"})],1)],1),e._v(" "),s("el-form-item",{attrs:{prop:"password"}},[s("el-input",{attrs:{type:"password",placeholder:"password"},nativeOn:{keyup:function(r){return!r.type.indexOf("key")&&e._k(r.keyCode,"enter",13,r.key,"Enter")?null:e.submitForm()}},model:{value:e.param.password,callback:function(r){e.$set(e.param,"password",r)},expression:"param.password"}},[s("el-button",{attrs:{slot:"prepend",icon:"el-icon-lx-lock"},slot:"prepend"})],1)],1),e._v(" "),s("div",{staticClass:"login-btn"},[s("el-button",{attrs:{type:"primary"},on:{click:function(r){return e.submitForm()}}},[e._v("Login")])],1),e._v(" "),s("p",{staticClass:"login-tips"},[e._v("Tips : 用户名和密码随便填。")])],1)],1)])},staticRenderFns:[]};var n=s("VU/8")(t,a,!1,function(e){s("QQK9")},"data-v-144781ec",null);r.default=n.exports},QQK9:function(e,r){}});
//# sourceMappingURL=5.eeb1e906ae831e06f418.js.map