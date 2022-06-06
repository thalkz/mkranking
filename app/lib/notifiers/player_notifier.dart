import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/models/player.dart';

class PlayerNotifier with ChangeNotifier {
  final int playerId;
  Player player = Player.empty();

  PlayerNotifier({required this.playerId}) {
    getPlayer();
  }

  Future<void> getPlayer() async {
    player = await Api.getPlayer(playerId);
    notifyListeners();
  }

  Future<void> deletePlayer() async {
    await Api.deletePlayer(playerId);
  }
}
