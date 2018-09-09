## 获取个人信息
**Get: /user/info**

### 请求参数
```
{
    "uId":"xxxx" //用标识
}
```

### 返回参数
```
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI0ZThhOTY2MTRkNTU2NWIwOWJkY2EzMDQxYzFhZmU2MiIsImV4cCI6MTUzODE4ODExNCwianRpIjoiMSJ9.a4ACiS2kTkr2mWeC3cPpD43LgIZIoVaH-LtyunP0DcM", //用户标识jwt
        "account": "y0KjesQn7Q", //账号
        "uuid": "4e8a96614d5565b09bdca3041c1afe62", //用户标识
        "nickName": "y0KjesQn7Q", //昵称
        "sex": 0, //性别
        "name": "", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP" //城市
    }
}
```

***

## 注册
**Post: /user/register**

### 请求参数
```
{
    "account":"xxxx", //账号
    "password": "xxx" //密码
}
```

### 返回参数
```
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI0ZThhOTY2MTRkNTU2NWIwOWJkY2EzMDQxYzFhZmU2MiIsImV4cCI6MTUzODE4ODExNCwianRpIjoiMSJ9.a4ACiS2kTkr2mWeC3cPpD43LgIZIoVaH-LtyunP0DcM", //用户标识jwt
        "account": "y0KjesQn7Q", //账号
        "uuid": "4e8a96614d5565b09bdca3041c1afe62", //用户标识
        "nickName": "y0KjesQn7Q", //昵称
        "sex": 0, //性别
        "name": "", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP" //城市
    }
}
```

***

## 登录
**Post: /user/login**

### 请求参数
```
{
    "account":"xxxx", //账号
    "phone":"138xxxx", //手机号
    "email":"10500@qq.com", //邮箱
    "password": "xxx" //密码（必须）
}
```
请求参数 account，phone，email 必须有其一

### 返回参数
```
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI0ZThhOTY2MTRkNTU2NWIwOWJkY2EzMDQxYzFhZmU2MiIsImV4cCI6MTUzODE4ODExNCwianRpIjoiMSJ9.a4ACiS2kTkr2mWeC3cPpD43LgIZIoVaH-LtyunP0DcM", //用户标识jwt
        "account": "y0KjesQn7Q", //账号
        "uuid": "4e8a96614d5565b09bdca3041c1afe62", //用户标识
        "nickName": "y0KjesQn7Q", //昵称
        "sex": 0, //性别
        "name": "", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP" //城市
    }
}
```