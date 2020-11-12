FROM python:3.8-alpine

RUN pip install pipenv
WORKDIR /app
COPY Pipfile* ./
RUN pipenv lock --requirements > requirements.txt
RUN pip install -r requirements.txt
COPY ./ ./
RUN cd src
WORKDIR /app/src

CMD ["uvicorn", "main:app", "--reload"]