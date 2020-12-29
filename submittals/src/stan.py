import asyncio
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN


async def get_nats_connection():
    nats_connection = NATS()
    await nats_connection.connect(servers=["http://nats-srv:4222"])
    return nats_connection


async def get_stan_connection(nats_connection: NATS):
    stan_connection = STAN()
    await stan_connection.connect("opuscm", "submittals", nats=nats_connection)
    print("Connected to NATS")
    return stan_connection


async def handle_requests():
    nats_connection = await get_nats_connection()
    stan_connection = await get_stan_connection(nats_connection)

    async def cb(msg):
        print(msg)
        await stan_connection.ack(msg)

    sub = await stan_connection.subscribe("test", manual_acks=True, start_at="first", cb=cb, ack_wait=5000)


async def run(loop):
    # Use borrowed connection for NATS then mount NATS Streaming
    # client on top.
    nc = NATS()
    await nc.connect(io_loop=loop, servers=["http://nats-srv:4222"])

    # Start session with NATS Streaming cluster.
    sc = STAN()
    await sc.connect("opuscm", "submittals", nats=nc)

    # Synchronous Publisher, does not return until an ack
    # has been received from NATS Streaming.
    await sc.publish("hi", b'hello')
    await sc.publish("hi", b'world')

    total_messages = 0
    future = asyncio.Future(loop=loop)

    async def cb(msg):
        nonlocal future
        nonlocal total_messages
        print("Received a message (seq={}): {}".format(msg.seq, msg.data))
        total_messages += 1

    # Subscribe to get all messages since beginning.
    sub = await sc.subscribe("test", start_at='first', cb=cb)
