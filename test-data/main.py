# Use this script only in dev mode

import json
import datetime
import os
import random
from pymongo import MongoClient


mongoURL = "mongodb://spainews:spainews@mongo:27017/spainews"

categories = ["sports", "finance", "politics"]

def load_data() -> list:
    news = []
    for dir, _, files in os.walk("data"):
        for file in files:
            abs_path = os.path.join(dir, file)
            with open(abs_path, "rb") as file_obj:
                contents = file_obj.read()
                obj = json.loads(contents.decode("utf-8", errors="ignore"))
                # 2022-08-16 17:13:00
                # date_download
                # date_modify
                # date_publish
                # obj["image_url"] = f"/images/{random.choice(categories)}-{random.randint(1, 4)}.jpg"
                # obj["date_download"] = datetime.datetime.strptime(obj["date_download"], '%Y-%m-%d %H:%M:%S')
                # obj["date_modify"] = datetime.datetime.strptime(obj["date_modify"], '%Y-%m-%d %H:%M:%S')
                obj["date_publish"] = datetime.datetime.strptime(obj["date_publish"], '%Y-%m-%d %H:%M:%S')
                article = {
                    "title": obj["title"],
                    "description": obj["description"],
                    "maintext": obj["maintext"],
                    "authors": obj["authors"],
                    "category": random.choice(categories),
                    "date_publish": obj["date_publish"],
                    "source_domain": obj["source_domain"],
                    "url": obj["url"]
                }
                news.append(article)
    return news


def main():
    news = load_data()
    print(len(news))
    client = MongoClient(mongoURL)
    db = client["spainews"]
    collection = db["news"]
    collection.insert_many(news)


if __name__ == "__main__":
    main()
