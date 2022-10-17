import train_dataset


def analysis(article: dict[str:any]) -> dict[str:any]:
    body = article['maintext'] if 'maintext' in article.keys() else article['description']
    category = train_dataset.check_news_type(body)
    article["category"] = category
    return article
