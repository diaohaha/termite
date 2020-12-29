<template>
	<div>
		<div class="container">
			<ul class="innertab">
				<li class="innertab" v-for="(tab,index) in tabs" @click="toggle(index,tab.view, tab.type)" :class="{active:active==index}">
					{{tab.type}}
				</li>
			</ul>
			<hr>
			<component style="padding-top: 20px" :is="currentView"></component>
		</div>
	</div>
</template>

<script>
    import workDetail from '../common/WorkDetail'
    import workInstance from '../common/WorkInstance'
    export default {
        name: 'Work',
        data() {
            return {
                active:0,
                currentType: 'detail',
                currentView:'workDetail',
                tabs:[
                    {
                        type: 'detail',
                        view: 'workDetail'
                    },
                    {
                        type: 'instance',
                        view: 'workInstance'
                    }
                ],
            };
        },
        methods:{
            toggle(i,v, t){
                console.log("toggle", i, v, t)
                this.active=i;
                this.currentView=v;
                this.currentType=t;
            }
        },
        components:{
            workDetail,
            workInstance
        }
    };
</script>

<style>
	ul.innertab{
		width:200px;
		display:flex;
		padding-bottom: 10px;
	}
	ul li.innertab{
		width:300px;
		height:20px;
		padding-left: 10px;
		padding-right: 10px;
		display: inline-flex;
		border-right:1px solid #ddd;
		justify-content: center;
		align-items: center;
		cursor:pointer
	}
	ul li.active{
		color: #00a854;
	}
</style>
