package model

import (
	"github.com/diaohaha/termite/utils"
	"fmt"
	"log"
)

const (
	STATE_WAITING = 0
	STATE_RUNING  = 1
	STATE_FINISH  = 2
)

type DagConfig struct {
	Name        string
	Dependences []string
}
type DagNode struct {
	Children []*DagNode
	Parents  []*DagNode
	Config   DagConfig
	Status   int
}

func (node *DagNode) PrintDag() {
	node.Each(func(node *DagNode) bool {
		log.Println(fmt.Sprintf("%s", node.Config.Name))
		return true
	})
	return
}

func (node *DagNode) Each(callback func(node *DagNode) bool) {
	todo := utils.SetFactory()
	done := utils.SetFactory()
	todo.Push(node.Config.Name, node)
	for todo.Len() > 0 {
		tmp := todo.RandPop().(*DagNode)
		if tmp != nil {
			done.Push(tmp.Config.Name, tmp)
			if !callback(tmp) {
				continue
			}
			for _, v := range tmp.Children {
				if !done.Check(v.Config.Name) {
					todo.Push(v.Config.Name, v)
				}
			}
		}

	}
	return
}
func (node *DagNode) UpdateStatus(nodeStatus map[string]int) {
	node.Each(func(node *DagNode) bool {
		if _, isOK := nodeStatus[node.Config.Name]; isOK {
			node.Status = nodeStatus[node.Config.Name]
			//log.Println(fmt.Sprintf("%s:%d",node.Config.Name,node.Status))
		} else {
			// 如果没有传入Status，默认为WAITING
			if node.Config.Name == "__start" {
				node.Status = STATE_FINISH
			} else {
				// 支持工作流平滑升级
				node.Status = STATE_FINISH
				//node.Status = STATE_WAITING
			}

		}
		return true
	})
	return
}
func (node *DagNode) GetReadyNodes() []string {
	tmp := []string{}
	node.Each(func(node *DagNode) bool {
		//log.Println(fmt.Sprintf("%s",node.Config.Name))
		if node.Status == STATE_FINISH {
			//只有父节点存在完成的任务，才继续向下查找
			return true
		}
		if node.Status == STATE_WAITING {
			//只有当前节点是waiting状态，才会去检查是否是ready任务
			for _, p := range node.Parents {
				if p.Status == STATE_FINISH {
					continue
				}
				//父节点有非完成状态，终止查找
				return false
			}
			//所有父节点为完成状态，加入当前节点为ready
			tmp = append(tmp, node.Config.Name)
			return false
		}
		return false
	})
	return tmp
}
func InitDag(configList *[]DagConfig) *DagNode {
	tmpMap := make(map[string]*DagNode)
	roots := make([]*DagNode, 0)
	for _, v := range *configList {
		var currentNode *DagNode = nil
		var isOK bool = false
		if currentNode, isOK = tmpMap[v.Name]; !isOK {
			currentNode = new(DagNode)
			tmpMap[v.Name] = currentNode
		}
		currentNode.Config = v

		for _, dependence := range v.Dependences {
			var node *DagNode = nil
			if node, isOK = tmpMap[dependence]; !isOK {
				node = new(DagNode)
				tmpMap[dependence] = node
			}
			//log.Println(fmt.Sprintf("%s->%s",node.Config.Name,currentNode.Config.Name))
			node.Children = append(node.Children, currentNode)
			currentNode.Parents = append(currentNode.Parents, node)
		}
	}

	for _, v := range *configList {
		if len(tmpMap[v.Name].Parents) == 0 {
			//找到所有入度为0的根节点
			roots = append(roots, tmpMap[v.Name])
		}
	}
	//判环
	//log.Println(*configList)
	//log.Println("roots is :")
	//log.Println(roots)
	tmpRoots := roots
	deleteNodes := map[string]bool{}
	var isRoot bool = false
	dagNodeCount := 0
	for len(tmpRoots) > 0 {
		dagNodeCount++
		root := tmpRoots[0]
		log.Println(root.Config.Name)
		tmpRoots = tmpRoots[1:]
		deleteNodes[root.Config.Name] = true
		//log.Println(root.Config.Name+" children count:"+fmt.Sprintf("%d",len(root.Children)))
		for _, child := range root.Children {
			//log.Println("child:"+child.Config.Name)
			if _, isOK := deleteNodes[child.Config.Name]; isOK {
				//如果该节点已经被删除，跳过不检查
				continue
			}
			isRoot = true
			for _, p := range child.Parents {
				//log.Println("child's parent:"+p.Config.Name)
				//遍历父节点，如果所有父节点均已被删除，加入tmpRoots
				if _, isOK := deleteNodes[p.Config.Name]; !isOK {
					isRoot = false
				}
			}
			if isRoot {
				tmpRoots = append(tmpRoots, child)
			}
		}
	}
	//log.Println(fmt.Sprintf("%d-%d",dagNodeCount,len(*configList)))
	if dagNodeCount < len(*configList) {
		return nil
	}
	if len(roots) > 1 {
		start := new(DagNode)
		start.Config.Name = "__start"
		start.Status = STATE_FINISH
		start.Children = roots

		for _, v := range roots {
			v.Parents = append(v.Parents, start)
		}
		return start
	}
	return roots[0]
}
