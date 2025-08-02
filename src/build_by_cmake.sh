#!/bin/bash

# https://cmake.org/cmake/help/latest/manual/cmake-toolchains.7.html#id24

rm -rf /Users/lihanguang/dev_kit/src_code/flutter_src/native_add/.cxx
rm -rf /Users/lihanguang/dev_kit/src_code/flutter_src/native_add/example/build/native_add/intermediates/cxx

/Users/lihanguang/.mise/installs/android-sdk/19.0/cmake/3.22.1/bin/cmake \
        -H/Users/lihanguang/dev_kit/src_code/flutter_src/native_add/src \
        -DCMAKE_SYSTEM_NAME=Android \
        -DCMAKE_EXPORT_COMPILE_COMMANDS=ON \
        -DCMAKE_SYSTEM_VERSION=24 \
        -DANDROID_PLATFORM=android-24 \
        -DANDROID_ABI=arm64-v8a \
        -DCMAKE_ANDROID_ARCH_ABI=arm64-v8a \
        -DANDROID_NDK=/Users/lihanguang/.mise/installs/android-sdk/19.0/ndk/27.2.12479018 \
        -DCMAKE_ANDROID_NDK=/Users/lihanguang/.mise/installs/android-sdk/19.0/ndk/27.2.12479018 \
        -DCMAKE_TOOLCHAIN_FILE=/Users/lihanguang/.mise/installs/android-sdk/19.0/ndk/27.2.12479018/build/cmake/android.toolchain.cmake \
        -DCMAKE_MAKE_PROGRAM=/Users/lihanguang/.mise/installs/android-sdk/19.0/cmake/3.22.1/bin/ninja \
        -DCMAKE_LIBRARY_OUTPUT_DIRECTORY=/Users/lihanguang/dev_kit/src_code/flutter_src/native_add/example/build/native_add/intermediates/cxx/Debug/1y39vy3s/obj/arm64-v8a \
        -DCMAKE_RUNTIME_OUTPUT_DIRECTORY=/Users/lihanguang/dev_kit/src_code/flutter_src/native_add/example/build/native_add/intermediates/cxx/Debug/1y39vy3s/obj/arm64-v8a \
        -DCMAKE_BUILD_TYPE=Debug \
        -B/Users/lihanguang/dev_kit/src_code/flutter_src/native_add/android/.cxx/Debug/1y39vy3s/arm64-v8a \
        -GNinja

cd /Users/lihanguang/dev_kit/src_code/flutter_src/native_add/android/.cxx/Debug/1y39vy3s/arm64-v8a

#ninja "native_add"
#ninja "nativeffi"
ninja "nativeffi" "native_add"
#ninja "all"

ls -al /Users/lihanguang/dev_kit/src_code/flutter_src/native_add/example/build/native_add/intermediates/cxx/Debug/1y39vy3s/obj/arm64-v8a