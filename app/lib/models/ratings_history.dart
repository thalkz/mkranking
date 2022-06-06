class History {
  final List<int> playerIds;
  final List<DateTime> dates;
  final List<List<double>> dataRows;

  History.fromJson(Map<String, dynamic> json)
      : playerIds = List.from(json['player_ids'] ?? []),
        dates = List.from(json['dates'] ?? []).map((entry) => DateTime.parse(entry)).toList(),
        dataRows = List.from(json['data_rows'] ?? [], growable: false)
            .map((list) => List.from(list ?? [], growable: false).map<double>((value) => value?.toDouble()).toList())
            .toList();

  factory History.empty() => History.fromJson({});
}
