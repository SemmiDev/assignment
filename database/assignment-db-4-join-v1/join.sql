-- TODO: answer here

-- r.quantity * r.quantity
SELECT
    r.id AS order_id,
    u.fullname,
    u.email,
    r.product_name,
    r.unit_price,
    r.quantity,
    r.order_date
FROM
    users u
INNER JOIN
    orders r
ON
    u.id = r.user_id
WHERE
    u.status = 'active' AND
    (r.unit_price * r.quantity > 500000 OR
    r.quantity > 20);
