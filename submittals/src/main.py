import asyncio
from fastapi import FastAPI
from stan.aio.client import Client as STAN, Subscription
from nats.aio.client import Client as NATS
from .nats_client import NatsClient

app = FastAPI()


@app.on_event('startup')
async def startup():
    loop = asyncio.get_event_loop()
    client = NatsClient(loop)
    await client.connect()
    await client.subscribe(on_message)


@app.get("/submittals")
async def read_submittals():
    return {"message": "this is the submittals endpoint"}


@app.get("/")
async def read_root():
    return {"message": "this is the root route of submittals!"}


def on_message(msg):
    print(msg.seq, msg.data)
