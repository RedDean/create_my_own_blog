/*
 *  @author: reddean
 *  @created-time: 2018/06/10
 *  @description: 路由注册器  // router register
 */

package controller

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(
			route.Method,
			route.Path,
			route.HandleFunc,
		)
	}

	return router
}