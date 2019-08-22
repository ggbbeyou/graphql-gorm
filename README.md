# graphql-gorm

使用gqlgen和gorm自动生成GraphQL接口代码 [gqlgen](https://gqlgen.com) and [gorm](https://gorm.io)

## 依赖包说明
使用go mod 进行依赖包管理，首先进行一下依赖安装
```
go get -u github.com/maiguangyang/graphql
go get -u golang.org/x/tools/cmd/goimports
go mod edit -replace=github.com/satori/go.uuid@v1.2.0=github.com/satori/go.uuid@master
go mod tidy
```
安装过程如果出现无法安装的依赖包，请自行翻墙、设置代理或者从github上找到对应包，clone到对应目录。

## 设置代理（二选一）
```
export GOPROXY="https://athens.azurefd.net"
export GOPROXY="https://goproxy.io"
```

## 开发说明

- 设置数据库连接
在根目录makefile文件里面进行设置
```
DATABASE_URL=mysql://'root:123456@tcp(localhost:3306)/graphql?charset=utf8mb4&parseTime=True&loc=Local' PORT=80 go run *.go
```

- 生成新的GraphQL接口代码
编辑根目录 model.graphql 文件

```
type User {
  email: String @column(gorm: "type:varchar(64) comment '用户邮箱地址';NOT NULL;default:0;") @validator(required: "true", type: "email")
  age: Int
  firstName: String
  lastName: String

  tasks: [Task!]! @relationship(inverse:"assignee")
}

type Task {
  title: String
  completed: Boolean
  dueDate: Time

  assignee: User @relationship(inverse:"tasks")
}
```
@column：参数为gorm，主要用来定义表结构到一些信息
@validator：字段校验，required是否必填、type正则匹配校验


------------
如果对model.graphql做了修改，必须运行以下命令，生成最新的GraphQL接口代码
```
make generate
或
go run github.com/maiguangyang/graphql
```
然后运行
```
make run
```

## 新增一个自定义方法
- 1.在model.graphql增加
```
extend type Mutation {
  token(input: String): String!
}
```
- 2. 在src/resolver.go 增加
```
func (r *MutationResolver) Token(ctx context.Context, input *string) (string, error) {
  return "token", nil
}
```
