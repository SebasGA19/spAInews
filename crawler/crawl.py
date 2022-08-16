import datetime

from newsplease import NewsPlease

date_format = '%Y/%m/%d'

def scrap(urls: list[str]) -> list[dict[str:any]]:
    articles = NewsPlease.from_urls(urls, timeout=10)
    result = []
    for article in articles.values():
        new_entry = {key: value for key, value in map(
            lambda item: (item[0], item[1]) if type(item[1]) != datetime.datetime 
            else (item[0], item[1].strftime(date_format)), 
            filter(lambda item: not item[0].startswith("__"), vars(article).items()))}
        result.append(new_entry)
    return result


if __name__ == "__main__": # Try the script
    articles = ["https://www.rt.com/news/560956-russian-defense-ministry-uk-provocation/", "https://www.nytimes.com/2022/08/16/business/economy/covid-pandemic-fraud.html"]
    result = scrap(articles)
    print(result)