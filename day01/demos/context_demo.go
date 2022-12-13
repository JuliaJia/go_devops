package demos

import (
	"context"
	"fmt"
	"time"
)

func Demo1() {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	//defer cancel()
	//time.Sleep(time.Second * 2)
	time.Sleep(500 * time.Millisecond)
	defer cancel()
	err := timeoutCtx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("It is Canceled!")
	case context.DeadlineExceeded:
		fmt.Println("It is Deadline!")
	default:
		fmt.Println(err)
	}
}

func Demo2() {
	ctx := context.Background()
	dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	childCtx := context.WithValue(dlCtx, "key", 123)
	cancel()
	err := childCtx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("It is Canceled!")
	case context.DeadlineExceeded:
		fmt.Println("It is Deadline!")
	default:
		fmt.Println(err)
	}
}

func Demo3() {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "key1", 123)
	ccCtx := context.WithValue(childCtx, "key2", 124)

	err := ccCtx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("It is Canceled!")
	case context.DeadlineExceeded:
		fmt.Println("It is Deadline!")
	default:
		fmt.Println(err)
	}
	fmt.Println(ccCtx.Value("key1"))
	fmt.Println(childCtx.Value("key2"))
}
func Demo4() {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "map", map[string]string{})
	ccCtx := context.WithValue(childCtx, "key1", 124)
	m := ccCtx.Value("map").(map[string]string)
	m["key1"] = "new124"
	val := childCtx.Value("key1")
	fmt.Println(val)
	val = childCtx.Value("map")
	fmt.Println(val)
}

func Demo5() {
	bg := context.Background()
	timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	subCtx, cancel2 := context.WithTimeout(timeoutCtx, 3*time.Second)
	go func() {
		<-subCtx.Done()
		fmt.Println("timeout")
	}()
	time.Sleep(2 * time.Second)
	cancel1()
	cancel2()
}

func Demo6() {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	end := make(chan struct{}, 1)
	go func() {
		MyBusiness()
		end <- struct{}{}
	}()
	ch := timeoutCtx.Done()
	select {
	case <-ch:
		fmt.Println("Timeout!")
	case <-end:
		fmt.Println("Done!")
	}
}

func MyBusiness() {
	time.Sleep(2 * time.Second)
	fmt.Print("MyBusiness!")
}

func Demo7() {

}
