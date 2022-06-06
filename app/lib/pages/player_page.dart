import 'package:flutter/material.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/notifiers/player_notifier.dart';
import 'package:kart_app/widgets/character_icon.dart';
import 'package:kart_app/widgets/race_tile.dart';
import 'package:provider/provider.dart';

class PlayerPage extends StatelessWidget {
  final int id;

  PlayerPage({Key? key, required this.id}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider<PlayerNotifier>(
      create: (context) => PlayerNotifier(playerId: id),
      builder: (context, _) => Scaffold(
        appBar: AppBar(
          title: Consumer<PlayerNotifier>(
            builder: (context, notifier, _) => Text(notifier.player.name),
          ),
          actions: [
            IconButton(
              onPressed: () async {
                await context.read<PlayerNotifier>().deletePlayer();
                context.read<AppNotifier>().refreshAll();
                Navigator.pop(context);
              },
              icon: Icon(Icons.delete_outline),
            ),
          ],
        ),
        body: Consumer<PlayerNotifier>(
          builder: (context, notifier, _) => ListView(
            padding: const EdgeInsets.all(16.0),
            children: [
              CharacterIcon(icon: notifier.player.icon, size: 120.0),
              SizedBox(height: 16.0),
              Center(child: Text('#${notifier.player.rank}', style: Theme.of(context).textTheme.headline5)),
              SizedBox(height: 16.0),
              Center(child: Text('${notifier.player.rating.toInt()} pts')),
              Center(child: Text('${notifier.player.racesCount} courses')),
              SizedBox(height: 16),
              ...notifier.races.map((race) => RaceTile(race: race)).toList()
            ],
          ),
        ),
      ),
    );
  }
}
