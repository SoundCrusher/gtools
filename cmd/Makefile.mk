APP_NAME = 'gtools'
APP_ICON = '../static/icon/soldering-svg.png'

macos:
	${GOPATH}/bin/fyne package --os darwin --name ${APP_NAME} --icon ${APP_ICON} --appVersion 0.0.1
	rm -rf ../dist/ && mkdir ../dist/
	mv ./${APP_NAME}.app ../dist/${APP_NAME}.app