# 配置文件存储到 consul

把`testdata/configuration/configs`目录下的文件，按文件名存储到consul

路径前缀：${app.project_name}/${app.server_name}/${app.server_env}/${app.version}；例如：

* go-micro-saas/ping-service/production/v1.0.0/app.yaml
* go-micro-saas/ping-service/production/v1.0.0/mysql.yaml
* go-micro-saas/ping-service/production/v1.0.0/filename.yaml

```shell
go run testdata/configuration/main.go
```