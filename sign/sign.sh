# server key&crt
HOST=127.0.0.1

openssl req -new \
	-newkey rsa:2048 \
	-keyout server.pem.key -out server.pem.csr \
	-nodes -x509 \
	-days 3650 \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=Me/OU=Me/CN=me.org" \
	-addext "subjectAltName = IP:$HOST"

openssl x509 -outform der -in server.pem.csr -out certificate.der.cer
