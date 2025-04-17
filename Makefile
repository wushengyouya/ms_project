captcha:
	curl -X POST -d mobile=16614411567 http://127.0.0.1:8080/project/login/getCaptcha
register:
	curl -X POST http://127.0.0.1:8080/project/login/getCaptcha