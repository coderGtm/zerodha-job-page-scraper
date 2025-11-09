# Zerodha Job Page Scraper

This is a very simple script I wrote to be notified whenever there is a job posting on the [Careers page](https://careers.zerodha.com/) of Zerodha.

## How it works?
It simply searches for this text on the Careers page: "There are no job openings currently."
If the text is not found, it means (most probably) that there are new job openings at Zerodha and the notification is fired.

## Configuration
Simply change the `ntfyTopic` variable in the code to your desired topic name and subscribe to it using the ntfy app or any other method mentioned on the [ntfy website](https://ntfy.sh/).

Add a new cron job to run this script at your desired frequency. For example, to run it every day at 9 AM, you can add the following line to your crontab:

```bash
0 9 * * * /usr/local/go/bin/go run /path/to/this/script/scraper.go
```

## Tools used
- [Colly](https://github.com/gocolly/colly) for scraping the page
- [ntfy.sh](https://ntfy.sh/) for notifying