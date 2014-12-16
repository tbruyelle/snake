sudo docker run -v $GOPATH/src:/src mobile /bin/bash -c 'cd /src/github.com/tbruyelle/snake && ./make.bash'
if [[ $? -eq 0  ]] 
then
	adb uninstall com.kamosoft.snake
	adb install bin/nativeactivity-debug.apk
fi
