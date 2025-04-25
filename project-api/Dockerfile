FROM alpine
WORKDIR /Initial
COPY ./target/project-api .
COPY ./config/config-docker.yaml .
RUN  mkdir config && mv config-docker.yaml config/config.yaml && apk update \
&& mkdir -p /logs/debug /logs/info /logs/error /logs/warn \
&& apk add --no-cache curl \
&& apk add --no-cache netcat-openbsd \
&& rm -rf /var/cache/apk/*  # 清理缓存（--no-cache 已自动处理，此步可选）\
&& chmod 0777 project-api

EXPOSE 80
ENTRYPOINT ["./project-api"]

# 健康检查
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl --fail http://localhost/ping || exit 1