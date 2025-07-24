from typing import Union

from fastapi import FastAPI
from fastapi.responses import StreamingResponse
import asyncio

# run at port 8081
# fastapi dev api.py --port 8081
app = FastAPI()

messages = [
    "message 1",
    "message 2",
    "message 3",
    "message 4",
    "message 5",
    "message 6",
    "message 7",
    "message 8",
]


async def stream_messages():
    for message in messages:
        yield f"data: {message}\n\n"
        await asyncio.sleep(2)


@app.get("/events")
def send_events():
    return StreamingResponse(stream_messages(), media_type="text/event-stream")
