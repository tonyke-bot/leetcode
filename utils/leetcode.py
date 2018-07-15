import gzip
import json
import urllib.request
import pickle


def graphql(query):
    csrf = "61DlaZIPNSvNIKWqSv0tvt9NF0bF7wG4dS98BQE5zgp4sbu00VZzmenGq0vUcWJY"
    headers = {
        "accept-encoding": "gzip",
        "content-type": "application/json",
        "referer": "https://leetcode.com/problems/",
        "user-agent": "LeetCode/1.0.0",
        "x-csrftoken": csrf,
        "cookie": "csrftoken=" + csrf,
    }
    url = "https://leetcode.com/graphql"
    post_data = json.dumps(dict(query=query)).encode("utf8")
    req = urllib.request.Request(url, data=post_data, headers=headers)
    req.method = "POST"
    furl = urllib.request.urlopen(req)

    gzipper = gzip.GzipFile(fileobj=furl)

    result = json.loads(gzipper.read().decode("utf8"))["data"]
    return result


def get_problems(update=False):
    cache_path = "./problems.cache"
    if not update:
        try:
            return pickle.load(open(cache_path, "rb"))
        except Exception as exp:
            print("loading problems cache fail:", exp)
            pass

    result = graphql("query {\n"
                     "    questions: allQuestions {\n"
                     "        id: questionId\n"
                     "        title: questionTitle\n"
                     "        slug: questionTitleSlug\n"
                     "        level: difficulty\n"
                     "        tags: topicTags { name }\n"
                     "        code: codeDefinition\n"
                     "        test_case: sampleTestCase\n"
                     "    }\n"
                     "}")["questions"]

    for p in result:
        # flatten tags
        p["tags"] = [t["namename"] for t in p["tags"]]

        # solution template
        codes = json.loads(p["code"])
        p["code"] = None
        for code in codes:
            if code["value"] == "python3":
                p["code"] = code["defaultCode"]
                break

    problems = dict([(problem["id"], problem) for problem in result])
    pickle.dump(problems, open(cache_path, "wb"))
    return problems


if __name__ == "__main__":
    print(get_problems(True))
