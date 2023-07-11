package router

import (
	"fmt"
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

type RouterNP interface {
	parseRouter() []*Router
}

func (rt *Router) parseRouter() []*Router {
	return []*Router{rt}
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

func (rt *RouterGroup) parseRouter() []*Router {
	return *rt
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

// 路由组,支持重复路由搭建
func UrlGroup(pre string, routerNPs ...RouterNP) *RouterGroup {
	var AddrList RouterGroup
	for np := range routerNPs {
		npAddrList := routerNPs[np].parseRouter()
		for i := range npAddrList {
			r := npAddrList[i]
			r.path = pre + r.path
		}
		AddrList = append(AddrList, npAddrList...)
	}
	return &AddrList
}

// 设置中断
func SetBreak(r *http.Request) {
	r.Header.Set("middleware-break", "1")
}

// 输出生成的路由信息
func Debug() {
	fmt.Println("----路由信息-----")
	routers := FilterMultiple()
	for i := range routers {
		r := routers[i]
		fmt.Println("注册路径:", r.path)
	}
	fmt.Println("----路由调试end----")
}

// 过滤重复路由以实现函数覆盖
func FilterMultiple() map[string]*Router {
	routers := make(map[string]*Router)
	for i := range routerMap {
		r := routerMap[i]
		routers[r.path] = r
	}
	return routers
}

func Init() {
	routers := FilterMultiple()
	for i := range routers {
		r := routers[i]
		http.HandleFunc(r.path, r.next)
	}
}
