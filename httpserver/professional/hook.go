package professional

/**
  @Description: 钩子函数，通过不同场景建立不同的钩子函数
  @author jun.hai
  @date 2022年1月7日 下午14:06:21
*/

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Hook func(ctx context.Context) error

func CloseServerHook(servers ...IServer) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		doneChan := make(chan struct{})
		wg.Add(len(servers))

		for _, s := range servers {
			go func(svc IServer) {
				err := svc.Shutdown(ctx)
				if err != nil {
					fmt.Printf("server shutdown error: %v \n", err)
				}
				time.Sleep(time.Second)
				wg.Done()
			}(s)
		}

		go func() {
			wg.Wait()
			doneChan <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			fmt.Printf("closing servers timeout \n")
			return errors.New("the hook timeout")
		case <-doneChan:
			fmt.Printf("close all servers \n ")
			return nil
		}
	}
}
