openssl genrsa -out ca.key 4096
openssl req -new -x509 -key ca.key -sha256 -subj "/C=RU/O=GOPHKEEPER/CN=CA" -days 36500 -out ca.pem

openssl genrsa -out server.key 4096
openssl req -new -key server.key -out server.csr -config server.cfg -extensions  v3_req
openssl x509 -req -in server.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out server.pem -days 36500 -sha256 -extfile server.cfg -extensions  v3_req