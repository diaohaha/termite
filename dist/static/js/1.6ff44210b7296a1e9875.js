webpackJsonp([1],{"//Fk":function(t,e,n){t.exports={default:n("U5ju"),__esModule:!0}},"1Yoh":function(t,e){
/*!
 * Determine if an object is a Buffer
 *
 * @author   Feross Aboukhadijeh <https://feross.org>
 * @license  MIT
 */
t.exports=function(t){return null!=t&&null!=t.constructor&&"function"==typeof t.constructor.isBuffer&&t.constructor.isBuffer(t)}},"21It":function(t,e,n){"use strict";var r=n("FtD3");t.exports=function(t,e,n){var o=n.config.validateStatus;!o||o(n.status)?t(n):e(r("Request failed with status code "+n.status,n.config,null,n.request,n))}},"2KxR":function(t,e){t.exports=function(t,e,n,r){if(!(t instanceof e)||void 0!==r&&r in t)throw TypeError(n+": incorrect invocation!");return t}},"3fs2":function(t,e,n){var r=n("RY/4"),o=n("dSzd")("iterator"),i=n("/bQp");t.exports=n("FeBl").getIteratorMethod=function(t){if(void 0!=t)return t[o]||t["@@iterator"]||i[r(t)]}},"5VQ+":function(t,e,n){"use strict";var r=n("cGG2");t.exports=function(t,e){r.forEach(t,function(n,r){r!==e&&r.toUpperCase()===e.toUpperCase()&&(t[e]=n,delete t[r])})}},"7GwW":function(t,e,n){"use strict";var r=n("cGG2"),o=n("21It"),i=n("DQCr"),a=n("oJlt"),s=n("GHBc"),c=n("FtD3");t.exports=function(t){return new Promise(function(e,u){var f=t.data,l=t.headers;r.isFormData(f)&&delete l["Content-Type"];var d=new XMLHttpRequest;if(t.auth){var p=t.auth.username||"",h=t.auth.password||"";l.Authorization="Basic "+btoa(p+":"+h)}if(d.open(t.method.toUpperCase(),i(t.url,t.params,t.paramsSerializer),!0),d.timeout=t.timeout,d.onreadystatechange=function(){if(d&&4===d.readyState&&(0!==d.status||d.responseURL&&0===d.responseURL.indexOf("file:"))){var n="getAllResponseHeaders"in d?a(d.getAllResponseHeaders()):null,r={data:t.responseType&&"text"!==t.responseType?d.response:d.responseText,status:d.status,statusText:d.statusText,headers:n,config:t,request:d};o(e,u,r),d=null}},d.onabort=function(){d&&(u(c("Request aborted",t,"ECONNABORTED",d)),d=null)},d.onerror=function(){u(c("Network Error",t,null,d)),d=null},d.ontimeout=function(){u(c("timeout of "+t.timeout+"ms exceeded",t,"ECONNABORTED",d)),d=null},r.isStandardBrowserEnv()){var m=n("p1b6"),v=(t.withCredentials||s(t.url))&&t.xsrfCookieName?m.read(t.xsrfCookieName):void 0;v&&(l[t.xsrfHeaderName]=v)}if("setRequestHeader"in d&&r.forEach(l,function(t,e){void 0===f&&"content-type"===e.toLowerCase()?delete l[e]:d.setRequestHeader(e,t)}),t.withCredentials&&(d.withCredentials=!0),t.responseType)try{d.responseType=t.responseType}catch(e){if("json"!==t.responseType)throw e}"function"==typeof t.onDownloadProgress&&d.addEventListener("progress",t.onDownloadProgress),"function"==typeof t.onUploadProgress&&d.upload&&d.upload.addEventListener("progress",t.onUploadProgress),t.cancelToken&&t.cancelToken.promise.then(function(t){d&&(d.abort(),u(t),d=null)}),void 0===f&&(f=null),d.send(f)})}},"82Mu":function(t,e,n){var r=n("7KvD"),o=n("L42u").set,i=r.MutationObserver||r.WebKitMutationObserver,a=r.process,s=r.Promise,c="process"==n("R9M2")(a);t.exports=function(){var t,e,n,u=function(){var r,o;for(c&&(r=a.domain)&&r.exit();t;){o=t.fn,t=t.next;try{o()}catch(r){throw t?n():e=void 0,r}}e=void 0,r&&r.enter()};if(c)n=function(){a.nextTick(u)};else if(!i||r.navigator&&r.navigator.standalone)if(s&&s.resolve){var f=s.resolve(void 0);n=function(){f.then(u)}}else n=function(){o.call(r,u)};else{var l=!0,d=document.createTextNode("");new i(u).observe(d,{characterData:!0}),n=function(){d.data=l=!l}}return function(r){var o={fn:r,next:void 0};e&&(e.next=o),t||(t=o,n()),e=o}}},BUHG:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=n("gyMJ"),o={name:"iflows",data:function(){return{query:{},tableData:[],multipleSelection:[],delList:[],editVisible:!1,copyVisible:!1,copyParams:{src_work_key:"",work_key:""},deleteParams:{work_key:""},pageTotal:0,form:{},idx:-1,id:-1}},created:function(){this.getData()},methods:{getData:function(){var t=this;Object(r.f)(this.query).then(function(e){t.tableData=e.data,t.pageTotal=e.data.length})},jump:function(t){this.$router.push("iworks/"+t)},handleSearch:function(){this.$set(this.query,"pageIndex",1),this.getData()},handleDelete:function(t,e){var n=this;this.deleteParams.work_key=e.key,this.$confirm("确定要删除吗？","提示",{type:"warning"}).then(function(){Object(r.d)(n.deleteParams).then(function(e){"A0001"==e.code?(n.$message.success("删除成功"),n.tableData.splice(t,1)):n.$message.error(e.msg)})}).catch(function(){})},handleSelectionChange:function(t){this.multipleSelection=t},delAllSelection:function(){var t=this.multipleSelection.length,e="";this.delList=this.delList.concat(this.multipleSelection);for(var n=0;n<t;n++)e+=this.multipleSelection[n].name+" ";this.$message.error("删除了"+e),this.multipleSelection=[]},handleEdit:function(t,e){console.log("edit",t,e),this.idx=t,this.form=e,this.editVisible=!0},handleCopy:function(t,e){this.copyParams.work_key=e.key,this.copyParams.src_work_key=e.key,this.copyVisible=!0},saveEdit:function(){this.editVisible=!1,this.$message.success("修改第 "+(this.idx+1)+" 行成功"),this.$set(this.tableData,this.idx,this.form)},saveCopy:function(){var t=this;Object(r.b)(this.copyParams).then(function(e){"A0001"==e.code?(t.$message.success("复制成功"),t.getData()):t.$message.error(e.msg)}),this.copyVisible=!1},handlePageChange:function(t){this.$set(this.query,"pageIndex",t),this.getData()}}},i={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("div",{staticClass:"container"},[n("div",{staticClass:"handle-box"},[n("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"work key"},model:{value:t.query.name,callback:function(e){t.$set(t.query,"name",e)},expression:"query.name"}}),t._v(" "),n("el-button",{attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.handleSearch}},[t._v("search")]),t._v(" "),n("el-button",{staticClass:"handle-add mr10",staticStyle:{float:"right"},attrs:{type:"primary",icon:"el-icon-circle-plus"}},[t._v("添加")])],1),t._v(" "),n("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:t.tableData,border:"","header-cell-class-name":"table-header"},on:{"selection-change":t.handleSelectionChange}},[n("el-table-column",{attrs:{type:"selection",width:"55",align:"center"}}),t._v(" "),n("el-table-column",{attrs:{prop:"Key",label:"key"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("a",{attrs:{href:"javascript:void(0)"},on:{click:function(n){return t.jump(e.row.key)}}},[t._v(t._s(e.row.key))])]}}])}),t._v(" "),n("el-table-column",{attrs:{prop:"name",label:"name"}}),t._v(" "),n("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"text",icon:"el-icon-copy-document"},on:{click:function(n){return t.handleCopy(e.$index,e.row)}}},[t._v("复制")]),t._v(" "),n("el-button",{staticClass:"red",attrs:{type:"text",icon:"el-icon-delete"},on:{click:function(n){return t.handleDelete(e.$index,e.row)}}},[t._v("删除")])]}}])})],1),t._v(" "),n("div",{staticClass:"pagination"},[n("el-pagination",{attrs:{background:"",layout:"total, prev, pager, next","current-page":t.query.pageIndex,"page-size":t.query.pageSize,total:t.pageTotal},on:{"current-change":t.handlePageChange}})],1)],1),t._v(" "),n("el-dialog",{attrs:{title:"编辑",visible:t.editVisible,width:"30%"},on:{"update:visible":function(e){t.editVisible=e}}},[n("el-form",{ref:"form",attrs:{model:t.form,"label-width":"70px"}},[n("el-form-item",{attrs:{label:"key"}},[n("el-input",{attrs:{disabled:""},model:{value:t.form.key,callback:function(e){t.$set(t.form,"key",e)},expression:"form.key"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"name"}},[n("el-input",{model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}})],1)],1),t._v(" "),n("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:function(e){t.editVisible=!1}}},[t._v("取 消")]),t._v(" "),n("el-button",{attrs:{type:"primary"},on:{click:t.saveEdit}},[t._v("确 定")])],1)],1),t._v(" "),n("el-dialog",{attrs:{title:"复制",visible:t.copyVisible,width:"30%"},on:{"update:visible":function(e){t.copyVisible=e}}},[n("el-form",{ref:"copyParams",attrs:{model:t.copyParams,"label-width":"70px"}},[n("el-form-item",{attrs:{label:"key"}},[n("el-input",{model:{value:t.copyParams.work_key,callback:function(e){t.$set(t.copyParams,"work_key",e)},expression:"copyParams.work_key"}})],1)],1),t._v(" "),n("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:function(e){t.copyVisible=!1}}},[t._v("取 消")]),t._v(" "),n("el-button",{attrs:{type:"primary"},on:{click:t.saveCopy}},[t._v("确 定")])],1)],1)],1)},staticRenderFns:[]};var a=n("VU/8")(o,i,!1,function(t){n("maGs")},"data-v-6d2b7475",null);e.default=a.exports},CXw9:function(t,e,n){"use strict";var r,o,i,a,s=n("O4g8"),c=n("7KvD"),u=n("+ZMJ"),f=n("RY/4"),l=n("kM2E"),d=n("EqjI"),p=n("lOnJ"),h=n("2KxR"),m=n("NWt+"),v=n("t8x9"),y=n("L42u").set,g=n("82Mu")(),w=n("qARP"),b=n("dNDb"),x=n("iUbK"),_=n("fJUb"),k=c.TypeError,C=c.process,P=C&&C.versions,j=P&&P.v8||"",S=c.Promise,R="process"==f(C),E=function(){},D=o=w.f,T=!!function(){try{var t=S.resolve(1),e=(t.constructor={})[n("dSzd")("species")]=function(t){t(E,E)};return(R||"function"==typeof PromiseRejectionEvent)&&t.then(E)instanceof e&&0!==j.indexOf("6.6")&&-1===x.indexOf("Chrome/66")}catch(t){}}(),A=function(t){var e;return!(!d(t)||"function"!=typeof(e=t.then))&&e},O=function(t,e){if(!t._n){t._n=!0;var n=t._c;g(function(){for(var r=t._v,o=1==t._s,i=0,a=function(e){var n,i,a,s=o?e.ok:e.fail,c=e.resolve,u=e.reject,f=e.domain;try{s?(o||(2==t._h&&G(t),t._h=1),!0===s?n=r:(f&&f.enter(),n=s(r),f&&(f.exit(),a=!0)),n===e.promise?u(k("Promise-chain cycle")):(i=A(n))?i.call(n,c,u):c(n)):u(r)}catch(t){f&&!a&&f.exit(),u(t)}};n.length>i;)a(n[i++]);t._c=[],t._n=!1,e&&!t._h&&U(t)})}},U=function(t){y.call(c,function(){var e,n,r,o=t._v,i=q(t);if(i&&(e=b(function(){R?C.emit("unhandledRejection",o,t):(n=c.onunhandledrejection)?n({promise:t,reason:o}):(r=c.console)&&r.error&&r.error("Unhandled promise rejection",o)}),t._h=R||q(t)?2:1),t._a=void 0,i&&e.e)throw e.v})},q=function(t){return 1!==t._h&&0===(t._a||t._c).length},G=function(t){y.call(c,function(){var e;R?C.emit("rejectionHandled",t):(e=c.onrejectionhandled)&&e({promise:t,reason:t._v})})},N=function(t){var e=this;e._d||(e._d=!0,(e=e._w||e)._v=t,e._s=2,e._a||(e._a=e._c.slice()),O(e,!0))},L=function(t){var e,n=this;if(!n._d){n._d=!0,n=n._w||n;try{if(n===t)throw k("Promise can't be resolved itself");(e=A(t))?g(function(){var r={_w:n,_d:!1};try{e.call(t,u(L,r,1),u(N,r,1))}catch(t){N.call(r,t)}}):(n._v=t,n._s=1,O(n,!1))}catch(t){N.call({_w:n,_d:!1},t)}}};T||(S=function(t){h(this,S,"Promise","_h"),p(t),r.call(this);try{t(u(L,this,1),u(N,this,1))}catch(t){N.call(this,t)}},(r=function(t){this._c=[],this._a=void 0,this._s=0,this._d=!1,this._v=void 0,this._h=0,this._n=!1}).prototype=n("xH/j")(S.prototype,{then:function(t,e){var n=D(v(this,S));return n.ok="function"!=typeof t||t,n.fail="function"==typeof e&&e,n.domain=R?C.domain:void 0,this._c.push(n),this._a&&this._a.push(n),this._s&&O(this,!1),n.promise},catch:function(t){return this.then(void 0,t)}}),i=function(){var t=new r;this.promise=t,this.resolve=u(L,t,1),this.reject=u(N,t,1)},w.f=D=function(t){return t===S||t===a?new i(t):o(t)}),l(l.G+l.W+l.F*!T,{Promise:S}),n("e6n0")(S,"Promise"),n("bRrM")("Promise"),a=n("FeBl").Promise,l(l.S+l.F*!T,"Promise",{reject:function(t){var e=D(this);return(0,e.reject)(t),e.promise}}),l(l.S+l.F*(s||!T),"Promise",{resolve:function(t){return _(s&&this===a?S:this,t)}}),l(l.S+l.F*!(T&&n("dY0y")(function(t){S.all(t).catch(E)})),"Promise",{all:function(t){var e=this,n=D(e),r=n.resolve,o=n.reject,i=b(function(){var n=[],i=0,a=1;m(t,!1,function(t){var s=i++,c=!1;n.push(void 0),a++,e.resolve(t).then(function(t){c||(c=!0,n[s]=t,--a||r(n))},o)}),--a||r(n)});return i.e&&o(i.v),n.promise},race:function(t){var e=this,n=D(e),r=n.reject,o=b(function(){m(t,!1,function(t){e.resolve(t).then(n.resolve,r)})});return o.e&&r(o.v),n.promise}})},DQCr:function(t,e,n){"use strict";var r=n("cGG2");function o(t){return encodeURIComponent(t).replace(/%40/gi,"@").replace(/%3A/gi,":").replace(/%24/g,"$").replace(/%2C/gi,",").replace(/%20/g,"+").replace(/%5B/gi,"[").replace(/%5D/gi,"]")}t.exports=function(t,e,n){if(!e)return t;var i;if(n)i=n(e);else if(r.isURLSearchParams(e))i=e.toString();else{var a=[];r.forEach(e,function(t,e){null!==t&&void 0!==t&&(r.isArray(t)?e+="[]":t=[t],r.forEach(t,function(t){r.isDate(t)?t=t.toISOString():r.isObject(t)&&(t=JSON.stringify(t)),a.push(o(e)+"="+o(t))}))}),i=a.join("&")}if(i){var s=t.indexOf("#");-1!==s&&(t=t.slice(0,s)),t+=(-1===t.indexOf("?")?"?":"&")+i}return t}},DUeU:function(t,e,n){"use strict";var r=n("cGG2");t.exports=function(t,e){e=e||{};var n={};return r.forEach(["url","method","params","data"],function(t){void 0!==e[t]&&(n[t]=e[t])}),r.forEach(["headers","auth","proxy"],function(o){r.isObject(e[o])?n[o]=r.deepMerge(t[o],e[o]):void 0!==e[o]?n[o]=e[o]:r.isObject(t[o])?n[o]=r.deepMerge(t[o]):void 0!==t[o]&&(n[o]=t[o])}),r.forEach(["baseURL","transformRequest","transformResponse","paramsSerializer","timeout","withCredentials","adapter","responseType","xsrfCookieName","xsrfHeaderName","onUploadProgress","onDownloadProgress","maxContentLength","validateStatus","maxRedirects","httpAgent","httpsAgent","cancelToken","socketPath"],function(r){void 0!==e[r]?n[r]=e[r]:void 0!==t[r]&&(n[r]=t[r])}),n}},EqBC:function(t,e,n){"use strict";var r=n("kM2E"),o=n("FeBl"),i=n("7KvD"),a=n("t8x9"),s=n("fJUb");r(r.P+r.R,"Promise",{finally:function(t){var e=a(this,o.Promise||i.Promise),n="function"==typeof t;return this.then(n?function(n){return s(e,t()).then(function(){return n})}:t,n?function(n){return s(e,t()).then(function(){throw n})}:t)}})},FtD3:function(t,e,n){"use strict";var r=n("t8qj");t.exports=function(t,e,n,o,i){var a=new Error(t);return r(a,e,n,o,i)}},GHBc:function(t,e,n){"use strict";var r=n("cGG2");t.exports=r.isStandardBrowserEnv()?function(){var t,e=/(msie|trident)/i.test(navigator.userAgent),n=document.createElement("a");function o(t){var r=t;return e&&(n.setAttribute("href",r),r=n.href),n.setAttribute("href",r),{href:n.href,protocol:n.protocol?n.protocol.replace(/:$/,""):"",host:n.host,search:n.search?n.search.replace(/^\?/,""):"",hash:n.hash?n.hash.replace(/^#/,""):"",hostname:n.hostname,port:n.port,pathname:"/"===n.pathname.charAt(0)?n.pathname:"/"+n.pathname}}return t=o(window.location.href),function(e){var n=r.isString(e)?o(e):e;return n.protocol===t.protocol&&n.host===t.host}}():function(){return!0}},"JP+z":function(t,e,n){"use strict";t.exports=function(t,e){return function(){for(var n=new Array(arguments.length),r=0;r<n.length;r++)n[r]=arguments[r];return t.apply(e,n)}}},KCLY:function(t,e,n){"use strict";(function(e){var r=n("cGG2"),o=n("5VQ+"),i={"Content-Type":"application/x-www-form-urlencoded"};function a(t,e){!r.isUndefined(t)&&r.isUndefined(t["Content-Type"])&&(t["Content-Type"]=e)}var s,c={adapter:(void 0!==e&&"[object process]"===Object.prototype.toString.call(e)?s=n("7GwW"):"undefined"!=typeof XMLHttpRequest&&(s=n("7GwW")),s),transformRequest:[function(t,e){return o(e,"Accept"),o(e,"Content-Type"),r.isFormData(t)||r.isArrayBuffer(t)||r.isBuffer(t)||r.isStream(t)||r.isFile(t)||r.isBlob(t)?t:r.isArrayBufferView(t)?t.buffer:r.isURLSearchParams(t)?(a(e,"application/x-www-form-urlencoded;charset=utf-8"),t.toString()):r.isObject(t)?(a(e,"application/json;charset=utf-8"),JSON.stringify(t)):t}],transformResponse:[function(t){if("string"==typeof t)try{t=JSON.parse(t)}catch(t){}return t}],timeout:0,xsrfCookieName:"XSRF-TOKEN",xsrfHeaderName:"X-XSRF-TOKEN",maxContentLength:-1,validateStatus:function(t){return t>=200&&t<300}};c.headers={common:{Accept:"application/json, text/plain, */*"}},r.forEach(["delete","get","head"],function(t){c.headers[t]={}}),r.forEach(["post","put","patch"],function(t){c.headers[t]=r.merge(i)}),t.exports=c}).call(e,n("W2nU"))},L42u:function(t,e,n){var r,o,i,a=n("+ZMJ"),s=n("knuC"),c=n("RPLV"),u=n("ON07"),f=n("7KvD"),l=f.process,d=f.setImmediate,p=f.clearImmediate,h=f.MessageChannel,m=f.Dispatch,v=0,y={},g=function(){var t=+this;if(y.hasOwnProperty(t)){var e=y[t];delete y[t],e()}},w=function(t){g.call(t.data)};d&&p||(d=function(t){for(var e=[],n=1;arguments.length>n;)e.push(arguments[n++]);return y[++v]=function(){s("function"==typeof t?t:Function(t),e)},r(v),v},p=function(t){delete y[t]},"process"==n("R9M2")(l)?r=function(t){l.nextTick(a(g,t,1))}:m&&m.now?r=function(t){m.now(a(g,t,1))}:h?(i=(o=new h).port2,o.port1.onmessage=w,r=a(i.postMessage,i,1)):f.addEventListener&&"function"==typeof postMessage&&!f.importScripts?(r=function(t){f.postMessage(t+"","*")},f.addEventListener("message",w,!1)):r="onreadystatechange"in u("script")?function(t){c.appendChild(u("script")).onreadystatechange=function(){c.removeChild(this),g.call(t)}}:function(t){setTimeout(a(g,t,1),0)}),t.exports={set:d,clear:p}},Mhyx:function(t,e,n){var r=n("/bQp"),o=n("dSzd")("iterator"),i=Array.prototype;t.exports=function(t){return void 0!==t&&(r.Array===t||i[o]===t)}},"NWt+":function(t,e,n){var r=n("+ZMJ"),o=n("msXi"),i=n("Mhyx"),a=n("77Pl"),s=n("QRG4"),c=n("3fs2"),u={},f={};(e=t.exports=function(t,e,n,l,d){var p,h,m,v,y=d?function(){return t}:c(t),g=r(n,l,e?2:1),w=0;if("function"!=typeof y)throw TypeError(t+" is not iterable!");if(i(y)){for(p=s(t.length);p>w;w++)if((v=e?g(a(h=t[w])[0],h[1]):g(t[w]))===u||v===f)return v}else for(m=y.call(t);!(h=m.next()).done;)if((v=o(m,g,h.value,e))===u||v===f)return v}).BREAK=u,e.RETURN=f},QaAC:function(t,e){},"RY/4":function(t,e,n){var r=n("R9M2"),o=n("dSzd")("toStringTag"),i="Arguments"==r(function(){return arguments}());t.exports=function(t){var e,n,a;return void 0===t?"Undefined":null===t?"Null":"string"==typeof(n=function(t,e){try{return t[e]}catch(t){}}(e=Object(t),o))?n:i?r(e):"Object"==(a=r(e))&&"function"==typeof e.callee?"Arguments":a}},TNV1:function(t,e,n){"use strict";var r=n("cGG2");t.exports=function(t,e,n){return r.forEach(n,function(n){t=n(t,e)}),t}},U5ju:function(t,e,n){n("M6a0"),n("zQR9"),n("+tPU"),n("CXw9"),n("EqBC"),n("jKW+"),t.exports=n("FeBl").Promise},W2nU:function(t,e){var n,r,o=t.exports={};function i(){throw new Error("setTimeout has not been defined")}function a(){throw new Error("clearTimeout has not been defined")}function s(t){if(n===setTimeout)return setTimeout(t,0);if((n===i||!n)&&setTimeout)return n=setTimeout,setTimeout(t,0);try{return n(t,0)}catch(e){try{return n.call(null,t,0)}catch(e){return n.call(this,t,0)}}}!function(){try{n="function"==typeof setTimeout?setTimeout:i}catch(t){n=i}try{r="function"==typeof clearTimeout?clearTimeout:a}catch(t){r=a}}();var c,u=[],f=!1,l=-1;function d(){f&&c&&(f=!1,c.length?u=c.concat(u):l=-1,u.length&&p())}function p(){if(!f){var t=s(d);f=!0;for(var e=u.length;e;){for(c=u,u=[];++l<e;)c&&c[l].run();l=-1,e=u.length}c=null,f=!1,function(t){if(r===clearTimeout)return clearTimeout(t);if((r===a||!r)&&clearTimeout)return r=clearTimeout,clearTimeout(t);try{r(t)}catch(e){try{return r.call(null,t)}catch(e){return r.call(this,t)}}}(t)}}function h(t,e){this.fun=t,this.array=e}function m(){}o.nextTick=function(t){var e=new Array(arguments.length-1);if(arguments.length>1)for(var n=1;n<arguments.length;n++)e[n-1]=arguments[n];u.push(new h(t,e)),1!==u.length||f||s(p)},h.prototype.run=function(){this.fun.apply(null,this.array)},o.title="browser",o.browser=!0,o.env={},o.argv=[],o.version="",o.versions={},o.on=m,o.addListener=m,o.once=m,o.off=m,o.removeListener=m,o.removeAllListeners=m,o.emit=m,o.prependListener=m,o.prependOnceListener=m,o.listeners=function(t){return[]},o.binding=function(t){throw new Error("process.binding is not supported")},o.cwd=function(){return"/"},o.chdir=function(t){throw new Error("process.chdir is not supported")},o.umask=function(){return 0}},XmWM:function(t,e,n){"use strict";var r=n("cGG2"),o=n("DQCr"),i=n("fuGk"),a=n("xLtR"),s=n("DUeU");function c(t){this.defaults=t,this.interceptors={request:new i,response:new i}}c.prototype.request=function(t){"string"==typeof t?(t=arguments[1]||{}).url=arguments[0]:t=t||{},(t=s(this.defaults,t)).method=t.method?t.method.toLowerCase():"get";var e=[a,void 0],n=Promise.resolve(t);for(this.interceptors.request.forEach(function(t){e.unshift(t.fulfilled,t.rejected)}),this.interceptors.response.forEach(function(t){e.push(t.fulfilled,t.rejected)});e.length;)n=n.then(e.shift(),e.shift());return n},c.prototype.getUri=function(t){return t=s(this.defaults,t),o(t.url,t.params,t.paramsSerializer).replace(/^\?/,"")},r.forEach(["delete","get","head","options"],function(t){c.prototype[t]=function(e,n){return this.request(r.merge(n||{},{method:t,url:e}))}}),r.forEach(["post","put","patch"],function(t){c.prototype[t]=function(e,n,o){return this.request(r.merge(o||{},{method:t,url:e,data:n}))}}),t.exports=c},bRrM:function(t,e,n){"use strict";var r=n("7KvD"),o=n("FeBl"),i=n("evD5"),a=n("+E39"),s=n("dSzd")("species");t.exports=function(t){var e="function"==typeof o[t]?o[t]:r[t];a&&e&&!e[s]&&i.f(e,s,{configurable:!0,get:function(){return this}})}},cGG2:function(t,e,n){"use strict";var r=n("JP+z"),o=n("1Yoh"),i=Object.prototype.toString;function a(t){return"[object Array]"===i.call(t)}function s(t){return null!==t&&"object"==typeof t}function c(t){return"[object Function]"===i.call(t)}function u(t,e){if(null!==t&&void 0!==t)if("object"!=typeof t&&(t=[t]),a(t))for(var n=0,r=t.length;n<r;n++)e.call(null,t[n],n,t);else for(var o in t)Object.prototype.hasOwnProperty.call(t,o)&&e.call(null,t[o],o,t)}t.exports={isArray:a,isArrayBuffer:function(t){return"[object ArrayBuffer]"===i.call(t)},isBuffer:o,isFormData:function(t){return"undefined"!=typeof FormData&&t instanceof FormData},isArrayBufferView:function(t){return"undefined"!=typeof ArrayBuffer&&ArrayBuffer.isView?ArrayBuffer.isView(t):t&&t.buffer&&t.buffer instanceof ArrayBuffer},isString:function(t){return"string"==typeof t},isNumber:function(t){return"number"==typeof t},isObject:s,isUndefined:function(t){return void 0===t},isDate:function(t){return"[object Date]"===i.call(t)},isFile:function(t){return"[object File]"===i.call(t)},isBlob:function(t){return"[object Blob]"===i.call(t)},isFunction:c,isStream:function(t){return s(t)&&c(t.pipe)},isURLSearchParams:function(t){return"undefined"!=typeof URLSearchParams&&t instanceof URLSearchParams},isStandardBrowserEnv:function(){return("undefined"==typeof navigator||"ReactNative"!==navigator.product&&"NativeScript"!==navigator.product&&"NS"!==navigator.product)&&"undefined"!=typeof window&&"undefined"!=typeof document},forEach:u,merge:function t(){var e={};function n(n,r){"object"==typeof e[r]&&"object"==typeof n?e[r]=t(e[r],n):e[r]=n}for(var r=0,o=arguments.length;r<o;r++)u(arguments[r],n);return e},deepMerge:function t(){var e={};function n(n,r){"object"==typeof e[r]&&"object"==typeof n?e[r]=t(e[r],n):e[r]="object"==typeof n?t({},n):n}for(var r=0,o=arguments.length;r<o;r++)u(arguments[r],n);return e},extend:function(t,e,n){return u(e,function(e,o){t[o]=n&&"function"==typeof e?r(e,n):e}),t},trim:function(t){return t.replace(/^\s*/,"").replace(/\s*$/,"")}}},cWxy:function(t,e,n){"use strict";var r=n("dVOP");function o(t){if("function"!=typeof t)throw new TypeError("executor must be a function.");var e;this.promise=new Promise(function(t){e=t});var n=this;t(function(t){n.reason||(n.reason=new r(t),e(n.reason))})}o.prototype.throwIfRequested=function(){if(this.reason)throw this.reason},o.source=function(){var t;return{token:new o(function(e){t=e}),cancel:t}},t.exports=o},dIwP:function(t,e,n){"use strict";t.exports=function(t){return/^([a-z][a-z\d\+\-\.]*:)?\/\//i.test(t)}},dNDb:function(t,e){t.exports=function(t){try{return{e:!1,v:t()}}catch(t){return{e:!0,v:t}}}},dVOP:function(t,e,n){"use strict";function r(t){this.message=t}r.prototype.toString=function(){return"Cancel"+(this.message?": "+this.message:"")},r.prototype.__CANCEL__=!0,t.exports=r},dY0y:function(t,e,n){var r=n("dSzd")("iterator"),o=!1;try{var i=[7][r]();i.return=function(){o=!0},Array.from(i,function(){throw 2})}catch(t){}t.exports=function(t,e){if(!e&&!o)return!1;var n=!1;try{var i=[7],a=i[r]();a.next=function(){return{done:n=!0}},i[r]=function(){return a},t(i)}catch(t){}return n}},fJUb:function(t,e,n){var r=n("77Pl"),o=n("EqjI"),i=n("qARP");t.exports=function(t,e){if(r(t),o(e)&&e.constructor===t)return e;var n=i.f(t);return(0,n.resolve)(e),n.promise}},fuGk:function(t,e,n){"use strict";var r=n("cGG2");function o(){this.handlers=[]}o.prototype.use=function(t,e){return this.handlers.push({fulfilled:t,rejected:e}),this.handlers.length-1},o.prototype.eject=function(t){this.handlers[t]&&(this.handlers[t]=null)},o.prototype.forEach=function(t){r.forEach(this.handlers,function(e){null!==e&&t(e)})},t.exports=o},gyMJ:function(t,e,n){"use strict";var r=n("//Fk"),o=n.n(r),i=n("mtWM"),a=n.n(i).a.create({baseURL:""+window.location.origin,timeout:5e3});a.interceptors.request.use(function(t){return t},function(t){return console.log(t),o.a.reject()}),a.interceptors.response.use(function(t){if(200===t.status)return t.data;o.a.reject()},function(t){return console.log(t),o.a.reject()});var s=a;n.d(e,"e",function(){return c}),n.d(e,"f",function(){return u}),n.d(e,"g",function(){return f}),n.d(e,"h",function(){return l}),n.d(e,"a",function(){return d}),n.d(e,"b",function(){return p}),n.d(e,"c",function(){return h}),n.d(e,"d",function(){return m});var c=function(t){return s({url:"/api/config/flow/query/",method:"post",data:t})},u=function(t){return s({url:"/api/config/work/query/",method:"post",data:t})},f=function(t){return s({url:"/api/config/flow/update/",method:"post",data:t})},l=function(t){return s({url:"/api/config/work/update/",method:"post",data:t})},d=function(t){return s({url:"/api/config/flow/copy/",method:"post",data:t})},p=function(t){return s({url:"/api/config/work/copy/",method:"post",data:t})},h=function(t){return s({url:"/api/config/flow/delete/",method:"post",data:t})},m=function(t){return s({url:"/api/config/work/delete/",method:"post",data:t})}},iUbK:function(t,e,n){var r=n("7KvD").navigator;t.exports=r&&r.userAgent||""},"jKW+":function(t,e,n){"use strict";var r=n("kM2E"),o=n("qARP"),i=n("dNDb");r(r.S,"Promise",{try:function(t){var e=o.f(this),n=i(t);return(n.e?e.reject:e.resolve)(n.v),e.promise}})},knuC:function(t,e){t.exports=function(t,e,n){var r=void 0===n;switch(e.length){case 0:return r?t():t.call(n);case 1:return r?t(e[0]):t.call(n,e[0]);case 2:return r?t(e[0],e[1]):t.call(n,e[0],e[1]);case 3:return r?t(e[0],e[1],e[2]):t.call(n,e[0],e[1],e[2]);case 4:return r?t(e[0],e[1],e[2],e[3]):t.call(n,e[0],e[1],e[2],e[3])}return t.apply(n,e)}},maGs:function(t,e){},msXi:function(t,e,n){var r=n("77Pl");t.exports=function(t,e,n,o){try{return o?e(r(n)[0],n[1]):e(n)}catch(e){var i=t.return;throw void 0!==i&&r(i.call(t)),e}}},mtWM:function(t,e,n){t.exports=n("tIFN")},oJlt:function(t,e,n){"use strict";var r=n("cGG2"),o=["age","authorization","content-length","content-type","etag","expires","from","host","if-modified-since","if-unmodified-since","last-modified","location","max-forwards","proxy-authorization","referer","retry-after","user-agent"];t.exports=function(t){var e,n,i,a={};return t?(r.forEach(t.split("\n"),function(t){if(i=t.indexOf(":"),e=r.trim(t.substr(0,i)).toLowerCase(),n=r.trim(t.substr(i+1)),e){if(a[e]&&o.indexOf(e)>=0)return;a[e]="set-cookie"===e?(a[e]?a[e]:[]).concat([n]):a[e]?a[e]+", "+n:n}}),a):a}},p1b6:function(t,e,n){"use strict";var r=n("cGG2");t.exports=r.isStandardBrowserEnv()?{write:function(t,e,n,o,i,a){var s=[];s.push(t+"="+encodeURIComponent(e)),r.isNumber(n)&&s.push("expires="+new Date(n).toGMTString()),r.isString(o)&&s.push("path="+o),r.isString(i)&&s.push("domain="+i),!0===a&&s.push("secure"),document.cookie=s.join("; ")},read:function(t){var e=document.cookie.match(new RegExp("(^|;\\s*)("+t+")=([^;]*)"));return e?decodeURIComponent(e[3]):null},remove:function(t){this.write(t,"",Date.now()-864e5)}}:{write:function(){},read:function(){return null},remove:function(){}}},pBtG:function(t,e,n){"use strict";t.exports=function(t){return!(!t||!t.__CANCEL__)}},pxG4:function(t,e,n){"use strict";t.exports=function(t){return function(e){return t.apply(null,e)}}},qARP:function(t,e,n){"use strict";var r=n("lOnJ");t.exports.f=function(t){return new function(t){var e,n;this.promise=new t(function(t,r){if(void 0!==e||void 0!==n)throw TypeError("Bad Promise constructor");e=t,n=r}),this.resolve=r(e),this.reject=r(n)}(t)}},qRfI:function(t,e,n){"use strict";t.exports=function(t,e){return e?t.replace(/\/+$/,"")+"/"+e.replace(/^\/+/,""):t}},t8qj:function(t,e,n){"use strict";t.exports=function(t,e,n,r,o){return t.config=e,n&&(t.code=n),t.request=r,t.response=o,t.isAxiosError=!0,t.toJSON=function(){return{message:this.message,name:this.name,description:this.description,number:this.number,fileName:this.fileName,lineNumber:this.lineNumber,columnNumber:this.columnNumber,stack:this.stack,config:this.config,code:this.code}},t}},t8x9:function(t,e,n){var r=n("77Pl"),o=n("lOnJ"),i=n("dSzd")("species");t.exports=function(t,e){var n,a=r(t).constructor;return void 0===a||void 0==(n=r(a)[i])?e:o(n)}},tIFN:function(t,e,n){"use strict";var r=n("cGG2"),o=n("JP+z"),i=n("XmWM"),a=n("DUeU");function s(t){var e=new i(t),n=o(i.prototype.request,e);return r.extend(n,i.prototype,e),r.extend(n,e),n}var c=s(n("KCLY"));c.Axios=i,c.create=function(t){return s(a(c.defaults,t))},c.Cancel=n("dVOP"),c.CancelToken=n("cWxy"),c.isCancel=n("pBtG"),c.all=function(t){return Promise.all(t)},c.spread=n("pxG4"),t.exports=c,t.exports.default=c},w7KR:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=n("mvHQ"),o=n.n(r),i=n("gyMJ"),a={name:"flowdetail",data:function(){return{workData:{},query:{work_key:this.$route.params.workKey},form:{work_name:"",work_key:"",work_desc:"",config:""},options:{mode:"code",indentation:4,search:!0},config_editor:null}},methods:{getWorkDetailData:function(){var t=this;Object(i.f)(this.query).then(function(e){if(t.workData=e.data[0],console.log("work Data: ",t.workData),t.workData){t.form.work_key=t.workData.key,t.form.work_name=t.workData.name,t.form.work_desc=t.workData.desc;var n=document.getElementById("config_editor");t.config_editor=new t.$jsoneditor(n,t.options),t.config_editor.set(t.workData.config)}})},onSubmit:function(){var t=this;this.form.config=o()(this.config_editor.get()),Object(i.h)(this.form).then(function(e){"A0001"==e.code?t.$message.success("更新成功！"):t.$message.error(e.msg)})}},created:function(){this.getWorkDetailData()}},s={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"form-box"},[n("el-form",{ref:"form",attrs:{model:t.form,"label-width":"80px"}},[n("el-form-item",{attrs:{label:"工作key"}},[n("el-input",{attrs:{disabled:!0},model:{value:t.form.work_key,callback:function(e){t.$set(t.form,"work_key",e)},expression:"form.work_key"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"工作名称"}},[n("el-input",{model:{value:t.form.work_name,callback:function(e){t.$set(t.form,"work_name",e)},expression:"form.work_name"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"描述"}},[n("el-input",{attrs:{type:"textarea",rows:"5"},model:{value:t.form.work_desc,callback:function(e){t.$set(t.form,"work_desc",e)},expression:"form.work_desc"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"工作配置"}},[n("div",{staticStyle:{height:"500px"},attrs:{id:"config_editor"}})]),t._v(" "),n("el-form-item",[n("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("更新")])],1)],1)],1)},staticRenderFns:[]},c={name:"Work",data:function(){return{active:0,currentType:"detail",currentView:"flowdetail",tabs:[{type:"detail",view:"flowdetail"}]}},methods:{toggle:function(t,e,n){console.log("toggle",t,e,n),this.active=t,this.currentView=e,this.currentType=n}},components:{flowdetail:n("VU/8")(a,s,!1,null,null,null).exports}},u={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("div",{staticClass:"container"},[n("ul",{staticClass:"innertab"},t._l(t.tabs,function(e,r){return n("li",{staticClass:"innertab",class:{active:t.active==r},on:{click:function(n){return t.toggle(r,e.view,e.type)}}},[t._v("\n\t\t\t\t"+t._s(e.type)+"\n\t\t\t")])}),0),t._v(" "),n("hr"),t._v(" "),n(t.currentView,{tag:"component",staticStyle:{"padding-top":"20px"}})],1)])},staticRenderFns:[]};var f=n("VU/8")(c,u,!1,function(t){n("QaAC")},null,null);e.default=f.exports},"xH/j":function(t,e,n){var r=n("hJx8");t.exports=function(t,e,n){for(var o in e)n&&t[o]?t[o]=e[o]:r(t,o,e[o]);return t}},xLtR:function(t,e,n){"use strict";var r=n("cGG2"),o=n("TNV1"),i=n("pBtG"),a=n("KCLY"),s=n("dIwP"),c=n("qRfI");function u(t){t.cancelToken&&t.cancelToken.throwIfRequested()}t.exports=function(t){return u(t),t.baseURL&&!s(t.url)&&(t.url=c(t.baseURL,t.url)),t.headers=t.headers||{},t.data=o(t.data,t.headers,t.transformRequest),t.headers=r.merge(t.headers.common||{},t.headers[t.method]||{},t.headers||{}),r.forEach(["delete","get","head","post","put","patch","common"],function(e){delete t.headers[e]}),(t.adapter||a.adapter)(t).then(function(e){return u(t),e.data=o(e.data,e.headers,t.transformResponse),e},function(e){return i(e)||(u(t),e&&e.response&&(e.response.data=o(e.response.data,e.response.headers,t.transformResponse))),Promise.reject(e)})}}});
//# sourceMappingURL=1.6ff44210b7296a1e9875.js.map