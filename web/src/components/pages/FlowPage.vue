<template>
    <div>
        <div class="container">
			<el-row>
				<el-col :span="20">
					<ul class="innertab">
						<li class="innertab" v-for="(tab,index) in tabs" @click="toggle(index,tab.view, tab.type)" :class="{active:active==index}">
							{{tab.type}}
						</li>
					</ul>
				</el-col>
				<el-col :span="3" style="text-align: right">
					<h5>scheduling switch</h5>
				</el-col>
				<el-col :span="1" style="text-align: right;">
					<el-switch
						v-model="iswitch"
						@change="updateFlowSwitch"
						active-color="#13ce66"
						inactive-color="#ff4949">
					</el-switch>
				</el-col>
			</el-row>
            <hr>
            <component style="padding-top: 20px" :is="currentView"></component>
        </div>
    </div>
</template>

<script>
import flowdetail from '../common/FlowDetail'
import flowdag from '../common/FlowDag'
import flowinstance from '../common/FlowInstance'
import flowInsAdd from '../common/FlowInsCreate'
import { queryFlowConfigs } from "../../api/index";
import { updateFlowSwitch } from "../../api/index";
export default {
    name: 'Flow',
    data() {
        return {
            iswitch: true,
            active:0,
            query: {
                "vflow": this.$route.params.flowKey
            },
            upuery: {
                "vflow": this.$route.params.flowKey,
                "switch": 1
            },
            currentType: 'detail',
            currentView:'flowdetail',
            tabs:[
                {
                    type:'details',
                    view: 'flowdetail'
                },
                {
                    type: 'dag',
                    view: 'flowdag'
                },
                {
                    type: 'instance',
                    view: 'flowinstance'
                },
                {
                    type: 'create',
                    view: 'flowInsAdd'
                }
            ],
        };
    },
    methods:{
        getFlowSwitch() {
            queryFlowConfigs(this.query).then(res => {
                console.log("get switch data", res.data.configs[0], res.data.configs[0].switch)
                if (res.data.configs[0].switch == 1) {
                    this.iswitch = true
				} else {
                    this.iswitch = false
				}
            });
        },
		updateFlowSwitch() {
            console.log(this.iswitch)
            if (this.iswitch == true) {
                this.upuery.switch = 1
			} else {
                this.upuery.switch = 0
			}
            updateFlowSwitch(this.upuery).then(res => {
                if (res.code == 'A0001') {
                    this.$message.success("更新成功!")
                } else {
                    this.$message.error("更新失败!")
                }
			})
		},
        toggle(i,v, t){
            console.log("toggle", i, v, t)
            this.active=i;
            this.currentView=v;
            this.currentType=t;
        },
    },
    components:{
        flowdag,
        flowdetail,
		flowinstance,
        flowInsAdd,
    },
    created() {
        this.getFlowSwitch();
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
