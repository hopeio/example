package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hopeio/cherry/context/ginctx"
	errorsi "github.com/hopeio/cherry/utils/errors"
	"github.com/hopeio/pick"
)

type UserService struct{}

func (*UserService) Service() (string, string, []gin.HandlerFunc) {
	return "用户相关", "/api/v1/user", []gin.HandlerFunc{Log}
}

type Object struct {
	Id int `json:"id,omitempty"`
}

type User struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Gender int    `json:"gender,omitempty"`
}

func (*UserService) Get(ctx *ginctx.Context, req *Object) (*User, error) {
	pick.Api(func() {
		pick.Get("/:id").
			Title("用户详情").
			CreateLog("1.0.0", "jyb", "2024/04/16", "创建").End()
	})
	// dao
	return &User{
		Id:     req.Id,
		Name:   "test",
		Gender: 1,
	}, nil
}

type Req struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (*UserService) GetErr(ctx *ginctx.Context, req *Req) (*User, error) {
	pick.Api(func() {
		pick.Post("/err/:id").
			Title("用户详情返回错误").
			CreateLog("1.0.0", "jyb", "2024/04/16", "创建").End()
	})
	fmt.Println(req.Name)
	// dao
	return nil, &errorsi.ErrRep{
		Code:    1,
		Message: "error",
	}
}

type Signup struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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
