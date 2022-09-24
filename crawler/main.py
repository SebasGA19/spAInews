from analysis import analysis
from database_store import insert_article
import scrap_articles
import obtain_rss


def main() -> None:
    for url in obtain_rss.get_urls():
        article = scrap_articles.scrap(url)
        if article is None:
            continue
        article = analysis(article)
        insert_article(article)
        


if __name__ == "__main__":  # Try the script
    main()
