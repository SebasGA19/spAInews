# TODO: Mario implementa esto!!!!!!!!!!!

import random


categories = ["sports", "finance", "politics"]

def analysis(article: dict[str:any]) -> dict[str:any]:
    article["category"] = random.choice(categories) # TODO: FIX ME
    return article
