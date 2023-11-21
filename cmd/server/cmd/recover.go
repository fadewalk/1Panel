package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(recoverCmd)
}

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "恢复面板正常运行",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !isRoot() {
			fmt.Println("请使用 sudo 1pctl recover 或者切换到 root 用户")
			return nil
		}
		db, err := loadDBConn()
		if err != nil {
			return fmt.Errorf("init my db conn failed, err: %v \n", err)
		}
		if err := db.Exec("PRAGMA journal_mode = DELETE;").Error; err != nil {
			return err
		}
		if err := db.Exec("PRAGMA wal_checkpoint(FULL);").Error; err != nil {
			return err
		}

		fmt.Println("面板恢复成功！")
		return nil
	},
}
