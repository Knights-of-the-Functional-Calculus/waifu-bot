# waifu-bot
## Installation
### Quick Start
I highly recommend that you set up [docker-compose](https://docs.docker.com/compose/install/). I put a lot of care into writing the configuration files to make it easy for you all <3

Rename the **docker-compose-dev.yml.dist** file to **docker-compose-dev.yml** and add modify the file, adding the relavent IDs and API tokens. The environment can be built relatively painlessly using the command **docker-compose -f docker-compose-dev.yml up --build** thereafter.

### Slow Start
All the necessary sources are in **src/**. I built the project using [Go v1.9](https://golang.org/dl/). Don't forget to set the environment variables. You're on your own :P

## Functionality

### Records Audio
This functionality was written to make verifying the audio easier and independent from an off-site API.

### Sends Audio to API via POST request
Right now the bot is programmed to send off chunked audio to Wit ai, but should be adaptable for any sort of POST request. Just replace URLs and read up on audio formatting and REST if you're a newb.

## Why Wit?
Essentially all the other voice APIs sucked. They weren't user-friendly and they didn't have a clean REST implementation that I could easily find. With Wit.ai I could make a POST request, go into a control panel, and verify that it was sent and whether the audio was correct. With Google, that option was non-existant, and set up with Amazon was so contrived that I didn't bother.

## Why Go?
I'm a control freak who likes to know where shit is going and gets a hard on for strongly typed languages. It's a great networking language that promotes clean code. I would've used python for popularity's sake, but the discord API for that lacks audio receiving support.

## Contact
If you would like to help me train my waifu, contact me through <j4qfrost@gmail.com> or <https://discord.gg/hzanRSz>
