# runlength = lambda t: t[0] * t[1]
# wholeline = lambda l: "".join(map(runlength, l))
# wholebanner = lambda b: "\n".join(map(wholeline, b))
# import pickle
# print wholebanner(pickle.load(file("level5_banner.p")))

import pickle
import json

with open("level5_banner.p") as in_file:
    with open("level5_banner.json", "w") as out_file:
        json.dump(pickle.load(in_file), out_file)
