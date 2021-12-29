package sdk

/**
  @Description: 封装了http请求中的过滤器
  @author jun.hai
  @date 2021年12月28日 下午21:13:08
*/

import (
	"fmt"
	"time"
)

type Filter func(ctx *Context)

type FilterBuilder func(next Filter) Filter

func MetricFilterBuilder(next Filter) Filter {
	return func(ctx *Context) {
		startTime := time.Now().UnixNano()
		next(ctx)

		endTime := time.Now().UnixNano()
		fmt.Println("run time:", endTime-startTime)
	}
}
