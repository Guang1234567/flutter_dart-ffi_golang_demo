NDK_GO_ARCH_x86 := 386
NDK_GO_ARCH_x86_64 := amd64
NDK_GO_ARCH_arm := arm
NDK_GO_ARCH_arm64 := arm64
NDK_GO_ARCH_mips := mipsx
NDK_GO_ARCH_mips64 := mips64x

CLANG_FLAGS := --target=$(ANDROID_LLVM_TRIPLE)

$(info --- Printing CLANG_FLAGS from makefile.mk ---)
$(info CLANG_FLAGS: $(CLANG_FLAGS))

export CGO_CFLAGS := $(CLANG_FLAGS) $(CFLAGS)
export CGO_LDFLAGS := $(CLANG_FLAGS) $(LDFLAGS) -Wl,-soname=${SONAME}
export CC := $(ANDROID_C_COMPILER)
export GOARCH := $(NDK_GO_ARCH_$(ANDROID_ARCH_NAME))
export GOOS := android
export CGO_ENABLED := 1
export GOTRACEBACK := all


# 目标文件名
TARGET = ${BUILDDIR}/${SONAME}

# 动态查找所有.c文件（包括子目录）
SRCS := $(shell find . -name "*.go")

BASE_SONAME := $(shell echo "${SONAME}" | sed 's/\.so$$//')
# 将.so文件转换为对应的.h文件（保持目录结构）
HEADER := $(TARGET:.so=.h)
LOCAL_HEADER := ./$(BASE_SONAME).h
DART_BINDING_FILE := ../lib/native_add_golang_bindings_generated.dart


# 默认目标
.PHONY: all
all: $(DART_BINDING_FILE)

$(DART_BINDING_FILE): $(LOCAL_HEADER)
	@echo "生成 dart-go binding 文件..."
	cd .. && dart run ffigen --config "ffigen_golang.yaml"

$(LOCAL_HEADER): $(HEADER)
	@echo "复制头文件到本地..."
	cp $< $@

$(TARGET) $(HEADER): $(SRCS)
	@echo "go 编译工具相关信息..."
	go version
	which go
	@echo "正在编译 Go 共享库..."
	go build -o "$(BUILDDIR)/${SONAME}" -buildmode c-shared -a -v
	@echo "编译完成: $(TARGET) 和 $(HEADER)"

# 清理生成的文件
.PHONY: clean
clean:
	rm -f $(TARGET)
	rm -f $(HEADER)
	rm -f ./$(BASE_SONAME).h
	rm -f ./$(BASE_SONAME).so
	rm -f ../lib/native_add_golang_bindings_generated.dart
	@echo "清理完成"

.DELETE_ON_ERROR: