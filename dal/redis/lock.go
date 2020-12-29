package redis

import (
	"fmt"
	"github.com/go-errors/errors"
	"time"
)

func Lock(kid int64) (err error) {
	key := fmt.Sprintf("lock_%d", kid)
	//exist, err := Conn.Exists("key").Result()
	//if err != nil {
	//	return err
	//}
	//if exist == 1 {
	//	err = errors.New("exist")
	//	return
	//}
	//tmpKey := fmt.Sprintf("lock10_%d", kid)
	//_ = Conn.Set(tmpKey, 1, 100*time.Second).Err()
	res, err := Conn.SetNX(key, 1, 1*time.Second).Result()
	if err != nil {
		return
	} else {
		if res == false {
			// 设置失败
			return errors.New("already exist")
		} else {
			return nil
		}
	}
}

func UnLock(kid int64) (err error) {
	key := fmt.Sprintf("lock_%d", kid)
	Conn.Del(key)
	return
}
