import 'package:flutter/material.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/widgets/race_tile.dart';
import 'package:provider/provider.dart';

class RacesTab extends StatelessWidget {
  const RacesTab({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return RefreshIndicator(
      onRefresh: context.read<AppNotifier>().refreshAll,
      child: Consumer<AppNotifier>(
        builder: (context, notifier, _) => ListView.builder(
          padding: const EdgeInsets.only(top: 16.0, left: 16.0, right: 16.0, bottom: 100.0),
          itemCount: notifier.races.length,
          itemBuilder: (_, i) => RaceTile(
            race: notifier.races[i],
          ),
        ),
      ),
    );
  }
}
