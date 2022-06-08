# OPENSSLDIR=/usr/lib/ssl
HOST=127.0.0.1
# server key&crt
openssl req -new \
	-newkey rsa:2048 \
	-nodes -x509 \
	-days 3650 \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=Me/OU=Me/CN=me.org" \
	-addext "subjectAltName = IP:$HOST" \
	-keyout server.key -out server.csr

# client key&crt
# openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -subj "/C=CN/ST=Beijing/L=Beijing/O=Me/OU=Me/CN=me.org" -keyout c.key -out c.csr

