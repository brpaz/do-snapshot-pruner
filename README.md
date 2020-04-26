# DigitalOcean snapshot pruner

> Command Line tool to automatically delete old volume snapshots on [DigitalOcean](https://digitalocean.com).

[![CI Status](https://img.shields.io/github/workflow/status/brpaz/do-snapshot-pruner/CI?color=orange&label=actions&logo=github&logoColor=orange&style=for-the-badge)](https://github.com/brpaz/do-snapshot-pruner/actions)
[![Codacy grade](https://img.shields.io/codacy/grade/2cac107e656b4507888fc13ba2fe4702?style=for-the-badge)](https://www.codacy.com/manual/brpaz/do-snapshot-pruner)
![Code coverage](https://img.shields.io/codacy/coverage/2cac107e656b4507888fc13ba2fe4702?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/brpaz/do-snapshot-pruner?style=for-the-badge)](https://goreportcard.com/report/github.com/brpaz/do-snapshot-pruner)

[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## Motivation

[DigitalOcean](https://digitalocean.com) doesn´t provide an automatic way to delete old snapshots. This tool provides a simple way to automate that.

## Install

This recommended way is to download the latest binary for your system from the [Releases](https://github.com/brpaz/do-snapshot-pruner/releases) page.

You can also use our [Docker Image](https://hub.docker.com/repository/docker/brpaz/do-snapshot-pruner).

## Usage

```sh
do-snapshot-pruner prune -t <do_token> -d <days>
```

Or with docker:

```sh
docker run -it brpaz/do-snapshot-pruner:latest -t <do_token> -d 5
```

### Command Flags

| Flag | Description                                                                                                                                                            | Default |
| ---- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------- |
| -d   | Specifies the number of days whose older snapshots will be deleted                                                                                                     | 7 days  |
| -t   | DitigalOcean API token - You can get yours [here](https://cloud.digitalocean.com/account/api/tokens). Optionally you can also set the "DO_TOKEN" envrionment variable. |         |
| -r   | The types of snapshots to delete. you can filter by "volume" or "droplet                                                                                               | all     |  |

## Run tests

```sh
make test
```

## Author

👤 **Bruno Paz**

  * Website: [brunopaz.dev](https://brunopaz.dev)
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

Copyright © 2020 [Bruno Paz](https://github.com/brpaz).

This project is [MIT](https://opensource.org/licenses/MIT) licensed.
