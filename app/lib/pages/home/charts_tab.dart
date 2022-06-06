import 'package:flutter/material.dart';
import 'package:flutter_charts/flutter_charts.dart';
import 'package:kart_app/models/ratings_history.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:provider/provider.dart';

class ChartsTab extends StatefulWidget {
  const ChartsTab({Key? key}) : super(key: key);

  @override
  State<ChartsTab> createState() => _ChartsTabState();
}

class _ChartsTabState extends State<ChartsTab> {
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

  ChartData _getChartData(RatingsHistory history) {
    return ChartData(
      dataRows: history.dataRows,
      xUserLabels: history.dates.map((date) => date.toString()).toList(),
      dataRowsLegends: history.playerNames,
      chartOptions: chartOptions,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(left: 8.0, right: 8.0, bottom: 80.0, top: 16.0),
      child: SizedBox(
        height: double.infinity,
        width: double.infinity,
        child: Consumer<AppNotifier>(builder: (context, notifier, _) {
          final chartData = _getChartData(notifier.history);
          return LineChart(
            painter: LineChartPainter(
              lineChartContainer: LineChartTopContainer(
                chartData: chartData,
              ),
            ),
          );
        }),
      ),
    );
  }
}
