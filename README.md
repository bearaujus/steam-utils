# STEAM-UTILS

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/bearaujus/steam-utils/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/bearaujus/steam-utils)](https://goreportcard.com/report/github.com/bearaujus/steam-utils)

Sets of utilities for managing your Steam

## Download

To get started with steam-utils, visit the release page at:

```shell
https://github.com/bearaujus/steam-utils/releases
```
From there, download the binary that matches your operating system and architecture listed under the desired release version.

## Usage

### A. Set Auto-Update Behavior

Control how games in your Steam library are updated:

```shell
steam-utils library set-auto-update [command]
```

Available Commands:

- 0 : Always keep all games updated automatically
- 1 : Update a game only when you launch it

### B. Set Background Downloads Behavior

Manage background download preferences for all collections in your Steam library:

```shell
steam-utils library set-background-downloads [command]
```

Available Commands:

- 0 : Follow your global Steam settings
- 1 : Always allow background downloads
- 2 : Never allow background downloads

## TODO

- Add more features to the `Library` category
- Support features for `Store` category
- Support features for `Community` category
- Support features for `Profile` category

## Donation

### Saweria

Support the project via Saweria: [https://saweria.co/bearaujus](https://saweria.co/bearaujus)

### Crypto

You can also donate via cryptocurrency:

```text
0xc6c3B0B41Ee6B544bF0F86Ec4065BCe9C8ccB599
```

## License

This project is licensed under the MIT License - see
the [LICENSE](https://github.com/bearaujus/berror/blob/master/LICENSE) file for details.
