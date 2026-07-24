package com.smartvpn

import android.app.Service
import android.content.Intent
import android.os.IBinder
import androidx.work.ExistingPeriodicWorkPolicy
import androidx.work.PeriodicWorkRequestBuilder
import androidx.work.WorkManager
import java.util.concurrent.TimeUnit

class TunService : Service() {

    override fun onCreate() {
        super.onCreate()

        // Запускаем автообновление каждые 15 минут
        val workRequest =
            PeriodicWorkRequestBuilder<AutoUpdateWorker>(15, TimeUnit.MINUTES)
                .build()

        WorkManager.getInstance(this).enqueueUniquePeriodicWork(
            "auto_update_subs",
            ExistingPeriodicWorkPolicy.UPDATE,
            workRequest
        )
    }

    override fun onBind(intent: Intent?): IBinder? {
        return null
    }
}