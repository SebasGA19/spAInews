from distutils.filelist import findall
import json
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
##CNN NEWS SOLO
def noticia_cnn_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    try:
        titulo = soup.find('h1',class_="pg-headline").get_text()
        autor = soup.find('span',class_="metadata__byline__author").get_text()
        fecha = soup.find('p',class_="update-time").get_text()
        cuerpo = soup.find('div',class_="l-container").get_text()
        descripcion = cuerpo.split(".")[0]
        categoria =""
        link = url
        dominio = "CNN"
        lenguaje = "en"
    except:
        print("")
    p = {
            "authors":autor,
            "date_publish":fecha,
            "description":descripcion,
            "category":"",
            "language":lenguaje,
            "source_domain":dominio,
            "maintext":cuerpo,
            "url":link,
            "title":titulo
        }   
    data_principal.append(p)
   
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
##noticia_cnn_solo('https://edition.cnn.com/2022/08/17/business/dodge-electric-muscle-car/index.html')

##Huffpost
##Huffpost solo
def noticia_Huffpost_solo(url):
    resultado = requests.get(url)
    contenido = resultado.text
    soup = BeautifulSoup(contenido,'html.parser')
    try:
        titulo = soup.find('h1',class_="headline").get_text()
        autor = soup.find('span',class_="entry-wirepartner__byline").get_text()
        fecha = soup.find('time').get_text()
        cuerpo = soup.find('div',class_="entry__content-list-container js-cet-unit-buzz_body").get_text()
        descripcion = cuerpo.split(".")[0]
        lenguaje = "en"
        dominio = "HuffPost"
        link = url
        
    except:
        titulo = soup.find('h1',class_="headline").get_text()
        autor = soup.find('a',class_="js-entry-link cet-internal-link")['data-vars-item-name']
        fecha = soup.find('time')['datetime']
        cuerpo = soup.find('section',class_="entry__content-list js-entry-content js-cet-subunit").get_text()
        descripcion = cuerpo.split(".")[0]
        lenguaje = "en"
        dominio = "HuffPost"
        link = url
    p = {
            "authors":autor,
            "date_publish":fecha,
            "description":descripcion,
            "category":"",
            "language":lenguaje,
            "source_domain":dominio,
            "maintext":cuerpo,
            "url":link,
            "title":titulo
        }        

    data_principal.append(p)

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
        titulo = soup.find('h1',class_="smart-header__hed smart-header__hed--size-2").get_text()
        autor_sub = soup.find('div',class_="contributor__meta")
        autor = autor_sub.findChild("a").get_text()
        fecha = soup.find('time')['datetime']
        cuerpo = soup.find('div',class_="article__body-components").get_text()
        descripcion = cuerpo.split(".")[0]
        lenguaje ="en"
        dominio ="Vice"
        link = url 
    except:
        print("No se pudo")
    p = {
            "authors":autor,
            "date_publish":fecha,
            "description":descripcion,
            "category":"",
            "language":lenguaje,
            "source_domain":dominio,
            "maintext":cuerpo,
            "url":link,
            "title":titulo
        }    
    data_principal.append(p)

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
        titulo = soup.find('h1',class_="story-title").get_text()
        autor = soup.find('span',class_="un-news-full-width scald-credit").get_text()
        fecha = soup.find('div',class_="field-item even").get_text()
        cuerpo_sub = soup.findAll('p')
        cuerpo = ""
        for i in range(len(cuerpo_sub)):
            cuerpo = cuerpo + cuerpo_sub[i].get_text()
        descripcion = cuerpo.split(".")[0]
        lenguaje = "en"
        dominio = "OnuNews"
        link = url
    except:
        titulo = soup.find('h2',class_="story-title quote-text").get_text()
        autor = soup.find('span',class_="un-news-feature scald-credit").get_text()
        fecha = soup.find('div',class_="field-item even").get_text()
        cuerpo_sub = soup.findAll('p')
        cuerpo = ""
        for i in range(len(cuerpo_sub)):
            cuerpo = cuerpo + cuerpo_sub[i].get_text()
        descripcion = cuerpo.split(".")[0]
        lenguaje = "en"
        dominio = "OnuNews"
        link = url
    p = {
            "authors":autor,
            "date_publish":fecha,
            "description":descripcion,
            "category":"",
            "language":lenguaje,
            "source_domain":dominio,
            "maintext":cuerpo,
            "url":link,
            "title":titulo
        }   
    data_principal.append(p)
        
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
    titulo = soup.find('h1',class_="titulo").get_text()
    contenedor_autor = soup.find('div',class_="author_data")
    autor = contenedor_autor.findChild("span",class_="who-modulo who")
    fecha = contenedor_autor.findChild("span",class_="publishedAt").get_text()
    cuerpo_sub = soup.findAll('p',class_="contenido")
    cuerpo = ""
    for i in range(len(cuerpo_sub)):
        cuerpo = cuerpo + cuerpo_sub[i].get_text()

    descripcion = cuerpo.split(".")[0]
    lenguaje ="en"
    dominio = "TheTime"
    link = url
    p = {
            "authors":autor,
            "date_publish":fecha,
            "description":descripcion,
            "category":"",
            "language":lenguaje,
            "source_domain":dominio,
            "maintext":cuerpo,
            "url":link,
            "title":titulo
        }  
    data_principal.append(p)

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
    data_principal.append(p)
 
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

main()
leer()