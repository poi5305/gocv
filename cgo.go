// +build !customenv,!openvino

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.
/*
# Download prebuilt
curl -L 'https://github.com/poi5305/android-opencv4-gnustl-prebuilt/releases/download/v4.1.2/OpenCV-android-sdk-4.1.2-prebuilt.zip' > OpenCV-android-sdk-4.1.2-prebuilt.zip
unzip OpenCV-android-sdk-4.1.2-prebuilt.zip
*/

/*
#cgo android CXXFLAGS: -I${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/jni/include
#cgo android LDFLAGS: -lopencv_stitching -lopencv_highgui  -lopencv_dnn -lopencv_video -lopencv_videoio  -lopencv_ml -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_photo -lopencv_imgproc -lopencv_core
#cgo android,arm LDFLAGS: -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/libs/armeabi-v7a -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/3rdparty/armeabi-v7a -llog -lz -ljnigraphics -lm
#cgo android,arm LDFLAGS: -llibtiff -llibpng -llibwebp -lquirc -llibprotobuf -lIlmImf -lcpufeatures -llibjasper -littnotify -ltegra_hal -llibjpeg-turbo
#cgo android,arm64 LDFLAGS: -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/libs/arm64-v8a -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/3rdparty/arm64-v8a -llog -lz -ljnigraphics -lmediandk -lm
#cgo android,arm64 LDFLAGS: -llibtiff -llibpng -llibwebp -lquirc -llibprotobuf -lIlmImf -lcpufeatures -llibjasper -littnotify -ltegra_hal -llibjpeg-turbo
#cgo android,386 LDFLAGS: -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/libs/x86 -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/3rdparty/x86 -llog -lz -lm -latomic
#cgo android,386 LDFLAGS: -lippiw -lippicv -llibtiff -llibpng -llibwebp -lquirc -llibprotobuf -lIlmImf -lcpufeatures -llibjasper -littnotify -llibjpeg-turbo
#cgo android,amd64 LDFLAGS: -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/libs/x86_64 -L${SRCDIR}/OpenCV-android-sdk-4.1.2/sdk/native/3rdparty/x86_64 -llog -lz -lm -latomic -ljnigraphics -lmediandk
#cgo android,amd64 LDFLAGS: -lippiw -lippicv -llibtiff -llibpng -llibwebp -lquirc -llibprotobuf -lIlmImf -lcpufeatures -llibjasper -littnotify -llibjpeg-turbo
#cgo !windows,!android pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core412 -lopencv_face412 -lopencv_videoio412 -lopencv_imgproc412 -lopencv_highgui412 -lopencv_imgcodecs412 -lopencv_objdetect412 -lopencv_features2d412 -lopencv_video412 -lopencv_dnn412 -lopencv_xfeatures2d412 -lopencv_plot412 -lopencv_tracking412 -lopencv_img_hash412 -lopencv_calib3d412
*/
import "C"
