import json
import requests
import re
import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from textblob import Word
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, cohen_kappa_score, confusion_matrix



def clean_str(string):
    """
    Tokenization/string cleaning for datasets.
    Original taken from https://github.com/yoonkim/CNN_sentence/blob/master/process_data.py
    """
    string = re.sub(r"\'s", "", string)
    string = re.sub(r"\'ve", "", string)
    string = re.sub(r"n\'t", "", string)
    string = re.sub(r"\'re", "", string)
    string = re.sub(r"\'d", "", string)
    string = re.sub(r"\'ll", "", string)
    string = re.sub(r",", "", string)
    string = re.sub(r"!", " ! ", string)
    string = re.sub(r"\(", "", string)
    string = re.sub(r"\)", "", string)
    string = re.sub(r"\?", "", string)
    string = re.sub(r"'", "", string)
    string = re.sub(r"[^A-Za-z0-9(),!?\'\`]", " ", string)
    string = re.sub(r"[0-9]\w+|[0-9]","", string)
    string = re.sub(r"\s{2,}", " ", string)
    return string.strip().lower()

data = pd.read_csv('./dataset/dataset.csv')
x = data['news'].tolist()
y = data['type'].tolist()

for index,value in enumerate(x):
    print ("processing data:",index)
    x[index] = ' '.join([Word(word).lemmatize() for word in clean_str(value).split()])

vect = TfidfVectorizer(stop_words='english',min_df=2)
X = vect.fit_transform(x)
Y = np.array(y)

print ("no of features extracted:",X.shape[1])

X_train, X_test, y_train, y_test = train_test_split(X, Y, test_size=0.20, random_state=42)


print ("train size:", X_train.shape)
print ("test size:", X_test.shape)

model = RandomForestClassifier(n_estimators=300, max_depth=150,n_jobs=1)
model.fit(X_train, y_train)
y_pred = model.predict(X_test)

def check_news_type(news_article):  # function to accept raw string and to provide news type(class)
    news_article = [' '.join([Word(word).lemmatize() for word in clean_str(news_article).split()])]
    features = vect.transform(news_article)
    return str(model.predict(features)[0])

max_page = 20

def download() -> list[dict]:
    for page in range(1, max_page):
        for article in json.loads(requests.get(f"http://spainews.bucaramanga.upb.edu.co/api/news/{page}").text)["news"]:
            yield article


def main():
    articles = list(download())
  
    for i in range(len(articles)):
       
     try:
         print(i)
         user = (articles)[i]['maintext']
         print(articles[i]['title'])
         output = check_news_type(user)
         print(output)
     except:
         print(i)
         print("mal formato")
        
    
     with open("news.json", "w") as f:
        json.dump(articles, f)
if __name__ == "__main__":
    main()
    




    


