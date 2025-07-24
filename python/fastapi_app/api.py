from typing import Union

from fastapi import FastAPI
from fastapi.responses import StreamingResponse
import httpx
import asyncio
import requests

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
    # Method 1: Using httpx for async streaming (Recommended)
    async with httpx.AsyncClient() as client:
        async with client.stream("GET", "http://localhost:8081/events") as response:
            async for line in response.aiter_lines():
                if line:
                    yield f"data: {line}\n\n"


async def stream_messages_method2():
    # Method 2: Using httpx with manual chunking
    async with httpx.AsyncClient() as client:
        async with client.stream("GET", "http://localhost:8081/events") as response:
            async for chunk in response.aiter_bytes():
                if chunk:
                    # Process chunks manually
                    yield f"data: {chunk.decode()}\n\n"


def stream_messages_sync():
    # Method 3: Using requests for synchronous streaming (not recommended for FastAPI)
    response = requests.get("http://localhost:8081/events", stream=True)
    for line in response.iter_lines():
        if line:
            yield f"data: {line}\n\n"


@app.get("/events")
async def send_events():
    return StreamingResponse(stream_messages(), media_type="text/event-stream")


@app.get("/events-sync")
def send_events_sync():
    return StreamingResponse(stream_messages_sync(), media_type="text/event-stream")


@app.get("/events-chunked")
async def send_events_chunked():
    return StreamingResponse(stream_messages_method2(), media_type="text/event-stream")
