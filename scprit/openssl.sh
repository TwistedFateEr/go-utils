#!/usr/bin/env bash


ca_sub="isecnet.com"
grpc_server_sub="isecnet.server.com"
grpc_client_sub="isecnet.client.com"
download_https_sub="isecnet.download.com"
partner_https_sub="isecnet.partner.com"

case $1 in
"commit")
  git add .
  git commit -m "$2"
  git pull
  git push
;;
"docker")
  docker run --rm -ti -v `pwd`:/opt/ccs golang_build:v1.0.0 sh -c "cd /opt/ccs && $2"
;;
"webopenssl")
# 生成.key  私钥文件
openssl genrsa -out key/web.key 2048

# 生成.csr 证书签名请求文件
openssl req -new -key key/web.key -out key/web.csr \
	-subj "/CN=$download_https_sub" \
	-reqexts SAN \
	-config <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.isecnet.com,DNS:*.isecnet.x.com"))

# 签名生成.crt 证书文件
openssl x509 -req -days 3650 \
   -in key/web.csr -out key/web.crt \
   -CA key/ca.crt -CAkey key/ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.isecnet.com,DNS:*.isecnet.x.com"))
;;
"openssl")
# 生成.key  私钥文件
openssl genrsa -out key/ca.key 2048

# 生成.csr 证书签名请求文件
openssl req -new -key key/ca.key -out key/ca.csr  -subj "/CN=$ca_sub"

# 自签名生成.crt 证书文件
openssl req -new -x509 -days 3650 -key key/ca.key -out key/ca.crt  -subj "/CN=$ca_sub"

# 生成.key  私钥文件
openssl genrsa -out key/server.key 2048

# 生成.csr 证书签名请求文件
openssl req -new -key key/server.key -out key/server.csr \
	-subj "/CN=$grpc_server_sub" \
	-reqexts SAN \
	-config <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.isecnet.com,DNS:*.isecnet.x.com"))

# 签名生成.crt 证书文件
openssl x509 -req -days 3650 \
   -in key/server.csr -out key/server.crt \
   -CA key/ca.crt -CAkey key/ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.isecnet.com,DNS:*.isecnet.x.com"))


# 生成.key  私钥文件
openssl genrsa -out key/client.key 2048

# 生成.csr 证书签名请求文件
openssl req -new -key key/client.key -out key/client.csr \
	-subj "/CN=$grpc_client_sub" \
	-reqexts SAN \
	-config <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.isecnet.com,DNS:*.isecnet.x.com"))

# 签名生成.crt 证书文件
openssl x509 -req -days 3650 \
   -in key/client.csr -out key/client.crt \
   -CA key/ca.crt -CAkey key/ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.isecnet.com,DNS:*.isecnet.x.com"))

./shell.sh rsa
./shell.sh webopenssl
;;
"rsa")
  openssl genrsa -out key/rsa.key 2048
  openssl rsa -in key/rsa.key -pubout -out key/rsa.pub
;;
"restart")
  docker-compose down
  rm -rf testdata/*
  make
  export LICENSEADDR="license:25535"
  export DOWNLOADADDR="download:443"
  export MYSQLADDR="mysql:3306"
  export REDISADDE="redis:6379"
  docker-compose config
  docker-compose up -d
;;
esac

docker-compose down
docker-compose up -d

