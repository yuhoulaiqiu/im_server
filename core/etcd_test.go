package core

import (
	"context"
	"fmt"
	"testing"
)

func TestInitEtcd(t *testing.T) {
	client := InitEtcd("localhost:2379")
	res, err := client.Put(context.Background(), "auth_api", "127.0.0.1:20021")
	fmt.Println(res, err)
	getResponse, err := client.Get(context.Background(), "auth_api")
	fmt.Println(getResponse, err)
	if err == nil && len(getResponse.Kvs) > 0 {
		fmt.Println(string(getResponse.Kvs[0].Value))
	}
}
