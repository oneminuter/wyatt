## 获取个人信息
**Get: /user/info**

### 请求参数
```shell
{
    "uId":"string" //用户标识
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiYzQ5ZTBhZmQ4MGZhNDIzZWY0NjU0MzYyNTJmMjgyNiIsImV4cCI6MTU0NTczNDIyMCwianRpIjoiMSIsImlhdCI6MX0.1tGsLd094pImTZYCCwPbdgG4fusvV0uj-qD_OiAFI1M", //token
        "account": "1234", //账号
        "uuid": "bc49e0afd80fa423ef465436252f2826", //uuid
        "email": "1050086935@qq.com", //邮箱
        "nickName": "  一分钟", //昵称
        "sex": 1, //性别
        "name": "小林", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "signature": "each youth will be old", //个性签名
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP", //城市
        "avaliable": 0, //可用积分
        "level": 0 //等级
    }
}
```

***

## 注册
**Post: /user/register**

### 请求参数
```shell
{
    "account":"string", //账号
    "password": "string" //密码
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiYzQ5ZTBhZmQ4MGZhNDIzZWY0NjU0MzYyNTJmMjgyNiIsImV4cCI6MTU0NTczNDIyMCwianRpIjoiMSIsImlhdCI6MX0.1tGsLd094pImTZYCCwPbdgG4fusvV0uj-qD_OiAFI1M", //token
        "account": "1234", //账号
        "uuid": "bc49e0afd80fa423ef465436252f2826", //uuid
        "email": "1050086935@qq.com", //邮箱
        "nickName": "  一分钟", //昵称
        "sex": 1, //性别
        "name": "小林", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "signature": "each youth will be old", //个性签名
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP", //城市
        "avaliable": 0, //可用积分
        "level": 0 //等级
    }
}
```

***

## 创建临时账号
**Post: /user/temp/create**

### 请求参数
无

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiYzQ5ZTBhZmQ4MGZhNDIzZWY0NjU0MzYyNTJmMjgyNiIsImV4cCI6MTU0NTczNDIyMCwianRpIjoiMSIsImlhdCI6MX0.1tGsLd094pImTZYCCwPbdgG4fusvV0uj-qD_OiAFI1M", //token
        "account": "1234", //账号
        "uuid": "bc49e0afd80fa423ef465436252f2826", //uuid
        "email": "1050086935@qq.com", //邮箱
        "nickName": "  一分钟", //昵称
        "sex": 1, //性别
        "name": "小林", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "signature": "each youth will be old", //个性签名
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP", //城市
        "avaliable": 0, //可用积分
        "level": 0 //等级
    }
}
```

***

## 登录
**Post: /user/login**

### 请求参数
```shell
{
    "account":"string", //账号
    "phone": "string", //手机号
    "email":"string", //邮箱
    "password": "string" //密码（必须）
}
```
请求参数 account，phone，email 必须有其一

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiYzQ5ZTBhZmQ4MGZhNDIzZWY0NjU0MzYyNTJmMjgyNiIsImV4cCI6MTU0NTczNDIyMCwianRpIjoiMSIsImlhdCI6MX0.1tGsLd094pImTZYCCwPbdgG4fusvV0uj-qD_OiAFI1M", //token
        "account": "1234", //账号
        "uuid": "bc49e0afd80fa423ef465436252f2826", //uuid
        "email": "1050086935@qq.com", //邮箱
        "nickName": "  一分钟", //昵称
        "sex": 1, //性别
        "name": "小林", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "signature": "each youth will be old", //个性签名
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP", //城市
        "avaliable": 0, //可用积分
        "level": 0 //等级
    }
}
```

***

## 修改用户信息
**Post: /user/info/modify**

### 请求参数
```shell
{
    "phone":"string", //手机号
    "nickname": "string", //昵称
    "sex":"string", //性别
    "name": "string", //名字
    "email":"string", //邮箱
    "avatarUrl":"string", //头像
    "signature":"string", //个性签名
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiYzQ5ZTBhZmQ4MGZhNDIzZWY0NjU0MzYyNTJmMjgyNiIsImV4cCI6MTU0NTczNDIyMCwianRpIjoiMSIsImlhdCI6MX0.1tGsLd094pImTZYCCwPbdgG4fusvV0uj-qD_OiAFI1M", //token
        "account": "1234", //账号
        "uuid": "bc49e0afd80fa423ef465436252f2826", //uuid
        "email": "1050086935@qq.com", //邮箱
        "nickName": "  一分钟", //昵称
        "sex": 1, //性别
        "name": "小林", //名字
        "avatarUrl": "http://blog.oneminuter.com/favicon.ico", //头像
        "signature": "each youth will be old", //个性签名
        "country": "XX", //国家
        "province": "XX", //省份
        "city": "内网IP", //城市
        "avaliable": 0, //可用积分
        "level": 0 //等级
    }
}
```

