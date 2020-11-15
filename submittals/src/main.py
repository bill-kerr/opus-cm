import asyncio
from fastapi import FastAPI
from stan.aio.client import Client as STAN
from .nats_client import run


app = FastAPI()


@app.on_event('startup')
async def startup():
    await connect_nats()


@app.get("/submittals")
async def read_submittals():
    return {"message": "this is the submittals endpoint"}


@app.get("/")
async def read_root():
    return {"message": "this is the root route of submittals!"}


async def connect_nats():
    loop = asyncio.get_event_loop()
    future = asyncio.Future[STAN](loop=loop)
    run_task = run(loop, future)
    sc = await asyncio.wait_for(future, None, loop=loop)
    print("Connected to NATS: ", sc)
    loop.create_task(run_task)
