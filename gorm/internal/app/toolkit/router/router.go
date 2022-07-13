package router

import (
	"net/http"
)

// use example:
// addr("/user",nihao).use(login)
// addrGroup("/user",
//		addr("/login",nihao)
//)

// 支持分组路由以及中间间，分组路由可直接全局添加中间价，提供默认中间件
// 注册只梳理关系，最后init挂数

var routerMap []*Router

type Router struct {
	path       string
	middleware []AddrRouter
}

type AddrRouter func(http.ResponseWriter, *http.Request)

type RouterGroup []*Router

func (rt *Router) Use(mw AddrRouter) *Router {
	rt.middleware = append(rt.middleware, mw)
	return rt

}

func (rt *Router) next(w http.ResponseWriter, r *http.Request) {
	for index := len(rt.middleware) - 1; index > -1; index-- {
		mwBreak := r.Header.Get("middleware-break")
		if mwBreak != "" {
			return
		}
		rt.middleware[index](w, r)
	}
}

func (routerGroup RouterGroup) Use(mw AddrRouter) RouterGroup {
	for i := range routerGroup {
		j := routerGroup[i]
		(*j).middleware = append((*j).middleware, mw)
	}
	return routerGroup
}

// 单路由
func Url(path string, handlerFunc AddrRouter) *Router {
	return __url(path, handlerFunc)
}

func __url(path string, handlerFunc AddrRouter) *Router {
	router := &Router{
		path:       path,
		middleware: []AddrRouter{handlerFunc},
	}

	routerMap = append(routerMap, router)
	return router
}

// 路由组
func UrlGroup(pre string, AddrList ...*Router) RouterGroup {
	for i := range AddrList {
		r := AddrList[i]
		r.path = pre + r.path
	}
	return AddrList
}

// 设置中断
func SetBreak(r *http.Request) {
	r.Header.Set("middleware-break", "1")
}

func Init() {
	for i := range routerMap {
		r := routerMap[i]
		http.HandleFunc(r.path, r.next)
	}
}
