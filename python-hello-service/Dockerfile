# ------------------------------
# Stage 1: Build dependencies
# ------------------------------
FROM python:alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies (using pip cache for faster builds)
COPY requirements.txt .

RUN python3 -m venv /app/venv

# RUN --mount=type=cache,target=/root/.cache/pip . /app/venv/bin/activate && \
#     python3 -m ensurepip --upgrade && \
#     python3 -m pip install -r /app/requirements.txt

RUN . /app/venv/bin/activate && \
    python3 -m ensurepip --upgrade && \
    python3 -m pip install -r /app/requirements.txt

# Copy application source code
COPY . .

# ------------------------------
# Stage 2: Production image
# ------------------------------
FROM python:alpine

# Set working directory
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app /app

# Set environment variables
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

# Expose application port (modify as needed)
EXPOSE 8000

# Create a user with a known UID/GID within range 10000-20000.
# This is required by Choreo to run the container as a non-root user.
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 10014 \
    "choreo"
# Use the above created unprivileged user
USER 10014

# Start the Python backend (modify for Flask, FastAPI, or Django)
# uvicorn service:app --host 0.0.0.0 --port 8000
CMD ["/bin/sh", "-c", ". /app/venv/bin/activate && uvicorn service:app --host 0.0.0.0 --port 8000"]
    