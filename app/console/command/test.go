package command

import (
	"fmt"
	"github.com/jjonline/golang-backend/app/model"
	"github.com/spf13/cobra"
)

// init version子命令
func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "test", // 子命令名称
		Short: "测试命令", // 子命令简短说明
		Long:  "测试命令：请在下方Run方法书写测试代码，测试代码不要提交到代码库", // 子命令完整说明
		Run: func(cmd *cobra.Command, args []string) {
			// FindByPrimary
			var ad model.Approval
			err := ad.FindByPrimary(8, &ad)
			fmt.Println(err)
			fmt.Println(ad)

			// FindByWhere
			var ad1 model.Approval
			err = ad1.FindByWhere([]model.Where{{"u8_user_id", "=", "02010001"}}, &ad1)
			fmt.Println(err)
			fmt.Println(ad1)

			// ListByWhere
			var ad2 []model.Approval
			err = ad1.ListByWhere([]model.Where{{"u8_user_id", "=", "02010001"}}, &ad2, "")
			fmt.Println(err)
			fmt.Println(ad2)

			// Paginate
			var ad3 []model.Approval
			var total int64
			err = ad1.Paginate(
				[]model.Where{{"u8_user_id", "=", "02010001"}},
				&ad3,
				&total,
				2,
				10,
				"",
				)
			fmt.Println(err)
			fmt.Println(ad3)
			fmt.Println(total)
		},
	})
}