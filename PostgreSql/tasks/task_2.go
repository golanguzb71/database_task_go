package tasks

import (
	"database/sql"
	"fmt"
)

type ReportResult struct {
	UserName              string
	TotalDiscountReceived float64
	TotalAmountOrdered    float64
	GoodsReceived         string
	BranchName            string
}

func ReportQuery(db *sql.DB) ([]ReportResult, error) {
	query := `
        SELECT
            u.name AS user_name,
            SUM(p.income_price - p.outcome_price) AS total_discount_received,
            SUM(p.outcome_price) AS total_amount_ordered,
            STRING_AGG(p.name, ', ') AS goods_received,
            b.name AS branch_name
        FROM
            orders o
        JOIN
            users u ON o.user_id = u.id
        JOIN
            order_items oi ON o.id = oi.order_id
        JOIN
            product p ON oi.product_id = p.id
        JOIN
            branch b ON o.branch_id = b.id
        GROUP BY
            u.id, b.name;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %v", err)
	}
	defer rows.Close()

	var results []ReportResult
	for rows.Next() {
		var result ReportResult
		if err := rows.Scan(&result.UserName, &result.TotalDiscountReceived, &result.TotalAmountOrdered, &result.GoodsReceived, &result.BranchName); err != nil {
			return nil, fmt.Errorf("row scan error: %v", err)
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return results, nil
}
