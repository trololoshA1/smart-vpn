package com.smartvpn

import android.content.Context
import androidx.work.Worker
import androidx.work.WorkerParameters
import mobile.Mobile

class AutoUpdateWorker(appContext: Context, workerParams: WorkerParameters)
    : Worker(appContext, workerParams) {

    override fun doWork(): Result {
        try {
            // Обновляем все подписки
            val subs = Mobile.getSubscriptions()
            for (i in subs.indices) {
                try {
                    Mobile.updateSubscription(i)
                } catch (_: Exception) {}
            }
        } catch (_: Exception) {
            return Result.retry()
        }

        return Result.success()
    }
}