package command

import (
	"fmt"
	"github.com/jjonline/golang-backend/app/model"
	"github.com/spf13/cobra"
)

// init version子命令
func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "test",                              // 子命令名称
		Short: "测试命令",                              // 子命令简短说明
		Long:  "测试命令：请在下方Run方法书写测试代码，测试代码不要提交到代码库", // 子命令完整说明
		Run: func(cmd *cobra.Command, args []string) {
			var data model.Test
			err := model.TestModel.FindByPrimary(1, &data);
			fmt.Println(err)
			fmt.Println(data)

			var one model.Test
			one.Name = "j"
			one.Type = 1
			err1 := model.TestModel.InsertOne(&one)
			fmt.Println(err1)
			fmt.Printf("%#v", one)
		},
	})
}
