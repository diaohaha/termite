webpackJsonp([2],{BUHG:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=a("gyMJ"),n={name:"iflows",data:function(){return{query:{},tableData:[],multipleSelection:[],delList:[],editVisible:!1,copyVisible:!1,copyParams:{src_work_key:"",work_key:""},deleteParams:{work_key:""},pageTotal:0,form:{},idx:-1,id:-1}},created:function(){this.getData()},methods:{getData:function(){var t=this;Object(s.i)(this.query).then(function(e){t.tableData=e.data,t.pageTotal=e.data.length})},jump:function(t){this.$router.push("iworks/"+t)},handleSearch:function(){this.$set(this.query,"pageIndex",1),this.getData()},handleDelete:function(t,e){var a=this;this.deleteParams.work_key=e.key,this.$confirm("确定要删除吗？","提示",{type:"warning"}).then(function(){Object(s.e)(a.deleteParams).then(function(e){"A0001"==e.code?(a.$message.success("删除成功"),a.tableData.splice(t,1)):a.$message.error(e.msg)})}).catch(function(){})},handleSelectionChange:function(t){this.multipleSelection=t},delAllSelection:function(){var t=this.multipleSelection.length,e="";this.delList=this.delList.concat(this.multipleSelection);for(var a=0;a<t;a++)e+=this.multipleSelection[a].name+" ";this.$message.error("删除了"+e),this.multipleSelection=[]},handleEdit:function(t,e){console.log("edit",t,e),this.idx=t,this.form=e,this.editVisible=!0},handleCopy:function(t,e){this.copyParams.work_key=e.key,this.copyParams.src_work_key=e.key,this.copyVisible=!0},saveEdit:function(){this.editVisible=!1,this.$message.success("修改第 "+(this.idx+1)+" 行成功"),this.$set(this.tableData,this.idx,this.form)},saveCopy:function(){var t=this;Object(s.c)(this.copyParams).then(function(e){"A0001"==e.code?(t.$message.success("复制成功"),t.getData()):t.$message.error(e.msg)}),this.copyVisible=!1},handlePageChange:function(t){this.$set(this.query,"pageIndex",t),this.getData()}}},i={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticClass:"container"},[a("div",{staticClass:"handle-box"},[a("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"work key"},model:{value:t.query.name,callback:function(e){t.$set(t.query,"name",e)},expression:"query.name"}}),t._v(" "),a("el-button",{attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.handleSearch}},[t._v("search")]),t._v(" "),a("el-button",{staticClass:"handle-add mr10",staticStyle:{float:"right"},attrs:{type:"primary",icon:"el-icon-circle-plus"}},[t._v("添加")])],1),t._v(" "),a("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:t.tableData,border:"","header-cell-class-name":"table-header"},on:{"selection-change":t.handleSelectionChange}},[a("el-table-column",{attrs:{type:"selection",width:"55",align:"center"}}),t._v(" "),a("el-table-column",{attrs:{prop:"Key",label:"key"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("a",{attrs:{href:"javascript:void(0)"},on:{click:function(a){return t.jump(e.row.key)}}},[t._v(t._s(e.row.key))])]}}])}),t._v(" "),a("el-table-column",{attrs:{prop:"name",label:"name"}}),t._v(" "),a("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{type:"text",icon:"el-icon-copy-document"},on:{click:function(a){return t.handleCopy(e.$index,e.row)}}},[t._v("复制")]),t._v(" "),a("el-button",{staticClass:"red",attrs:{type:"text",icon:"el-icon-delete"},on:{click:function(a){return t.handleDelete(e.$index,e.row)}}},[t._v("删除")])]}}])})],1),t._v(" "),a("div",{staticClass:"pagination"},[a("el-pagination",{attrs:{background:"",layout:"total, prev, pager, next","current-page":t.query.pageIndex,"page-size":t.query.pageSize,total:t.pageTotal},on:{"current-change":t.handlePageChange}})],1)],1),t._v(" "),a("el-dialog",{attrs:{title:"编辑",visible:t.editVisible,width:"30%"},on:{"update:visible":function(e){t.editVisible=e}}},[a("el-form",{ref:"form",attrs:{model:t.form,"label-width":"70px"}},[a("el-form-item",{attrs:{label:"key"}},[a("el-input",{attrs:{disabled:""},model:{value:t.form.key,callback:function(e){t.$set(t.form,"key",e)},expression:"form.key"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"name"}},[a("el-input",{model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}})],1)],1),t._v(" "),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(e){t.editVisible=!1}}},[t._v("取 消")]),t._v(" "),a("el-button",{attrs:{type:"primary"},on:{click:t.saveEdit}},[t._v("确 定")])],1)],1),t._v(" "),a("el-dialog",{attrs:{title:"复制",visible:t.copyVisible,width:"30%"},on:{"update:visible":function(e){t.copyVisible=e}}},[a("el-form",{ref:"copyParams",attrs:{model:t.copyParams,"label-width":"70px"}},[a("el-form-item",{attrs:{label:"key"}},[a("el-input",{model:{value:t.copyParams.work_key,callback:function(e){t.$set(t.copyParams,"work_key",e)},expression:"copyParams.work_key"}})],1)],1),t._v(" "),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(e){t.copyVisible=!1}}},[t._v("取 消")]),t._v(" "),a("el-button",{attrs:{type:"primary"},on:{click:t.saveCopy}},[t._v("确 定")])],1)],1)],1)},staticRenderFns:[]};var o=a("VU/8")(n,i,!1,function(t){a("maGs")},"data-v-6d2b7475",null);e.default=o.exports},enD8:function(t,e){},maGs:function(t,e){},"uus/":function(t,e){},w7KR:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=a("mvHQ"),n=a.n(s),i=a("gyMJ"),o={name:"workDetail",data:function(){return{workData:{},query:{work_key:this.$route.params.workKey},form:{work_name:"",work_key:"",work_desc:"",config:""},options:{mode:"code",indentation:4,search:!0},config_editor:null}},methods:{getWorkDetailData:function(){var t=this;Object(i.i)(this.query).then(function(e){if(t.workData=e.data[0],console.log("work Data: ",t.workData),t.workData){t.form.work_key=t.workData.key,t.form.work_name=t.workData.name,t.form.work_desc=t.workData.desc;var a=document.getElementById("config_editor");t.config_editor=new t.$jsoneditor(a,t.options),t.config_editor.set(t.workData.config)}})},onSubmit:function(){var t=this;this.form.config=n()(this.config_editor.get()),Object(i.o)(this.form).then(function(e){"A0001"==e.code?t.$message.success("更新成功！"):t.$message.error(e.msg)})}},created:function(){this.getWorkDetailData()}},r={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"form-box"},[a("el-form",{ref:"form",attrs:{model:t.form,"label-width":"80px"}},[a("el-form-item",{attrs:{label:"工作key"}},[a("el-input",{attrs:{disabled:!0},model:{value:t.form.work_key,callback:function(e){t.$set(t.form,"work_key",e)},expression:"form.work_key"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"工作名称"}},[a("el-input",{model:{value:t.form.work_name,callback:function(e){t.$set(t.form,"work_name",e)},expression:"form.work_name"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"描述"}},[a("el-input",{attrs:{type:"textarea",rows:"5"},model:{value:t.form.work_desc,callback:function(e){t.$set(t.form,"work_desc",e)},expression:"form.work_desc"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"工作配置"}},[a("div",{staticStyle:{height:"500px"},attrs:{id:"config_editor"}})]),t._v(" "),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("更新")])],1)],1)],1)},staticRenderFns:[]},l={name:"workInstance",data:function(){return{pageTotal:0,tableData:[],query:{cid:"",vwork:this.$route.params.workKey,page_index:1,page_size:15},expands:[]}},methods:{tableRowClassName:function(t){var e=t.row;t.rowIndex;return console.log("row is:",e),"超时"===e.state?"warning-row":"异常"===e.state?"failed-row":"失败"===e.state?"warning-row":"初始化"===e.state?"success-row":"已下发"===e.state?"success-row":"执行中"===e.state?"success-row":"成功"===e.state?"success-row":"延期调度"===e.state?"success-row":""},expandContext:function(t,e,a){Array.prototype.remove=function(t){var e=this.indexOf(t);e>-1&&this.splice(e,1)},this.expands=[],this.expands.indexOf(t.flow_id)<0?this.expands.push(t.flow_id):this.expands.remove(t.flow_id),this.expands.indexOf(t.flow_id)<0?this.expands.push(t.flow_id):this.expands.remove(t.flow_id)},getWorkInstanceData:function(){var t=this;console.log(this.$route.params.workkey),Object(i.h)(this.query).then(function(e){t.tableData=e.data.works,t.pageTotal=e.data.count})},handleSearch:function(){this.$set(this.query,"page_index",1),this.getWorkInstanceData()},handlePageChange:function(t){this.$set(this.query,"page_index",t),this.getWorkInstanceData()}},created:function(){this.getWorkInstanceData()}},c={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticClass:"container"},[a("div",{staticClass:"handle-box"},[a("el-input",{staticClass:"handle-input",staticStyle:{width:"300px"},attrs:{placeholder:"cid"},model:{value:t.query.cid,callback:function(e){t.$set(t.query,"cid",e)},expression:"query.cid"}}),t._v(" "),a("el-button",{attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.handleSearch}},[t._v("search")])],1),t._v(" "),a("br"),t._v(" "),a("el-table",{ref:"multipleTable",staticClass:"table",staticStyle:{width:"100%"},attrs:{data:t.tableData,border:"","row-class-name":t.tableRowClassName,"header-cell-class-name":"table-header","row-key":"work_id","expand-row-keys":t.expands},on:{"row-click":t.expandContext}},[a("el-table-column",{attrs:{type:"expand"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-row",{staticClass:"row-bg",attrs:{type:"flex",justify:"space-around"}},[a("el-col",{attrs:{span:2}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",{staticStyle:{"font-weight":"lighter"}},[t._v("创建时间:")])])]),t._v(" "),a("el-col",{attrs:{span:4}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",[t._v(t._s(e.row.create_time))])])]),t._v(" "),a("el-col",{attrs:{span:2}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",{staticStyle:{"font-weight":"lighter"}},[t._v("更新时间:")])])]),t._v(" "),a("el-col",{attrs:{span:4}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",[t._v(t._s(e.row.update_time))])])]),t._v(" "),a("el-col",{attrs:{span:2}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",{staticStyle:{"font-weight":"lighter"}},[t._v("对象ID:")])])]),t._v(" "),a("el-col",{attrs:{span:4}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",[t._v(t._s(e.row.cid))])])]),t._v(" "),a("el-col",{attrs:{span:2}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",{staticStyle:{"font-weight":"lighter"}},[t._v("任务名称:")])])]),t._v(" "),a("el-col",{attrs:{span:4}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",[t._v(t._s(e.row.vwork))])])])],1),t._v(" "),a("el-row",{staticClass:"row-bg",attrs:{type:"flex",justify:"space-around"}},[a("el-col",{attrs:{span:2}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",{staticStyle:{"font-weight":"lighter"}},[t._v("任务输出:")])])]),t._v(" "),a("el-col",{attrs:{span:22}},[a("div",{staticClass:"grid-content bg-purple"},[a("pre",[t._v(t._s(e.row.output))])])])],1),t._v(" "),a("el-row",{staticClass:"row-bg",attrs:{type:"flex",justify:"space-around"}},[a("el-col",{attrs:{span:2}},[a("div",{staticClass:"grid-content bg-purple"},[a("span",{staticStyle:{"font-weight":"lighter"}},[t._v("异常信息:")])])]),t._v(" "),a("el-col",{attrs:{span:22}},[a("div",{staticClass:"grid-content bg-purple"},[a("pre",[t._v(t._s(e.row.error))])])])],1)]}}])}),t._v(" "),a("el-table-column",{attrs:{prop:"work_id",label:"work_id"}}),t._v(" "),a("el-table-column",{attrs:{prop:"create_time",label:"create_time"}}),t._v(" "),a("el-table-column",{attrs:{prop:"update_time",label:"update_time"}}),t._v(" "),a("el-table-column",{attrs:{prop:"cid",label:"cid"}}),t._v(" "),a("el-table-column",{attrs:{prop:"state",label:"state"}})],1),t._v(" "),a("div",{staticClass:"pagination"},[a("el-pagination",{attrs:{background:"",layout:"total, prev, pager, next","current-page":t.query.page_index,"page-size":t.query.page_size,total:t.pageTotal},on:{"current-change":t.handlePageChange}})],1)],1)])},staticRenderFns:[]};var d={name:"Work",data:function(){return{active:0,currentType:"detail",currentView:"workDetail",tabs:[{type:"detail",view:"workDetail"},{type:"ins",view:"workInstance"}]}},methods:{toggle:function(t,e,a){console.log("toggle",t,e,a),this.active=t,this.currentView=e,this.currentType=a}},components:{workDetail:a("VU/8")(o,r,!1,null,null,null).exports,workInstance:a("VU/8")(l,c,!1,function(t){a("enD8")},null,null).exports}},u={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticClass:"container"},[a("ul",{staticClass:"innertab"},t._l(t.tabs,function(e,s){return a("li",{staticClass:"innertab",class:{active:t.active==s},on:{click:function(a){return t.toggle(s,e.view,e.type)}}},[t._v("\n\t\t\t\t"+t._s(e.type)+"\n\t\t\t")])}),0),t._v(" "),a("hr"),t._v(" "),a(t.currentView,{tag:"component",staticStyle:{"padding-top":"20px"}})],1)])},staticRenderFns:[]};var p=a("VU/8")(d,u,!1,function(t){a("uus/")},null,null);e.default=p.exports}});
//# sourceMappingURL=2.22a6087371bb5b1a03b1.js.map