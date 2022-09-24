from urllib.parse import urlparse

import requests
from newsplease import NewsPlease

import common

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


def scrap(article_url: str) -> dict[str:any]:
    try:
        response = requests.get(article_url, headers=common.base_headers)
        if response.status_code != 200:
            return None
        article_html = response.text
        article = vars(NewsPlease.from_html(article_html))
        return {
            "title": article["title"],
            "description": article["description"],
            "maintext": article["maintext"],
            "authors": article["authors"],
            "date_publish": article["date_publish"],
            "source_domain": urlparse(article_url).netloc,
            "url": article_url
        }
    except Exception as e:
        print(type(e), e)
    return None
