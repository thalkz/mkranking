class Player {
  final int id;
  final String name;
  final double rating;
  final int rank;
  final int icon;
  final int racesCount;

  Player.fromJson(Map<String, dynamic> json)
      : id = json['id'] ?? -1,
        name = json['name'] ?? '[unknown]',
        rank = json['rank'] ?? 0,
        rating = json['rating']?.toDouble() ?? 0.0,
        icon = json['icon'] ?? 0,
        racesCount = json['races_count'] ?? 0;

  factory Player.empty() => Player.fromJson({});

  @override
  int get hashCode => id.hashCode;

  @override
  bool operator ==(other) {
    return (other is Player && other.id == id);
  }
}
