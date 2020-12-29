// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css'; // 默认主题
import 'jsoneditor/dist/jsoneditor.min.css'
import DAGBoard from 'dag-board'
import jsoneditor from 'jsoneditor'

Vue.prototype.$jsoneditor = jsoneditor

Vue.use(DAGBoard)


Vue.config.productionTip = false;
Vue.use(ElementUI, {
	size: 'small'
});

Vue.config.productionTip = false

/* eslint-disable no-new */
// new Vue({
//   el: '#app',
//   router,
//   components: { App },
//   template: '<App/>'
// })
new Vue({
	router,
	render: h => h(App)
}).$mount('#app');
