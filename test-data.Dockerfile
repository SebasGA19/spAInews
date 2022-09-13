FROM python

WORKDIR /test-data
CMD pip install -r /test-data/requiriments.txt && python3 /test-data/main.py