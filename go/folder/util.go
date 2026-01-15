// 用于定义私有方法和公有方法的示例
package folder

import "fmt"

func PublicFunc() {
	fmt.Println("This is a public function")
	privateFunc()
}

func privateFunc() {
	fmt.Println("This is a private function")
}
