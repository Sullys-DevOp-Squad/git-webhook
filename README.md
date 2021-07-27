# Git Webhook Example

## Overview 

This repository contains an example of a web service written in Go that allows for the creation of branch policies when a new repository is created for a particular user or organization. When a new repository is created, this web service also creates an issue ticket to inform the team/individual that the branch policies have been set.

## Endpoints Available

* /protectrepo
  * This endpoint will allow for you to have a webhook that automatically sets branch policies on the `main` branch when a new repository is created.

## Prerequisites

All you will need for your workstation is:

* Go 1.16.x
* NPM

## Running the webhook locally

First, clone the repository

```bash
git clone https://github.com/Sullys-DevOp-Squad/git-webook-example
```

Once you have the repository cloned, to run the service you can specify:

```bash
go run main.go 
```
This will run the web service with a temporary binary on the OS you are targeting.

To compile the project, all you need to do is run:

```bash
go build .
```

This will produce an executable binary where you can run on a web server of your choice.

You will need to also ensure your Github PAT token is stored as an environment variable where you're running the server. You can set the PAT token for Windows and Unixlike environment as followed:

Windows (via Powershell)
```powershell
$Env:ACCESS_KEY = "YOUR_PAT_TOKEN_VALUE"
```

Linux/OSX
```bash
export ACCESS_KEY=YOUR_PAT_TOKEN_VALUE 
```

You will then need to be able to allow for Github to access your local running service, over a public IP. I highly reccomend using [LocalTunnel](https://github.com/localtunnel/localtunnel#readme), which you can install through npm:

```bash
npm install -g localtunnel
```

From there, if you're web service is running, you can then use LocalTunnel to make a public proxy:

```bash
lt --port 80
```

You will then get an output URL such as this:

![image](https://user-images.githubusercontent.com/22037844/126569772-17aea072-0153-4cb8-9016-bb2134164590.png)

In which you can then plug into your webhook as follows:

![image](https://user-images.githubusercontent.com/22037844/126569848-4e7b53c3-f21c-4444-82f2-e67f38941c3e.png)
