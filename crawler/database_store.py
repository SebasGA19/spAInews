import sys
import os

from pymongo import MongoClient
from dotenv import dotenv_values

if len(sys.argv) != 2:
    print(f"Usage: {sys.argv[0]} config.env", file=sys.stderr)
    os.exit(1)

env = dotenv_values(sys.argv[1])
mongo_host = env["MONGO_HOST"]
mongo_port = env["MONGO_PORT"]
mongo_username = env["MONGO_USER"]
mongo_password = env["MONGO_PASS"]
mongo_database = env["MONGO_DB"]
mongoURL = f"mongodb://{mongo_username}:{mongo_password}@{mongo_host}:{mongo_port}/{mongo_database}"

client = MongoClient(mongoURL)
db = client["spainews"]
collection = db["news"]


def insert_article(article: dict[str:any]) -> None:
    ref = article.copy()
    del ref["category"]
    collection.update_one(ref, {"$set": article}, upsert=True)
