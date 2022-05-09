## 错误码

### 示例代码
```go
package errno

import "github.com/bitrainforest/nuwas/errno"

var (
	OK                     = errno.NewError(0, "操作成功")
	ParamValidationErr     = errno.NewError(100, "参数验证不正确")
	RecordNotFound         = errno.NewError(101, "记录未找到")
	RecordNotFoundWithMore = errno.NewError(101, "记录未找到，失败原因 %s")
)

```