# DigitalOcean snapshot pruner

> Command Line tool to automatically delete old volume snapshots on [DigitalOcean](https://digitalocean.com).

[![GitHub release](https://img.shields.io/github/v/release/brpaz/do-snapshot-pruner?style=for-the-badge)](https://github.com/brpaz/do-snapshot-pruner/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/brpaz/do-snapshot-pruner)](https://goreportcard.com/report/github.com/brpaz/do-snapshot-pruner)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## Motivation

[DigitalOcean](https://digitalocean.com) doesn´t provide an automatic way to delete old snapshots.


## Install

This recommended way is to download the latest binary for your system from the [Releases](https://github.com/brpaz/do-snapshot-pruner/releases) page.

You can also use our [Docker Image](https://hub.docker.com/repository/docker/brpaz/do-snapshot-pruner).


## Usage

```sh
do-snapshot-pruner -n <days> -t <do_token>
```

Where:

* "n" -> The number of days before current date to delete. For Exemple if you see this value with "3", the tool will delete all the snaphots older than 3 days.
* "t" -> The DigitalOcean API token. You can get yours [here](https://cloud.digitalocean.com/account/api/tokens). Optionally you can also set the "DO_TOKEN" envrionment variable.

Or with docker:

```sh
docker run -it brpaz/do-snapshot-pruner:latest -t <do_token> -n 5
```

## Run tests

```sh
make test
```

## Author

👤 **Bruno Paz**

* Website: [brunopaz.net](https://brunopaz.net)
* Github: [@brpaz](https://github.com/brpaz)
* Twitter: [@brunopaz88](https://twitter.com/brunopaz88)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com).

## Show your support

If this project have been useful for you, I would be grateful to have your support.

Give a ⭐️ to the project, or just:

<a href="https://www.buymeacoffee.com/Z1Bu6asGV" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## 📝 License

Copyright © 2019 [Bruno Paz](https://github.com/brpaz).

This project is [MIT](https://opensource.org/licenses/MIT) licensed.
