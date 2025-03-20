FROM golang:alpine
RUN apk add --no-cache chromium nss freetype freetype-dev \
    harfbuzz ttf-freefont
ENV CHROME_PATH="/usr/bin/chromium-browser"
COPY . /app
WORKDIR /app
RUN go build -o myapp .
CMD ["./myapp"]
