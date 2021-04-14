package handler

import (
	"greet/libs"
	"net/http"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

type ParamsReq struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
func GetUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

	    client := libs.HttpClient{}
		client.SetConfig(ctx.Config.Sms)
		var uri string = "admin-api/login"
		paramsReq :=  ParamsReq{}
		paramsReq.Password="password"
		paramsReq.Username="admin"
		reqs , err := client.HttpPost(uri,paramsReq)
		//l := logic.NewGetUserLogic(r.Context(), ctx)
		//resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, reqs)
		}
	}
}
