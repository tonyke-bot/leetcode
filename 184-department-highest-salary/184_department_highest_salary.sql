-- Source: https://leetcode.com/problems/department-highest-salary

-- Table: Employee
-- | Id | Name  | Salary | DepartmentId |
-- | -- | ----- | ------ | ------------ |
-- | 1  | Joe   | 70000  | 1            |
-- | 2  | Henry | 80000  | 2            |
-- | 3  | Sam   | 60000  | 2            |
-- | 4  | Max   | 90000  | 1            |

-- Table: Department
-- | Id | Name  |
-- | -- | ----- |
-- | 1  | IT    |
-- | 2  | Sales |

SELECT
    `d`.`Name` AS `Department`,
    `e`.`Name` AS `Employee`,
    `e`.`Salary` AS `Salary`
FROM (
    SELECT `DepartmentId`, MAX(`Salary`) as `Salary`
    FROM `Employee`
    GROUP BY `DepartmentId`
) AS `m`
INNER JOIN `Department` as `d`
    ON `d`.`Id` = `m`.`DepartmentId`
LEFT JOIN `Employee` as `e`
    ON `e`.`Salary` = `m`.`Salary` AND `m`.`DepartmentId` = `e`.`DepartmentId`;
