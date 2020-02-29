build:
	go mod tidy
	#检测race时需打开以下注释
	#go build --ldflags "-extldflags -static" -o run -race cmd/main/main.go 
	go build --ldflags "-extldflags -static" -o ucm_dev cmd/main/main.go
	nohup ./ucm_dev > /dev/null 2>&1 &
# 自动提示
# 要在vim中自动提示，请先运行make autocompletor
autocompletor:
	go install ./pkg/...
