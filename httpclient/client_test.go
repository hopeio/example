package client

import (
	"github.com/davecgh/go-spew/spew"
	httpi "github.com/hopeio/cherry/utils/net/http"
	"github.com/hopeio/cherry/utils/net/http/client"
	clientv2 "github.com/hopeio/cherry/utils/net/http/client/v2"
	"testing"
)

func TestUserList(t *testing.T) {
	req := client.New().AddHeader("Content-Type", "application/json").LogLevel(client.LogLevelInfo).Request("GET", "http://localhost:8080/api/v1/user")
	var res httpi.ResData[UserListRes]
	err := req.Do(&Page{1, 2}, &res)
	if err != nil {
		t.Log(err)
	}
	spew.Dump(res)
}

func TestUserListV2(t *testing.T) {
	req := clientv2.NewRequest[httpi.ResData[UserListRes]]("GET", "http://localhost:8080/api/v1/user")
	res, err := req.Do(&Page{1, 2})
	if err != nil {
		t.Log(err)
	}
	spew.Dump(res)
}
