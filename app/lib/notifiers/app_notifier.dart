import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/models/race.dart';
import 'package:kart_app/models/ratings_history.dart';

class AppNotifier with ChangeNotifier {
  final GlobalKey<ScaffoldMessengerState> scaffoldKey;
  List<Player> players = [];
  List<Race> races = [];
  RatingsHistory history = RatingsHistory.empty();

  AppNotifier({required this.scaffoldKey});

  void _showSnackBar(String text) {
    scaffoldKey.currentState?.showSnackBar(SnackBar(content: Text(text)));
  }

  Future<void> refreshPlayers() async {
    try {
      players = await Api.getAllPlayers();
      notifyListeners();
    } catch (error, trace) {
      debugPrint('$trace');
      _showSnackBar(error.toString());
    }
  }

  Future<void> refreshRaces() async {
    try {
      races = await Api.getAllRaces();
      notifyListeners();
    } catch (error) {
      _showSnackBar(error.toString());
    }
  }

  List<Player> getPlayersFromResults(List<int> results) {
    final List<Player> selected = [];
    for (final id in results) {
      selected.add(players.firstWhere((player) => player.id == id, orElse: () => Player.empty()));
    }
    return selected;
  }

  Future<void> refreshCharts() async {
    try {
      history = await Api.getRatingsHistory();
      notifyListeners();
    } catch (error) {
      _showSnackBar(error.toString());
    }
  }
}