***

## 修改用户账号

**Post /user/account/modify**

### 请求参数
```shell
{
    "account":"string" //新账号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 修改登录密码

**Post /user/password/modify**

### 请求参数
```shell
{
    "oldPassword":"string", //旧密码
    "newPassword":"string", //新密码
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 所有社区列表
**Get /community/list/all**

### 请求参数
```shell
{
    "page": "int", //页码
    "limit": "int" //查询条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "createdAt": 1538035982, //创建时间
            "cId": "CMT.1234567890", //社区流水号
            "logo": "http://oneminuter.com/favicon.ico", //社区logo
            "name": "一分钟社区", //社区名
            "desc": "一分钟社区", //社区介绍
            "joinNum": 0, //加入人数
            "articleNum": 0 //文章话题数
        }
    ]
}
```

***

## 加入社区
**Post /community/join**

### 请求参数
```shell
{
    "cId":"string" //社区流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 退出社区
**Post /community/exit**

### 请求参数
```shell
{
    "cId": "string" //社区流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 加入的社区列表

**Get /community/list/my**

### 请求参数
```shell
{
    "page": "int", //页码
    "limit": "int" //查询条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "createdAt": 1538035982, //创建时间
            "cId": "CMT.1234567890", //社区流水号
            "logo": "http://oneminuter.com/favicon.ico", //logo
            "name": "一分钟社区", //社区名
            "desc": "一分钟社区", //简介
            "joinNum": 1, //加入人数
            "articleNum": 0 //文章数
        }
    ]
}
```

***

## 创建社区

**Post /community/create**

### 请求参数
```shell
{
    "name": "string", // 社区名
    "desc": "string" // 社区简介
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 修改社区

**Post /community/modify**

### 请求参数
```shell
{
    "cId": "string", //社区流水号id, 必须
    "logo": "string", // 社区logo，非必须
    "name": "string", // 社区名，非必须
    "desc": "string" // 社区简介，非必须
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 删除社区
删除为软删除，改变社区的状态

**Post /community/delete**

### 请求参数
```shell
{
    "cId": "string" // 社区流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 添加社区管理员

**Post /community/manager/add**

### 请求参数
```shell
{
    "cId": "string", //社区流水号
    "account": "string" //被添加目标用户的账号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 删除社区管理员

**Post /community/manager/add**

### 请求参数
```shell
{
    "cId": "string", //社区流水号
    "account": "string" //被删除目标用户的账号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 添加话题

**Post /topic/add**

### 请求参数
```shell
{
    "cId": "string" //社区流水号
    "title": "string" //标题
    "desc": "string" //简介或者内容
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 话题列表

**Get /topic/list**

### 请求参数
```shell
{
    "cId": "string" //社区流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "tId": "TP.1539533494", //话题流水号
            "title": "这是话题标题", //标题
            "desc": "这是话题内容", //内容或简介
            "cId": "CMT.1538754033", //所属社区流水号
            "creatorAccount": "1234", //发布者账号
            "creatorAvatarUrl": "http://blog.oneminuter.com/favicon.ico", //发布者头像
            "createTime": 1539533494, //创建时间
            "viewedNum": 0, //浏览量
            "zanNum": 0, //点赞数
            "commentNum": 0 //评论数
        }
    ]
}
```
***

## 修改话题

**Post /topic/modify**

### 请求参数
```shell
{
    "tId": "string", //话题流水号
    "title": "string", //标题
    "desc": "string" //内容或简介
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 话题详情

**Get /topic/detail**

### 请求参数
```shell
{
    "tId": "string" //话题流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "tId": "TP.12.1539792891", //话题流水号
        "title": "标题标题标题标题", //标题
        "desc": "修改后的内容1", //内容
        "cId": "CMT.1538754033", //所属社区流水号
        "creatorAccount": "1234", //发布者账号
        "creatorAvatarUrl": "http://blog.oneminuter.com/favicon.ico", //发布者头像
        "createTime": 1539792891, //发布时间
        "viewedNum": 0, //浏览量
        "zanNum": 0, //点赞数
        "commentNum": 0 //评论数量
    }
}
```

***

## 删除话题

**Post /topic/delete**

### 请求参数
```shell
{
    "tId": "string" //话题流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 收藏话题

**Post /topic/collect/add**

### 请求参数
```shell
{
    "tId":"string" //话题流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 取消收藏话题

**Post /topic/collect/cancel**

### 请求参数
```shell
{
    "tId":"string" //话题流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 收藏话题列表

**Post /topic/collect/list**

### 无请求参数

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
           "tId": "TP.1539792891", //话题流水号
           "title": "标题标题标题标题", //标题
           "desc": "修改后的内容1", //内容
           "cId": "CMT.1538754033", //所属社区流水号
           "creatorAccount": "1234", //发布者账号
           "creatorAvatarUrl": "http://blog.oneminuter.com/favicon.ico", //发布者头像
           "createTime": 1539792891, //发布时间
           "viewedNum": 0, //浏览量
           "zanNum": 0, //点赞数
           "commentNum": 0 //评论数量
           "commentNum": 0
        }
    ]
}
```

***

## 添加评论

**Post /comment/add**

### 请求参数
```shell
{
    "articleId":"string" //文章流水号
    "content":"string" //评论类容
    "replyCid":"string" //回复流水号，非必须参数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 评论列表

**Get /comment/list**

### 请求参数
```shell
{
    "articleId"："string", //文章流水号
    "page":"int", //页码
    "limit":"int" //条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "cid": "CM.1.1541220122", //评论流水号
            "userAccount": "123", //评论者账号
            "userAvatarUrl": "http://blog.oneminuter.com/favicon.ico", //评论者头像
            "createdAt": 1541220123, //评论时间
            "content": "这是评论内容", //评论内容
            "replyCid": "" //回复id
        }
    ]
}
```

***

## 删除评论

**Post /comment/delete**

### 请求参数
```shell
{
    "articleId": "string" //文章流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 点赞

**Post /zan/add**

### 请求参数
```shell
{
    "sourceFlowId":"string" //点赞对象的流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 取消点赞

**Post /zan/cancel**

### 请求参数
```shell
{
    "sourceFlowId":"string" //点赞对象的流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 获得的点赞列表

**Get /zan/list**

### 请求参数
```shell
{
    "page":"int" //请求页码
    "limit":"int" //请求条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "userAccount": "123", //点赞用户账号
            "userAvatarUrl": "http://blog.oneminuter.com/favicon.ico", //点赞用户头像
            "sourceFlowId": "TP.2.1541326047", //点赞的资源信息流水号
            "classify": "话题", //点赞资源的分类
            "createdAt": 1541347056 //点赞时间
        }
    ]
}
```

***

## 订阅用户 - 成为其粉丝
**Post /fans/follow**

### 请求参数
```shell
{
    "userAccount": "string" //用户账号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 取消订阅用户
**Post /fans/cancel**

### 请求参数
```shell
{
    "userAccount": "string" //用户账号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 粉丝列表
**Get /fans/list**

### 请求参数
```shell
{
    "page":"int", //请求页码
    "limit":"int" //请求条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "account": "1234", //用户账号
            "avatarrl": "http://blog.oneminuter.com/favicon.ico", //用户头像
            "nickName": "1234", //用户昵称
            "sex": 0, //用户性别
            "signature": "" //个性签名
        }
    ]
}
```

***

## 订阅列表
**Get /follow/list**

### 请求参数
```shell
{
    "page":"int", //请求页码
    "limit":"int" //请求条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "account": "1234", //用户账号
            "avatarrl": "http://blog.oneminuter.com/favicon.ico", //用户头像
            "nickName": "1234", //用户昵称
            "sex": 0, //用户性别
            "signature": "" //个性签名
        }
    ]
}
```

***

## 消息列表

**Get /message/list**

### 请求参数
```shell
{
    "page":"int", //请求页码
    "limit":"int" //请求条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "mId": "MG.1.1541692635", //消息流水号
            "msgType": "system", //消息类型
            "content": "这是一个测试消息", //消息内容
            "isViewed": 0, //是否查看过
            "createdAt": 1541692635 //消息创建时间
        }
    ]
}
```

***

## 消息详情

**Get /message/detaill**

### 请求参数
```shell
{
    "mId": "string" //消息流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "mId": "MG.1.1541692635", //消息流水号
        "msgType": "system", //消息类型
        "content": "这是一个测试消息", //消息内容
        "isViewed": 0, //是否查看过
        "createdAt": 1541692635 //消息创建时间戳
    }
}
```

***

## 消息查看回调

**Get /message/viewed**

### 请求参数
```shell
{
    "mId": "string" //消息流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 删除消息

**Get /message/delete**

### 请求参数
```shell
{
    "mId": "string" //消息流水号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 用户建议

**Post /advise/add**

### 请求参数
```shell
{
    "content":"string", //建议内容, 必须
    "phone":"string", //建议人手机号， 非必须
    "email":"string", //建议人邮箱，非必须
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 用户建议列表

**Post /advise/list**

### 请求参数
```shell
{
    "page":"int", //页码
    "limit":"int", //请求条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "content": "建议123", //建议内容
            "status": 0, //处理状态
            "remark": "", //处理备注
            "createdAt": 1543074241 //处理时间
        }
    ]
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 举报

**Post /tip/add**

### 请求参数
```shell
{
    "sourceFlowId":"string", //举报内容流水号id
    "reason":"string", //举报原因
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 添加故事

**Post /story/add**

### 请求参数
```shell
{
    "title":"string", //故事标题
    "desc":"string", //标题
    "classify":"string", //分类
    "coverImg":"http://blog.oneminuter.com/favicon.ico" //封面图
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 故事列表

**Get /story/list**

### 请求参数
```shell
{
    "page":"int", //请求页码
    "limit":"int", //请求条数
    "userAccour":"string", //作者账号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "storyId": "S.1.1543076097", //故事流水号id
            "title": "一个人的故事", //故事标题
            "desc": "", //故事简介
            "classify": "", //分类
            "coverImg": "", //封面图
            "author":"oneminuter", //作者账号
            "avatarUrl":"", //作者头像
            "viewedNum": 0, //浏览量
            "zanNum": 0, //点赞数
            "commentNum": 0 //评论数
        }
    ]
}
```

***

## 故事修改

**Post /story/modify**

### 请求参数
```shell
{
    "storyId":"string", //故事流水号id， 必须
    "title":"string", //故事标题，非必须
    "desc":"string", //简介，非必须
    "coverImg":"string", //封面图链接，非必须
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 添加角色

**Post /role/add**

### 请求参数
```shell
{
    "storyId":"string", //为哪个故事添加角色，该故事的流水号id，该参数必须
    "nickname":"string", //角色昵称，该参数必须
    "sex":"int", //角色性别，非必须
    "avatarUrl": "string", //角色头像
    "introduce": "string", //角色介绍
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 修改角色信息

**Post /role/modify**

### 请求参数
```shell
{
    "rolerId":"string", //角色流水号id，该参数必须
    "nickname":"string", //角色昵称，该参数必须
    "sex":"int", //角色性别，非必须
    "avatarUrl": "string", //角色头像
    "introduce": "string", //角色介绍
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 删除角色

**Post /role/delete**

### 请求参数
```shell
{
    "rolerId":"string", //角色流水号id
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": null
}
```

***

## 故事角色列表

**Get /role/list**

### 请求参数
```shell
{
    "storyId":"string", //角色流水号id
    "page":"int", //请求页码
    "limit":"int", //请求条数
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [{
        "rolerId": "R.1.1543076097", //角色流水号
        "avatarUrl":"http://blog.oneminuter.com/favicon.ico", //角色头像
        "nickname":"oneminuter", //角色昵称
        "sex":1, //角色性别
        "introduce":"你不懂我，我不怪你", //角色介绍
    }]
}
```

***

## 故事角色信息

**Get /role/info**

### 请求参数
```shell
{
    "rolerId":"string", //角色流水号id
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": {
        "rolerId": "R.1.1543076097", //角色流水号
        "avatarUrl":"http://blog.oneminuter.com/favicon.ico", //角色头像
        "nickname":"oneminuter", //角色昵称
        "sex":1, //角色性别
        "introduce":"你不懂我，我不怪你", //角色介绍
    }
}
```

***

## 系列故事列表

**Get /story/series/list**

### 请求参数
```shell
{
    "page":"int", //请求页码
    "limit":"int", //请求条数
    "userAccour":"string", //作者账号
}
```

### 返回参数
```shell
{
    "errCode": 200,
    "errMsg": "",
    "data": [
        {
            "seriesId": "S.1.1543076097", //系列流水号id
            "title": "一个人的故事", //系列标题
            "desc": "", //系列简介
            "classify": "", //分类
            "coverImg": "", //封面图
            "author":"oneminuter", //作者账号
            "avatarUrl":"", //作者头像
            "viewedNum": 0, //浏览量
            "zanNum": 0, //点赞数
            "commentNum": 0 //评论数
        }
    ]
}
```
