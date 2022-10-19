import os
import re
import pickle

import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from textblob import Word
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier

model_filename = 'model.pickle'


class Analysis(object):
    def __init__(self, vect, model) -> None:
        self.vect = vect
        self.model = model


def clean_str(string):
    if type(string) not in (str, bytes):
        return ""
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
    string = re.sub(r"[0-9]\w+|[0-9]", "", string)
    string = re.sub(r"\s{2,}", " ", string)
    return string.strip().lower()


def init():
    if not os.path.exists(model_filename):
        print("Training model")
        data = pd.read_csv('./dataset/dataset.csv')
        x = data['news'].tolist()
        y = data['type'].tolist()

        for index, value in enumerate(x):
            x[index] = ' '.join([Word(word).lemmatize()
                                for word in clean_str(value).split()])

        vect = TfidfVectorizer(stop_words='english', min_df=2)
        X = vect.fit_transform(x)
        Y = np.array(y)

        X_train, X_test, y_train, y_test = train_test_split(
            X, Y, test_size=0.20, random_state=42)
        model = RandomForestClassifier(n_estimators=300, max_depth=150, n_jobs=1)
        model.fit(X_train, y_train)
        analysis = Analysis(vect, model)
        with open(model_filename, 'wb') as model_file:
            pickle.dump(analysis, model_file)
        return analysis
    with open(model_filename, 'rb') as model_file:
        analysis = pickle.load(model_file)
        return analysis


analysis = init()

# function to accept raw string and to provide news type(class)
def check_news_type(news_article):
    news_article = [' '.join([Word(word).lemmatize()
                             for word in clean_str(news_article).split()])]
    features = analysis.vect.transform(news_article)
    return str(analysis.model.predict(features)[0])
