from cmath import e
from scrapper import noticia_cnn_solo , noticia_Huffpost_solo ,noticia_vice_solo , noticia_un_solo,noticia_tiempo_solo,noticia_today_solo

def scrap(article_url):
    try:
        if(str.__contains__(article_url,"edition.cnn")):
            return noticia_cnn_solo(article_url)    
        elif(str.__contains__(article_url,"vice.com")):
            return noticia_vice_solo(article_url)
        elif(str.__contains__(article_url,"huffpost.com")):
            return noticia_Huffpost_solo(article_url)
        elif(str.__contains__(article_url,"news.un.org")):
            return noticia_un_solo(article_url)
        elif(str.__contains__(article_url,"eltiempo.com")):
            return noticia_tiempo_solo(article_url)
        elif(str.__contains__(article_url,"usatoday.com")):
            return noticia_today_solo(article_url)
        else:
            return None
    except:
        print(e)
   
#p = print(scrap("https://edition.cnn.com/2022/09/30/business/nike-stock-discounts/index.html"))

    





    