#stage 01 -building
FROM golang AS builder

ENV GOPATH=/go
ENV GO111MODULE=on

ENV APP_NAME=JobCreator
ENV GIT_URL=https://github.com/dylan96dashintha/JobCreator
ENV BRANCH=master
WORKDIR /tmp

# Clone the repository from Git
RUN apt-get update && \
    apt install git && \
    git clone ${GIT_URL} && \
    cd ${APP_NAME} && \
    git checkout ${BRANCH}


RUN cd ${APP_NAME} && \
    go clean && \
    go mod download && \
    go get github.com/JobCreator/app/http && \
    go build -o ${APP_NAME}

RUN chmod +x ${APP_NAME}



#stage2  - run
FROM alpine
RUN apk add --update --no-cache tzdata bash libc6-compat

ENV APP_NAME=JobCreator
WORKDIR /app

#COPY build binary
COPY --from=builder /tmp/${APP_NAME}/${APP_NAME} .

##COPY config config
#COPY --from=builder /build/${APP_NAME}/config config


ENTRYPOINT ./${APP_NAME}

