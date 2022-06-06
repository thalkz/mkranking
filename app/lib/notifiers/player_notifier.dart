import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/models/race.dart';

class PlayerNotifier with ChangeNotifier {
  final int playerId;
  Player player = Player.empty();
  List<Race> races = [];

  PlayerNotifier({required this.playerId}) {
    getPlayer();
    getPlayerRaces();
  }

  Future<void> getPlayer() async {
    player = await Api.getPlayer(playerId);
    notifyListeners();
  }

    Future<void> getPlayerRaces() async {
    races = await Api.getPlayerRaces(playerId);
    notifyListeners();
  }

  Future<void> deletePlayer() async {
    await Api.deletePlayer(playerId);
  }
}
