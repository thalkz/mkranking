class RatingsHistory {
  final List<String> playerNames;
  final List<DateTime> dates;
  final List<List<double>> dataRows;

  RatingsHistory.fromJson(Map<String, dynamic> json)
      : playerNames = List.from(json['player_names'] ?? []),
        dates = List.from(json['dates'] ?? []).map((entry) => DateTime.parse(entry)).toList(),
        dataRows = List.from(json['data_rows'] ?? [], growable: false)
            .map((list) => List.from(list ?? [], growable: false).map<double>((value) => value?.toDouble()).toList())
            .toList();

  factory RatingsHistory.empty() => RatingsHistory.fromJson({});
}
