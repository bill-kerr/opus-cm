import asyncio
from fastapi import FastAPI
from .nats_client import run

app = FastAPI()


@app.on_event('startup')
async def startup():
    pass


@app.get("/")
def read_root():
    return {"message": "this is the root route of submittals!"}


@app.get("/submittals")
def read_submittals():
    return {"message": "this is the submittals endpoint"}
