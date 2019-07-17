/* Source: https://leetcode.com/problems/combine-two-tables */

SELECT
    `p`.`FirstName`,
    `p`.`LastName`,
    `addr`.`City`,
    `addr`.`State`
FROM `Person` AS `p`
LEFT JOIN `Address` AS `addr`
    ON `p`.`PersonId` = `addr`.`PersonId`;


/*
Test Case:
{"headers": {"Person": ["PersonId", "LastName", "FirstName"], "Address": ["AddressId", "PersonId", "City", "State"]}, "rows": {"Person": [[1, "Wang", "Allen"]], "Address": [[1, 2, "New York City", "New York"]]}}
*/
