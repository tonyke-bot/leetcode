-- Source: https://leetcode.com/problems/nth-highest-salary
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    DECLARE M INT;
    SET M = N-1;
    RETURN IFNULL((
        SELECT DISTINCT `Salary`
        FROM `Employee`
        ORDER BY `Salary` DESC
        LIMIT M, 1), NULL
    );
END

/*
Test Case:
{"headers": {"Employee": ["Id", "Salary"]}, "argument": 2, "rows": {"Employee": [[1, 100], [2, 200], [3, 300]]}}
*/
