-- Source: https://leetcode.com/problems/trips-and-users

-- Table: Trips
-- | Id | Client_Id | Driver_Id | City_Id | Status              | Request_at |
-- | -- | --------- | --------- | ------- | ------------------- | ---------- |
-- | 1  | 1         | 10        | 1       | completed           | 2013-10-01 |
-- | 2  | 2         | 11        | 1       | cancelled_by_driver | 2013-10-01 |
-- | 3  | 3         | 12        | 6       | completed           | 2013-10-01 |
-- | 4  | 4         | 13        | 6       | cancelled_by_client | 2013-10-01 |
-- | 5  | 1         | 10        | 1       | completed           | 2013-10-02 |
-- | 6  | 2         | 11        | 6       | completed           | 2013-10-02 |
-- | 7  | 3         | 12        | 6       | completed           | 2013-10-02 |
-- | 8  | 2         | 12        | 12      | completed           | 2013-10-03 |
-- | 9  | 3         | 10        | 12      | completed           | 2013-10-03 |
-- | 10 | 4         | 13        | 12      | cancelled_by_driver | 2013-10-03 |

-- Table: Users
-- | Users_Id | Banned | Role   |
-- | -------- | ------ | ------ |
-- | 1        | No     | client |
-- | 2        | Yes    | client |
-- | 3        | No     | client |
-- | 4        | No     | client |
-- | 10       | No     | driver |
-- | 11       | No     | driver |
-- | 12       | No     | driver |
-- | 13       | No     | driver |

SELECT
    `t`.`Request_at` AS `Day`,
    ROUND(COUNT(IF(`Status` <> 'completed', `t`.`Id`, NULL))/COUNT(`t`.`Id`), 2) AS `Cancellation Rate`
FROM `Trips` AS `t`
RIGHT JOIN `Users` AS `c` ON `c`.`Banned` <> 'Yes' AND `t`.`Client_Id` = `c`.`Users_Id`
RIGHT JOIN `Users` AS `d` ON `d`.`Banned` <> 'Yes' AND `t`.`Driver_Id` = `d`.`Users_Id`
WHERE `t`.`Request_at` BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY `t`.`Request_at`
ORDER BY `t`.`Request_at`;
