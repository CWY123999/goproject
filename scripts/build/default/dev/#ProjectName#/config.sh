# 参数说明:
#   RegistryHost (必须) 镜像仓库地址, 多数情况下无需配置, 默认值为 registry.cn-beijing.aliyuncs.com
#   ImageNamespace (必须) 镜像仓库命名空间名称, 通常是以 / 分隔的一个字符串, 表示 命名空间/镜像名
#   PortList (非必须) 服务要暴露的端口, 多个端口以半角都好分隔
#   DeployDomain (必须) 部署接口服务域名, 由运维配置提供
#   DeployID (必须) 部署ID, 由运维配置提供

RegistryHost="registry.jiaxianghudong.com"
ImageNamespace="jiaxiang"
PortList="8080,8081"
DeployDomain="https://k8sapp.jiaxianghudong.com"
DeployID="12345678-abcd-efgh-ijkl-ae02319ce195"

