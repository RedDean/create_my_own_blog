/*
 *  @author: reddean
 *  @created-time: 2018/06/10
 *  @description: 路由信息定义
 */
package controller

import "github.com/julienschmidt/httprouter"

type Route struct {
	RouteName    string
	Method       string
	Path         string
	HandleFunc   httprouter.Handle
}

type Routes []Route

// 路由注册
// register route
var routes = Routes{
	{"Index",
	 "GET",
	 "/",
	 Index,
	},
	{"Register",
	 "POST",
	 "/user/register",
	 Register,
	},
	{"signIn",
	 "POST",
	 "/user/sigin",
	 SignIn,
	},
	{"signOut",
	 "GET",
	 "/signout",
	 signOut,
	},
	{"createComment",
	 "POST",
	 "/comments/create",
	 createComment,
	},
	{"removeComment",
	 "GET",
	 "/comments/:ID/remove",
	 removeComment,
	},
	{"viewArticle",
	 "GET",
	 "/article/:ID",
	 viewArticle,
	},
	{"createArticle",
	 "POST",
	 "/article/:ID/remove",
	 createArticle,
	},
	{"editArticle",
	 "POST",
	 "/article/:ID/edit",
	 editArticle,
	},
	{"removeArticle",
	 "GET",
	 "/article/:ID/remove",
	 removeArticle,
	},
}











