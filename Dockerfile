FROM golang:1.24.4-alpine3.22 AS build
WORKDIR /app
RUN apk add --no-cache npm
COPY package*.json .
RUN npm i
COPY . .
RUN npx @tailwindcss/cli -i ./input.css -o ./static/output.css
RUN go run main.go > static/index.html

FROM caddy:2.10.0-alpine
COPY --from=build /app/static/ /usr/share/caddy/
