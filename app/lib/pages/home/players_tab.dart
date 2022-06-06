import 'package:flutter/material.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/pages/create_player_page.dart';
import 'package:kart_app/pages/player_page.dart';
import 'package:kart_app/widgets/character_icon.dart';
import 'package:provider/provider.dart';

class PlayersTab extends StatelessWidget {
  const PlayersTab({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return RefreshIndicator(
      onRefresh: context.read<AppNotifier>().refreshPlayers,
      child: Consumer<AppNotifier>(
        builder: (context, notifier, _) => ListView.builder(
          padding: const EdgeInsets.only(top: 16.0, left: 16.0, right: 16.0, bottom: 100.0),
          itemCount: notifier.players.length + 1,
          itemBuilder: (_, i) {
            if (i == notifier.players.length) {
              return _CreatePlayerButton();
            } else {
              return _RankingTile(player: notifier.players[i]);
            }
          },
        ),
      ),
    );
  }
}

class _RankingTile extends StatelessWidget {
  final Player player;

  const _RankingTile({Key? key, required this.player}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListTile(
      onTap: () => Navigator.push(context, MaterialPageRoute(builder: (_) => PlayerPage(id: player.id))),
      title: Text(player.name),
      subtitle: Text(player.rating.toInt().toString()),
      leading: CharacterIcon(icon: player.icon),
      trailing: Text(
        player.rank.toString(),
        style: Theme.of(context).textTheme.headline5,
      ),
    );
  }
}

class _CreatePlayerButton extends StatelessWidget {
  const _CreatePlayerButton({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: TextButton(
          onPressed: () => Navigator.push(context, MaterialPageRoute(builder: (_) => CreatePlayerPage())),
          child: Text('Ajouter un joueur'),
        ),
      ),
    );
  }
}
