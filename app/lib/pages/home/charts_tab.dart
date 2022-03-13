import 'package:flutter/material.dart';
import 'package:flutter_charts/flutter_charts.dart';
import 'package:kart_app/data/api.dart';

import '../../models/ratings_history.dart';

class ChartsTab extends StatefulWidget {
  const ChartsTab({Key? key}) : super(key: key);

  @override
  State<ChartsTab> createState() => _ChartsTabState();
}

class _ChartsTabState extends State<ChartsTab> {
  RatingsHistory? _history;

  Future<void> _refreshCharts() async {
    try {
      final history = await Api.getRatingsHistory();
      setState(() {
        _history = history;
      });
    } catch (error) {
      ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
    }
  }

  @override
  Widget build(BuildContext context) {
    _refreshCharts(); // TODO Call on refresh

    final chartOptions = const ChartOptions(
      dataContainerOptions: DataContainerOptions(
        gridLinesColor: Colors.black12,
        startYAxisAtDataMinRequested: true,
      ),
      yContainerOptions: YContainerOptions(
        isYContainerShown: true,
        isYGridlinesShown: true,
      ),
      legendOptions: LegendOptions(
        legendColorIndicatorWidth: 8.0,
      ),
    );
    final chartData = ChartData(
      dataRows: _history?.dataRows ?? [],
      xUserLabels: _history?.dates.map((date) => date.toString()).toList() ?? [],
      dataRowsLegends: _history?.playerNames ?? [],
      chartOptions: chartOptions,
    );
    final lineChartContainer = LineChartTopContainer(
      chartData: chartData,
    );

    return Padding(
      padding: const EdgeInsets.only(left: 8.0, right: 8.0, bottom: 80.0, top: 16.0),
      child: SizedBox(
        height: double.infinity,
        width: double.infinity,
        child: LineChart(
          painter: LineChartPainter(
            lineChartContainer: lineChartContainer,
          ),
        ),
      ),
    );
  }
}
