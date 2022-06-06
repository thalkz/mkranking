import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/models/player.dart';

class SubmitResultsNotifier with ChangeNotifier {
  Map<Player, bool> selectedPlayers = {};
  List<Player> rankedParticipants = [];
  Map<Player, double> ratingsDiff = {};

  SubmitResultsNotifier({required List<Player> players}) {
    for (final player in players) {
      selectedPlayers.putIfAbsent(player, () => false);
    }
  }

  void _showSnackBar(String text) {
    // TODO Display errors
    // scaffoldKey.currentState?.showSnackBar(SnackBar(
    //   content: Text(text),
    //   backgroundColor: Colors.red,
    // ));
  }

  Future<void> submitResults() async {
    try {
      final rankedIds = rankedParticipants.map((player) => player.id).toList();
      final response = await Api.submitResults(rankedIds);
      print(response.ratingsDiff);
      ratingsDiff = selectedPlayers.map((player, _) => MapEntry(player, response.ratingsDiff[player.id] ?? 0));
      print(ratingsDiff);
      notifyListeners();
    } catch (error) {
      _showSnackBar(error.toString());
      rethrow;
    }
  }

  void confirmParticipants() {
    rankedParticipants = selectedPlayers.entries.where((entry) => entry.value).map((entry) => entry.key).toList();
    notifyListeners();
  }

  void reorderParticipants(int oldIndex, int newIndex) {
    if (oldIndex < newIndex) {
      newIndex -= 1;
    }
    final player = rankedParticipants.removeAt(oldIndex);
    rankedParticipants.insert(newIndex, player);
    notifyListeners();
  }

  void selectPlayer(Player player, bool? selected) {
    selectedPlayers[player] = selected ?? false;
    notifyListeners();
  }
}
