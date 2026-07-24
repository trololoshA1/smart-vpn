import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class AddSubscriptionScreen extends StatefulWidget {
  static const platform = MethodChannel("smartvpn/tun");

  @override
  State<AddSubscriptionScreen> createState() => _AddSubscriptionScreenState();
}

class _AddSubscriptionScreenState extends State<AddSubscriptionScreen> {
  final nameCtrl = TextEditingController();
  final urlCtrl = TextEditingController();
  bool autoUpdate = true;
  int interval = 10;

  Future<void> add() async {
    await AddSubscriptionScreen.platform.invokeMethod("addSubscription", {
      "name": nameCtrl.text,
      "url": urlCtrl.text,
      "auto": autoUpdate,
      "interval": interval,
    });

    Navigator.pop(context);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Добавить подписку")),
      body: Padding(
        padding: EdgeInsets.all(20),
        child: Column(
          children: [
            TextField(
              controller: nameCtrl,
              decoration: InputDecoration(labelText: "Название"),
            ),
            TextField(
              controller: urlCtrl,
              decoration: InputDecoration(labelText: "URL"),
            ),
            SwitchListTile(
              title: Text("Автообновление"),
              value: autoUpdate,
              onChanged: (v) => setState(() => autoUpdate = v),
            ),
            Row(
              children: [
                Text("Интервал (мин): "),
                DropdownButton<int>(
                  value: interval,
                  items: [5, 10, 15, 30, 60]
                      .map((e) => DropdownMenuItem(value: e, child: Text("$e")))
                      .toList(),
                  onChanged: (v) => setState(() => interval = v!),
                ),
              ],
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: add,
              child: Text("Добавить"),
            )
          ],
        ),
      ),
    );
  }
}