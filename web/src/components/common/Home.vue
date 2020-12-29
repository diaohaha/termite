<template>
	<div class="wrapper">
		<v-header></v-header>
		<div class="content-box">
			<div class="content">
				<transition name="move" mode="out-in">
					<router-view></router-view>
				</transition>
			</div>
		</div>
	</div>
</template>

<script>
import vHeader from "./Header";
export default {
	data() {
		return {};
	},
	components: {
		vHeader
	}
};
</script>

<style></style>
