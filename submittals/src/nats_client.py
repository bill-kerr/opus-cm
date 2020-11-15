import asyncio
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN


async def run(loop: asyncio.AbstractEventLoop, future_sc: asyncio.Future[STAN]):
    # Use borrowed connection for NATS then mount NATS Streaming
    # client on top.
    nc = NATS()
    await nc.connect(io_loop=loop, servers=["http://nats-srv:4222"])

    # Start session with NATS Streaming cluster.
    sc = STAN()
    await sc.connect("opuscm", "submittals", nats=nc)
    future_sc.set_result(sc)

    # Synchronous Publisher, does not return until an ack
    # has been received from NATS Streaming.
    # await sc.publish("hi", b'hello')
    # await sc.publish("hi", b'world')
    # total_messages = 0
    # future = asyncio.Future(loop=loop)

    # async def cb(msg):
    #     nonlocal future
    #     nonlocal total_messages
    #     print(f"Received a message (seq={msg.seq}): {msg.data}")
    #     total_messages += 1
    #     if total_messages >= 3:
    #         future.set_result(None)

    # Subscribe to get all messages since beginning.
    future = asyncio.Future(loop=loop)
    sub = await sc.subscribe("test", start_at='first', cb=build_callback(future))
    await asyncio.wait_for(future, 5, loop=loop)
    print("future complete test!")

    # Stop receiving messages
    await sub.unsubscribe()

    # Close NATS Streaming session
    await sc.close()

    # We are using a NATS borrowed connection so we need to close manually.
    await nc.close()


def build_callback(future: asyncio.Future):
    total_messages = 0

    async def callback(msg):
        nonlocal future
        nonlocal total_messages
        print(f"Received a message (seq={msg.seq}): {msg.data}")
        total_messages += 1
        if total_messages >= 3:
            future.set_result(None)
    return callback
