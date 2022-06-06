import 'package:flutter/material.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/models/race.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/widgets/character_icon.dart';
import 'package:provider/provider.dart';
import 'package:timeago/timeago.dart' as timeago;

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
          itemBuilder: (_, i) => _RaceTile(
            race: notifier.races[i],
            players: notifier.getPlayersFromResults(notifier.races[i].results),
          ),
        ),
      ),
    );
  }
}

class _RaceTile extends StatelessWidget {
  final Race race;
  final List<Player> players;

  const _RaceTile({Key? key, required this.race, required this.players}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 8.0),
      child: InkWell(
        onTap: () {},
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Align(alignment: Alignment.bottomRight, child: Text(timeago.format(race.date))),
              Wrap(
                spacing: 4.0,
                children: players
                    .map((player) => Column(
                          children: [
                            Stack(
                              children: [
                                CharacterIcon(icon: player.icon, size: 72.0),
                                Text('${players.indexOf(player) + 1}', style: Theme.of(context).textTheme.headline5)
                              ],
                            ),
                            SizedBox(
                              width: 72.0,
                              child: Center(child: Text(player.name, overflow: TextOverflow.ellipsis, maxLines: 1)),
                            ),
                          ],
                        ))
                    .toList(),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
