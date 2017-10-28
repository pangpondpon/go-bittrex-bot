# go-bittrex-bot
Go Bittrex Bot is the auto alert library that help you watch the price of crypto currency from bittrex.com

## Usage
1. Copy `config.json.example` to `config.json` and update it to match your need.
2. Run the script using this command `./go-bittrex-bot`. Any alert will go to your Slack, according to the `web_hook_url` that you setup in `config.json` file
3. It's highly recommended that you should add this to your cron script so it run every minute, see example below

```
// Open cron job using 'crontab -e' command
// After run 'crontab -e' add this to cronjob file
 * * * * * ./path/to/go-bittrex-bot
```

## Config file explanation
The explanation is already commented in `config.json.example`

```json
{
  "credentials": {
    // Get Credential from Bittrex Api Key page in bittrex.com
    // https://bittrex.com/Manage#sectionApi
    "api_key": "API_KEY",
    "api_secret": "API_SECRET"
  },
  "pairs": [
    {
      "symbol": "BTC-XEM",

       // If above is true, the code will alert if 'current price >= threshold'
       // If above is false, the code will alert if 'current price <= threshold'
      "above": true,

      "threshold": 1.0 // Threshold is in dollar
    },
    {
      "symbol": "BTC-ADA",
      "above": true,
      "threshold": 0.1
    }
  ],
  "slack": {
    "web_hook_url": "",
    "user_name": "",
    "icon_emoji": ":dollar:",
  }
}
```