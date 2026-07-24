package com.smartvpn

import io.flutter.embedding.android.FlutterActivity
import io.flutter.embedding.engine.FlutterEngine
import io.flutter.plugin.common.MethodChannel

class MainActivity : FlutterActivity() {

    private val CHANNEL = "smartvpn/tun"
    private var tunService: TunService? = null

    override fun configureFlutterEngine(flutterEngine: FlutterEngine) {
        super.configureFlutterEngine(flutterEngine)

        MethodChannel(flutterEngine.dartExecutor.binaryMessenger, CHANNEL)
            .setMethodCallHandler { call, result ->
                when (call.method) {
                    "startTun" -> {
                        tunService = TunService()
                        val fd = tunService!!.startTun()
                        result.success(fd)
                    }
                    "startCore" -> {
                        val fd = call.argument<Int>("fd")!!
                        mobile.NewAndroidTun(fd)
                        result.success(true)
                    }
                }
            }
    }
}