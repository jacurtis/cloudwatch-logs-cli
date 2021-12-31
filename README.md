# CloudWatch Logs Utility

A simple utility for working with CloudWatch Logs. AWS should probably build this themselves, but since they won't, I am here to save the day.

## Convert Timestamps

This is currently the only available function. It takes that nasty unix timestamps function (which is in milliseconds btw) and makes it human readable.

### Usage:

```bash
cloudwatch-logs convert-timestamps [--timezone <timezone>] [--rename <newfile.csv>]
cloudwatch-logs convert-timestamps [-t <timezone>] [-n <newfile.csv>]
cloudwatch-logs convert-timestamps -t America/Denver -n humanreadable.csv
```

### Options

**Timezone**

Allows you to set a timezone to convert into. This defaults to `America/Chicago` because that is where I work out of, so I didn't want to type it in everytime. A more sensible default would be local/autodetection. I could change this trivially down the road.

Use `--timezone` or `-z` followed by the timezone.

**Rename**

The default action is to overwrite your current file. But if you pass a filname, then it will write a new file to that filename instead.