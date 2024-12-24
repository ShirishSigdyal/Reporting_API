-- Sales report query
SELECT 
    p.category, 
    SUM(oi.price * oi.quantity) AS total_sales
FROM 
    order_items AS oi
INNER JOIN 
    products AS p ON oi.product_id = p.id
INNER JOIN 
    orders AS o ON oi.order_id = o.id
WHERE 
    o.order_date >= $1 AND o.order_date <= $2
GROUP BY 
    p.category;

-- Customer behavior report query
SELECT 
    COUNT(*) AS total_customers,
    AVG(o_count) AS average_orders,
    CASE 
        WHEN lifetime_value < 100 THEN 'Low'
        WHEN lifetime_value BETWEEN 100 AND 500 THEN 'Medium'
        ELSE 'High'
    END AS segment
FROM 
    (SELECT 
        c.id, c.lifetime_value, COUNT(o.id) AS o_count
     FROM 
        customers AS c
     LEFT JOIN 
        orders AS o ON c.id = o.customer_id
     WHERE 
        c.signup_date >= $1 AND c.signup_date <= $2
     GROUP BY 
        c.id) AS subquery
GROUP BY 
    segment;
