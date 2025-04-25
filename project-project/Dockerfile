FROM alpine
WORKDIR /Initial
COPY ./target/project-project .
COPY ./config/config-docker.yaml .
RUN  mkdir config && mv config-docker.yaml config/config.yaml \
&& mkdir -p /logs/debug /logs/info /logs/error /logs/warn \
&& apk add --no-cache curl \
&& apk add --no-cache netcat-openbsd \
&& rm -rf /var/cache/apk/*  # 清理缓存（--no-cache 已自动处理，此步可选）\
&& chmod 0777 project-project


EXPOSE 8081 8882
ENTRYPOINT ["./project-project"]

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl --fail http://localhost:8081/ping || exit 1