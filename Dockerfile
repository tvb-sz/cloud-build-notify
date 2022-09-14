FROM alpine:3.16

# 简介说明
LABEL Maintainer="team tvb-sz <nmg-sz@tvb.com>" Description="Google Cloud Build notify Sender Image"

# install bash & tzData & set TimeZone HK
RUN apk update && apk add --no-cache bash tzdata \
    && cp /usr/share/zoneinfo/Asia/Hong_Kong /etc/localtime \
    && echo "Asia/Hong_Kong" > /etc/timezone \
    && apk del tzdata \
    && rm -rf /var/cache/apk/*

# 定义工作目录
WORKDIR /srv/

# 拷贝可执行二进制文件
COPY notify /srv/notify

# 赋予二进制文件可执行权限
RUN chmod +x /srv/notify

# 执行入口点，外部通过参数指定执行的子命令和秘钥等信息
ENTRYPOINT ["/srv/notify"]
