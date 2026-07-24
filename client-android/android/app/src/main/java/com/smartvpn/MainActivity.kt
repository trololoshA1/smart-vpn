package com.smartvpn

import android.content.Intent
import io.flutter.embedding.android.FlutterActivity
import io.flutter.embedding.engine.FlutterEngine
import io.flutter.plugin.common.MethodChannel
import mobile.Mobile

class MainActivity: FlutterActivity() {
    private val CHANNEL = "smartvpn/tun"

    override fun configureFlutterEngine(flutterEngine: FlutterEngine) {
        super.configureFlutterEngine(flutterEngine)

        // Запуск фонового сервиса
        startService(Intent(this, TunService::class.java))

        MethodChannel(flutterEngine.dartExecutor.binaryMessenger, CHANNEL).setMethodCallHandler { call, result ->
            when (call.method) {

                "getSubs" -> {
                    val subs = Mobile.getSubscriptions()
                    result.success(subs)
                }

                "addSubscription" -> {
                    val name = call.argument<String>("name")!!
                    val url = call.argument<String>("url")!!
                    val auto = call.argument<Boolean>("auto")!!
                    val interval = call.argument<Int>("interval")!!

                    Mobile.addSubscription(name, url, auto, interval)
                    result.success(true)
                }

                "removeSubscription" -> {
                    val index = call.argument<Int>("index")!!
                    Mobile.removeSubscription(index)
                    result.success(true)
                }

                "editSubscription" -> {
                    val index = call.argument<Int>("index")!!
                    val name = call.argument<String>("name")!!
                    val url = call.argument<String>("url")!!
                    val auto = call.argument<Boolean>("auto")!!
                    val interval = call.argument<Int>("interval")!!

                    Mobile.editSubscription(index, name, url, auto, interval)
                    result.success(true)
                }

                else -> result.notImplemented()
            }
        }
    }
}