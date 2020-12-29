package dag

import (
	Dag "github.com/diaohaha/termite/dag/model"
	"fmt"
	"log"
	"reflect"
)

func DeepCopyDagNode(dagNodeSrc Dag.DagNode, dagNodeMap map[string]*Dag.DagNode) (dagNodeDst Dag.DagNode, retDagNodeMap map[string]*Dag.DagNode) {
	retDagNodeMap = dagNodeMap
	log.Println("dag node:", dagNodeSrc.Config.Name, len(dagNodeSrc.Children), len(dagNodeSrc.Parents))
	//time.Sleep(time.Second * 1)
	dagNodeDst.Status = dagNodeSrc.Status
	dagNodeDst.Config = dagNodeSrc.Config
	retDagNodeMap[dagNodeDst.Config.Name] = &dagNodeDst
	if len(dagNodeSrc.Children) > 0 {
		for _, child := range dagNodeSrc.Children {
			if _, ok := retDagNodeMap[child.Config.Name]; ok {
				deepCopyChild := retDagNodeMap[child.Config.Name]
				dagNodeDst.Children = append(dagNodeDst.Children, deepCopyChild)
			} else {
				deepCopyChild, tmpNodeMap := DeepCopyDagNode(*child, retDagNodeMap)
				retDagNodeMap = tmpNodeMap
				dagNodeDst.Children = append(dagNodeDst.Children, &deepCopyChild)
			}
		}
	}
	if len(dagNodeSrc.Parents) > 0 {
		for _, parent := range dagNodeSrc.Parents {
			if _, ok := retDagNodeMap[parent.Config.Name]; ok {
				deepCopyParent := retDagNodeMap[parent.Config.Name]
				dagNodeDst.Children = append(dagNodeDst.Children, deepCopyParent)
			} else {
				deepCopyParent, tmpNodeMap := DeepCopyDagNode(*parent, retDagNodeMap)
				retDagNodeMap = tmpNodeMap
				dagNodeDst.Parents = append(dagNodeDst.Parents, &deepCopyParent)
			}
		}
	}
	return dagNodeDst, retDagNodeMap
}

func CompareNode(dagNodeA Dag.DagNode, dagNodeB Dag.DagNode) (equal bool, err error) {
	if &dagNodeA == &dagNodeB {
		equal = false
		return
	}

	return
}

func PrintNode(dagNode *Dag.DagNode) {
	fmt.Println("------------------start------------------")
	dagMap := map[string]*Dag.DagNode{}
	dagMap = getDagMap(dagNode, dagMap)
	for name, dagAddr := range dagMap {
		dag := reflect.ValueOf(dagAddr)
		fmt.Println("name: ", name, "addr:", dag.Pointer())
	}
	fmt.Println("-------------------end-------------------")
}

func getDagMap(node *Dag.DagNode, dagMap map[string]*Dag.DagNode) (retDagMap map[string]*Dag.DagNode) {
	retDagMap = dagMap
	if _, ok := dagMap[node.Config.Name]; !ok {
		retDagMap[node.Config.Name] = node
	} else {
		return retDagMap
	}

	for _, child := range node.Children {
		if _, ok := dagMap[child.Config.Name]; !ok {
			retDagMap = getDagMap(child, retDagMap)
		}
	}
	for _, parent := range node.Parents {
		if _, ok := dagMap[parent.Config.Name]; !ok {
			retDagMap = getDagMap(parent, retDagMap)
		}
	}

	return retDagMap
}
