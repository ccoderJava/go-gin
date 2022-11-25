# go-gin-examples

+ Gin: Golang web framework
+ beego-validation: beego form validation
+ gorm: ORM framework
+ com: common functions
+ conf: config file

## interfaces

### 标签

+ 获取标签列表： `GET /tags`
+ 新建标签： `POST /tags`
+ 更新指定标签： `PUT /tags/:id`
+ 删除指定标签： `DELETE /tags/:id`



### 文章
+ 获取多个文章： `GET /articles`
+ 获取单个文章： `GET /articles/:id`
+ 新建文章： `POST /articles`
+ 更新指定文章： `PUT /articles/:id`
+ 删除指定文章： `DELETE /articles/:id`

### 文件处理

+ 上传图片： `POST /upload`
+ 获取图片： `GET /upload/:path`