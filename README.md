<!-- TOC -->
  * [介绍](#介绍)
  * [技术栈](#技术栈)
  * [接口](#接口)
    * [标签](#标签)
    * [文章](#文章)
  * [MySQL容器启动](#MySQL容器启动)
  * [Dockerfile](#Dockerfile)
    * [golang镜像](#golang镜像)
      * [构建镜像](#构建镜像)
      * [构建镜像](#构建镜像)
    * [scratch镜像](#scratch镜像)
      * [编译可执行文件](#编译可执行文件)
      * [打包镜像](#打包镜像)
      * [运行镜像](#运行镜像)
<!-- TOC -->

## 介绍

本项目是一个简单的博客系统，主要用于学习golang的web开发。项目使用golang开发，使用MySQL数据库，项目使用docker进行部署，使用docker-compose进行编排。

## 技术栈

+ Gin: Golang web framework
+ beego-validation: beego form validation
+ gorm: ORM framework
+ com: common functions
+ conf: config file

## 接口

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


## MySQL容器启动

由于需要使用到MySQL服务。此处启动一个容器名称为`mysql_gin`的MySQL容器。
在app.ini文件中使用MySQL地址时就是用容器的主机名`mysql_gin:3306`
```bash
docker run --name mysql_gin -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -v /data/docker-mysql:/var/lib/mysql -d mysql
```

## Dockerfile

### golang镜像

```dockerfile
FROM golang:latest
MAINTAINER ccoderJava "congccoder@gmail.com"

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go-gin

COPY . /go-gin

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./go-gin"]
```

#### 构建镜像

```bash
docker build -t gin-blog-docker .
```
运行日志
```
 go-gin git:(main) ✗ docker build -t gin-blog-docker .
[+] Building 28.2s (9/9) FINISHED                                                                                                                                           
 => [internal] load build definition from Dockerfile                                                                                                                   0.0s
 => => transferring dockerfile: 233B                                                                                                                                   0.0s
 => [internal] load .dockerignore                                                                                                                                      0.0s
 => => transferring context: 2B                                                                                                                                        0.0s
 => [internal] load metadata for docker.io/library/golang:latest                                                                                                       2.3s
 => [1/4] FROM docker.io/library/golang:latest@sha256:dc76ef03e54c34a00dcdca81e55c242d24b34d231637776c4bb5c1a8e8514253                                                 0.0s
 => [internal] load build context                                                                                                                                      0.1s
 => => transferring context: 900.48kB                                                                                                                                  0.1s
 => CACHED [2/4] WORKDIR /go-gin                                                                                                                                       0.0s
 => [3/4] COPY . /go-gin                                                                                                                                               0.8s
 => [4/4] RUN go build .                                                                                                                                              23.1s
 => exporting to image                                                                                                                                                 1.4s 
 => => exporting layers                                                                                                                                                1.4s 
 => => writing image sha256:85b81e05147f2cd47dff3f8a37a16b48157d3cc47be4c91ab559a01d637f4ddc                                                                           0.0s 
 => => naming to docker.io/library/gin-blog-docker 
```

#### 运行镜像

```bash
 docker run --link mysql_gin:mysql_gin --name gin-blog  -p 8000:8000 gin-blog-docker 
```
运行日志
```
➜  go-gin git:(main) ✗ docker run --link mysql_gin:mysql_gin --name gin-blog  -p 8000:8000 gin-blog-docker                               
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /auth                     --> go-gin/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> go-gin/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> go-gin/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> go-gin/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> go-gin/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> go-gin/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> go-gin/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> go-gin/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> go-gin/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> go-gin/routers/api/v1.DeleteArticle (4 handlers)
```

### scratch镜像
```dockerfile
FROM scratch
MAINTAINER ccoderJava "congccoder@gmail.com"

WORKDIR /go-gin

COPY . /go-gin

EXPOSE 8000

CMD ["./go-gin"]
```

#### 编译可执行文件

编译所生成的可执行文件会依赖一些库，并且是动态链接。所以需要将这些库一起打包到镜像中。
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin . 
```

#### 打包镜像

```bash
docker build -t gin-blog-docker-scratch .  
```
```
[+] Building 2.3s (6/6) FINISHED                                                                                                                                            
 => [internal] load build definition from Dockerfile                                                                                                                   0.0s
 => => transferring dockerfile: 36B                                                                                                                                    0.0s
 => [internal] load .dockerignore                                                                                                                                      0.0s
 => => transferring context: 2B                                                                                                                                        0.0s
 => [internal] load build context                                                                                                                                      1.4s
 => => transferring context: 28.46MB                                                                                                                                   1.4s
 => CACHED [1/2] WORKDIR /go-gin                                                                                                                                       0.0s
 => [2/2] COPY . /go-gin                                                                                                                                               0.4s
 => exporting to image                                                                                                                                                 0.2s
 => => exporting layers                                                                                                                                                0.2s
 => => writing image sha256:d0277b07633b2c30b73ad15620094e58a1f039daffcfd5f794f6e1bf061003ef                                                                           0.0s
 => => naming to docker.io/library/gin-blog-docker-scratch  
```

#### 运行镜像

```bash
 docker run --link mysql_gin:mysql_gin --name gin-blog  -p 8000:8000 gin-blog-docker-scratch
```
运行日志
```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
2022-12-02T05:06:21.600530133Z  - using env:	export GIN_MODE=release
2022-12-02T05:06:21.600534883Z  - using code:	gin.SetMode(gin.ReleaseMode)
2022-12-02T05:06:21.600538342Z 
2022-12-02T05:06:21.600866258Z [GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
2022-12-02T05:06:21.600875967Z [GIN-debug] GET    /auth                     --> go-gin/routers/api.GetAuth (3 handlers)
2022-12-02T05:06:21.600881550Z [GIN-debug] GET    /api/v1/tags              --> go-gin/routers/api/v1.GetTags (4 handlers)
2022-12-02T05:06:21.600886383Z [GIN-debug] POST   /api/v1/tags              --> go-gin/routers/api/v1.AddTag (4 handlers)
2022-12-02T05:06:21.600892050Z [GIN-debug] PUT    /api/v1/tags/:id          --> go-gin/routers/api/v1.EditTag (4 handlers)
2022-12-02T05:06:21.600897800Z [GIN-debug] DELETE /api/v1/tags/:id          --> go-gin/routers/api/v1.DeleteTag (4 handlers)
2022-12-02T05:06:21.600903133Z [GIN-debug] GET    /api/v1/articles          --> go-gin/routers/api/v1.GetArticles (4 handlers)
2022-12-02T05:06:21.600908050Z [GIN-debug] GET    /api/v1/articles/:id      --> go-gin/routers/api/v1.GetArticle (4 handlers)
2022-12-02T05:06:21.600913842Z [GIN-debug] POST   /api/v1/articles          --> go-gin/routers/api/v1.AddArticle (4 handlers)
2022-12-02T05:06:21.600919717Z [GIN-debug] PUT    /api/v1/articles/:id      --> go-gin/routers/api/v1.EditArticle (4 handlers)
2022-12-02T05:06:21.600926008Z [GIN-debug] DELETE /api/v1/articles/:id      --> go-gin/routers/api/v1.DeleteArticle (4 handlers)
```
