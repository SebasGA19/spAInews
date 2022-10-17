import json
import requests
import train_dataset


max_page = 2


def download() -> list[dict]:
    for page in range(1, max_page):
        for article in json.loads(requests.get(f"http://spainews.bucaramanga.upb.edu.co/api/news/{page}").text)["news"]:
            yield article


def main():
    articles = download()
    for article in articles:
        body = article['maintext'] if 'maintext' in article.keys() else article['description']
        output = train_dataset.check_news_type(body)
        print(article['title'], output)


if __name__ == "__main__":
    main()
