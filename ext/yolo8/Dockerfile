FROM python:3.10.11-slim-buster

RUN apt-get update && apt-get install -y libgl1 libglib2.0-0

COPY requirements.txt /app/
RUN pip install -r /app/requirements.txt

COPY . /app/
WORKDIR /app

EXPOSE 5000

CMD ["python3", "main.py"]
