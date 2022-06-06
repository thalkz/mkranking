import 'dart:math';

import 'package:flutter/material.dart';
List<Color> defaultLineColors(int dataRowsCount) {
  List<Color> _rowsColors = List.empty(growable: true);

  if (dataRowsCount >= 1) {
    _rowsColors.add(Colors.yellow);
  }
  if (dataRowsCount >= 2) {
    _rowsColors.add(Colors.green);
  }
  if (dataRowsCount >= 3) {
    _rowsColors.add(Colors.blue);
  }
  if (dataRowsCount >= 4) {
    _rowsColors.add(Colors.black);
  }
  if (dataRowsCount >= 5) {
    _rowsColors.add(Colors.grey);
  }
  if (dataRowsCount >= 6) {
    _rowsColors.add(Colors.orange);
  }
  if (dataRowsCount > 6) {
    for (int i = 6; i < dataRowsCount; i++) {
      int colorHex = Random().nextInt(0xFFFFFF);
      int opacityHex = 0xFF;
      _rowsColors.add(Color(colorHex + (opacityHex * pow(16, 6)).toInt()));
    }
  }
  return _rowsColors;
}
