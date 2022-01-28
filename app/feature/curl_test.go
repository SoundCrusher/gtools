package feature

import "testing"

func Test_matchCurl(t *testing.T) {
	buf := `
		curl -k -L -X POST "http://www.chacuo.net/"  \
		 -H "Host: \"www.chacuo.net"  \
		 -H "Connection: keep-alive"  \
		 -H "Accept: */*"  \
		 -H "Origin: http://www.chacuo.net"  \
		 -H "X-Requested-With: XMLHttpRequest"  \
		 -H "User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.62 Safari/537.36"  \
		 -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8"  \
		 -H "Referer: http://24mail.chacuo.net/"  \
		 -H "Accept-Language: zh-CN,zh;q=0.9"  \
		 --data "data=你好，chacuo.net&hello=test"  \
		 --compressed
	`
	matchCurl(buf)
}
