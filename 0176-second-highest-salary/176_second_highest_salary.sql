/* Source: https://leetcode.com/problems/second-highest-salary */
-- SELECT
--     `Salary` as `SecondHighestSalary`
-- FROM `Employee`
-- ORDER BY `Salary` DESC
-- LIMIT 1, 1;
SELECT IFNULL((
    SELECT DISTINCT `Salary`
    FROM `Employee`
    ORDER BY `Salary` DESC
    LIMIT 1, 1), NULL
) AS `SecondHighestSalary`

/*
Test Case:
{"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100], [2, 200], [3, 300]]}}
*/
