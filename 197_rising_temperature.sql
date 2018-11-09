-- Source: https://leetcode.com/problems/rising-temperature

-- Table: Weather
-- | Id | RecordDate | Temperature |
-- | -- | ---------- | ----------- |
-- | 1  | 2015-01-01 | 10          |
-- | 2  | 2015-01-02 | 25          |
-- | 3  | 2015-01-03 | 20          |
-- | 4  | 2015-01-04 | 30          |
SELECT `b`.`Id`
FROM `Weather` AS `a`
INNER JOIN `Weather` AS `b`
    ON `a`.`Temperature` < `b`.`Temperature`
        AND DATE_ADD(`a`.`RecordDate`, INTERVAL 1 DAY) = `b`.`RecordDate`;
