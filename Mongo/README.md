# Initializing the database
You can get sample data from "example_data"
Create 1 million rows for "orders" and "order_products"
Create appropriate relations for tables


Requirements:
1. Insert 100 TPS
2. Query 100 ms
3. Write benchmark tests for each function

Tasks:
1. Inserting data from file should work max=1s
2. Make a report query on the 
    - total discount received by each user, 
    - the total amount of the order, 
    - which goods they received, 
    - from which branch 
they received them 

3. Create a query of total income and total expenses for each branch.
- total income = sum(product.outcome_price) - sum(order.discount_price) - sum(product.income_price)
- total_expenses = sum(order.discount_price)
4. Create API for each function