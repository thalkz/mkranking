class RatingsHistory {
  final List<String> playerNames;
  final List<DateTime> dates;
  final List<List<double>> dataRows;

  RatingsHistory.fromJson(Map<String, dynamic> json)
      : playerNames = List.from(json['player_names']),
        dates = List.from(json['dates']).map((entry) => DateTime.parse(entry)).toList(),
        dataRows = (json['data_rows'] as List<dynamic>)
            .map((list) => (list as List<dynamic>).map<double>((value) => value?.toDouble()).toList())
            .toList();

  factory RatingsHistory.empty() => RatingsHistory.fromJson({});
}
