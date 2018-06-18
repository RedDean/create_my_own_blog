### 如何处理json

参考资料:

[JSON处理](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md)

#### 1.1 序列化

#### 1.2 反序列化

##### 1.2.1 数据结构已知

&emsp;&emsp;**Decoder:**
``` goalng
   type test_struct struct {
       Test string
   }

   func test(rw http.ResponseWriter, req *http.Request) {
       decoder := json.NewDecoder(req.Body)
       var t test_struct
       err := decoder.Decode(&t)
       if err != nil {
           panic(err)
       }
       log.Println(t.Test)
   }
```
##### 1.2.2 数据结构未知

&emsp;&emsp;**官方做法**:

``` golang

    s := []byte(str)
    var f interface{}

    err := json.Unmarshal(s, &f)
    if err != nil {
        fmt.Println("Unmarshall Error!")
    }
```

&emsp;&emsp;这时f 存储的类型为:

``` golang
    f = map[string]interface{}{
        "aaa": "xxx",
        "bbb": 5,
        "ccc": []interface{}{
            "ddd",
            "eee"
        }
    }
```

&emsp;&emsp;使用断言访问数据

``` golang
   m := f.(map[string]interface{})

   for k,v := range m {
       switch vv := v.(type) {
           case string:
                fmt.Println(k, "is string", vv)
           case int:
                fmt.Println(k, "is int", vv)
           case []interface{}:
                fmt.Println(k, "is array:")
                for i, u := range vv {
                    fmt.Println(i,u)
                }
           default:
                fmt.Println("Unkown Type")
           }
    }
```

&emsp;&emsp;**simplejson:**

``` golang
   js, err := NewJson([]byte(`{
   	"test": {
   		"array": [1, "2", 3],
   		"int": 10,
   		"float": 5.150,
   		"bignum": 9223372036854775807,
   		"string": "simplejson",
   		"bool": true
   	}
   }`))

   arr, _ := js.Get("test").Get("array").Array()
   i, _ := js.Get("test").Get("int").Int()
   ms := js.Get("test").Get("string").MustString()
```
&emsp;&emsp; 以上两种方式都需要直接读取请求体。
***













