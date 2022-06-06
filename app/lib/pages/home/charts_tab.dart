
import 'package:flutter/material.dart';
import 'package:flutter_charts/flutter_charts.dart';
import 'package:kart_app/models/ratings_history.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/utils/chart_utils.dart';
import 'package:provider/provider.dart';

class ChartsTab extends StatefulWidget {
  const ChartsTab({Key? key}) : super(key: key);

  @override
  State<ChartsTab> createState() => _ChartsTabState();
}

class _ChartsTabState extends State<ChartsTab> {
  final chartOptions = ChartOptions(
    dataContainerOptions: DataContainerOptions(
      yTransform: (num value) => value - 1000.0,
    ),
    yContainerOptions: YContainerOptions(
      isYContainerShown: false,
      isYGridlinesShown: false,
    ),
    legendOptions: LegendOptions(
      legendColorIndicatorWidth: 8.0,
    ),
    lineChartOptions: LineChartOptions(
      hotspotInnerRadius: 1.0,
      hotspotOuterRadius: 0.0,
      lineStrokeWidth: 2.0,
      hotspotInnerPaintColor: Colors.black,
    ),
  );

  ChartData _getChartData(History history) {
    return ChartData(
      dataRows: history.dataRows,
      xUserLabels: history.dates.map((date) => '${date.day}/${date.month}').toList(),
      dataRowsLegends: history.playerIds.map((id) => context.read<AppNotifier>().getPlayer(id).name).toList(),
      chartOptions: chartOptions,
      dataRowsColors: defaultLineColors(history.dataRows.length)
    );
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(left: 8.0, right: 8.0, bottom: 80.0, top: 16.0),
      child: Consumer<AppNotifier>(builder: (context, notifier, _) {
        if (notifier.history.playerIds.isEmpty) return Text("No history");
        final chartData = _getChartData(notifier.history);
        return SizedBox(
          height: double.infinity,
          width: double.infinity,
          child: LineChart(
            painter: LineChartPainter(
              lineChartContainer: LineChartTopContainer(
                chartData: chartData,
              ),
            ),
          ),
        );
      }),
    );
  }
}
