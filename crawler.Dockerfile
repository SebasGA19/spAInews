FROM python

COPY crawler /crawler
WORKDIR /crawler
RUN pip install -r ./requirements.txt
RUN python3 -c "import nltk; nltk.download('all')"

CMD python3 ./main.py /configs/crawler.env