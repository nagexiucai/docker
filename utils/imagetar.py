#!/usr/bin/env python
# -*- coding: utf-8 -*-


import tarfile

tf = tarfile.open("./resources/f5d1f028-5031-4405-a7b5-48243811907b.tar")
tf.extract(tf.getmember("manifest.json"), path="./resources")
tf.extract(tf.getmember("repositories"), path="./resources")
tf.close()


import jsond


tf = tarfile.open("./resources/f5d1f028-5031-4405-a7b5-48243811907b.tar", "a")

# with open("./resources/manifest.json", mode="rb") as jf:
#     tf.addfile(tf.gettarinfo(fileobj=jf), jf)
tf.add(name="./resources/manifest.json", arcname="manifest.json")
tf.add(name="./resources/repositories", arcname="repositories")

tf.close()

tf = tarfile.open("./resources/f5d1f028-5031-4405-a7b5-48243811907b.tar")
tf.extractall(path="./resources/f5d1f028-5031-4405-a7b5-48243811907b")
tf.close()
