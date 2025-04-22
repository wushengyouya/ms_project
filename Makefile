captcha:
	curl -X POST -d mobile=15545556677 http://127.0.0.1/project/login/getCaptcha
register:
	curl -X POST -d 'account=123456&email=100aaa@qq.com&name=mikasa&password=123456&password2=123456&mobile=15545556677&captcha=123456' \
		http://127.0.0.1/project/login/register
login:
	curl -X POST -d 'account=mikasa&password=123456' http://127.0.0.1/project/login
index:
	curl -X POST -H 'Authorization:bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU5MjA3NzAsInRva2VuIjoiMTAwMyJ9.oTeUpww5B_CTIqNdlI9LGaBriDx_2Sabt_xpc3_L3Ug' \
		 http://127.0.0.1/project/index
selfList:
	curl -X POST -H 'Authorization:bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU5MjA3NzAsInRva2VuIjoiMTAwMyJ9.oTeUpww5B_CTIqNdlI9LGaBriDx_2Sabt_xpc3_L3Ug' \
		 http://127.0.0.1/project/selfList
getOrgList:
	curl -X POST -H 'Authorization:bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU5MjA3NzAsInRva2VuIjoiMTAwMyJ9.oTeUpww5B_CTIqNdlI9LGaBriDx_2Sabt_xpc3_L3Ug' \
		 http://127.0.0.1/project/organization/_getOrgList