# Women Who Go Workshop Oct 13th, 2018
## Intro to Go: Web Crawler
This workshop was put on by Zach Badgett. This is a step-by-step walk through walktrough of what was covered at the worshop.

## Install Go
Follow this link to install the Go binary on your computer. [download](https://golang.org/dl/) Follow the instructions to setup you go path as an environment variable.

## Make a Repo
The next step is to create a directory. Following go convention we will add this folder to the GOPATH.

`mkdir -p $GOPATH/src/github.com/soypete/web-crawler`

This makes the all of the directories in the path.

## Start Project
We want to keep this simple. Directions are [here](https://github.com/Soypete/Example-Web-Crawler/tree/master/hello-world#hello-world)

## Build an HTTP client
First step is to verify that we can access information on a desired site. To set up the Basic client go [here]()

There are lost of ways to customize the clients that we use to connect to sites and endpoints. [Here]() we have a made some simple change to the crawler.

## Parse the HTML
We want to be able to access specific information found on our site. The steps are found [here]()

## Find Embedded Links
Using our parser we search for all of the embedded links found on the webpage. Directions [here]()

## Using Multiple Functions to Clean Code
In the following steps we want to make the code more readable by using multiple functions to divide logic into steps. [Go here]()

## Using Go Routines and Go Channels
Go Routines and Channels and a major part of what make Go unique. [HERE]() is a basic implementaion of a channel and routine used in the crawler.

---
In the main directory I have provided an expanded web crawler example. This example uses multiple routines to speed up the process.

