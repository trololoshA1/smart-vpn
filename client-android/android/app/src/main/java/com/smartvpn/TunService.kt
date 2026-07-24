package com.smartvpn

import android.app.PendingIntent
import android.app.Service
import android.content.Intent
import android.net.VpnService
import android.os.Build
import android.os.ParcelFileDescriptor
import androidx.work.ExistingPeriodicWorkPolicy
import androidx.work.PeriodicWorkRequestBuilder
import androidx.work.WorkManager
import mobile.Mobile
import java.io.FileInputStream
import java.io.FileOutputStream
import java.nio.ByteBuffer
import java.util.concurrent.TimeUnit

class TunService : VpnService() {

    private var tunInterface: ParcelFileDescriptor? = null
    private var running = false

    override fun onCreate() {
        super.onCreate()

        // Автообновление подписок каждые 15 минут
        val workRequest =
            PeriodicWorkRequestBuilder<AutoUpdateWorker>(15, TimeUnit.MINUTES)
                .build()

        WorkManager.getInstance(this).enqueueUniquePeriodicWork(
            "auto_update_subs",
            ExistingPeriodicWorkPolicy.UPDATE,
            workRequest
        )
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        startVpn()
        return Service.START_STICKY
    }

    private fun startVpn() {
        if (running) return
        running = true

        val builder = Builder()

        builder.setSession("SmartVPN")
        builder.addAddress("10.0.0.2", 32)
        builder.addDnsServer("1.1.1.1")
        builder.addDnsServer("8.8.8.8")

        builder.addRoute("0.0.0.0", 0)

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) {
            builder.setMetered(false)
        }

        tunInterface = builder.establish()

        Thread {
            val input = FileInputStream(tunInterface!!.fileDescriptor)
            val output = FileOutputStream(tunInterface!!.fileDescriptor)

            val buffer = ByteArray(32767)

            while (running) {
                val length = input.read(buffer)
                if (length > 0) {
                    val packet = buffer.copyOfRange(0, length)

                    try {
                        val processed = Mobile.handlePacket(packet)
                        if (processed != null) {
                            output.write(processed)
                        }
                    } catch (_: Exception) {}
                }
            }
        }.start()
    }

    override fun onDestroy() {
        running = false
        tunInterface?.close()
        super.onDestroy()
    }
}