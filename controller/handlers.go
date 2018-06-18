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
	"encoding/json"
	"log"
)



// 返回首页
func Index (w http.ResponseWriter,
	        r *http.Request, p httprouter.Params){

}

func Register (w http.ResponseWriter,
	r *http.Request, p httprouter.Params){
	var t interface{}	
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		log.Println("decoded failure!")
	}
	log.Println(t)
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

	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"result":"yes"}`))
	//fmt.Fprintf(w, "article ID is %s", p.ByName("ID"))
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