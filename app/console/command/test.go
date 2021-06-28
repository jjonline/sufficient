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
			var test model.Test
			test.Name = ""
			test.Type = 1
			err := test.InsertOne(&test)
			fmt.Println(test)
			fmt.Println(err)

			var test1 model.Test
			test1.ID = 10
			test1.Name = "update-name-by-model"
			rows, err1 := test1.UpdateOne(&test1)
			fmt.Println(rows)
			fmt.Println(err1)

			var test2 model.Test
			rows2, err2 := test.UpdateByWhere(&test2, []model.Where{{
				Field: "id",
				Op:    ">=",
				Value: 5,
			}}, map[string]interface{}{"name": "updated Name once", "type": 3})
			fmt.Println(err2)
			fmt.Println(rows2)

			res3, err3 := test2.DeleteOne(&test2, 11)
			fmt.Println(res3)
			fmt.Println(err3)

			res4, err4 := test2.DeleteOne(&test2, "11")
			fmt.Println(res4)
			fmt.Println(err4)

			res5, err5 := test2.DeleteByWhere(&test2, []model.Where{
				{Field: "id",
					Op:    ">=",
					Value: 1},
			})
			fmt.Println(res5)
			fmt.Println(err5)
		},
	})
}
