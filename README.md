# Kart - who's the best Mario Kart player ?

Kart is a simple web-app to record Mario Kart races and attribute an elo score to all participants. This project is simply a small side-project to have fun & learn new things, but also to put an end to the discussions of who the best player actually is.

## How it works

- Users create a character with a name, a [Mario Kart avatar](https://www.mariowiki.com/Mario_Kart_8_Deluxe), and a starting [elo score of 1000
- Races results are entered in a few clics (only the real players are taken into account, bots are ignored)
- After each race, players will earn or loose elo points, based on their result in the race
- The ranking is updated after each race. More elo points = better ranking.

## Ranking & score

Kart uses the [Elo rating system](https://en.wikipedia.org/wiki/Elo_rating_system) in order to compute scores. Elo is used in many competitive sports, like in chess and tennis. The basic principle of the Elo ranking system is that winning against a player with a bigger score earns more points than beating a player with less score. The system aims to make an accurate estimation of the skills of each players, based on match history.

Elo usually only works for two player games (1 vs 1), but [this article](https://towardsdatascience.com/developing-a-generalized-elo-rating-system-for-multiplayer-games-b9b495e87802) show how to generalize this system for multiplayer games, with an arbitrary number of participants for each match.

## Project structure

This project has 3 main parts :
- `app`, a Flutter app that can be compiled to an Android or iOS app, but also as a web-app
- `server`, the backend for this project, written in Go
- `reverse_proxy` that contains Nginx configs for handeling incomming requests to the server

The whole project can be deployed using docker-compose. The docker-compose config also starts a Posgresql database, using the official Postgres image.

## How to use

- Clone the project
- Create a `.env` file at the root of the project based on the `.env-example` file, with your configs
- Deploy using `docker-compose build` and `docker-compose up`

## License 

[MIT](LICENSE)