FROM golang:1.24.2 AS interpreter

WORKDIR /interpreter
COPY go.mod .
COPY go.sum* .
RUN go mod download
COPY . .
RUN GOOS=js GOARCH=wasm go build -o interpreter.wasm


FROM node:22.15.0-alpine

WORKDIR /frontend
COPY /editor/package.json .
RUN npm install
COPY /editor .
COPY --from=interpreter interpreter/interpreter.wasm /frontend/editor/static/wasm/
RUN npm run build





