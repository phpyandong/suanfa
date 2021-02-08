package errgrp

import (
	"fmt"
	"context"
	"u2pppw/basic/sync/errgroup"
	"errors"
)

func main()  {
	g := new(errgroup.Group)
	var a,b,c []int
	//调用广告
	g.Go(func(ctx context.Context, cancel context.CancelFunc) error {
		a = []int{0}
		return nil
	})
	//调用AI
	g.Go(func(ctx context.Context, cancel context.CancelFunc) error {
		b = []int{1}
		return errors.New("ai" )
	})
	//调用运营平台
	g.Go(func(ctx context.Context, cancel context.CancelFunc) error {
		c = []int{1}
		return errors.New("ai")
	})

	if err := g.Wait();err == nil{
		fmt.Println("get all success")
	}
	fmt.Println()
}