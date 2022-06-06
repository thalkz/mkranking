import 'package:flutter/material.dart';
import 'package:kart_app/models/race.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/widgets/character_icon.dart';
import 'package:provider/provider.dart';
import 'package:timeago/timeago.dart' as timeago;

class RaceTile extends StatelessWidget {
  final Race race;

  const RaceTile({Key? key, required this.race}) : super(key: key);

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
                children: race.results.map((playerId) {
                  final player = context.read<AppNotifier>().getPlayer(playerId);
                  return Column(
                    children: [
                      Stack(
                        children: [
                          CharacterIcon(icon: player.icon, size: 72.0),
                          Text(playerId.toString(), style: Theme.of(context).textTheme.headline5)
                        ],
                      ),
                      SizedBox(
                        width: 72.0,
                        child: Center(child: Text(player.name, overflow: TextOverflow.ellipsis, maxLines: 1)),
                      ),
                    ],
                  );
                }).toList(),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
