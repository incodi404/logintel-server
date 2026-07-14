FROM python:3.11-slim as base

WORKDIR /app

COPY requirements.txt .

RUN --mount=type=cache,target=/root/.cache/pip \pip install --no-cache-dir -r requirements.txt

COPY main.py . logging_config.py . database.py . /app/

FROM base as dev
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]