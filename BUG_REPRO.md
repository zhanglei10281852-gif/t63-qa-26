# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

维修单完工后订单状态已完成，但车辆仍显示维修中，无法再次排班。请修复完工流程的车辆状态恢复。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/t63-qa-26
- 仓库地址：https://github.com/zhanglei10281852-gif/t63-qa-26.git
- parent SHA：88967c3954d0c6d62fa71486591b404ecc8812ab

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/t63-qa-26.git bug-repro
cd bug-repro
git checkout --detach 88967c3954d0c6d62fa71486591b404ecc8812ab
go test ./internal/httpapi -run TestCompletedServiceReturnsVehicleToDispatchPool -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestCompletedServiceReturnsVehicleToDispatchPool -count=1
--- FAIL: TestCompletedServiceReturnsVehicleToDispatchPool (0.69s)
    maintenance_release_test.go:20: vehicle status=maintenance
FAIL
FAIL	sanitation-operations/internal/httpapi	0.699s
FAIL

```

stderr：

```text
warning: internal/httpapi/maintenance_release_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644
warning: internal/httpapi/maintenance_release_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644

```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestCompletedServiceReturnsVehicleToDispatchPool -count=1
--- FAIL: TestCompletedServiceReturnsVehicleToDispatchPool (1.25s)
    maintenance_release_test.go:20: vehicle status=maintenance
FAIL
FAIL	sanitation-operations/internal/httpapi	1.463s
FAIL

```

stderr：

```text
warning: internal/httpapi/maintenance_release_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644
warning: internal/httpapi/maintenance_release_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644

```

## 通过条件

在触发条件下，定向测试 TestCompletedServiceReturnsVehicleToDispatchPool 应通过，相关包、全量测试、竞态测试和构建检查均通过；回退 gold 唯一修复后定向测试重新失败。
