# allHandsOnDeck

allHandsOnDeck is a go-based package for sending you Alerts upon an error occuring in your Application.

> allHandsOnDeck is a custom package I made while working on another Project i've been working on; and I thought the concept was good enough to turn it into it's own Package


### Current Support

This Package Currently Supports sending notifications through a few Notification Carriers.
To Use Each Service Read their respective Section

- [Twillio](https://twillio.com)

### Upcoming Support

- [AWS SNS](https://aws.amazon.com/sns/)
- [GMAIL API](https://developers.google.com/gmail/api/quickstart/go)

### Use Cases

- Health Checking API's to be notified upon Downtime in Business Critical Infrastructure.
- Processing Large Amounts of Data which may take a while to complete and you want to be notified incase of a Critical Error.

### Usage
All Carrier require different Environment Variables so please read further below for each of their [respective useages](#environment-variables).

```go
import (
    "github.com/joho/godotenv"

    "github.com/KodeyThomas/allHandsOnDeck"
)


godotenv.Load("environmentVariables.env")

importantData, err := functionThatMayReturnAnError()
if err != nil {
	sentCorrectly := allHandsOnDeck.SendTwillioErrorAlert(err)

    if(sentCorrectly) {
        // Notification Sent Correctly
    } else {
        // Notification Not Sent Correctly
    }
}
```

### Environment Variables

Please read this section for your respective Carrier.

#### Twillio

- `ACCOUNT_SID`
- `TWILLIO_AUTH_TOKEN`
- `NUMBER_TO`
- `NUMBER_FROM`

You can find all of these [HERE](https://twillio.com).

### License

This Project is Licensed Under the MIT License which you can find [HERE](LICENSE)