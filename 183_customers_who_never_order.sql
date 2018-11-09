-- Source: https://leetcode.com/problems/customers-who-never-order
SELECT
    `c`.`Name` AS `Customers`
FROM `Customers` as `c`
LEFT JOIN (SELECT
            `CustomerId` AS `Id`,
            COUNT(`Id`) as `Times`
        FROM `Orders`
        GROUP BY `CustomerId`) AS `f`
    ON `c`.`Id` = `f`.`Id`
WHERE `f`.`Id` IS NULL;

/*
Test Case:
{"headers": {"Customers": ["Id", "Name"], "Orders": ["Id", "CustomerId"]}, "rows": {"Customers": [[1, "Joe"], [2, "Henry"], [3, "Sam"], [4, "Max"]], "Orders": [[1, 3], [2, 1]]}}
*/
