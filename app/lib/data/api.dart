import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:kart_app/models/player.dart';
import 'package:kart_app/models/race.dart';
import 'package:kart_app/models/ratings_history.dart';
import 'package:kart_app/models/submit_results_response.dart';

const String? host = String.fromEnvironment("SERVER_HOST");
const String? port = String.fromEnvironment("SERVER_PORT");

class Api {
  static _post(String endpoint, {Map? body}) async {
    var response = await http.post(
      Uri.parse('$host:$port/$endpoint'),
      body: body != null ? json.encode(body) : null,
      headers: {
        'Content-type': 'application/json',
        'Accept': 'application/json',
      },
    );

    if (response.statusCode == 500) {
      final res = Map<String, dynamic>.from(jsonDecode(response.body));
      print('[http] $endpoint ${res["status"]} — ${res["error"]}');
      throw Exception("$endpoint ${res["status"]} ${res["error"]}");
    } else if (response.statusCode == 200) {
      final res = Map<String, dynamic>.from(jsonDecode(response.body));
      print('[http] $endpoint ${res["status"]}');
      return res['data'];
    } else {
      print('[http] $endpoint ${response.body}');
      throw Exception("$endpoint — ${response.body}");
    }
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

  static Future<List<Race>> getPlayerRaces(int id) async {
    final data = await _post('getPlayerRaces', body: {
      'id': id,
    });
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

  static Future<SubmitResultsResponse> submitResults(List<int> results) async {
    final response = await _post('submitResults', body: {
      'ranking': results,
    });
    return SubmitResultsResponse.fromJson(response);
  }

  static Future<History> getHistory(List<Player> players) async {
    final ids = players.map((player) => player.id).toList();
    final response = await _post('getHistory', body: {
      'ids': ids,
    });
    return History.fromJson(response);
  }
}
