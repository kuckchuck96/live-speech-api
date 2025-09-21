# Live Speech API

## Steps to run
- `cd ./speech-api`
- `make run`
- `cd ./asr-api`
- Create virtual env and run `pip install -r ./requirements.txt`
- `make run`
- To test web socket connection `python ./ws_server.py`
- Optionally to create new audio use `python ./record.py`
