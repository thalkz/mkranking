import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/models/race.dart';
import 'package:kart_app/pages/create_player_page.dart';
import 'package:kart_app/pages/home/charts_tab.dart';
import 'package:kart_app/pages/player_page.dart';
import 'package:kart_app/pages/submit_results_page.dart';
import 'package:kart_app/widgets/character_icon.dart';
import 'package:timeago/timeago.dart' as timeago;

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _page = 0;
  List<Player> _players = [];
  List<Race> _races = [];

  void _updatePage(int page) {
    setState(() {
      _page = page;
    });
  }

  Future<void> _refreshAllPlayers() async {
    try {
      final players = await Api.getAllPlayers();
      setState(() {
        _players = players;
      });
    } catch (error, trace) {
      debugPrint('$trace');
      ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
    }
  }

  Future<void> _refreshAllRaces() async {
    try {
      final races = await Api.getAllRaces();
      setState(() {
        _races = races;
      });
    } catch (error) {
      ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(error.toString())));
    }
  }

  List<Player> _getPlayersFromResults(List<int> results) {
    final List<Player> selected = [];
    for (final id in results) {
      selected.add(_players.firstWhere((player) => player.id == id, orElse: () => Player.empty()));
    }
    return selected;
  }

  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      _refreshAllPlayers();
      _refreshAllRaces();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Mario Kart'),
        actions: [
          IconButton(
            icon: Icon(Icons.refresh),
            onPressed: _page == 0 ? _refreshAllPlayers : _refreshAllRaces,
          ),
        ],
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
        return RefreshIndicator(
          onRefresh: _refreshAllPlayers,
          child: ListView.builder(
              padding: const EdgeInsets.all(16.0),
              itemCount: _players.length + 1,
              itemBuilder: (_, i) {
                if (i == _players.length) {
                  return _CreatePlayerButton();
                } else {
                  return _RankingTile(player: _players[i]);
                }
              }),
        );
      case 1:
        return RefreshIndicator(
          onRefresh: _refreshAllRaces,
          child: ListView.builder(
              padding: const EdgeInsets.all(16.0),
              itemCount: _races.length,
              itemBuilder: (_, i) {
                return _RaceTile(
                  race: _races[i],
                  players: _getPlayersFromResults(_races[i].results),
                );
              }),
        );
      case 2:
        return ChartsTab();
      default:
        return SizedBox();
    }
  }
}

class _RaceTile extends StatelessWidget {
  final Race race;
  final List<Player> players;

  const _RaceTile({Key? key, required this.race, required this.players}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 8.0),
      child: InkWell(
        onTap: () {},
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Align(alignment: Alignment.bottomRight, child: Text(timeago.format(race.date))),
              Wrap(
                spacing: 4.0,
                children: players
                    .map((player) => Column(
                          children: [
                            Stack(
                              children: [
                                CharacterIcon(icon: player.icon, size: 72.0),
                                Text('${players.indexOf(player) + 1}', style: Theme.of(context).textTheme.headline5)
                              ],
                            ),
                            SizedBox(
                              width: 72.0,
                              child: Center(child: Text(player.name, overflow: TextOverflow.ellipsis, maxLines: 1)),
                            ),
                          ],
                        ))
                    .toList(),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class _RankingTile extends StatelessWidget {
  final Player player;

  const _RankingTile({Key? key, required this.player}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListTile(
      onTap: () => Navigator.push(context, MaterialPageRoute(builder: (_) => PlayerPage(id: player.id))),
      title: Text(player.name),
      subtitle: Text(player.rating.toInt().toString()),
      leading: CharacterIcon(icon: player.icon),
      trailing: Text(
        player.rank.toString(),
        style: Theme.of(context).textTheme.headline5,
      ),
    );
  }
}

class _CreatePlayerButton extends StatelessWidget {
  const _CreatePlayerButton({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: TextButton(
          onPressed: () => Navigator.push(context, MaterialPageRoute(builder: (_) => CreatePlayerPage())),
          child: Text('Ajouter un joueur'),
        ),
      ),
    );
  }
}
