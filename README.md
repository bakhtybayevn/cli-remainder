# GO-cli-reminder
Go CLI Reminder is a command-line tool for setting and displaying reminders at specific times. It allows you to schedule reminders in the future and receive notifications with custom messages.
# Features
1. Set reminders by specifying the time (HH:MM) and a text message.
2. Display reminders at the scheduled time.
3. Countdown timer to show how much time is left until the reminder.
# Prerequisites
Before using this tool, make sure you have the following installed on your system:
1. Go (to build and run the project)
2. Required Go dependencies (see installation instructions below)
# Installation
1. Clone the repository to your local machine:
git clone https://github.com/bakhtybayevn/cli-reminder.git
2. Change the working directory to the project folder:
cd cli-reminder
3. Build the project:
go build
4. Install the required Go dependencies:
go mod tidy
# Usage
To set a reminder, use the following format:
./go-cli-reminder <hh:mm> <text message>
Replace <hh:mm> with the time you want the reminder to trigger (in 24-hour format), and <text message> with your reminder message.

Example:
./go-cli-reminder 15:30 "Meeting with John"

The tool will display a countdown timer, indicating how much time is left until the reminder. At the specified time, it will show a notification with your reminder message.
# Acknowledgments
This project uses the following third-party libraries:
1. gen2brain/beeep: A cross-platform notification library for Go.
2. olebedev/when: A natural language date and time parser.
# Contributing
Contributions are welcome! If you have ideas for improvements, open an issue or submit a pull request.
