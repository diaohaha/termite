<template>
    <div class="svg">
        <DAGBoard :DataAll="svgData"></DAGBoard>
    </div>
</template>

<script>
    import { fetchFlows } from "../../api/index"
    import { fetchWorksConfig } from "../../api/index"
    export default {
        name: 'flowdag',
        data() {
            return {
                workData: {},
                flowData: {},
                svgData: {"edges": [], "nodes": []},
                dquery: {
                    "flow_key": this.$route.params.flowKey
                },
				wquery: {},
            };
        },
        methods: {
            getWorkData() {
                fetchWorksConfig(this.wquery).then(res => {
                    res.data.forEach(
                        workConfig => {
                            this.workData[workConfig.key] = workConfig
						}
					)
				})
			},
            getFlowDagData() {
                fetchFlows(this.dquery).then(res => {
                    console.log("this is res")
                    console.log(res)
                    this.flowData = res.data[0];
                    console.log("this is flow data")
                    console.log(this.flowData)
                    if (this.flowData) {
                        this.svgData["nodes"] = []
                        let svgDataNodes = []
                        let svgDataEdges = []
                        let nodeIdMap = {}
                        let svgNodeYLevel = {}
                        let svgNodeXLevel = {}
                        let myPointSet = this.pointSet
						this.flowData.config.works.forEach(
						    work => {
						        svgNodeYLevel[work] = 1
							}
						)
						let loop = this.flowData.config.works.length
						do {
						    console.log("do loop ...")
						    loop = loop - 1
                            for (var dagKey in this.flowData.config.dags) {
                                // nodeTo: dagKey
                                let max = 1
                                this.flowData.config.dags[dagKey].dependences.forEach(
                                    nodeFrom => {
                                        if (max < svgNodeYLevel[nodeFrom]) {
                                            max = svgNodeYLevel[nodeFrom]
                                        }
                                    }
                                );
                                if (max + 1 > svgNodeYLevel[dagKey]) {
                                    svgNodeYLevel[dagKey] = max + 1
                                }
                            };
						}while (loop > 0)
                        let yLevelCount = {}
                        console.log("this is svgNodeYLevel");
                        console.log(svgNodeYLevel)

                        for (var key in svgNodeYLevel) {
                            yLevelCount[svgNodeYLevel[key]] = 1
                        }
                        for (var key in svgNodeYLevel) {
                            svgNodeXLevel[key] = yLevelCount[svgNodeYLevel[key]]
                            yLevelCount[svgNodeYLevel[key]] += 1
                        }
                        console.log("this is svgNodeXLevel");
                        console.log(svgNodeXLevel)
                        let unitY = 100;
                        let unitX = 250;
                        let floatY = 20;
                        let floatX = 40;
                        let workDataTmp = this.workData
						console.log("workData", this.workData)
                        console.log("workDataTmp", workDataTmp)
						console.log("flowData.config", this.flowData.config.works)
                        this.flowData.config.works.forEach(
                            function(item, index) {
                                console.log("item:", item, "workDataTmp[item]", workDataTmp[item])
                                nodeIdMap[item] = index
                                svgDataNodes.push({
                                    "id": index,
                                    "name": workDataTmp[item].name,
                                    "pos_x": svgNodeXLevel[item] * unitX + Math.floor(Math.random() * floatX),
                                    "pos_y": svgNodeYLevel[item] * unitY + Math.floor(Math.random() * floatY),
                                    "type": "XGBoostOp",
                                    "in_ports": [
                                        0,
                                    ],
                                    "out_ports": [
                                        0,
                                    ]
                                })
                            }
                        );
                        for (var dagKey in this.flowData.config.dags) {
                            var item = this.flowData.config.dags[dagKey]
                            item.dependences.forEach(
                                jtem => {
                                    svgDataEdges.push({
                                        "dst_input_idx": 0,
                                        "dst_node_id": nodeIdMap[dagKey],
                                        "id": 0,
                                        "src_node_id": nodeIdMap[jtem],
                                        "src_output_idx": 0,
                                        "type": "active"
                                    })
                                }
                            );
                        };
                        console.log("svgDataNodes", svgDataNodes)
                        this.svgData["nodes"] = svgDataNodes
                        this.svgData["edges"] = svgDataEdges
                    }
                    console.log("this is svg data")
                    console.log(this.svgData)
                });
            },
        },
        created() {
            // this.getWorkData().then(
            //     this.getFlowDagData()
			// )
			this.getWorkData()
			setTimeout(this.getFlowDagData, 100)
        },
    };
</script>

<style>
</style>
