import Vue from "vue";
import Router from "vue-router";

const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
	return originalPush.call(this, location).catch(err => err)
}

Vue.use(Router);

export default new Router({
	routes: [
		{
			path: "/",
			redirect: "/dashboard"
		},
		{
			path: "/",
			component: () =>
				import(
					/* webpackChunkName: "home" */ "../components/common/Home.vue"
				),
			meta: { title: "Termite" },
			children: [
				{
					path: "/dashboard",
					component: () =>
						import(
							/* webpackChunkName: "dashboard" */ "../components/pages/Dashboard.vue"
						),
					meta: { title: "admin" }
				},
				{
					path: '/iflows',
					component: () => import(/* webpackChunkName: "flows" */ '../components/pages/Flows.vue'),
					meta: { title: 'Flows' }
				},
				{
					path: '/iworks',
					component: () => import(/* webpackChunkName: "works" */ '../components/pages/Works.vue'),
					meta: { title: 'Works' }
				},
				{
					path: '/iflows/:flowKey',
					component: () => import(/* webpackChunkName: "flows" */ '../components/pages/FlowPage.vue'),
					meta: { title: 'Flow' }
				},
				{
					path: '/iworks/:workKey',
					component: () => import(/* webpackChunkName: "works" */ '../components/pages/WorkPage.vue'),
					meta: { title: 'Work' }
				}
			]
		},
		{
			path: '/login',
			component: () => import(/* webpackChunkName: "login" */ '../components/pages/Login.vue')
		},
		{
			path: '*',
			redirect: '/404'
		}
	]
});
