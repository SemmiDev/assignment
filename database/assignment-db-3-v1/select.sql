-- TODO: answer here

SELECT
    id,
    CONCAT(first_name, ' ', last_name) as fullname,
    split_part(exam_id, '-', 1) AS class,
    (bahasa_indonesia + bahasa_inggris + matematika + ipa) / 4 AS average_score
FROM final_scores
WHERE
    exam_status = 'pass' AND (fee_status = 'full' OR fee_status = 'installment')
ORDER BY average_score DESC
LIMIT 5;
