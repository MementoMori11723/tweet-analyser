FROM golang:1.23-bullseye
RUN apt-get update && apt-get install -y wget gnupg && \
    wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | gpg --dearmor > /usr/share/keyrings/google-linux-keyring.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/google-linux-keyring.gpg] http://dl.google.com/linux/chrome/deb/ stable main" | tee /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*
ENV CHROME_BIN=/usr/bin/google-chrome
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .
ARG API_KEY
ENV API_KEY=$API_KEY
ARG PORT
ENV PORT=$PORT
EXPOSE $PORT
CMD ["./main"]
