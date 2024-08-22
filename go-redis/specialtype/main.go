package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go-redis/redis/v8"
)

func main() {
	cli := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	ctx := context.Background()
	if err := cli.Ping(ctx).Err(); err != nil {
		fmt.Printf("Failed to connect redis: %v \n", err)
		return
	}
	// OpBitMap(cli)
	OpHyperLogLog(cli)
}

func OpBitMap(cli *redis.Client) {
	ctx := context.Background()
	key := "id:1:2024:06"
	//六月份签到
	for i := 0; i < 30; i += 7 {
		if ret := cli.SetBit(ctx, key, int64(i), 1); ret.Err() != nil {
			fmt.Println(ret.Err())
		} else {
			fmt.Println("====================")
			fmt.Println(ret.Args()...)
			fmt.Println(ret.Result())
		}
	}

	//签到了多少天
	if ret := cli.BitCount(ctx, key, &redis.BitCount{Start: 0, End: 30}); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("====================")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

	//第一天签到是什么时候
	if ret := cli.BitPos(ctx, key, 1); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("====================")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

	//第一次没签到是什么时候
	if ret := cli.BitPos(ctx, key, 0); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("====================")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

	for i := 1; i <= 4; i++ {
		for j := 0; j < i; j++ {
			cli.SetBit(ctx, fmt.Sprintf("sign:2024:06:0%d", i), int64(j), 1)
		}
	}

	//有多少人谁连续四天签到
	cli.BitOpAnd(ctx, "sign_temp", "sign:2024:06:01", "sign:2024:06:02", "sign:2024:06:03", "sign:2024:06:04")
	if ret := cli.BitCount(ctx, "sign_temp", &redis.BitCount{Start: 0, End: 4}); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("====================")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}
}

func OpHyperLogLog(cli *redis.Client) {
	ctx := context.Background()
	key := "pf:online"
	arr := make([]any, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(20)
	}
	//加入在线人数
	if ret := cli.PFAdd(ctx, key, arr...); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("===========")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

	//查看当前有多少在线人数
	if ret := cli.PFCount(ctx, key); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("===========")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

	key_1 := "pf:online:1"
	arr_1 := make([]any, 10)
	for i := 0; i < 10; i++ {
		arr_1[i] = rand.Intn(20)
	}
	//加入在线人数
	if ret := cli.PFAdd(ctx, key_1, arr_1...); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("===========")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

	key_2 := "pf:online:merge"
	if ret := cli.PFMerge(ctx, key_2, key, key_1); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("===========")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

	//查看当前有多少在线人数
	if ret := cli.PFCount(ctx, key_2); ret.Err() != nil {
		fmt.Println(ret.Err())
	} else {
		fmt.Println("===========")
		fmt.Println(ret.Args()...)
		fmt.Println(ret.Result())
	}

}
