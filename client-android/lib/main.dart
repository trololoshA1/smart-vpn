import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class MainScreen extends StatefulWidget {
  static const platform = MethodChannel("smartvpn/tun");

  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  String status = "Отключено";

  Future<void> startVpn() async {
    try {
      await MainScreen.platform.invokeMethod("startVpn");
      setState(() => status = "VPN запущен");
    } catch (e) {
      setState(() => status = "Ошибка: $e");
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("SmartVPN")),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text("Статус: $status", style: TextStyle(fontSize: 20)),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: startVpn,
              child: Text("Запустить VPN"),
            )
          ],
        ),
      ),
    );
  }
}