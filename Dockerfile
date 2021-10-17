# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gobuffalo/buffalo:v0.17.3 as builder

ENV GO111MODULE on
ENV GOPROXY http://proxy.golang.org

RUN mkdir -p /src/callofduty
WORKDIR /src/callofduty

# this will cache the npm install step, unless package.json changes
ADD package.json .
RUN npm install --no-progress
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

ADD . .
RUN buffalo build --static -o /bin/callofduty

FROM alpine:3.14

WORKDIR /bin/

COPY --from=builder /bin/callofduty .

# Port to expose the applicaiton
EXPOSE 3000

# Run the migrations before running the binary:
CMD /bin/callofduty migrate; /bin/callofduty
