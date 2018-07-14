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

    problems = get_problems()
    for name in names:
        question_id = int(os.path.splitext(name)[0])
        problem = problems.get(question_id)
        if not problem:
            print("didn't find question id {}".format(question_id))
            exit(1)

        slug = problem.slug
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
        "|ID|Name|Difficulty|Source Code",
        "|--|---|-----|---------|",
    ]

    problems = get_problems()

    local_problems = []
    for name in os.listdir("."):
        if not re.match("^\d{3}-[a-zA-Z\-]*?\.py$", name):
            continue

        question_id = int(name[:3])
        problem = problems.get(question_id)

        if not problem:
            print("didn't find question id {}".format(question_id))
            exit(1)

        local_problems.append(
            dict(
                id=problem.id,
                title=problem.title,
                level=problem.level,
                slug=problem.slug,
                file=name,
            ))

    local_problems = sorted(local_problems, key=lambda x: x["id"])
    link_template = \
        "[{title}](https://leetcode.com/problems/{slug}/description/)"
    overview_tables_lines += [
        ("|{id}|" + link_template + "|{level}|[{file}](./{file})|").format(**p)
        for p in local_problems
    ]
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


if __name__ == "__main__":
    commands = {
        "pre-commit": before_commit,
    }
    parser = argparse.ArgumentParser()
    parser.add_argument("command", choices=commands.keys())
    args = parser.parse_args()

    commands[args.command]()
