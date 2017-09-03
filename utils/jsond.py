#!/usr/bin/env python
# -*- coding: utf-8 -*-

import uuid
import json


data = {}
x = uuid.uuid1()


with open("./resources/manifest.json", "r") as jf:
    data = json.load(jf, encoding="utf-8")
    print type(data), data
    for _ in data:
        _["RepoTags"] = ["%s:latest" % x]

with open("./resources/manifest.json", "w") as jf:
    json.dump(data, jf, encoding="utf-8")


with open("./resources/repositories", "r") as jf:
    data = json.load(jf, encoding="utf-8")
    it = None
    for k, v in data.iteritems():
        if v.get("latest") is not None:
            it = k
    if it is not None:
        data[unicode(x)] = data.pop(it)


with open("./resources/repositories", "w") as jf:
    json.dump(data, jf, encoding="utf-8")


print x
