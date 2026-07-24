import 'package:flutter/material.dart';

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
  bool connected = false;

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
                setState(() {
                  connected = !connected;
                });
              },
              child: Text(connected ? "Отключить" : "Подключить"),
            ),
          ],
        ),
      ),
    );
  }
}