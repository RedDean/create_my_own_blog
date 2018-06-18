/*
 *  @author: reddean
 *  @created-time: 2018/06/10
 *  @description: 路由处理方法实现
 */
package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	//"io/ioutil"
)

// 返回首页
func Index (w http.ResponseWriter,
	        r *http.Request, p httprouter.Params){

}

func Register (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){
	// 方法一:
    // r.ParseForm()
	// str := r.PostForm.Get("name")
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)

	// 方法二: str := r.PostFormValue("name")

	/* 方法三:
	ctx,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("post error !")
	} */

	//fmt.Fprintf(w, "request body is %s", str)
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "article ID is %s", p.ByName("ID"))
}

func createArticle (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "article ID is %s", p.ByName("ID"))
}

func editArticle (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}

func removeArticle (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){

}