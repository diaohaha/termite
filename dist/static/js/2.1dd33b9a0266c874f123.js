webpackJsonp([2],{MpTN:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o={data:function(){return{activeIndex:"1",name:"linxin"}},computed:{username:function(){var t=localStorage.getItem("ms_username");return t||this.name},onRoutes:function(){return this.$route.path.includes("iflows")?"/iflows":this.$route.path.includes("iworks")?"/iworks":this.$route.path.includes("dashboard")?"/dashboard":this.$route.path}},methods:{handleCommand:function(t){"loginout"==t&&(localStorage.removeItem("ms_username"),this.$router.push("/login"))}}},s={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"header"},[n("div",{staticClass:"logo"},[t._v("Termite")]),t._v(" "),n("div",{staticClass:"menu"},[n("el-menu",{staticClass:"el-menu-demo",attrs:{"default-active":t.onRoutes,mode:"horizontal","background-color":"#545c64","text-color":"#fff","active-text-color":"#ffd04b",router:""}},[n("el-menu-item",{attrs:{index:"/dashboard"}},[t._v("Home")]),t._v(" "),n("el-menu-item",{attrs:{index:"/iworks"}},[t._v("Work")]),t._v(" "),n("el-menu-item",{attrs:{index:"/iflows"}},[t._v("Flow")])],1)],1),t._v(" "),n("div",{staticClass:"header-right"},[n("div",{staticClass:"header-user-con"},[n("el-dropdown",{staticClass:"user-name",attrs:{trigger:"click"},on:{command:t.handleCommand}},[n("span",{staticClass:"el-dropdown-link"},[t._v("\n\t\t\t\t\t"+t._s(t.username)+"\n\t\t\t\t\t"),n("i",{staticClass:"el-icon-caret-bottom"})]),t._v(" "),n("el-dropdown-menu",{attrs:{slot:"dropdown"},slot:"dropdown"},[n("el-dropdown-item",{attrs:{divided:"",command:"loginout"}},[t._v("退出登录")])],1)],1)],1)])])},staticRenderFns:[]};var a={data:function(){return{}},components:{vHeader:n("VU/8")(o,s,!1,function(t){n("Spmz")},"data-v-6ec127b7",null).exports}},r={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"wrapper"},[e("v-header"),this._v(" "),e("div",{staticClass:"content-box"},[e("div",{staticClass:"content"},[e("transition",{attrs:{name:"move",mode:"out-in"}},[e("router-view")],1)],1)])],1)},staticRenderFns:[]};var i=n("VU/8")(a,r,!1,function(t){n("zvJh")},null,null);e.default=i.exports},Spmz:function(t,e){},zvJh:function(t,e){}});
//# sourceMappingURL=2.1dd33b9a0266c874f123.js.map