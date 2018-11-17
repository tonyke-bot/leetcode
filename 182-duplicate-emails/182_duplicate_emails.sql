-- Source: https://leetcode.com/problems/duplicate-emails
SELECT `Email`
FROM `Person`
GROUP BY `Email`
HAVING COUNT(`Email`) > 1;

/*
Test Case:
{"headers": {"Person": ["Id", "Email"]}, "rows": {"Person": [[1, "a@b.com"], [2, "c@d.com"], [3, "a@b.com"]]}}
*/
