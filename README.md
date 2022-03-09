# Raise Money For Partners In Health

For the past decade, the Vlogbrothers and Nerdfigtheria continuously support Partners In Health as a beneificiary of the annual [Project for Awesome](http://www.projectforawesome.com). In 2019, Hank and John started an inititive to fund a new maternal care facility in Sierra Leon to help reduce maternal morality in the Kona district, where Partners In Health works tirelessly to improve healthcare.

## But first, some background

Read about the Green family connection to Partners In Health on the [Partners In Health website](https://www.pih.org/vlogbrothers-support-maternal-health). Find more about the [fundraiser for this materity center in this Vlogbrothers video on YouTube](https://www.youtube.com/watch?v=DwDjsNFHVhQ).

## This Bot

Allows Twitter users to raise money for Partners in Health. If you see a tweet so good you want to donate to Partners In Health on behalf of that tweeter, respond to the tweet, saying "Hey [@CharityYeti](https://twitter.com/charityyeti)" somewhere in the tweet text. We'll send you a link to donate to Partners in Health and will tweet back letting the original tweeter that you appreciated their tweet to support maternal health in Sierra Leon! [This tweet by Hank](https://twitter.com/hankgreen/status/1186824079120011264) is the basis of this project.

## Deployment

We built the Charity Yeti on [Docker](https://www.docker.com) using [docker-compose](https://docs.docker.com/compose/). To build, issue `docker-compose build ${serviceName}`, where `${serviceName}` is either `charityyeti-frontend`, `charityyeti-backend`, or both. You can spin up a backend by cloning the repo and issuing `docker-compose up backend` from the same directory as the `Dockerfile.backend` (the repo root). This exposes the port set as your environment variable `PORT`. See the dependecies section below for setting environment variables for local development easily. The frontend can also be started using `docker-compose up frontend` from repo root where the `Dockerfile.frontend` lives. The production deployment of Charity Yeti runs on [Kubernetes](https://kubernetes.io) hosted by [Google Cloud](https://cloud.google.com).

## Code of Conduct

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](code-of-conduct.md)

This project is by Nerdfighters, for Partners In Health. We would love open collaboration and accept contributions. We also expect that you not forget to be awesome, and adhere by the [Contributer Covenant](https://www.contributor-covenant.org). See `code-of-conduct.md` for full details.

## Contributing
To contribute, please open an issue so we can discuss. Once we address the issue at hand, fork the repo and make the necessary changes or updates, then create a pull request into this repo. `architecture.md` contains a [Mermaid diagram](https://github.com/mermaid-js/mermaid) of the general architecture Charity Yeti uses. If you make a change to the way Charity Yeti behaves, please also update the Mermaid sequence diagram.

## Dependencies
 - [GoDotEnv](https://github.com/joho/godotenv) and [env](https://github.com/caarlos0/env) to manage environment variables. 
 - [Anaconda](https://github.com/ChimeraCoder/anaconda) and [Go-Twitter](https://github.com/dghubble/go-twitter) for Twitter interactions.
 - [Gorilla/Mux](https://github.com/gorilla/mux) for backend HTTP server and routing.
 - [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) for interactions with the [MongoDB](https://www.mongodb.com) database.
 - [Zap Suggared Logging](https://github.com/uber-go/zap) for backend logging.
 - [PHP](https://www.php.net) powers the front end.
 - [Apache](https://httpd.apache.org) is the web server running the front end.
 - [Docker](https://docker.io) and [Docker Compose](https://docs.docker.com/compose/) for deployment and local development.
 - [Mermaid](https://github.com/mermaid-js/mermaid) for the sequence diagram.
 - [Kubernetes](https://kubernetes.io) hosted by Google Cloud is where Charity Yeti lives in production.