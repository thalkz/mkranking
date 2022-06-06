class SubmitResultsResponse {
  final Map<int, double> ratingsDiff;

  SubmitResultsResponse.fromJson(Map<String, dynamic> json)
      : ratingsDiff = Map<String, double>.from(json['rating_diff'] ?? {}).map((key, value) => MapEntry(int.parse(key), value));

  factory SubmitResultsResponse.empty() => SubmitResultsResponse.fromJson({});
}
