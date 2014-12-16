sudo docker run -v $GOPATH/src:/src mobile /bin/bash -c 'cd /src/github.com/tbruyelle/snake && ./make.bash'
adb uninstall com.kamosoft.snake
adb install bin/nativeactivity-debug.apk
