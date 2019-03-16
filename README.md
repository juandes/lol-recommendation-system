# League of Legends Recommendation System [![GoDoc](https://godoc.org/github.com/juandes/lol-recommendation-system?status.svg)](https://godoc.org/github.com/juandes/lol-recommendation-system)
## Overview
lol-recommendation-system, or **LoLReco** is recommender system built from scratch using Go for the game [League of Legends](https://na.leagueoflegends.com/) (LoL). LoL is a multiplayer online battle area game (MOBA) which features two teams made of five players each (in its default mode) competing against each other. In layman's terms, we could summarize the game's goal into the following sentence: each team has to reach the other team's hub (known as the Nexus) and destroy it before the opposing team slays yours. This task is not an easy one since each side has defensive structures, items, and minions that hinder every step you make. However, the real complexity of LoL lies in the fighters, or Champions as they called in game.

The recommendation system I created, deals with those champions, and its purpose is to recommend a complete team composition, that is the five champions each player should choose.

This system features a web service that provides the recommendations. Also, the project has been completely dockerized.

A complete explanation of the system is available at: [Building a League of Legends Champions recommender system in Go and how to deploy it in the cloud](https://towardsdatascience.com/building-a-league-of-legends-champions-recommender-system-in-go-and-how-to-deploy-it-in-the-cloud-1ee7a4fb55ee)

## The algorithm
This system employs a neighborhood-based collaborative filtering algorithm that uses k-NN (k-nearest neighbors) to find its nearest objects. At training time, because k-NN is a lazy learner, the model won't learn a target function. Instead, it will "memorize" data and generalize from it once a prediction request has been done to the system. At this prediction stage, the algorithm loops through every single training example from its search space and calculates the distance between them and the input feature vector. Once it has finished calculating the distances, they are sorted in descending order and the algorithm finally returns the top N items.

The algorithm's input is an incomplete team composition, in other words, a list of champions, made of either one, two, three or four of them. The output is a complete team composition made of the five champions the algorithm thinks are the most appropriate ones. For example, the input vector `[ashe, drmundo]` might return `[ashe, drmundo, masteryi, yasuo, zyra]`.

## Web service and Docker
One of the applications of this program (https://github.com/juandes/lol-recommendation-system/tree/master/cmd/service) is a web service that is used to serve the recommendations. By default, it will run on port `8080`. This can be modified with the  `--port` argument.
Besides this, this whole project is dockerized (https://hub.docker.com/r/juandes/lol-champions-recommender)

## What's included in this repo
This repo includes the recommendation engine, the training dataset, the Python code used to obtain the dataset, and R script that was used to analyze the prediction latency.

## How to use 
The easier to use the system is through the `Makefile`. The targets are:
- `make test` to run tests
- `make go-build` to build each application
- `make run-service-local` to start the web service on port 8080
- `make run-service-docker` to pull the Docker image and execute it

The service serves the recommendations on the endpoint `/recommend`. Its parameter is a JSON with the following structure

```
type PredictionInput struct {
	Champions   []string `json:"champions"`
	Intercept   bool     `json:"intercept"`
	Shuffle     bool     `json:"shuffle"`
	Serendipity bool     `json:"serendipity"`
}
```
A valid request looks like this:
`curl -d '{"champions":["ashe", "drmundo"], "intercept": false, "shuffle": false, "serendipity":false}' -H "Content-Type: application/json" -X POST http://localhost:8080/recommend`

The champions list have have at least one champion and four at most, and their names have to be written in lowercase and without any whitespace, e.g. _Jarvan IV_ has to be _jarvaniv_.

## Contribution
This system is a playground for testing and experimenting, is someone has a nice idea and would like to contribute, fix bugs or improving the system, I'd be super happy.

## Disclaimer
I should mention that this engine is by no means perfect nor complete and that it is a proof of concept on how we could build recommendation systems for a video game. League of Legends and the champions selection component of the game involves more complex processes such as the banning of champions (champions that can't be used by any player), and the order in which each player selects its champion. These are processes that require a high level of expertise, something that can't be that easily automated or learned.

