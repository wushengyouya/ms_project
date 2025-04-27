# 全局变量
TOKEN := eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU5MjA3NzAsInRva2VuIjoiMTAwMyJ9.oTeUpww5B_CTIqNdlI9LGaBriDx_2Sabt_xpc3_L3Ug
ADDR := http://127.0.0.1
captcha:
	curl -X POST -d mobile=15545556677 $(ADDR)/project/login/getCaptcha
register:
	curl -X POST -d 'account=123456&email=100aaa@qq.com&name=mikasa&password=123456&password2=123456&mobile=15545556677&captcha=123456' \
		$(ADDR)/project/login/register
login:
	curl -X POST -d 'account=mikasa&password=123456' $(ADDR)/project/login
index:
	curl -X POST -H 'Authorization:bearer ' \
		 $(ADDR)/project/index
selfList:
	curl -X POST -H 'Authorization:bearer $(TOKEN)' \
		 http://127.0.0.1/project/selfList
getOrgList:
	curl -X POST -H 'Authorization:bearer $(TOKEN)' \
		 $(ADDR)/project/organization/_getOrgList

# 生成自签名ssl证书   
ssl:
	openssl req -x509 -newkey rsa:4096 -nodes -keyout localhost.key -out localhost.crt -days 365 \
	  -subj "//CN=localhost" \
	  -addext "subjectAltName=DNS:localhost,IP:127.0.0.1"

check-ssl:
	openssl rsa -noout -modulus -in msproject-data/ssl/localhost.key | openssl md5
	openssl x509 -noout -modulus -in msproject-data/ssl/localhost.crt | openssl md5
# 验证自签证书是否正常
vk:
	curl -vk https://localhost