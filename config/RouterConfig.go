package config

import (
	"PoTest/handle/RecommendHandle"
	"github.com/emicklei/go-restful"
)

var recommend = new(RecommendHandle.Recommend)

const (
	PATH = "/api"
)

var wc = restful.NewContainer()


var ws = new(restful.WebService)

func routerConf() {
	recommendRouterConf()

}

func GetRouterContainer() *restful.Container {

	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path(PATH)
	routerConf()
	wc.Add(ws)
	return wc
}

func recommendRouterConf() {
	rootPath := "recommend/"
	methodPath := "postWeibo"
	ws.Route(ws.POST(rootPath + methodPath).To(recommend.PostWeibo))

	methodPath = "suggest"
	ws.Route(ws.GET(rootPath + methodPath).To(recommend.Suggest))


}


