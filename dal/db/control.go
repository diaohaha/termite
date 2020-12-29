package db

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/stable"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

func NodeHeartBeat(nodeId string, nodeType string) (err error) {
	// 更新节点信息
	var tnode dal.TermiteNode
	err = WTDB.Model(dal.TermiteNode{}).Where("node_id = ? and node_type = ?", nodeId, nodeType).Find(&tnode).Error
	if err == gorm.ErrRecordNotFound {
		tnode = dal.TermiteNode{
			NodeId:     nodeId,
			NodeType:   nodeType,
			ExpireTime: time.Now().Add(30 * time.Second),
		}
		err = WTDB.Save(&tnode).Error
	} else {
		if err != nil {
			stable.SchedulerLogger.Error("db error.", zap.Error(err))
			return err
		} else {
			//tnode.ExpireTime = time.Now().Add(30 * time.Second)
			//err = WTDB.Save(&tnode).Error
			err = WTDB.Model(dal.TermiteNode{}).Where("node_id = ?", nodeId).Update("expire_time", time.Now().Add(30*time.Second)).Error
			if err != nil {
				stable.SchedulerLogger.Error("db error.", zap.Error(err))
				return err
			}
		}
	}
	return
}

func GetLiveNodes(nodeType string) (tnodes []dal.TermiteNode, err error) {
	err = WTDB.Model(dal.TermiteNode{}).
		Where("node_type = ?", nodeType).
		Where("expire_time > ?", time.Now()).
		Find(&tnodes).Error
	if err != nil {
		stable.SchedulerLogger.Error("db error.", zap.Error(err))
		return
	}
	return
}

func DeleteExpireNodes() (err error) {
	err = WTDB.Where("expire_time < ?", time.Now()).Delete(dal.TermiteNode{}).Error
	return
}

func UpdateNodePartitions(nodeId string, partitions string) (err error) {
	var tnode dal.TermiteNode
	err = WTDB.Model(dal.TermiteNode{}).Where("node_id = ?", nodeId).Find(&tnode).Error
	if err != nil {
		return
	}
	//tnode.Partitions = partitions
	err = WTDB.Model(dal.TermiteNode{}).Where("node_id = ?", nodeId).Update("partitions", partitions).Error
	if err != nil {
		return
	}
	return
}

func QueryTermiteNodes(nodeId string) (tnode dal.TermiteNode, err error) {
	err = WTDB.Model(dal.TermiteNode{}).Where("node_id = ?", nodeId).Find(&tnode).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}
	if err != nil {
		return
	}
	return
}
