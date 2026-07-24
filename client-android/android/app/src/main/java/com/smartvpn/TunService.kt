package com.smartvpn

import android.net.VpnService
import android.os.ParcelFileDescriptor

class TunService : VpnService() {

    var tunFd: ParcelFileDescriptor? = null

    fun startTun(): Int {
        val builder = Builder()
        builder.addAddress("10.0.0.2", 32)
        builder.addRoute("0.0.0.0", 0)
        builder.setSession("SmartVPN")

        tunFd = builder.establish()
        return tunFd!!.fd
    }
}