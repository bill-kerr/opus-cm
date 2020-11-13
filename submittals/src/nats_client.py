import asyncio
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN


async def run(loop):
    nc = NATS()
    await nc.connect(io_loop=loop)

    sc = STAN()
    await sc.connect("opuscm", "submittals", nats=nc)
    print("Connected to NATS server.")

    future = asyncio.Future(loop=loop)

    await asyncio.wait_for(future, 1, loop=loop)

    await sc.close()
    await nc.close()
