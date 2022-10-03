from distutils.filelist import findall
import json
from operator import truediv
from bs4 import BeautifulSoup
import requests

##Páginas seleccionadas
##CNN 
##Huffpost
##Vice
##UN NEWS
##El tiempo --eng version 
##USA TODAY

##DESARROLLO
data_principal = []
##CNN
##CNN NEWS SOLO// BUSSINES
def noticia_cnn_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    try:
        cuerpo = soup.find('div',class_="article__content").get_text()
        p = {
        "authors":soup.find('span',class_="byline__name").get_text(),
        "date_publish":soup.find('div',class_="timestamp").get_text(),
        "description":cuerpo.split(".")[0],
        "category":"",
        "language":"en",
        "source_domain":"CNN",
        "maintext":cuerpo,
        "url":url,
        "title":soup.find('h1',class_="headline__text inline-placeholder").get_text()
    }
    except:
        p = ""
    return p   
    #data_principal.append(p)
   
##CNN NEWS LINK
def generar_cnn_varias(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    contenedor = soup.findAll('h3',class_="cd__headline")
    for i in range(0,len(contenedor)):
        try:
            var = contenedor[i].findChild("a")['href']
            var = "https://edition.cnn.com"+ var
            noticia_cnn_solo(var)
        except:
            print("NO SE PUDO")

#generar_cnn_varias('https://edition.cnn.com/business')
# noticia_cnn_solo('https://edition.cnn.com/2022/08/17/business/dodge-electric-muscle-car/index.html')

##Huffpost
##Huffpost solo
def noticia_Huffpost_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    try:
        cuerpo = soup.find('div',class_="entry__content-list-container js-cet-unit-buzz_body").get_text()
        p = {
            "authors":soup.find('span',class_="entry-wirepartner__byline").get_text(),
            "date_publish":soup.find('time').get_text(),
            "description":cuerpo.split(".")[0],
            "category":"",
            "language":"en",
            "source_domain":"HuffPost",
            "maintext":cuerpo,
            "url":url,
            "title":soup.find('h1',class_="headline").get_text()
        }
    except:
        try:
            cuerpo = soup.find('section',class_="entry__content-list js-entry-content js-cet-subunit").get_text()
            p = {
                "authors":soup.find('a',class_="js-entry-link cet-internal-link")['data-vars-item-name'],
                "date_publish":soup.find('time')['datetime'],
                "description":cuerpo.split(".")[0],
                "category":"",
                "language":"en",
                "source_domain":"HuffPost",
                "maintext":cuerpo,
                "url":url,
                "title":soup.find('h1',class_="headline").get_text()
            }  
        except:
            p = ""
    return p   
    #data_principal.append(p)

def generar_Huffpost_varias(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    contenedor = soup.findAll('a',class_="card__headline card__headline--long")
    for i in range(0,len(contenedor)):
        try:
            data = contenedor[i]['href']
            noticia_Huffpost_solo(data)
        except:
            print("NO SE PUDO")

#noticia_Huffpost_solo("https://www.huffpost.com/entry/never-have-i-ever-netflix-maitreyi-ramakrishnan_n_62fa49e5e4b045e6f6af1f32")
#generar_Huffpost_varias("https://www.huffpost.com/")

##VICE
##Vice solo
def noticia_vice_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    try:
        autor_sub = soup.find('div',class_="contributor__meta")
        cuerpo = soup.find('div',class_="article__body-components").get_text()
        p = {
            "authors":autor_sub.findChild("a").get_text(),
            "date_publish":soup.find('time')['datetime'],
            "description":cuerpo.split(".")[0],
            "category":"",
            "language":"en",
            "source_domain":"Vice",
            "maintext":cuerpo,
            "url":url,
            "title":soup.find('h1',class_="smart-header__hed smart-header__hed--size-2").get_text()
        }  
    except:
        p = ""
      
    return p   
    #data_principal.append(p)

##Vice varias
def generar_vice_varias(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    contenedor = soup.findAll('h3',class_="vice-card-hed vice-card-hed--light vice-card__vice-card-hed")
    for i in range(0,len(contenedor)):
        try:
            data = contenedor[i].findChild("a")["href"]
            data = "https://www.vice.com"+ data
            noticia_vice_solo(data)     
        except:
            print("NO SE PUDO")

#noticia_vice_solo("https://www.vice.com/en/article/59qjed/how-westworld-uses-multilingualism-to-explore-prejudice")
#generar_vice_varias("https://www.vice.com/en/topic/english?page=1")

##UN NEWS 
##UN NEWS SOLO 
def noticia_un_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    try:
        cuerpo_sub = soup.find('div',class_="clearfix text-formatted field field--name-field-text-column field--type-text-long field--label-hidden field__item")
        cuerpo_x = soup.findChildren('p')
        cuerpo = ""
        for i in range(len(cuerpo_x)):
            if(not(str.__contains__(cuerpo_x[i].get_text(),"Twitter")) and not(str.__contains__(cuerpo_x[i].get_text(),"Email")) and not(str.__contains__(cuerpo_x[i].get_text(),"Facebook")) and not(str.__contains__(cuerpo_x[i].get_text(),"Print"))):
                cuerpo = cuerpo + cuerpo_x[i].get_text() 
        p = {
            "authors":soup.find('div',class_="field__item").get_text(),
            "date_publish":soup.find('time',class_="datetime").get_text(),
            "description":cuerpo.split(".")[0],
            "category":"",
            "language":"en",
            "source_domain":"OnuNews",
            "maintext":cuerpo,
            "url":url,
            "title":soup.find('span',class_="field field--name-title field--type-string field--label-hidden").get_text()
        } 
    except:
            p = "sa"
          
    return p   
    #data_principal.append(p)
        
##UN VARIAS
def generar_un_varias(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    contenedor = soup.findAll('h3',class_="story-title")
    for i in range(0,len(contenedor)):
        try:
            data_1 = contenedor[i].findChild("a")["href"]
            data_1 = "https://news.un.org/" + data_1
            noticia_un_solo(data_1)
        except:
            print("NO SE PUDO")


#generar_un_varias("https://news.un.org/en/")
#noticia_un_solo('https://news.un.org//en/story/2022/08/1125522')

##EL TIEMPO
##El tiempo solo
def noticia_tiempo_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    contenedor_autor = soup.find('div',class_="author_data")
    cuerpo_sub = soup.findAll('p',class_="contenido")
    cuerpo = ""
    for i in range(len(cuerpo_sub)):
        cuerpo = cuerpo + cuerpo_sub[i].get_text()
    p = {
            "authors":contenedor_autor.findChild("span",class_="who-modulo who"),
            "date_publish":contenedor_autor.findChild("span",class_="publishedAt").get_text(),
            "description":cuerpo.split(".")[0],
            "category":"",
            "language":"en",
            "source_domain":"TheTime",
            "maintext":cuerpo,
            "url":url,
            "title":soup.find('h1',class_="titulo").get_text()
        }  
    return p   
    #data_principal.append(p)

#El tiempo Varias
def generar_tiempo_varias(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    contenedor = soup.findAll('h3',class_="title-container")
    for i in range(0,len(contenedor)):
        try:
            data_1 = contenedor[i].findChild("a",class_="title page-link")["href"]
            data_1 = "https://www.eltiempo.com/" + data_1
            noticia_tiempo_solo(data_1)
        except:
            print("NO SE PUDO")   

#noticia_tiempo_solo("https://www.eltiempo.com/cultura/gente/pablo-escobar-who-are-the-sons-of-the-famous-colombian-drug-lord-619763")
#generar_tiempo_varias("https://www.eltiempo.com/noticias/english-news")

##USA TODAY 
##Solo
def generar_today_varias(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    contenedor_links_secciones = soup.findAll('a',class_="gnt_m_sb_mn")
    for i in range(0,len(contenedor_links_secciones)):
        hijo_seccion = contenedor_links_secciones[i]["href"]
        hijo_seccion = "https://www.usatoday.com/"+hijo_seccion
        resultado_seccion = requests.get(hijo_seccion)
        contenido_seccion = resultado_seccion.text
        soup_seccion = BeautifulSoup(contenido_seccion,'html.parser')
        contenedor_noticias = soup_seccion.findAll('a',class_="gnt_m_flm_a")
        if(len(contenedor_noticias)>0):
            for i in range(0,len(contenedor_noticias)):
                try:
                    linx = "https://www.usatoday.com/"+contenedor_noticias[i]["href"]
                    noticia_today_solo(linx)
                except:
                    print("No se pudo")


def noticia_today_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    cuerpo = soup.find_all("p",class_="gnt_ar_b_p")
    cuerpo_principal = ""
    for q in cuerpo:
        cuerpo_principal = cuerpo_principal + "\n" + q.get_text()
    p = {
            "authors":soup.find("a",class_="gnt_ar_by_a").get_text(),
            "date_publish":soup.find("div",class_="gnt_ar_dt")["aria-label"],
            "description":cuerpo[0].get_text(),
            "category":"",
            "language":"en",
            "source_domain":"usatoday.com",
            "maintext":cuerpo_principal,
            "url":url,
            "title":soup.find("h1",class_="gnt_ar_hl").get_text()
        }  
    return p   
    #data_principal.append(p)
 
#generar_today_varias("https://www.usatoday.com/")
#noticia_today_solo("https://www.usatoday.com/story/news/politics/2022/09/19/pandemic-is-over-biden-comment-complicates-matters/10425981002/?gnt-cfr=1")

def main():
    generar_cnn_varias('https://edition.cnn.com/business')
    generar_Huffpost_varias("https://www.huffpost.com/")
    generar_vice_varias("https://www.vice.com/en/topic/english?page=1")
    generar_un_varias("https://news.un.org/en/")
    generar_tiempo_varias("https://www.eltiempo.com/noticias/english-news")
    generar_today_varias("https://www.usatoday.com/")
    with open("data.json","w") as f:
        json.dump(data_principal,f)

def leer():
    texto =""
    with open("data.json") as f1:
        texto = json.load(f1)
    print(texto[0])
    print(len(texto)) 

##main()
##leer()

##Funcionalidades 
#print(noticia_cnn_solo("https://edition.cnn.com/2022/08/17/business/dodge-electric-muscle-car/index.html"))
#print(noticia_Huffpost_solo("https://www.huffpost.com/entry/jair-bolsonaro-brazil-election-trump_n_63376ea2e4b0e376dbf6e52b"))
#print(noticia_vice_solo("https://www.vice.com/en/article/jmk9qy/we-asked-people-with-the-most-useless-sounding-university-degrees-if-they-regret-their-life-choices"))
#print(noticia_un_solo("https://news.un.org/en/story/2022/09/1129082")["title"])
#print(noticia_tiempo_solo("https://www.eltiempo.com/cultura/gente/epa-colombia-controversial-colombian-instagram-influencer-daneidy-barrera-619696")["maintext"])
#print(noticia_today_solo("https://www.usatoday.com/story/news/nation/2022/10/02/hurricane-ian-live-updates-death-toll/8159168001/")["title"])