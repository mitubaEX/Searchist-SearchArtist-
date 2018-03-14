import random
from flask import Flask
app = Flask(__name__)


@app.route('/')
def hello():
    f = open("./result.txt").read().split("\n")
    random.shuffle(f)
    return f[0]


if __name__ == "__main__":
    app.run(debug=True)
