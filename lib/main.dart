import 'dart:ffi' as ffi;
import 'package:ffi/ffi.dart';
import 'native_add_golang_bindings_generated.dart';

NativeAddGolangBindings _bindings = NativeAddGolangBindings(
  ffi.DynamicLibrary.open('libnativeffi.dylib'),
);

/// A wrapper function to call the Go `TestPlus` function.
///
/// This function handles the creation and memory management of the `GoString`.
int testPlus(int a, int b, String c) {
  // 1. 将 Dart String 转换为 C 风格的 UTF-8 字符串。
  //    `toNativeUtf8` 会在原生堆上分配内存，这块内存必须被释放。
  final cNative = c.toNativeUtf8();

  // 2. 为 GoString 结构体分配内存并填充其字段。
  //    这块内存也必须被释放。
  final goString = calloc<GoString>()
    ..ref.p = cNative.cast<ffi.Char>()
    ..ref.n = cNative.length; // .length 是不包含 \0 的字节长度

  try {
    // 3. 调用 Go 函数，传递结构体的引用。
    return _bindings.TestPlus(a, b, goString.ref);
  } finally {
    // 4. 使用 try...finally 来确保原生内存总是被释放，避免内存泄漏。
    //    必须释放为 C 字符串和 GoString 结构体分配的两块内存。
    calloc.free(cNative);
    calloc.free(goString);
  }
}

void main() {
  testPlus(1, 2, 'dart says hello to golang');
}