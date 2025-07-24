#!/usr/bin/env python3
"""
Examples of how to call streaming APIs from a client perspective
"""

import asyncio
import httpx
import requests
import json


async def test_httpx_streaming():
    """Method 1: Using httpx for async streaming (Recommended)"""
    print("=== Testing httpx streaming ===")
    async with httpx.AsyncClient() as client:
        async with client.stream("GET", "http://localhost:8080/events") as response:
            async for line in response.aiter_lines():
                if line:
                    print(f"Received: {line}")


async def test_httpx_chunked():
    """Method 2: Using httpx with manual chunking"""
    print("=== Testing httpx chunked streaming ===")
    async with httpx.AsyncClient() as client:
        async with client.stream("GET", "http://localhost:8080/events") as response:
            async for chunk in response.aiter_bytes():
                if chunk:
                    print(f"Received chunk: {chunk.decode()}")


def test_requests_streaming():
    """Method 3: Using requests for synchronous streaming"""
    print("=== Testing requests streaming ===")
    response = requests.get("http://localhost:8080/events", stream=True)
    for line in response.iter_lines():
        if line:
            print(f"Received: {line.decode()}")


async def test_with_timeout():
    """Method 4: Streaming with timeout"""
    print("=== Testing streaming with timeout ===")
    async with httpx.AsyncClient(timeout=30.0) as client:
        async with client.stream("GET", "http://localhost:8080/events") as response:
            async for line in response.aiter_lines():
                if line:
                    print(f"Received: {line}")
                    # Break after 5 messages for demo
                    if "message 5" in line:
                        break


async def test_with_headers():
    """Method 5: Streaming with custom headers"""
    print("=== Testing streaming with headers ===")
    headers = {
        "Accept": "text/event-stream",
        "Cache-Control": "no-cache",
        "User-Agent": "StreamingClient/1.0",
    }

    async with httpx.AsyncClient() as client:
        async with client.stream(
            "GET", "http://localhost:8080/events", headers=headers
        ) as response:
            async for line in response.aiter_lines():
                if line:
                    print(f"Received: {line}")


async def main():
    """Run all streaming tests"""
    print("Starting streaming API tests...")
    print("Make sure both api.py (port 8080) and api2.py (port 8081) are running!")
    print()

    # Test async methods
    try:
        await test_httpx_streaming()
        print()
        await test_httpx_chunked()
        print()
        await test_with_timeout()
        print()
        await test_with_headers()
        print()
    except Exception as e:
        print(f"Async test failed: {e}")

    # Test sync method
    try:
        test_requests_streaming()
    except Exception as e:
        print(f"Sync test failed: {e}")


if __name__ == "__main__":
    asyncio.run(main())
