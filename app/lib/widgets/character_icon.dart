import 'package:flutter/material.dart';
import 'package:kart_app/data/characters.dart';

class CharacterIcon extends StatelessWidget {
  final int icon;
  final double size;

  const CharacterIcon({Key? key, this.icon = 0, this.size = 40.0}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Image.asset(
      'res/characters/${characterAssets[icon]}',
      width: size,
      height: size,
    );
  }
}
