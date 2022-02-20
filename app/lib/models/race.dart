
class Race {
  final int id;
  final DateTime date;
  final List<int> results;

  Race.fromJson(Map<String, dynamic> json)
      : id = json['id'] ?? -1,
        date = DateTime.parse(json['date'] ?? "2000-01-01"),
        results = List.from(json['results'] ?? []);

  factory Race.empty() => Race.fromJson({});

  @override
  int get hashCode => id.hashCode;

  @override
  bool operator ==(other) {
    return (other is Race && other.id == id);
  }
}
