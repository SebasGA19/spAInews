FROM debian

RUN apt-get update -y && apt-get install python3-pip -y
COPY test-data /test-data
WORKDIR /test-data
RUN pip install -r /test-data/requiriments.txt
CMD python3 /test-data/main.py