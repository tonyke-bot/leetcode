-- Source: https://leetcode.com/problems/big-countries

-- Table: World
-- | name        | continent | area         | population    | gdp         |
-- | ----------- | --------- | ------------ | ------------- | ----------- |
-- | Afghanistan | Asia      | 652230       | 2.55001e+07   | 2.0343e+10  |
-- | Albania     | Europe    | 28748        | 2.831741e+06  | 1.296e+10   |
-- | Algeria     | Africa    | 2.381741e+06 | 3.71e+07      | 1.88681e+11 |
-- | Andorra     | Europe    | 468          | 78115         | 3.712e+09   |
-- | Angola      | Africa    | 1.2467e+06   | 2.0609294e+07 | 1.0099e+11  |

SELECT
    `name`,
    `population`,
    `area`
FROM
    `World`
WHERE
    `area` > 3000000
    OR `population` > 25000000;
