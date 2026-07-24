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
  static const platform = MethodChannel("smartvpn/tun");

  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  bool connected = false;

  Future<void> connectVPN() async {
    final fd = await MainScreen.platform.invokeMethod("startTun");
    await MainScreen.platform.invokeMethod("startCore", {"fd": fd});

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
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (_) => SubscriptionsScreen()),
                );
              },
              child: Text("Подписки"),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (_) => StatusScreen()),
                );
              },
              child: Text("Статус"),
            ),
          ],
        ),
      ),
    );
  }
}

class SubscriptionsScreen extends StatefulWidget {
  @override
  State<SubscriptionsScreen> createState() => _SubscriptionsScreenState();
}

class _SubscriptionsScreenState extends State<SubscriptionsScreen> {
  List subs = [];

  @override
  void initState() {
    super.initState();
    loadSubs();
  }

  Future<void> loadSubs() async {
    final result = await MainScreen.platform.invokeMethod("getSubs");
    setState(() {
      subs = List<Map>.from(result);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Подписки")),
      body: ListView.builder(
        itemCount: subs.length,
        itemBuilder: (context, i) {
          final s = subs[i];
          return ListTile(
            title: Text(s["name"]),
            subtitle: Text("${s["region"]} • ping: ${s["last_ping"]} ms"),
          );
        },
      ),
    );
  }
}

class StatusScreen extends StatefulWidget {
  @override
  State<StatusScreen> createState() => _StatusScreenState();
}

class _StatusScreenState extends State<StatusScreen> {
  String status = "Отключено";

  @override
  void initState() {
    super.initState();
    loadStatus();
  }

  Future<void> loadStatus() async {
    final result = await MainScreen.platform.invokeMethod("status");
    setState(() {
      status = result;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Статус VPN")),
      body: Center(
        child: Text(
          status,
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}