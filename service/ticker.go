package service

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Ticker struct {
	
}

var (
	tickerDefaultDuration = 1 * time.Second
)

// InitTicker 根据bucket数量启动相应ticker进行扫描
func InitTicker(ctx context.Context, bucketName string, s *Service)  {
	ticker := time.NewTicker(tickerDefaultDuration)
	defer func() {
		s.wg.Done()
		ticker.Stop()
	}()
	s.wg.Add(1)

	for {
		select {
		case <-ticker.C:
			tickHandler(bucketName)
		case <-ctx.Done(): // 等待上级通知
			log.Printf("context Done msg: %#v\n", ctx.Err())
			return
		default:
			//fmt.Printf("default#%s\n", bucketName)
		}
	}
}

func tickHandler(bucketName string)  {
	// 扫描bucket
	fmt.Printf("ticker handler bucket name#%s\n", bucketName)
}