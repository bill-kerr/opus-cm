import asyncio
from fastapi import FastAPI
from stan.aio.client import Client as STAN, Subscription
from nats.aio.client import Client as NATS
from .nats_client import NatsClient
from .stan import handle_requests

app = FastAPI()


@app.on_event('startup')
async def startup():
    loop = asyncio.get_event_loop()
    asyncio.ensure_future(handle_requests(), loop=loop)
    # client = NatsClient(loop)
    # await client.connect()
    # future = asyncio.Future()
    # future.add_done_callback(got_result)
    # loop.create_task(client.subscribe(on_message(future)))
    # print("Hello from before subscribe")
    # print("Hello from after subscribe")
    pass


@app.get("/")
async def read_submittals():
    return {"message": "this is the submittals endpoint"}


def on_message(future: asyncio.Future):
    async def cb(msg):
        print(msg.seq, msg.data)
        future.set_result(msg.data)
    return cb


def got_result(future: asyncio.Future):
    print(future.result())
