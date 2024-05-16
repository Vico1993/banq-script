# BanQ-Script

Small script to check availability of the BanQ machine.

Link: https://square.banq.qc.ca/reserver

## Table of Contents

- [Getting Started](#getting-started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

- Go (at least version 1.21)
- A valid Telegram bot ( look at the great bot father )

1. Create `.env`

```sh
# Bot token given by the @BotFather
TELEGRAM_BOT_TOKEN=<TELEGRAM_BOT>
```

2. Install all dependancies

```sh
make ensure_deps
```

## Deployment

Github action will take ownership of the deployment

## Usage

To use Otto, run the following command:

```sh
make build && ./bin/bot
```

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](./CONTRIBUTING.md) file for more information.

## License

This project is licensed under the [LICENSE](./LICENSE) file in the root directory of this repository.
