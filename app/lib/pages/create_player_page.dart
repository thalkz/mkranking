import 'package:flutter/material.dart';
import 'package:kart_app/data/api.dart';
import 'package:kart_app/data/characters.dart';
import 'package:kart_app/widgets/character_icon.dart';

class CreatePlayerPage extends StatefulWidget {
  CreatePlayerPage({Key? key}) : super(key: key);

  @override
  State<CreatePlayerPage> createState() => _CreatePlayerPageState();
}

class _CreatePlayerPageState extends State<CreatePlayerPage> {
  final _textController = TextEditingController();
  int _icon = 0;

  Future<void> _createPlayer() async {
    await Api.createPlayer(name: _textController.text, icon: _icon);
    Navigator.pop(context);
  }

  void _selectCharacter(int icon) {
    setState(() {
      _icon = icon;
    });
  }

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Créer un joueur')),
      body: ListView(
        padding: const EdgeInsets.only(top: 16.0, left: 16.0, right: 16.0, bottom: 100.0),
        children: [
          TextField(
            controller: _textController,
            decoration: InputDecoration(helperText: 'Nom'),
          ),
          SizedBox(height: 16.0),
          GridView.extent(
            shrinkWrap: true,
            physics: NeverScrollableScrollPhysics(),
            maxCrossAxisExtent: 80.0,
            children: List.generate(
                characterAssets.length,
                (icon) => InkWell(
                      onTap: () => _selectCharacter(icon),
                      child: Stack(
                        alignment: Alignment.center,
                        children: [
                          CharacterIcon(icon: icon),
                          if (icon == _icon) ...[
                            Positioned.fill(child: Container(color: Colors.black26)),
                            Icon(
                              Icons.check,
                              color: Colors.white,
                              size: 64.0,
                            )
                          ],
                        ],
                      ),
                    )),
          ),
          SizedBox(height: 80.0),
        ],
      ),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: _createPlayer,
        icon: Icon(Icons.check),
        label: Text('Créer'),
      ),
    );
  }
}
