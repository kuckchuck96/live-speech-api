from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
from contextlib import asynccontextmanager
from transformers import pipeline
import tempfile
import os


@asynccontextmanager
async def lifespan(app: FastAPI):
    # Perform any startup actions here
    print("🔄 Loading Whisper model...")
    app.state.asr_pipeline = pipeline("automatic-speech-recognition", model="dkunal96/desi-whisper", device="mps")
    print("✅ Whisper model loaded")
    yield
    # Perform any shutdown actions here
    del app.state.asr_pipeline
    print("🛑 Whisper model unloaded")

app = FastAPI(title="Whisper ASR API", version="1.0.0", lifespan=lifespan)


@app.get("/health")
async def health():
    return {"status": "ok"}


@app.post("/transcribe")
async def transcribe(request: Request):
  try:
    audio_bytes = await request.body()

    # Save to temporary file
    with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp:
      tmp.write(audio_bytes)
      tmp_path = tmp.name

    # Run transcription
    result = request.app.state.asr_pipeline(tmp_path)

    # Clean up
    os.remove(tmp_path)

    return JSONResponse(content={"text": result["text"]})
  except Exception as e:
    return JSONResponse(content={"error": str(e)}, status_code=500)
