# In-Charge 后端接口说明

## 返回值约定
``` js
{
    code: int,
    msg: string,
    data: obejct
}
```
- code 
状态码，见 [response.go](https://github.com/cildhdi/In-charge/blob/master/utils/response.go) 中的状态码定义。
- msg
状态码非 ```Ok``` 时对应的描述。
- data
成功时返回的数据，以下 Response 的格式均指 data 的格式。

### GET /status
服务器当前状态

### POST /api/send-code
发送验证码
#### Request
``` js
{
    phone: string
}
```
- phone: 电话号码，仅做 11 位字符串校验

### POST /api/login
登录
#### Request
``` js
{
    phone: string,
    code: string
}
```
- phone: 同上
- code: 短信验证码，仅做 4 位数字校验

#### Response
``` js
{
    token: string
}
```
- token: jwt 字符串，需要保存到 cookie/localStorage 中，后续调用用户相关接口时，需要以 ``` "Bearer ${token}" ``` 的格式放在 request header 的 Authorization 中，当 外层 code 不为 ``` Ok ``` 时请重新登录。