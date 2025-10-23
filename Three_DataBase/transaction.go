package main

import (
	"fmt"
	"go-base-practice/utils"
	"gorm.io/gorm"
)

// 事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
//  在事务中，需要先检查账户 A 的余额是否足够，
//    如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
//      并在 transactions 表中记录该笔转账信息。
//          如果余额不足，则回滚事务。

type Account struct {
	gorm.Model
	Name    string
	Balance float64
}

type Transaction struct {
	gorm.Model
	FormAccountId uint
	ToAccountId   uint
	Amount        float64
}

func main() {
	err := utils.DB().AutoMigrate(&Account{}, &Transaction{})
	if err != nil {
		return
	}
	// 创建账户A,B
	createAccount()

	// 开启事务
	tx := utils.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 模拟转账
	var accountA Account
	var accountB Account
	tx.Model(&Account{}).Debug().Where("Name = ? and Balance >= ?", "A", 100).Scan(&accountA)

	if accountA.Balance < 100 {
		// 回滚
		tx.Rollback()
		fmt.Print("账户A余额不足,终止交易!")
	}

	tx.Model(&Account{}).Where("Name = ?", "A").Update("Balance", gorm.Expr("Balance - ?", 100))
	tx.Model(&Account{}).Where("Name = ?", "B").Update("Balance", gorm.Expr("Balance + ?", 100))
	tx.Model(&Account{}).Debug().Where("Name = ? ", "B").Scan(&accountB)

	transaction := Transaction{
		FormAccountId: accountA.ID,
		ToAccountId:   accountB.ID,
		Amount:        100,
	}
	tx.Create(&transaction)

	// 提交事务
	tx.Commit()
}

func createAccount() {
	accountArr := []Account{
		{
			Name:    "A",
			Balance: 10,
		},
		{
			Name:    "B",
			Balance: 100,
		},
	}
	utils.DB().Create(&accountArr)
}
