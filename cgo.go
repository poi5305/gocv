// +build !customenv,!openvino

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo android CXXFLAGS: -I${SRCDIR}/OpenCV-android-sdk-3.4.3/sdk/native/jni/include
#cgo android,arm LDFLAGS: -L${SRCDIR}/OpenCV-android-sdk-3.4.3/sdk/native/libs/armeabi-v7a -lopencv_java3
#cgo android,arm64 LDFLAGS: -L${SRCDIR}/OpenCV-android-sdk-3.4.3/sdk/native/libs/arm64-v8a -lopencv_java3
#cgo android,386 LDFLAGS: -L${SRCDIR}/OpenCV-android-sdk-3.4.3/sdk/native/libs/x86 -lopencv_java3 -lm
#cgo !windows,!android pkg-config: opencv
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core412 -lopencv_face412 -lopencv_videoio412 -lopencv_imgproc412 -lopencv_highgui412 -lopencv_imgcodecs412 -lopencv_objdetect412 -lopencv_features2d412 -lopencv_video412 -lopencv_dnn412 -lopencv_xfeatures2d412 -lopencv_plot412 -lopencv_tracking412 -lopencv_img_hash412 -lopencv_calib3d412
*/
import "C"
