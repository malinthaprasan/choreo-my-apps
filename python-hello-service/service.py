
from fastapi import FastAPI
from typing import Dict

app = FastAPI()

@app.get("/")
async def root() -> Dict[str, str]:
    return {"message": "Hello World"}

@app.get("/hello/{name}")
async def say_hello(name: str) -> Dict[str, str]:
    return {"message": f"Hello {name}"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)

