package gen

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// 代码生成工具:直接修改原始Go文件,在每个函数开头添加打印方法名的代码
func main() {
	// 接收命令行参数:需要修改的业务文件路径
	var srcFile string
	var overwrite bool
	flag.StringVar(&srcFile, "src", "", "需要修改的业务文件路径(如./business/user.go)")
	flag.BoolVar(&overwrite, "overwrite", false, "是否覆盖已存在的文件")
	flag.Parse()
	if srcFile == "" {
		fmt.Fprintln(os.Stderr, "错误: 请指定-src参数,如:go run loggen.go -src=./business/user.go")
		os.Exit(1)
	}

	// 1. 读取原始文件内容
	_, err := os.ReadFile(srcFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: 读取文件失败:%v\n", err)
		os.Exit(1)
	}

	// 2. 解析业务文件的AST语法树
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, srcFile, nil, parser.AllErrors)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: 解析AST失败:%v\n", err)
		os.Exit(1)
	}

	// 3. 检查是否需要添加fmt导入
	needsFmtImport := !hasFmtImport(node)

	// 4. 遍历AST中的所有函数,修改函数体
	modified := false
	ast.Inspect(node, func(n ast.Node) bool {
		// 匹配函数定义(排除方法、结构体等)
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok || funcDecl.Recv != nil { // Recv!=nil是结构体方法,可按需扩展
			return true
		}

		// 检查函数体是否为空
		if funcDecl.Body == nil || len(funcDecl.Body.List) == 0 {
			return true
		}

		// 检查函数体第一行是否已经是打印方法名的代码
		if isLogStatement(funcDecl.Body.List[0]) {
			return true
		}

		// 在函数体开头插入打印方法名的代码
		funcName := funcDecl.Name.Name
		logStmt := &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "fmt"},
					Sel: &ast.Ident{Name: "Printf"},
				},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: "[LOG] 方法执行:%s\\n"},
					&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", funcName)},
				},
			},
		}

		// 将日志语句插入到函数体开头
		funcDecl.Body.List = append([]ast.Stmt{logStmt}, funcDecl.Body.List...)
		modified = true
		return true
	})

	// 5. 如果需要,添加fmt导入
	if needsFmtImport && modified {
		if node.Imports == nil {
			node.Imports = []*ast.ImportSpec{}
		}
		node.Imports = append(node.Imports, &ast.ImportSpec{
			Path: &ast.BasicLit{Kind: token.STRING, Value: `"fmt"`},
		})
	}

	// 6. 如果没有修改,直接退出
	if !modified {
		fmt.Printf("文件 %s 无需修改,所有函数已包含日志语句\n", srcFile)
		os.Exit(0)
	}

	// 7. 格式化修改后的代码
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		fmt.Fprintf(os.Stderr, "错误: 代码格式化失败:%v\n", err)
		os.Exit(1)
	}

	// 8. 写回修改后的代码到原始文件
	if _, err := os.Stat(srcFile); err == nil && !overwrite {
		fmt.Fprintf(os.Stderr, "错误: 文件 %s 已存在,使用 -overwrite 标志覆盖\n", srcFile)
		os.Exit(1)
	}
	if err := os.WriteFile(srcFile, buf.Bytes(), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "错误: 写入文件失败:%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("成功修改文件:%s,已在所有函数开头添加日志语句\n", srcFile)
}

// 以下为辅助函数

// hasFmtImport 检查文件是否已经导入了 fmt 包
func hasFmtImport(node *ast.File) bool {
	for _, imp := range node.Imports {
		if imp.Path.Value == `"fmt"` {
			return true
		}
	}
	return false
}

// isLogStatement 检查语句是否已经是打印方法名的日志语句
func isLogStatement(stmt ast.Stmt) bool {
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return false
	}

	callExpr, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return false
	}

	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	// 检查是否是 fmt.Printf 调用
	if selExpr.X.(*ast.Ident).Name != "fmt" || selExpr.Sel.Name != "Printf" {
		return false
	}

	// 检查参数数量
	if len(callExpr.Args) != 2 {
		return false
	}

	// 检查第一个参数是否是日志格式字符串
	formatStr, ok := callExpr.Args[0].(*ast.BasicLit)
	if !ok || formatStr.Kind != token.STRING {
		return false
	}

	return strings.Contains(formatStr.Value, "[LOG] 方法执行:")
}
