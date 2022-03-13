import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:kart_app/models/player.dart';
import 'package:kart_app/models/race.dart';
import 'package:kart_app/models/ratings_history.dart';

final String? url = String.fromEnvironment("SERVER_URL", defaultValue: 'http://localhost');

class Api {
  static _parseResponse(http.Response response) {
    final json = Map<String, dynamic>.from(jsonDecode(response.body));
    print("<- ${response.statusCode} $json");
    return json['data'];
  }

  static _post(String endpoint, {Map? body}) async {
    var response = await http.post(
      Uri.parse('$url/$endpoint'),
      body: body != null ? json.encode(body) : null,
      headers: {
        'Content-type': 'application/json',
        'Accept': 'application/json',
      },
    );
    print(response.request);
    return _parseResponse(response);
  }

  static Future<List<Player>> getAllPlayers() async {
    final data = await _post('getAllPlayers');
    return List.from(data).map((value) => Player.fromJson(value)).toList()
      ..sort((a, b) => b.rating.compareTo(a.rating));
  }

  static Future<List<Race>> getAllRaces() async {
    final data = await _post('getAllRaces');
    return List.from(data).map((value) => Race.fromJson(value)).toList()..sort((a, b) => b.date.compareTo(a.date));
  }

  static Future<Player> getPlayer(int id) async {
    final data = await _post('getPlayer', body: {
      'id': id,
    });
    return Player.fromJson(data);
  }

  static Future<void> deletePlayer(int id) async {
    await _post('deletePlayer', body: {
      'id': id,
    });
  }

  static Future<int> createPlayer({required String name, required int icon}) async {
    final data = await _post('createPlayer', body: {
      'name': name,
      'icon': icon,
    });
    return int.tryParse('$data') ?? 0;
  }

  static Future<void> submitResults(List<int> results) async {
    await _post('submitResults', body: {
      'ranking': results,
    });
    return;
  }

  static Future<RatingsHistory> getRatingsHistory() async {
    final response = await _post('getRatingsHistory', body: {});
    return RatingsHistory.fromJson(response);
  }
}
