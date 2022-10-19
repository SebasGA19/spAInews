import re

import requests
import feedparser

import common


pattern = re.compile(r'"https?://[0-9a-zA-Z_.-]+/[0-9a-zA-Z_./-]+\.?(?:rss|xml)"')

sources = [
    "https://edition.cnn.com/services/rss/",
    "https://www.rt.com/rss-feeds/", # TODO: Fix me
    "https://www.foxnews.com/story/foxnews-com-rss-feeds"
]


def __get_feeds() -> feedparser.util.FeedParserDict:
    for source in sources:
        try:
            response = requests.get(source, headers=common.base_headers)
            if response.status_code != 200:
                continue
            for rss_doc in (url[1:len(url)-1] for url in (result.group() for result in pattern.finditer(response.text))):
                feeds = feedparser.parse(rss_doc)
                yield from feeds["entries"]
        except Exception as e:
            print(type(e), e)


def get_urls() -> str:
    for feed in __get_feeds():
        yield feed["link"]


def main():
    for url in get_urls():
        print(url)


if __name__ == "__main__":
    main()