package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
var workCode int64 = 0

func getWorkId() int64 {
	return workCode % 1
}

func Init() (err error) {
	node, err = snowflake.NewNode(1)
	return
}

func GenIDInt() int64 {
	node, _ = snowflake.NewNode(getWorkId())
	return node.Generate().Int64()
}

func GenIDString() string {
	node, _ = snowflake.NewNode(1)
	id := node.Generate()
	return id.String()
}

func main() {
	if err := Init(); err != nil {
		fmt.Printf("全局id生成器初始化失败,%v\n", err)
		return
	}
	var userId = GenIDInt()
	fmt.Print(userId)

}
