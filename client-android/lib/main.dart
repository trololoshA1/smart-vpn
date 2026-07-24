import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class MainScreen extends StatefulWidget {
  static const platform = MethodChannel("smartvpn/tun");

  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  String status = "Отключено";
  Map? bestNode;

  Future<void> connect() async {
    try {
      final node = await MainScreen.platform.invokeMethod("bestNode");
      setState(() => bestNode = Map.from(node));

      await MainScreen.platform.invokeMethod("connect");
      setState(() => status = "Подключено");
    } catch (e) {
      setState(() => status = "Ошибка: $e");
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("SmartVPN")),
      body: Padding(
        padding: EdgeInsets.all(20),
        child: Column(
          children: [
            Text("Статус: $status", style: TextStyle(fontSize: 20)),
            SizedBox(height: 20),
            if (bestNode != null)
              Text(
                "Лучший узел:\n${bestNode!["name"]}\n${bestNode!["address"]}\n${bestNode!["region"]}",
                style: TextStyle(fontSize: 18),
              ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: connect,
              child: Text("Подключиться"),
            )
          ],
        ),
      ),
    );
  }
}