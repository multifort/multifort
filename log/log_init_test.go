package log

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//fmt.Println(viper.Get("db.default"))
	fmt.Println("test")
	t := &Test{
		Name: "multifort",
		Age:  12,
	}

	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("marshal is failed, err: %S", err.Error())
	}

	logger := InitLogger("./all.log", "info")

	for i := 0; i < 6; i++ {
		logger.Info(fmt.Sprintf("test log %c", i), zap.Int("line", 47))
		logger.Debug(fmt.Sprintf("debug log %c", i), zap.ByteString("level", data))
		logger.Info(fmt.Sprintf("Info log %c", i), zap.String("level", `{"a":"4","b":"5"}`))
		logger.Debug(fmt.Sprintf("debug log %c", i), zap.String("level", `{"a":"7","b":"8"}`))
	}

}
