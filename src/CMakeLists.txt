# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(native_add_library VERSION 0.0.1 LANGUAGES C)


SET(GO_LIB nativeffi)

add_custom_target(${GO_LIB} ALL
        WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}"
        COMMENT "Building ${GO_LIB} for ${ANDROID_ABI}"
        VERBATIM COMMAND make
        ANDROID_ARCH_NAME=${ANDROID_ARCH_NAME}
        ANDROID_C_COMPILER=${ANDROID_C_COMPILER}
        ANDROID_TOOLCHAIN_ROOT=${ANDROID_TOOLCHAIN_ROOT}
        ANDROID_LLVM_TRIPLE=${ANDROID_LLVM_TRIPLE}
        ANDROID_SYSROOT=${ANDROID_SYSROOT}
        CFLAGS=${CMAKE_C_FLAGS}\ -Wno-unused-command-line-argument
        LDFLAGS=${CMAKE_SHARED_LINKER_FLAGS}\ -Wl,-z,max-page-size=16384
        BUILDDIR=${CMAKE_LIBRARY_OUTPUT_DIRECTORY}
        SONAME=lib${GO_LIB}.so
)

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
