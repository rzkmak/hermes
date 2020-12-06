# Hermes

Save and sync your ssh config in private repository

*Disclaimer*
```
- Use with your own risk
- Never use public repository to save your config
```

## Installation

```
curl -L https://github.com/rzkmak/hermes/releases/download/0.0.1/hermes -o hermes -H 'Accept: application/octet-stream' && sudo chmod +x hermes && sudo mv hermes /usr/local/bin
```

## How to Use

- Create your private repository to save the config, add config there with format example in `example/config.yaml`
- Run `hermes init ${your_private_url_repository}`
- Run `hermes connect ${hostname}`

## How to Remove Local Config

- Run `hermes destroy`
