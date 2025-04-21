FROM python:3.9.22-alpine
WORKDIR /app
COPY . .
RUN pip install matplotlib
RUN pip install pika
CMD ["python",  "./main.py"]