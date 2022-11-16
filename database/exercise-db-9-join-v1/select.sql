-- TODO: answer here

SELECT
    r.id,
    s.fullname,
    s.class,
    s.status,
    r.study,
    r.score
FROM
    students s
INNER JOIN
    reports r
ON
    s.id = r.student_id
WHERE
    s.status = 'active' AND r.score < 70
ORDER BY
    r.score ASC;
