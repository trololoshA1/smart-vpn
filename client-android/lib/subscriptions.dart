import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'add_subscription.dart';
import 'edit_subscription.dart';

class SubscriptionsScreen extends StatefulWidget {
  static const platform = MethodChannel("smartvpn/tun");

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
    final result = await SubscriptionsScreen.platform.invokeMethod("getSubs");
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

          return Card(
            margin: EdgeInsets.all(10),
            child: InkWell(
              onTap: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (_) => EditSubscriptionScreen(index: i, sub: s),
                  ),
                ).then((_) => loadSubs());
              },
              child: Padding(
                padding: EdgeInsets.all(15),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      s["name"],
                      style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                    ),
                    SizedBox(height: 8),
                    Text("URL: ${s["url"]}"),
                    Text("Автообновление: ${s["auto_update"] ? "Включено" : "Выключено"}"),
                    Text("Интервал: ${s["update_interval"]} мин"),
                    Text("Последнее обновление: ${s["last_update"]}"),
                    Text("Узлов: ${s["nodes"].length}"),
                  ],
                ),
              ),
            ),
          );
        },
      ),

      floatingActionButton: FloatingActionButton(
        child: Icon(Icons.add),
        onPressed: () {
          Navigator.push(
            context,
            MaterialPageRoute(builder: (_) => AddSubscriptionScreen()),
          ).then((_) => loadSubs());
        },
      ),
    );
  }
}