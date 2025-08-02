import 'dart:async';
import 'dylib_utils.dart';
import 'native_add_golang_bindings_generated.dart';
import 'dart:isolate';

NativeAddGolangBindings _bindings = NativeAddGolangBindings(
  dlopenPlatformSpecific('nativeffi'),
);

/// A wrapper function to call the Go `TestPlus` function.
///
/// This function handles the creation and memory management of the `GoString`.
int sum(int a, int b, String c) {
  return c.toGoString((goStr) {
    // 3. 调用 Go 函数，传递 GoString。
    return _bindings.TestPlus(a, b, goStr);
  });
}

Future<int> sumAsync(int a, int b, String c) async {
  // int slowFib(int n) =>
  //     c.toGoString((goStr) => _bindings.TestPlusLongRunning(a, b, goStr));
  //
  // // Compute without blocking current isolate.
  // return Isolate.run(() => slowFib(40));

  return Isolate.run(
    () => c.toGoString((goStr) => _bindings.TestPlusLongRunning(a, b, goStr)),
  );

  // 千万不要如👇这样子写, 这是错误的, 会发生 Panic.
  // return c.toGoString((goStr) async {
  //   return Isolate.run(() => _bindings.TestPlus(a, b, goStr));
  // });
  /*
E/Go      (20594): panic: runtime error: growslice: len out of range
F/libc    (20594): Fatal signal 6 (SIGABRT), code -6 (SI_TKILL) in tid 20906 (DartWorker), pid 20594 (ive_add_example)
Process name is com.example.native_add_example, not key_process
*** *** *** *** *** *** *** *** *** *** *** *** *** *** *** ***
Build fingerprint: 'OnePlus/OnePlus9Pro_CH/OnePlus9Pro:11/RKQ1.201105.002/1638246235130:user/release-keys'
Revision: '0'
ABI: 'arm64'
Timestamp: 2025-08-07 17:10:16+0800
pid: 20594, tid: 20906, name: DartWorker  >>> com.example.native_add_example <<<
uid: 10618
signal 6 (SIGABRT), code -6 (SI_TKILL), fault addr --------
    x0  0000000000000000  x1  00000000000051aa  x2  0000000000000006  x3  0000000000000008
    x4  0000000000000001  x5  0000000000000002  x6  0000000000000000  x7  0000000000000000
    x8  0000000000000083  x9  0000000000000032  x10 0000000000000000  x11 0000007603311030
    x12 0000000008fa2e6e  x13 0000090c27ccea4e  x14 0036c19d302269fc  x15 0000000034155555
    x16 00000074930063a0  x17 0000090c27cc383a  x18 000000749249e000  x19 0000000000005072
    x20 0000007493106cb0  x21 0000004000180008  x22 0000000000000001  x23 0000007300000000
    x24 0000007300008081  x25 0000007493025000  x26 0000000000000000  x27 0000000000000010
    x28 0000004000182380  x29 00000040000528f8
    lr  00000074e019eb00  sp  0000004000052900  pc  00000074e01bfed8  pst 0000000080001000
backtrace:
      #00 pc 00000000000bfed8  /data/app/~~79gKgipQ0teiA4H1B65Crg==/com.example.native_add_example-HEsF8fxCDgizKlLY7pu2WQ==/base.apk!libnativeffi.so (offset 0x28c000) (BuildId: b4e9be180d3ce7df742d6185ae0a598bbec7e21e)
   */
}

Future<int> divideByZeroAsync() async {
  return _bindings.TestDivideByZero();
  //return Isolate.run(() => _bindings.TestDivideByZero());
}
