<template>
	<div class="header">
		<div class="logo">Termite</div>
		<div class="menu">
			<el-menu
				:default-active="onRoutes"
				class="el-menu-demo"
				mode="horizontal"
				background-color="#545c64"
				text-color="#fff"
				active-text-color="#ffd04b"
				router
			>
				<el-menu-item index="/dashboard">Home</el-menu-item>
				<el-menu-item index="/iworks">Work</el-menu-item>
				<el-menu-item index="/iflows">Flow</el-menu-item>
			</el-menu>
		</div>
		<div class="header-right">
			<div class="header-user-con">
				<!-- 用户名下拉菜单 -->
				<el-dropdown
					class="user-name"
					trigger="click"
					@command="handleCommand"
				>
					<span class="el-dropdown-link">
						{{ username }}
						<i class="el-icon-caret-bottom"></i>
					</span>
					<el-dropdown-menu slot="dropdown">
						<el-dropdown-item divided command="loginout"
							>退出登录</el-dropdown-item
						>
					</el-dropdown-menu>
				</el-dropdown>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	data() {
		return {
			activeIndex: "1",
			name: "linxin"
		};
	},
	computed: {
		username() {
			let username = localStorage.getItem("ms_username");
			return username ? username : this.name;
		},
        onRoutes() {
		    if (this.$route.path.includes("iflows")) {
		        return "/iflows";
			}
            else if (this.$route.path.includes("iworks")) {
                return "/iworks";
            }
            else if (this.$route.path.includes("dashboard")) {
                return "/dashboard"
			} else {
                return this.$route.path;
			}
        }
	},
	methods: {
		// 用户名下拉菜单选择事件
		handleCommand(command) {
			if (command == "loginout") {
				localStorage.removeItem("ms_username");
				this.$router.push("/login");
			}
		},
	}
};
</script>
<style scoped>
.header {
	position: relative;
	box-sizing: border-box;
	width: 100%;
	height: 60px;
	font-size: 22px;
	color: #fff;
}
.header .menu {
	float: left;
}
.header .logo {
	float: left;
	width: 20%;
	line-height: 60px;
	padding-left: 3%;
}
.header-right {
	float: right;
	padding-right: 50px;
}
.header-user-con {
	display: flex;
	height: 60px;
	align-items: center;
}
.btn-fullscreen {
	transform: rotate(45deg);
	margin-right: 5px;
	font-size: 24px;
}
.btn-bell,
.btn-fullscreen {
	position: relative;
	width: 30px;
	height: 30px;
	text-align: center;
	border-radius: 15px;
	cursor: pointer;
}
.btn-bell-badge {
	position: absolute;
	right: 0;
	top: -2px;
	width: 8px;
	height: 8px;
	border-radius: 4px;
	background: #f56c6c;
	color: #fff;
}
.btn-bell .el-icon-bell {
	color: #fff;
}
.user-name {
	margin-left: 10px;
}
.user-avator {
	margin-left: 20px;
}
.user-avator img {
	display: block;
	width: 40px;
	height: 40px;
	border-radius: 50%;
}
.el-dropdown-link {
	color: #fff;
	cursor: pointer;
}
.el-dropdown-menu__item {
	text-align: center;
}
</style>
