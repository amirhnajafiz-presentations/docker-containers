# Go Telegram Bot

This is a **Telegram** bot to get the latest NBA results.

## How to run?
To run the bot, first you need to create a bot in **Telegram**, then
save your _bot api token_ somewhere.

Now clone the project:
```shell
git clone https://github.com/amirhnajafiz/nba-bot.git
```

Now copy the configs:
```shell
cp ./configs/configs.yaml ./config.yaml
```

After that place your token in the config file:
```yaml
telegram:
  token: "[BOT API TOKEN]"
```

Now you can run the application by the following command:
```shell
make build
make run
```

You should see the following response:
```
Waiting for telegram bot to start ...
Bot started.
```

## Routes
Now you can test the bot in telegram by the following command:
- /test

And for getting NBA results, use the following commad:
- /view
