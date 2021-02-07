package main

import (
	"time"
	"context"
	"fmt"
)
const shortDuration = 1 * time.Second
func main(){
	d := time.Now().Add(shortDuration)
	ctx, cancelfun := context.WithDeadline(context.Background(),d)
	defer cancelfun()  //不是特别明白 不执行会泄露，不清楚为啥。要注意就行

	select {
		case <-time.After(2 * time.Second):
			fmt.Println("over sept:")
		case <-ctx.Done():
			fmt.Println("ctx done:",ctx.Err())
	}
}
