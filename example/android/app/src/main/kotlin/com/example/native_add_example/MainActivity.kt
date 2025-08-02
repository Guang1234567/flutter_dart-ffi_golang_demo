package com.example.native_add_example

import android.os.Bundle
import android.os.PersistableBundle
import android.util.Log
import android.widget.Toast
import io.flutter.embedding.android.FlutterActivity

import ru.ivanarh.jndcrash.NDCrashError;
import ru.ivanarh.jndcrash.NDCrash;
import ru.ivanarh.jndcrash.NDCrashUtils;
import ru.ivanarh.jndcrash.NDCrashUnwinder;


class MainActivity : FlutterActivity() {

    companion object {
        private const val TAG = "MainActivity"
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        ///*
        val mReportPath = filesDir.absolutePath + "/crash.txt"
        val unwinder = NDCrashUnwinder.stackscan
        val initResult = NDCrash.initializeInProcess(mReportPath, unwinder)
        Log.w(TAG, "In-process signal handler is initialized with result: $initResult unwinder: $unwinder")
        if (initResult != NDCrashError.ok) {
            Toast.makeText(
                applicationContext,
                "Couldn't initialize NDCrash with unwinder $unwinder in in-process mode, error: $initResult",
                Toast.LENGTH_SHORT
            ).show()
        }
        //*/

        super.onCreate(savedInstanceState)
    }

}