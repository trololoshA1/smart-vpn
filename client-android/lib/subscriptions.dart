import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

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
            child: ListTile(
              title: Text(s["name"]),
              subtitle: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text("URL: ${s["url"]}"),
                  Text("Автообновление: ${s["auto_update"] ? "Вкл" : "Выкл"}"),
                  Text("Интервал: ${s["update_interval"]} мин"),
                  Text("Узлов: ${s["nodes"].length}"),
                  Text("Последнее обновление: ${s["last_update"]}"),
                ],
              ),
              trailing: Icon(Icons.arrow_forward),
              onTap: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (_) => SubscriptionDetailsScreen(sub: s),
                  ),
                );
              },
            ),
          );
        },
      ),
    );
  }
}

class SubscriptionDetailsScreen extends StatelessWidget {
  final Map sub;

  SubscriptionDetailsScreen({required this.sub});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(sub["name"])),
      body: ListView(
        children: [
          ListTile(title: Text("URL: ${sub["url"]}")),
          ListTile(title: Text("Автообновление: ${sub["auto_update"]}")),
          ListTile(title: Text("Интервал: ${sub["update_interval"]} мин")),
          ListTile(title: Text("Узлов: ${sub["nodes"].length}")),
          Divider(),
          ...List.generate(sub["nodes"].length, (i) {
            final n = sub["nodes"][i];
            return ListTile(
              title: Text(n["name"]),
              subtitle: Text("${n["region"]} — ${n["last_ping"]} ms"),
            );
          })
        ],
      ),
    );
  }
}