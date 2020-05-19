package RecommendHandle

import (
	"PoTest/models/recommend"
	"PoTest/types"
	"PoTest/types/recommend"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"io/ioutil"
	"strconv"
)

type Recommend struct {
}


func (e *Recommend) PostWeibo(req *restful.Request, rsp *restful.Response) {

	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	reqData, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		panic("read")
	}
	postWeiboreq := new(recommendType.PostWeiboReq)
	if err := json.Unmarshal(reqData, postWeiboreq); err != nil {
		panic("unmarshal")
	}
	err = recommendModel.PostWeibo(postWeiboreq)
	if err != nil {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseFailHttpData(rsp, err.Error())
	} else {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseSuccessHttpData(rsp, "ok")
	}
}

func (e *Recommend) Suggest(req *restful.Request, rsp *restful.Response) {

	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	userId := req.Request.URL.Query().Get("userId")
	id,err:=strconv.Atoi(userId)
	if err!=nil{
		panic("string to int err")
	}
	data, err := recommendModel.Suggest(id)
	if err != nil {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseFailHttpData(rsp, err.Error())
	} else {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseSuccessHttpData(rsp, data)
	}

}
