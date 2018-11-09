-- Source: https://leetcode.com/problems/employees-earning-more-than-their-managers
SELECT
    `e`.`Name` as `Employee`
FROM `Employee` AS `e`
LEFT JOIN `Employee` AS `m`
    ON `e`.`ManagerId` = `m`.`Id`
WHERE `e`.`Salary` > `m`.`Salary`;

/*
Test Case:
{"headers": {"Employee": ["Id", "Name", "Salary", "ManagerId"]}, "rows": {"Employee": [[1, "Joe", 70000, 3], [2, "Henry", 80000, 4], [3, "Sam", 60000, null], [4, "Max", 90000, null]]}}
*/
