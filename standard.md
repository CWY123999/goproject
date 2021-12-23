# Golang编码规范
> 当前版本：v0.1，最后编辑日期：2020-01-14  

## 基础要求
* 所有代码必须经过`gofmt`格式化；
* 所有代码必须经过`golangci-lint`（配置为：[查看 .golangci.yml](https://git.jiaxianghudong.com/go/goproject/-/blob/master/scripts/.golangci.yml)）检测通过；
> https://github.com/golangci/golangci-lint  
* 对于非工具包的常规项目，其目录结构必须以 https://github.com/golang-standards/project-layout 为准，可使用 https://git.jiaxianghudong.com/go/goproject 自动生成该结构；
* 每行代码建议不超过`80`个字符，必须不能超过`120`个字符；
* 函数/方法体内的代码，除一些多字段`Struct/Map/Slice`的声明外，其体量要尽可能小，建议不超过`100`行；严禁出现偷懒导致单一函数/方法体内堆积大量可拆分、复用的代码；
* 项目的包管理工具**必须使用** Golang 自带的`go mod`；
* 缩进一律使用`tab`（一般为4个空格的宽度）；

## 注释
* 代码的任何部分皆**不允许**出现块注释（`/*……*/`），一律使用行注释（`//`）对代码进行说明，且每个`//`符号后要留一个空格，例：

```go
// this is a comment
```

* 包、类型、函数/方法的注释中，第一个单词必须是它的名称（严格遵守大小写），即：

```go
// fruit is a toolkit about fruits
package fruit

// Fruit define all fruits 
type Fruit struct {
	Price int
}

// Sell sell fruit
func (f *Fruit) Sell() {
	// something ……
}
```

* 除`main`外的所有包都应该至少在其一个包文件的`package`前对其进行注释介绍该包的作用及注意事项；
* 对于函数/方法声明中的`可变参数`（即：... parameters T），需在注释中对齐进行说明，其如果大于 1 的有限数量的情况下，应尽量对其每项元素进行解释；

## 声明约束
* 类型声明中的字段顺序应满足内存对齐要求；
* 宏定义、变量、类型、函数/方法的名称必须遵循驼峰式规范，禁止出现`_`在内的任何特殊符号；
* 对于名称中可能出现的专有缩写名词，要严格按照其国际规范命名。如：

```go
func parseUrl() // 错误
func parseURL() // 正确

func GetHtml() // 错误
func GetHTML() // 正确

func main() {
	var ApiXXX // 错误
	var APIXXX // 正确
}
```

* 变量名声明**不要**带有任何类型前缀，如`var Icount int`这样表示`int`类型的东西，而是应直接使用`count`为名；但`bool`类型应尽量带有`is`、`Is`、`has`、`Has`、`can`、`Can`、`allow`、`Allow`此类带有在自然语言中明确真假/是否的前缀；
* 接口命名应遵守`动词+er`的形式，如：Writer、Reader；
* 包导入应按其导入路径进行分组，顺序为：标准库、本项目自有包、团队内部包、第三方包，例：

```go
import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gsdkserver/internal/pkg/cfg"
	"gsdkserver/internal/pkg/crypto"
	"gsdkserver/internal/pkg/env"
	"gsdkserver/internal/pkg/user"

	"git.jiaxianghudong.com/go/alisls"
	"git.jiaxianghudong.com/go/db"
	"git.jiaxianghudong.com/go/utils"
	"git.jiaxianghudong.com/go/xlog"

	"github.com/hprose/hprose-golang/rpc"
	"github.com/nu7hatch/gouuid"
)
```

## 文本规范
* 除特殊需要给到终端或者用户的提示外，项目内部的日志和输出内容，句首单次的首字母`禁止使用大写`；
* 日志文本末尾不能带有标点符号（如：`.`或`。`）；
* 对于使用`Format`类函数（如：`fmt.Printf`）的场景，如果最终用于日志打印，那么其格式应为：

```go
fmt.Printf("some texts, var1: %v, var2: %s, var3: %d", []byte(1), "aaa", 123)
```

即：对于需格式化显示的变量，必须带有它的名称， 名称后直接跟一个半角冒号+空格`: `（这里不能用其他任何东西替代）；
* 对于输出中带有`error`字样的内容，不允许简写成`err`；
* 所有简拼、缩写文本必须遵守国际规范。如：`parameter`的缩写是`param`而不是`parm`；`service`/`server`可缩写为`srv`不能是`svr`；
* 汉字（不含符号）中间的英文字符、英文单词和数字的左右两边要加`1个`空格，例：

```
变量 test 左右两边要有空格，另外数字 123 的左右两边也有空格；但如果内容是以英文字符、英文单词和数字结尾的，其末尾不需要再有空格了，比如 end。
```

## 代码仓库
* 项目的正式代码或内容其命名的首字符不允许出现下划线`_`，以下划线`_`开头的文件/目录皆要在 git 中忽略；
* 任何项目的根目录必须带有`.gitignore`文件，要对自己使用 IDE 生成的项目文件或临时文件进行忽略，例：

```
/.idea/
.project*
.svn
Thumbs.db
ehthumbs.db
Desktop.ini
.DS_Store
.stignore
.stfolder
*.swp

# ignore with prefix of file name
_*

# apollo config or cache file
.*_default
```

* 项目必须带有自述文件`README.md`，并在其中说明项目的作用、使用方法、注意事项等内容；
* 需使用配置文件的项目，必须在`configs`目录下带上示例配置文件（无论是否使用 Apollo），并注释每项配置的作用；
* 严禁将编译后的二进制执行文件传至仓库，一般情况下，可在项目目录下创建一个`_bin`目录用于存放这些文件；

## 关于SQL
* SQL 中的关键字必须全部大写;
* 避免使用`SELECT *`查询所有字段，应在使用前明确自己需要哪些字段数据；
* 计算总行数应使用`SELECT COUNT(*)`而不是`SELECT COUNT(id)`等其他形式；
* 如无必要，查询中严禁使用`ORDER BY`子句；
* 严禁项目中使用ORM；

## 编码建议
* 对于一些格式化声明`error`类型的时候，应直接使用`fmt.Errorf`而不是`errors.New(fmt.Sprintf(…))`；
* 任何项目都应避免使用反射，高并发项目严禁使用反射；
* 对于需要反复声明/使用的变量，应尽量使用对象池进行复用，减轻GC负担，推荐标准库的`sync.Pool`；
* 所有的返回值（尤其是`error`类型）皆要处理，避免忽略；

## 常用包限定
**现有项目中指定功能如果不是使用下列规定包的，需尽快替换。**

* 日志输出：xlog 或 zap（一些极轻的项目可以使用标准库`log`或`fmt`）

> https://git.jiaxianghudong.com/go/xlog  
> https://github.com/uber-go/zap  