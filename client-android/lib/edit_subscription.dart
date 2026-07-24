import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class EditSubscriptionScreen extends StatefulWidget {
  final int index;
  final Map sub;

  EditSubscriptionScreen({required this.index, required this.sub});

  static const platform = MethodChannel("smartvpn/tun");

  @override
  State<EditSubscriptionScreen> createState() => _EditSubscriptionScreenState();
}

class _EditSubscriptionScreenState extends State<EditSubscriptionScreen> {
  late TextEditingController nameCtrl;
  late TextEditingController urlCtrl;
  late bool autoUpdate;
  late int interval;

  @override
  void initState() {
    super.initState();
    nameCtrl = TextEditingController(text: widget.sub["name"]);
    urlCtrl = TextEditingController(text: widget.sub["url"]);
    autoUpdate = widget.sub["auto_update"];
    interval = widget.sub["update_interval"];
  }

  Future<void> save() async {
    await EditSubscriptionScreen.platform.invokeMethod("editSubscription", {
      "index": widget.index,
      "name": nameCtrl.text,
      "url": urlCtrl.text,
      "auto": autoUpdate,
      "interval": interval,
    });

    Navigator.pop(context);
  }

  Future<void> delete() async {
    await EditSubscriptionScreen.platform.invokeMethod("removeSubscription", {
      "index": widget.index,
    });

    Navigator.pop(context);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Редактировать подписку")),
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
              onPressed: save,
              child: Text("Сохранить"),
            ),
            SizedBox(height: 10),
            ElevatedButton(
              onPressed: delete,
              style: ElevatedButton.styleFrom(backgroundColor: Colors.red),
              child: Text("Удалить"),
            )
          ],
        ),
      ),
    );
  }
}