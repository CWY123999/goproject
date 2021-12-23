## Usage:

### Makefile

> 项目构建脚本, 执行时会以当前脚本所在目录名构建对应的执行文件到`_bin/`目录下。

```shell
# 使用 golangci-lint 进行代码检查
make lint

# 运行测试脚本
make test

# 竞态检查
make race

# 构建当前操作系统到执行文件
make

# 构建 Linux 64 位系统执行文件
make linux

# 构建 MacOS 64 位系统执行文件
make macos
# 或
make darwin

# 构建 Windows 64 位系统执行文件
make windows

# 构建 Linux、MacOS、Windows 3个平台的 64 位执行文件
make all

# 构建 Linux 32 位系统执行文件
make linux32

# 构建 MacOS 32 位系统执行文件
make macos32
# 或
make darwin32

# 构建 Windows 32 位系统执行文件
make windows32

# 构建 Linux、MacOS、Windows 3个平台的 32 位执行文件
make all32

# 构建用于构建生产镜像的发布版本
make release

# 对于多个执行程序的情况，可以使用 name 参数进行指定, 其必须与 cmd 下的 main.go 所在的子目录名称一致. 如:
# make linux name=userapi
```

### .golangci.yml

> 通用 golangci-lint 代码检查工具配置文件, 这也是目前 CI 脚本中执行`golangci-lint run`的默认配置, 如果你要对 Golang 项目启用 CI, 则需将本文件下载至你项目对根目录中。

### .gitlab-ci.yml

> 是一个针对 Golang 项目的通用 GitLab-CI 脚本, 它依赖于`Makefile`和`.golangci.yml`文件。

#### 要启用`CI`，你需要：

**1. 将`.gitlab-ci.yml`放至项目根目录下；**

**2. 创建`CI`配置文件，位置是：`项目根目录/build/集群名称/环境标识/执行文件名/config.sh`，内容如下:**

> 注：默认的集群名称应为：default

```shell script
# 参数说明:
#   RegistryHost (必须) 镜像仓库地址, 多数情况下无需配置, 默认值为 registry.cn-beijing.aliyuncs.com
#   ImageNamespace (必须) 镜像仓库命名空间名称
#   PortList (非必须) 服务要暴露的端口, 多个端口以半角逗号分隔
#   DeployDomain (必须) 部署接口服务域名, 由运维配置提供
#   DeployID (必须) 部署ID, 由运维配置提供

RegistryHost="registry.cn-beijing.aliyuncs.com"
ImageNamespace="test"
PortList="8080,8081"
DeployDomain="https://k8sapp.jiaxianghudong.com"
DeployID="12345678-abcd-efgh-ijkl-ae02319ce195"
```
你可以根据项目需要配置多个环境、多个执行程序的CI配置。

**3. 放入程序配置文件**

执行程序自己的配置文件可以放在 `项目根目录/configs/集群名称/环境标识/执行文件名/` 目录下，数量不限，构建镜像时会自动将其拷贝至执行文件所在目录中。

**4. 触发CI流水线**

通过浏览器进入项目仓库中，点选左侧`CI/CD - 流水线`，之后点击右上的`运行流水线`，选择要部署的分之后填入`CLUSTER`、`ENV`、`CMD`、`JUMPTEST` 4 个变量即可开始。

考虑到项目会有不同集群、环境和多个执行程序的情况，默认所有构建工作都是需要手动在项目仓库中触发的。

**5. 一些注意事项**

大多数情况下，项目开发都会存在一个内网的测试环境，该`CI`配置同样支持此类场景。这需要由运维人员在内网运行一个`Gitlab-CI-Runner`并注册到项目或共享 Runner 中，并为其配置一个独立的触发`tags`（如：dev），之后在`运行流水线`时，`ENV`参数填入该`tags`（如：dev）即可由内网的 Runner 进行构建处理。