// Copyright (c) 2019, the Dart project authors.  Please see the AUTHORS file
// for details. All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

import 'dart:async' show FutureOr;
import 'dart:ffi';
import 'dart:io' show Platform;

import 'package:ffi/ffi.dart';

import 'native_add_golang_bindings_generated.dart';

String _platformPath(String name, String path) {
  if (Platform.isLinux || Platform.isAndroid || Platform.isFuchsia) {
    return "${path}lib$name.so";
  }
  if (Platform.isMacOS) return "${path}lib$name.dylib";
  if (Platform.isWindows) return "$path$name.dll";
  throw Exception("Platform not implemented");
}

DynamicLibrary dlopenPlatformSpecific(String name, {String path = ""}) {
  String fullPath = _platformPath(name, path);
  return DynamicLibrary.open(fullPath);
}

extension StringGoLangExt on String {
  R toGoString<R>(final R Function(GoString goString) block) {
    final String dartString = this;
    // 1. 将 Dart String 转换为 C 风格的 UTF-8 字符串。
    //    `toNativeUtf8` 会在原生堆上分配内存，这块内存必须被释放。
    final Pointer<Utf8> cString = dartString.toNativeUtf8();

    // 2. 为 GoString 结构体分配内存并填充其字段。
    //    这块内存也必须被释放。
    final Pointer<GoString> pGoString = calloc<GoString>()
      ..ref.p = cString
      ..ref.n = cString.length; // .length 是不包含 \0 的字节长度

    try {
      // 3. 调用 Go 函数，传递结构体的引用。
      return block(pGoString.ref);
    } finally {
      // 4. 使用 try...finally 来确保原生内存总是被释放，避免内存泄漏。
      //    必须释放为 C 字符串和 GoString 结构体分配的两块内存。
      calloc
        ..free(cString)
        ..free(pGoString);
    }
  }
}
