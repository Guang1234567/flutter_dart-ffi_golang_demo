# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.20)

project(native_add_library VERSION 0.0.1 LANGUAGES C)

# 1. 计算相关变量
set(GO_LIB        nativeffi)
set(GO_TARGET     ${CMAKE_LIBRARY_OUTPUT_DIRECTORY}/lib${GO_LIB}.so)
set(GO_HEADER     ${CMAKE_LIBRARY_OUTPUT_DIRECTORY}/lib${GO_LIB}.h)
set(LOCAL_HEADER  ${CMAKE_CURRENT_SOURCE_DIR}/lib${GO_LIB}.h)
set(DART_BINDING  ${CMAKE_CURRENT_SOURCE_DIR}/../lib/native_add_golang_bindings_generated.dart)

if(ANDROID_ABI STREQUAL "arm64-v8a")
    set(GOARCH arm64)
elseif(ANDROID_ABI STREQUAL "armeabi-v7a")
    set(GOARCH arm)
elseif(ANDROID_ABI STREQUAL "x86_64")
    set(GOARCH amd64)
elseif(ANDROID_ABI STREQUAL "x86")
    set(GOARCH 386)
endif()

# 2. 收集所有 .go 源
file(GLOB_RECURSE GO_SOURCES CONFIGURE_DEPENDS "*.go")

# 3. 生成 .so + .h
add_custom_command(
    OUTPUT  ${GO_TARGET} ${GO_HEADER}
    COMMAND ${CMAKE_COMMAND} -E echo "go 编译工具相关信息..."
    COMMAND go version
    COMMAND ${CMAKE_COMMAND} -E echo "正在编译 Go 共享库..."
    COMMAND ${CMAKE_COMMAND} -E env
            "CGO_CFLAGS=--target=${ANDROID_LLVM_TRIPLE} ${CMAKE_C_FLAGS} -Wno-unused-command-line-argument"
            "CGO_LDFLAGS=--target=${ANDROID_LLVM_TRIPLE} ${CMAKE_SHARED_LINKER_FLAGS} -Wl,-soname,lib${GO_LIB}.so -Wl,-z,max-page-size=16384"
            "CC=${ANDROID_C_COMPILER}"
            "GOARCH=${GOARCH}"
            "GOOS=android"
            "CGO_ENABLED=1"
            "GOTRACEBACK=all"
            # --------------------------------------
            go build -o ${GO_TARGET}
                     -buildmode c-shared -a -v
    COMMAND ${CMAKE_COMMAND} -E echo "编译完成: ${GO_TARGET} 和 ${GO_HEADER}"
    DEPENDS ${GO_SOURCES}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Building ${GO_LIB} for ${ANDROID_ABI} ..."
    VERBATIM
)

# 4. 复制头文件到源码树
add_custom_command(
    OUTPUT  ${LOCAL_HEADER}
    COMMAND ${CMAKE_COMMAND} -E echo "复制头文件到本地..."
    COMMAND ${CMAKE_COMMAND} -E copy_if_different ${GO_HEADER} ${LOCAL_HEADER}
    DEPENDS ${GO_HEADER}
    COMMENT "Copying header to source tree ..."
    VERBATIM
)

# 5. 生成 Dart FFI 绑定
add_custom_command(
    OUTPUT  ${DART_BINDING}
    COMMAND ${CMAKE_COMMAND} -E echo "生成 dart-go binding 文件..."
    COMMAND cd ${CMAKE_CURRENT_SOURCE_DIR}/.. &&
            dart run ffigen --config ffigen_golang.yaml
    DEPENDS ${LOCAL_HEADER}
    COMMENT "Generating Dart FFI bindings ..."
    VERBATIM
)

# 6. 顶层目标，默认 ALL
add_custom_target(${GO_LIB} ALL
    DEPENDS ${GO_TARGET} ${GO_HEADER} ${LOCAL_HEADER} ${DART_BINDING}
)

# 7. 让 cmake clean 时一并删除这些产物
set_property(DIRECTORY APPEND PROPERTY
             ADDITIONAL_CLEAN_FILES
             ${GO_TARGET} ${GO_HEADER} ${LOCAL_HEADER} ${DART_BINDING})

message("CMAKE_LIBRARY_OUTPUT_DIRECTORY: ${CMAKE_LIBRARY_OUTPUT_DIRECTORY}")


SET(C_LIB native_add)

add_library(${C_LIB} SHARED
  "native_add.c"
)

set_target_properties(${C_LIB} PROPERTIES
  PUBLIC_HEADER native_add.h
  OUTPUT_NAME "${C_LIB}"
)

target_compile_definitions(${C_LIB} PUBLIC DART_SHARED_LIB)

if (ANDROID)
  # Support Android 15 16k page size
  target_link_options(${C_LIB} PRIVATE "-Wl,-z,max-page-size=16384")
endif()
