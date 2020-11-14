from fastapi import FastAPI
from .nats_client import run

app = FastAPI()

# loop = asyncio.get_event_loop()
# loop.run_until_complete(run(loop))
# loop.close()


@app.get("/")
def read_root():
    return {"message": "this is the root route of submittals"}


@app.get("/submittals")
def read_submittals():
    return {"message": "this is the submittals endpoint"}


print(__name__)
if __name__ == "main":
    print(__package__)
    print(__name__)
