# 声明变量
Name="goorigin"
# target 命令 : prerequisites 前置条件
	# @command 命令，且该命令输出不会被打印
default:
	@swag init
	@if [ -f config.yaml ] ; then go build -o ${Name} main.go  ;else  echo "no config";fi
clean:
	@if [ -f ${Name} ] ; then rm ${Name}  ; fi
swag:
	@swag init
	echo "http://localhost:8008/swagger/index.html"



