/*
 *  @author: reddean
 *  @created-time: 2018/06/10
 *  @description: 路由处理方法实现
 */
package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// 返回首页
func Index (w http.ResponseWriter,
	        r *http.Request, p httprouter.Params){

}

func Register (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func SignIn (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func signOut (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func createComment (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func removeComment (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func viewArticle (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func createArticle (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func editArticle (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func removeArticle (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}