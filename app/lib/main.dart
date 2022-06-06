import 'package:flutter/material.dart';
import 'package:kart_app/notifiers/app_notifier.dart';
import 'package:kart_app/pages/home_page.dart';
import 'package:provider/provider.dart';
import 'package:timeago/timeago.dart' as timeago;

void main() {
  timeago.setLocaleMessages('fr', timeago.FrMessages());
  timeago.setDefaultLocale('fr');
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  MyApp({Key? key}) : super(key: key);

  final _scaffoldKey = GlobalKey<ScaffoldMessengerState>();

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider<AppNotifier>(
      create: (context) => AppNotifier(scaffoldKey: _scaffoldKey),
      child: MaterialApp(
        title: 'Flutter Demo',
        theme: ThemeData(
          primarySwatch: Colors.red,
        ),
        home: const HomePage(),
        scaffoldMessengerKey: _scaffoldKey,
      ),
    );
  }
}
