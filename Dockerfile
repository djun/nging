FROM alpine
RUN apk update
RUN apk upgrade

RUN wget -c https://dl.webx.top/nging/v4.1.5/nging_linux_amd64.tar.gz -O /home/nging_linux_amd64.tar.gz
RUN tar -zxvf /home/nging_linux_amd64.tar.gz -C /home

WORKDIR /home/nging_linux_amd64

# VOLUME [ "/home/nging_linux_amd64/data/cache", "/home/nging_linux_amd64/data/ftpdir", "/home/nging_linux_amd64/data/logs", "/home/nging_linux_amd64/data/sm2", "/home/nging_linux_amd64/myconfig", "/home/nging_linux_amd64/public" ]

ENTRYPOINT [ "./nging" ]
CMD [ "-p", "9999", "-c", "myconfig/config.yaml" ]

# docker build . -t "admpub/nging:latest" 
