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
    scaffoldKey.currentState?.showSnackBar(SnackBar(
      content: Text(text),
      backgroundColor: Colors.red,
    ));
  }

  Future<void> refreshAll() async {
    try {
      final results = await Future.wait([
        Api.getAllPlayers(),
        Api.getAllRaces(),
        Api.getRatingsHistory(),
      ]);
      players = results[0] as List<Player>;
      races = results[1] as List<Race>;
      history = results[2] as RatingsHistory;
      notifyListeners();
    } catch (error) {
      _showSnackBar(error.toString());
    }
  }

  Future<void> initPlayers() async {
    if (players.isNotEmpty) return;
    try {
      players = await Api.getAllPlayers();
      notifyListeners();
    } catch (error) {
      _showSnackBar(error.toString());
    }
  }

  Future<void> initRaces() async {
    if (races.isNotEmpty) return;
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

  Future<void> initCharts() async {
    if (history.playerNames.isNotEmpty) return;
    try {
      history = await Api.getRatingsHistory();
      notifyListeners();
    } catch (error) {
      _showSnackBar(error.toString());
    }
  }

  Future<void> createPlayer({required String name, required int icon}) async {
    try {
      await Api.createPlayer(name: name, icon: icon);
      await refreshAll();
    } catch (error) {
      _showSnackBar(error.toString());
    }
  }
}
