FROM amazonlinux:2

WORKDIR /

RUN ls -la

COPY bin/fargate-controller /bin/controller

ENTRYPOINT ["/bin/controller"]

#RUN  go build -a -o $work_dir/bin/controller $work_dir/services/$service_alias/cmd/controller/main.go

