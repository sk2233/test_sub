package main

import (
	"fmt"
)

/*
# 格式：git submodule add <共享仓库远程地址> <本地子目录路径>
git submodule add https://github.com/yourusername/shared-files.git shared

# 初始化子模块配置（读取 .gitmodules 并更新主仓库的 .git/config）
git submodule init

# 拉取子模块的文件内容（根据配置检出对应的 commit）
git submodule update

# 设置自动拉取
git config --local submodule.recurse true
*/

func main() {
	fmt.Println("日志 1: 程序开始执行")
	fmt.Println("日志 2: 初始化完成")
	fmt.Println("日志 3: 处理业务逻辑")
	fmt.Println("日志 4: 数据验证通过")
	fmt.Println("日志 5: 执行核心操作")
	fmt.Println("日志 6: 中间处理步骤")
	fmt.Println("日志 7: 继续处理中")
	fmt.Println("日志 8: 接近完成")
	fmt.Println("日志 9: 最后检查")
	fmt.Println("日志 10: 程序执行完成")
}
