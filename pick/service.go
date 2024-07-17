package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hopeio/context/ginctx"
	"github.com/hopeio/pick"
	errorsi "github.com/hopeio/utils/errors"
	"strconv"
)

type UserService struct{}

func (*UserService) Service() (string, string, []gin.HandlerFunc) {
	return "用户相关", "/api/v1/user", []gin.HandlerFunc{Log}
}

func (*UserService) Get(ctx *ginctx.Context, req *Object) (*User, error) {
	pick.Api(func() {
		pick.Get("/:id").
			Title("用户详情").
			CreateLog("1.0.0", "jyb", "2024/04/16", "创建").End()
	})
	// dao
	return &User{Id: req.Id, Name: "test", Gender: 1}, nil
}

func (*UserService) GetErr(ctx *ginctx.Context, req *Req) (*User, error) {
	pick.Api(func() {
		pick.Post("/err/:id").
			Title("用户详情返回错误").
			CreateLog("1.0.0", "jyb", "2024/04/16", "创建").End()
	})
	fmt.Println(req.Name)
	// dao
	return nil, &errorsi.ErrRep{Code: 1, Message: "error"}
}

func (*UserService) Signup(ctx *ginctx.Context, req *Signup) (*User, error) {
	pick.Api(func() {
		pick.Post("").
			Title("用户注册").
			CreateLog("1.0.0", "jyb", "2024/06/17", "创建").End()
	})
	fmt.Println(req.Name)
	// dao
	return &User{Id: 1, Name: req.Name}, nil
}

func (*UserService) List(ctx *ginctx.Context, req *Page) (*UserListRes, error) {
	pick.Api(func() {
		pick.Get("").
			Title("用户列表").
			CreateLog("1.0.0", "jyb", "2024/06/11", "创建").End()
	})
	var users []*User
	for i := range 10 {
		users = append(users, &User{Id: i + 1, Name: strconv.Itoa(i + 1)})
	}
	// dao
	return &UserListRes{Total: 10, Users: users}, nil
}

func (*UserService) BachUpdate(ctx *ginctx.Context, req *BachUpdate) (struct{}, error) {
	pick.Api(func() {
		pick.Put("").
			Title("用户批量修改").
			CreateLog("1.0.0", "jyb", "2024/06/11", "创建").End()
	})
	var users []*User
	for i := range 10 {
		users = append(users, &User{Id: i + 1, Name: strconv.Itoa(i + 1)})
	}
	// dao
	return struct{}{}, nil
}
