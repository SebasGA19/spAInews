FROM python

COPY crawler /crawler
WORKDIR /crawler
RUN pip install -r ./requirements.txt

CMD python3 ./main.py /configs/crawler.env