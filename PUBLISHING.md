# 发布指南

这个文档说明了如何将这个增强版的gdt-http库发布到GitHub和pkg.go.dev。

## 前提条件

1. 你需要有一个GitHub组织账户或用户账户，可以创建名为 `gdt` 的组织
2. 确保你的GitHub组织/用户名与go.mod中的包名匹配

## 发布步骤

### 1. 创建GitHub仓库

#### 选项1：如果你有 `gdt` 组织
1. 登录到 https://github.com
2. 切换到 `gdt` 组织
3. 点击右上角的 "+" 按钮，选择 "New repository"
4. 仓库名填写：`http`
5. 设置为 Public（公开仓库，这样pkg.go.dev才能访问）
6. 不要初始化README（因为我们已经有了）
7. 点击"Create repository"

#### 选项2：如果你需要创建 `gdt` 组织
1. 访问 https://github.com/organizations/new
2. 创建名为 `gdt` 的组织
3. 然后按照选项1的步骤创建仓库

### 2. 推送代码到GitHub

```bash
# 如果还没有初始化git仓库
git init

# 添加远程仓库
git remote add origin https://github.com/doingdd/http.git

# 添加所有文件
git add .

# 提交
git commit -m "Initial commit with custom headers support"

# 推送到GitHub
git push -u origin main
```

### 3. 创建版本标签

pkg.go.dev使用Git标签来识别版本。创建一个初始版本：

```bash
# 创建v1.0.0标签
git tag v1.0.0

# 推送标签
git push origin v1.0.0
```

### 4. 验证pkg.go.dev上的包

发布后，你可以通过以下方式验证：

1. 访问：https://pkg.go.dev/github.com/doingdd/http
2. 如果没有立即显示，可以手动触发：
   - 访问：https://proxy.golang.org/github.com/doingdd/http/@v/list
   - 或者访问：https://sum.golang.org/lookup/github.com/doingdd/http@v1.0.0

### 5. 其他用户如何使用你的包

其他开发者可以通过以下方式使用你的增强版gdt-http：

```bash
go get github.com/doingdd/http
```

在他们的代码中：

```go
import "github.com/doingdd/http"
```

## 重要提醒

1. **包名注意事项**：`github.com/doingdd/http` 这个名字很简洁，但需要你有对 `gdt` 这个GitHub组织的控制权
2. **版本管理**：建议使用语义化版本 (semantic versioning)
3. **文档**：确保README.md中包含清晰的安装和使用说明
4. **许可证**：确保包含适当的开源许可证

## 示例使用

用户可以在他们的测试文件中这样使用：

```yaml
# test.yaml
tests:
  - name: test with custom headers
    GET: /api/users
    headers:
      Authorization: Bearer token123
      Content-Type: application/json
    assert:
      status: 200
```

## 更新版本

当你有新的更改时：

```bash
# 提交更改
git add .
git commit -m "Add new features"
git push

# 创建新版本标签
git tag v1.1.0
git push origin v1.1.0
```

## 注意事项

1. 包名必须是公开的GitHub仓库
2. 模块路径必须与GitHub仓库路径匹配
3. pkg.go.dev会自动索引公开的Go模块
4. 首次发布可能需要几分钟才能在pkg.go.dev上显示 