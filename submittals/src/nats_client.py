import asyncio
from typing import List
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN, Subscription


class NatsClient:
    def __init__(self, loop: asyncio.AbstractEventLoop):
        self.__loop = loop
        self.__nc: NATS
        self.__sc: STAN
        self.__subs: List[Subscription] = []

    async def connect(self):
        self.__nc = NATS()
        await self.__nc.connect(io_loop=self.__loop, servers=["http://nats-srv:4222"])
        self.__sc = STAN()
        await self.__sc.connect("opuscm", "submittals", nats=self.__nc)
        print("Connected to NATS")

    def subscribe(self, on_message):
        return self.__sc.subscribe("test", start_at="first", cb=on_message)

    async def publish(self):
        await self.__sc.publish("test", b"this is a test from the submittals service")
