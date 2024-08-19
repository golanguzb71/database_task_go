package tasks

import "database/sql"

/*
Create a query of total income and total expenses for each branch.
- total income = sum(product.outcome_price) - sum(order.discount_price) - sum(product.income_price)
- total_expenses = sum(order.discount_price)
*/

type TotalIncomeExpense struct {
	BranchId      int64
	Total_expense float64
	Total_income  float64
}

func TotalIncomeExpenseFunc(db *sql.DB) []*TotalIncomeExpense {
	var sql = `
SELECT b.id,
       (SUM(p.outcome_price) - SUM(o.discount_price) - sum(p.income_price)) AS total_income,
       sum(o.discount_price)                                                AS total_expenses
FROM branch b
         join orders o on b.id = o.branch_id
         join order_items oi on o.id = oi.order_id
         join product p on oi.product_id = p.id
GROUP BY b.id
`
	rows, err := db.Query(sql)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var result []*TotalIncomeExpense
	for rows.Next() {
		var TIE TotalIncomeExpense
		err = rows.Scan(&TIE.BranchId, &TIE.Total_income, &TIE.Total_expense)
		if err != nil {
			return nil
		}
		result = append(result, &TIE)
	}
	return result
}
