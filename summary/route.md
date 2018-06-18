## 在go中如何组织路由

&emsp;&emsp;在设计博客项目的时候，我打算将其做成前后端分离。前后端使用RESTful api进行交互。
因此，接下来一个问题是如何设计接口与在go中组织路由。

### 1. RESTful api的设计

&emsp;&emsp;关于RESTful api的设计，可参考如下资料：

+ [阮一峰 RESTful API设计指南](http://www.ruanyifeng.com/blog/2014/05/restful_api.html)
+ [RESTful API 设计参考文献列表](https://github.com/aisuhua/restful-api-design-references)
+ [RESTful架构详解](https://kb.cnblogs.com/page/512047/)

&emsp;&emsp;在设计api的时候，涉及到几个问题:

- GET和POST请求有什么区别？
- 同样是更新资源，POST和PUT请求有什么区别？
- PUT和DELETE请求到底安不安全
- go中如何处理post请求

&emsp;&emsp;提到请求区别之前，有两个概念[^参考资料1]。
>
>- 安全：所谓安全的意味着该操作用于获取信息而非修改信息。换句话说，GET 请求一般不应产生副作用。
>就是说，它仅仅是获取资源信息，就像数据库查询一样，不会修改，增加数据，不会影响资源的状态。
>- 幂等：幂等的意味着对同一URL的多个请求应该返回同样的结果。

四种请求在幂等和安全方面的详解请参考这里 <https://kb.cnblogs.com/page/512047/>

#### 1.1 GET和POST请求有什么区别？

&emsp;&emsp;GET与POST的区别主要在以下三个方面：

>- GET请求的数据会附在URL之后（就是把数据放置在HTTP协议头中），以?分割URL和传输数据，参数之间以&相连。
POST把提交的数据则放置在是HTTP包的请求体中。
>- 因为GET是通过URL提交数据，那么GET可提交的数据量就跟URL的长度有直接关系了。
  而实际上，URL不存在参数上限的问题，HTTP协议规范没有对URL长度进行限制。这个限制是特定的浏览器及服务器对它的限制。IE对URL长度的限制是2083字节(2K+35)。
  对于其他浏览器，如Netscape、FireFox等，理论上没有长度限制，其限制取决于操作系统的支持。
  POST是没有大小限制的，HTTP协议规范也没有进行大小限制，起限制作用的是服务器的处理程序的处理能力。
>- POST的安全性要比GET的安全性高。注意：这里所说的安全性和上面GET提到的“安全”不是同个概念。上面“安全”的含义仅仅是不作数据修改，而这里安全的含义是真正的Security的含义。[^参考资料1]

#### 1.2 同样是更新资源，POST和PUT请求有什么区别？

用POST和PUT创建资源或更新资源的**区别**更多是**语义风格**上的。

二者在创建资源时的区别如下：

> POST和PUT在创建资源的区别在于，**所创建的资源的名称(URI)是否由客户端决定**。
  例如为我的博文增加一个java的分类，生成的路径就是分类名/categories/java，那么就可以采用PUT方法。不过很多人直接把POST、GET、PUT、DELETE直接对应上CRUD，例如在一个典型的rails实现的RESTful应用中就是这么做的。
  我认为，这是因为rails默认使用服务端生成的ID作为URI的缘故，而不少人就是通过rails实践REST的，所以很容易造成这种误解。[^参考资料2]

另外，**POST请求不是幂等的，PUT请求是幂等的**，即如果希望请求不产生副作用，每次结果都一样那应该选择PUT,否则应该选择POST.

#### 1.3 PUT和DELETE请求到底安不安全

&emsp;&emsp;在设计API的时候，我看到许多文章给出的例子都是GET,POST请求，难道是PUT,DELETE在安全性上有什么
限制吗？

&emsp;&emsp;实际上GET,PUT,POST,DELETE请求对HTTP来说没什么区别，分出这四种Method主要是为了语义上的区分，让程序可以更好地交互。
对于HTTP来说都没什么区别，那么这几者在安全性（**这里说的安全是指信息安全并非上文所指的产生副作用**）上也没有什么区别了。

&emsp;&emsp;不过我在查阅资料的时候，发现了一条[^参考资料3]：

> 历史问题。
  开启了WebDAV 的 IIS 允许匿名访问者直接通过 PUT 往服务器上上传文件，曾经导致了很多严重的漏洞，具体可以搜下 IIS put 。
> 此外 apache 默认允许 trace， 又导致了一批XSS。
> 这些历史问题给运维造成很不好的印象，所以干脆把除必须的http头外都禁掉了事。

&emsp;&emsp;一些资料说是因为有些浏览器不支持PUT，DELETE请求，然而这种说法也是错误的。IE浏览器也支持这两种请求，其他浏览器更不可能不支持了。
所以不论是安全性方面还是浏览器方面，两种请求都没有问题。
造成这种情况的原因很可能是这样：
>  1.很多人贪方便，更新资源时用了GET，因为用POST必须要到FORM（表单），这样会麻烦一点。
   2.对资源的增，删，改，查操作，其实都可以通过GET/POST完成，不需要用到PUT和DELETE。
   3.另外一个是，早期的Web MVC框架设计者们并没有有意识地将URL当作抽象的资源来看待和设计，所以导致一个比较严重的问题是传统的Web MVC框架基本上都只支持GET和POST两种HTTP方法，而不支持PUT和DELETE方法。

#### 1.4 go中如何处理post请求?

&emsp;&emsp;POST请求是将数据放在请求体中，因此想办法读取读取请求体即可。
有三种方法：

``` go
     // 方法一: 直接读取请求体
      ctx,err := ioutil.ReadAll(r.Body)
      if err != nil {
          fmt.Println("post error !")
      }
     // 方法二: 先解析请求体
      r.ParseForm()
      str := r.PostForm.Get("name")

     // 方法三: 直接读取
     str := r.PostFormValue("name")
```

&emsp;&emsp;测试POST请求可使用
>     curl -d "data" "host:port//path"
>     // 测试json
>     curl -X POST -d "{\"test\": \"that\"}" host:port/path



### 2. 在go中组织路由



[^参考资料1]: https://blog.csdn.net/mfe10714022/article/details/39692305
[^参考资料2]: https://kb.cnblogs.com/page/512047/
[^参考资料3]: https://www.zhihu.com/question/38770182