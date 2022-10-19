import sys

from pymongo import MongoClient
from dotenv import dotenv_values

if len(sys.argv) != 2:
    print(f"Usage: {sys.argv[0]} config.env", file=sys.stderr)
    exit(1)

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

# {
#   "title": "",
#   "description": "",
#   "maintext": "",
#   "authors": [],
#   "category": "finance",
#   "date_publish": "DD/MM/YYYY",
#   "source_domain": "abc.com",
#   "url": "http://abc.com/article"
# }

def insert_article(article: dict[str:any]) -> None:
    if collection.find_one(
        {},
        {
            "title": article["title"],
            "maintext": article["maintext"],
            "authors": article["authors"],
            "source_domain": article["source_domain"],
            "url": article["url"],
        }
    ) is not None:
        return
    collection.insert_one(article)
