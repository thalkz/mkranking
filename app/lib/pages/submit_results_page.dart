import 'package:flutter/material.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/widgets/character_icon.dart';
import 'package:provider/provider.dart';

import '../notifiers/submit_results_notifier.dart';

class SubmitResultsPage extends StatelessWidget {
  const SubmitResultsPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider<SubmitResultsNotifier>(
      create: (context) => SubmitResultsNotifier(players: context.read<AppNotifier>().players),
      child: Scaffold(
        appBar: AppBar(title: Text('Nouvelle course')),
        body: Consumer<SubmitResultsNotifier>(
          builder: ((context, notifier, _) {
            if (notifier.ratingsDiff.isNotEmpty) {
              return _ShowRatingsDiffWidget(ratingsDiff: notifier.ratingsDiff);
            } else if (notifier.rankedParticipants.isNotEmpty) {
              return _RankParticipantsList(participants: notifier.rankedParticipants);
            } else {
              return _SelectPlayersList(selectedPlayers: notifier.selectedPlayers);
            }
          }),
        ),
        floatingActionButton: Consumer<SubmitResultsNotifier>(builder: (context, notifier, _) {
          if (notifier.ratingsDiff.isNotEmpty) {
            return FloatingActionButton.extended(
              onPressed: () => Navigator.pop(context),
              icon: Icon(Icons.navigate_next),
              label: Text('Ok'),
            );
          } else if (notifier.rankedParticipants.isNotEmpty) {
            return FloatingActionButton.extended(
              onPressed: notifier.submitResults,
              icon: Icon(Icons.check),
              label: Text('Valider'),
            );
          } else {
            return FloatingActionButton(
              onPressed: notifier.confirmParticipants,
              child: Icon(Icons.navigate_next),
            );
          }
        }),
      ),
    );
  }
}

class _RankParticipantsList extends StatelessWidget {
  final List<Player> participants;

  const _RankParticipantsList({Key? key, required this.participants}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ReorderableListView(
      buildDefaultDragHandles: false,
      padding: const EdgeInsets.only(top: 16.0, left: 16.0, right: 16.0, bottom: 100.0),
      children: participants
          .map((player) => ListTile(
                key: Key('${player.id}'),
                leading: Text(
                  '${participants.indexOf(player) + 1}',
                  style: Theme.of(context).textTheme.headline4,
                ),
                title: Text(player.name),
                trailing: ReorderableDragStartListener(
                  index: participants.indexOf(player),
                  child: Icon(Icons.reorder_rounded),
                ),
              ))
          .toList(),
      onReorder: context.read<SubmitResultsNotifier>().reorderParticipants,
    );
  }
}

class _SelectPlayersList extends StatelessWidget {
  final Map<Player, bool> selectedPlayers;

  const _SelectPlayersList({Key? key, required this.selectedPlayers}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: const EdgeInsets.only(top: 16.0, left: 16.0, right: 16.0, bottom: 100.0),
      children: selectedPlayers.entries
          .map((entry) => CheckboxListTile(
                title: Row(
                  children: [
                    CharacterIcon(icon: entry.key.icon),
                    SizedBox(width: 8.0),
                    Expanded(child: Text(entry.key.name)),
                  ],
                ),
                value: entry.value,
                onChanged: (selected) => context.read<SubmitResultsNotifier>().selectPlayer(entry.key, selected),
              ))
          .toList(),
    );
  }
}

class _ShowRatingsDiffWidget extends StatelessWidget {
  final Map<Player, double> ratingsDiff;

  const _ShowRatingsDiffWidget({Key? key, required this.ratingsDiff}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: const EdgeInsets.only(top: 16.0, left: 16.0, right: 16.0, bottom: 100.0),
      children: ratingsDiff.entries
          .map((entry) => ListTile(
                title: Text(entry.key.name),
                subtitle: Text(entry.value.toInt().toString()),
                leading: CharacterIcon(icon: entry.key.icon),
              ))
          .toList(),
    );
  }
}