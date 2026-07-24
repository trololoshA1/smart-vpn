import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

void main() {
  runApp(SmartVPNApp());
}

class SmartVPNApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'SmartVPN',
      home: MainScreen(),
    );
  }
}

class MainScreen extends StatefulWidget {
  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  static const platform = MethodChannel("smartvpn/tun");

  bool connected = false;

  Future<void> connectVPN() async {
    final fd = await platform.invokeMethod("startTun");
    await platform.invokeMethod("startCore", {"fd": fd});

    setState(() {
      connected = true;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("SmartVPN")),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              connected ? "VPN подключен" : "VPN отключен",
              style: TextStyle(fontSize: 22),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                if (!connected) connectVPN();
              },
              child: Text("Подключить"),
            ),
          ],
        ),
      ),
    );
  }
}