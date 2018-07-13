import urllib.request
import json
from functools import lru_cache


class Problem:
    def __init__(self, id, title, slug, level):
        self.id = id
        self.title = title
        self.slug = slug
        self.level = level

    def __repr__(self):
        return "Problem {}: {} ({})".format(
            self.id, self.title, self.level)


def parse_difficulty_level(level):
    return {
        1: "Easy",
        2: "Medium",
        3: "Hard"
    }[level]


@lru_cache()
def get_problems():
    url = "https://leetcode.com/api/problems/all/"
    req = urllib.request.Request(url, headers={"User-Agent": "LeetCode/1.0.0"})
    f = urllib.request.urlopen(req)

    data = json.loads(f.read().decode("utf-8"))
    problems = dict()

    for problem in data["stat_status_pairs"]:
        stat = problem["stat"]
        id = stat["question_id"]
        problems[id] = Problem(
            id=stat["question_id"],
            title=stat["question__title"],
            slug=stat["question__title_slug"],
            level=parse_difficulty_level(problem["difficulty"]["level"]),
        )

    return problems


if __name__ == "__main__":
    problems = get_problems()
    print(problems[1])
