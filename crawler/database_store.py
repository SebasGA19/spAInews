from pymongo import MongoClient

mongoURL = "mongodb://spainews:spainews@127.0.0.1:27017/spainews"
client = MongoClient(mongoURL)
db = client["spainews"]
collection = db["news"]


def insert_article(article: dict[str:any]) -> None:
    ref = article.copy()
    del ref["category"]
    collection.update_one(ref, {"$set": article}, upsert=True)
