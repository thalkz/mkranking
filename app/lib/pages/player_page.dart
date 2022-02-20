import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/models/player.dart';
import 'package:kart_app/widgets/character_icon.dart';

class PlayerPage extends StatefulWidget {
  final int id;

  PlayerPage({Key? key, required this.id}) : super(key: key);

  @override
  State<PlayerPage> createState() => _PlayerPageState();
}

class _PlayerPageState extends State<PlayerPage> {
  Player _player = Player.empty();

  Future<void> _getPlayer() async {
    final player = await Api.getPlayer(widget.id);
    setState(() {
      _player = player;
    });
  }

  Future<void> _deletePlayer() async {
    await Api.deletePlayer(widget.id);
    Navigator.pop(context);
  }

  @override
  void initState() {
    super.initState();
    Future.microtask(_getPlayer);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(_player.name),
        actions: [
          IconButton(
            onPressed: _deletePlayer,
            icon: Icon(Icons.delete_outline),
          ),
        ],
      ),
      body: ListView(
        padding: const EdgeInsets.all(16.0),
        children: [
          CharacterIcon(icon: _player.icon, size: 120.0),
          SizedBox(height: 16.0),
          Center(child: Text('${_player.rank}', style: Theme.of(context).textTheme.headline5)),
          SizedBox(height: 16.0),
          Center(child: Text('${_player.rating.toInt()}')),
        ],
      ),
    );
  }
}
