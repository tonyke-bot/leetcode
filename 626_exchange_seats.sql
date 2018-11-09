-- Source: https://leetcode.com/problems/exchange-seats

-- Table: seat
-- | id | student |
-- | -- | ------- |
-- | 1  | Abbot   |
-- | 2  | Doris   |
-- | 3  | Emerson |
-- | 4  | Green   |
-- | 5  | Jeames  |

SELECT
    `a`.`id` AS `id`,
    IF(`a`.`id`%2 = 1, IF(`c`.`id` IS NULL, `a`.`student`, `c`.`student`), `b`.`student`) AS `student`
FROM `seat` AS `a`
LEFT JOIN `seat` AS `b` ON `a`.`id`%2 = 0 AND `a`.`id`-1 = `b`.`id`
LEFT JOIN `seat` AS `c` ON `a`.`id`%2 = 1 AND `a`.`id`+1 = `c`.`id`
ORDER BY `a`.`id`;
