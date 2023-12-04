package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// Common Declarations

type PutArgs struct {
	Key   string
	Value string
}

type PutReply struct{}

type GetArgs struct {
	Key string
}

type GetReply struct {
	Value string
}

// Client

func connect() *rpc.Client {
	// 创建到 RPC 服务器的 TCP 连接
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

// 函数 put 和 get 都是 client stub (客户端存根)
// 客户端存根负责将请求参数打包成网络消息发送给服务方
// 以及接收表示结果的消息并进行解码

func get(key string) string {
	client := connect()
	args := GetArgs{
		Key: key,
	}
	reply := GetReply{}
	err := client.Call("KV.Get", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
	return reply.Value
}

func put(key string, val string) {
	client := connect()
	args := PutArgs{
		Key:   key,
		Value: val,
	}
	reply := PutReply{}
	err := client.Call("KV.Put", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
}

// Server

type KV struct {
	mu   sync.Mutex
	data map[string]string
}

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	reply.Value = kv.data[args.Key]
	return nil
}

func (kv *KV) Put(args *PutArgs, reply *PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[args.Key] = args.Value
	return nil
}

func server() {
	// 需要声明一个某种类型的对象用于注册 RPC 服务
	kv := &KV{
		data: make(map[string]string),
	}
	rpcs := rpc.NewServer()
	rpcs.Register(kv)

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		for {
			// 等待连接请求
			// 并为每个连接请求创建一个 goroutine 来处理
			conn, err := l.Accept()
			if err == nil {
				go rpcs.ServeConn(conn)
			} else {
				break
			}
		}
		l.Close()
	}()
}

func main() {
	server()

	put("subject", "6.5840")
	fmt.Printf("Put(subject, 6.5840) done\n")
	fmt.Printf("get(subject) -> %s\n", get("subject"))
}
