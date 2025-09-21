import asyncio
import websockets
from json import loads

async def send_wav(file_path):
    binary_audio = open(file_path, "rb").read()
    uri = "ws://localhost:8080/stream"

    async with websockets.connect(uri) as websocket:
        # Send the WAV as binary
        await websocket.send(binary_audio)
        # Signal end of audio
        await websocket.send("END")

        # Receive transcription
        transcription: bytes = await websocket.recv()
        # print("Transcription:", str(transcription))
        decoded = transcription.decode("utf-8")
        result = loads(decoded)
        print("Transcription:", result["text"])


asyncio.run(send_wav("output.wav"))
