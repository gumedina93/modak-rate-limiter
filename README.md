# Modak Rate Limiter

[![Modak](https://uploads-ssl.webflow.com/6393514104d9044309a9e12e/6393514104d904e02da9e341_modakLogoUpdated.svg)](https://www.modakmakers.com/)

![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)

We have a Notification system that sends out email notifications of various types (supdatesupdate, daily news, project invitations, etc). We need to protect recipients from getting too many emails, either due to system errors or due to abuse, so let's limit the number of emails sent to them by implementing a rate-limited version of NotificationService.

## Features

- Allows concurrently sending notifications to users

## Installation

Rate Limiter requires [Go](https://go.dev/) to run.

Install the dependencies:

```sh
make install
```

## Run app

In order to run the application:

```sh
make run
```

## Test

In order to run tests:

```sh
make test
```
