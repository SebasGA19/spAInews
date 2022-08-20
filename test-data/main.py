# Use this script only in dev mode

import json
import datetime
import os
from pymongo import MongoClient


mongoURL = "mongodb://spainews:spainews@127.0.0.1:27017/spainews"


def load_data() -> list:
    news = []
    for dir, _, files in os.walk("data"):
        for file in files:
            abs_path = os.path.join(dir, file)
            with open(abs_path) as file_obj:
                contents = file_obj.read()
                obj = json.loads(contents)
                # 2022-08-16 17:13:00
                # date_download, date_modify, date_publish
                obj["date_download"] = datetime.datetime.strptime(obj["date_download"], '%Y-%m-%d %H:%M:%S')
                obj["date_modify"] = datetime.datetime.strptime(obj["date_modify"], '%Y-%m-%d %H:%M:%S')
                obj["date_publish"] = datetime.datetime.strptime(obj["date_publish"], '%Y-%m-%d %H:%M:%S')
                news.append(obj)
    return news


def main():
    news = load_data()
    client = MongoClient(mongoURL)
    db = client["spainews"]
    collection = db["news"]
    collection.insert_many(news)


if __name__ == "__main__":
    main()
