#!/usr/bin/env python3
"""
1. Rename filename
2. Fill README.md with problem info
"""
import argparse
import os
import os.path
import re

from utils.leetcode import get_problems

FORCE_UPDATE_PROBLEMS = False


def get_unnamed_file(directory):
    names = []
    for name in os.listdir(directory):
        if not name.endswith(".py"):
            continue

        if re.match("^\d*?\.py$", name):
            names.append(name)
    return names


def rename_filenames(directory):
    names = get_unnamed_file(directory)
    if not len(names):
        print("  no files need to be renamed")
        return

    problems = get_problems(FORCE_UPDATE_PROBLEMS)
    for name in names:
        question_id = int(os.path.splitext(name)[0])
        problem = problems.get(str(question_id))
        if not problem:
            print("didn't find question id {}".format(question_id))
            exit(1)

        slug = problem["slug"]
        new_filename = "{:03d}-{}.py".format(question_id, slug)
        print("  {} => {}".format(name, new_filename))
        os.renames(
            os.path.join(directory, name), os.path.join(
                directory, new_filename))


def update_readme():
    overview_start_tag = "<!-- OVERVIEW START -->"
    overview_end_tag = "<!-- OVERVIEW END -->"
    overview_tables_lines = [
        overview_start_tag,
        "#|Name|Difficulty|Solution|Tags",
        "-|----|----------|--------|----",
    ]

    problems = get_problems(FORCE_UPDATE_PROBLEMS)

    lines = []
    for name in os.listdir("."):
        if not re.match("^\d{3}-[a-zA-Z\-]*?\.py$", name):
            continue

        question_id = int(name[:3])
        problem = problems.get(str(question_id))

        if not problem:
            print("didn't find question id {}".format(question_id))
            exit(1)

        lines.append([
            problem["id"],
            "[{}](https://leetcode.com/problems/{}/description/)".format(
                problem["title"], problem["slug"]),
            problem["level"],
            "[{0}](./{0})".format(name),
            " ".join(["`" + tag + "`" for tag in problem["tags"]]),
        ])

    lines = sorted(lines, key=lambda x: int(x[0]))
    overview_tables_lines += ["|".join(line) for line in lines]
    overview_tables_lines.append(overview_end_tag)

    readme = open("./README.md", encoding="utf8").read()

    start = readme.find(overview_start_tag)
    end = readme.find(overview_end_tag, start)

    old = readme[start:end + len(overview_end_tag)]
    new = "\n".join(overview_tables_lines)

    open("./README.md", "w", encoding="utf8").write(readme.replace(old, new))


def before_commit():
    print("rename")
    rename_filenames(".")

    print("update README")
    update_readme()


def new_problem():
    problems = get_problems(FORCE_UPDATE_PROBLEMS)
    question_id = input('question id: ')

    problem = problems.get(question_id)
    if not problem:
        print("didn't find question id {}".format(question_id))
        exit(1)

    solution = problem["code"] or ""
    test_case = (problem["test_case"] or "").split("\n")
    print(test_case)

    problem_id = int(problem["id"])
    problem_slug = problem["slug"]

    filename = "{:03d}-{}.py".format(problem_id, problem_slug)
    content = \
        "\n\n" + \
        solution + \
        "\n\n" + \
        "if __name__ == '__main__':\n" + \
        "    solver = Solution()\n" + \
        "    # Test Case:\n" + \
        "\n".join(["    # " + line for line in test_case])

    open(filename, "w").write(content)


if __name__ == "__main__":
    commands = {
        "pre-commit": before_commit,
        "new": new_problem,
    }
    parser = argparse.ArgumentParser()
    parser.add_argument("command", choices=commands.keys())
    parser.add_argument(
        "-f",
        "--force-update",
        action="store_true",
        help="force update problems")
    args = parser.parse_args()

    commands[args.command]()
