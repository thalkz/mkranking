import 'package:flutter/material.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/pages/home/charts_tab.dart';
import 'package:kart_app/pages/home/players_tab.dart';
import 'package:kart_app/pages/submit_results_page.dart';
import 'package:provider/provider.dart';

import 'home/races_tab.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _page = 0;

  void _updatePage(int page) {
    if (_page == page) return;
    setState(() {
      _page = page;
    });
    switch (page) {
      case 0:
        context.read<AppNotifier>().refreshPlayers();
        break;
      case 1:
        context.read<AppNotifier>().refreshRaces();
        break;
      case 2:
        context.read<AppNotifier>().refreshCharts();
        break;
    }
  }

  @override
  void initState() {
    super.initState();
    Future.microtask(() => context.read<AppNotifier>().refreshPlayers());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Kart'),
      ),
      body: _buildPage(),
      floatingActionButton: FloatingActionButton.extended(
        label: Text('Nouvelle course'),
        icon: Icon(Icons.add),
        onPressed: (() => Navigator.push(context, MaterialPageRoute(builder: (_) => SubmitResultsPage()))),
      ),
      bottomNavigationBar: BottomNavigationBar(
        onTap: _updatePage,
        currentIndex: _page,
        items: [
          BottomNavigationBarItem(icon: Icon(Icons.star), label: 'Ranking'),
          BottomNavigationBarItem(icon: Icon(Icons.flag), label: 'Courses'),
          BottomNavigationBarItem(icon: Icon(Icons.line_axis), label: 'Stats')
        ],
      ),
    );
  }

  Widget _buildPage() {
    switch (_page) {
      case 0:
        return PlayersTab();
      case 1:
        return RacesTab();
      case 2:
        return ChartsTab();
      default:
        return SizedBox();
    }
  }
}
