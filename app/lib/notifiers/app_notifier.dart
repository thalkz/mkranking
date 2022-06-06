import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/models/race.dart';
import 'package:kart_app/models/ratings_history.dart';

class AppNotifier with ChangeNotifier {
  final GlobalKey<ScaffoldMessengerState> scaffoldKey;
  List<Player> players = [];
  List<Race> races = [];
  History history = History.empty();

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
      ]);
      players = results[0] as List<Player>;
      races = results[1] as List<Race>;
      history = await Api.getHistory(players);
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

  Future<void> initCharts() async {
    if (history.playerIds.isNotEmpty) return;
    try {
      history = await Api.getHistory(players);
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

  Player getPlayer(int playerId) {
    try {
      return players.firstWhere((player) => player.id == playerId);
    } catch (error) {
      return Player.empty();
    }
  }
}
