# Coin Flip

Personal Project

## Goal

The goal of this project was to practice Docker and Golang to create a simple web application.

## Project 

The website has two interactive elements: a flip button and a reset button.

The flip button "flips a coin" by generating a random number between 0 (inclusive) and 1 (exclusive). 
A value less than 0.5 is considered to be heads, while a value of 0.5 and greater is tails.
The result will increment the corresponding counter on the web page and update the percentage of heads.

The reset button will reset both counters back to 0.

Counter values are stored in a PostgreSQL database that is accessed by the Golang backend. 

Docker was used to create separate containers for the frontend and the backend.

## Development

Dockerfile.dev files and the docker-compose-dev.yml were used during the development of the application. 
System can be tested (assuming one has Docker CLI) using *docker-compose -f docker-compose-dev.yml up* and accessing *localhost:3000* on a browser.

## Production

Project was hosted temporarily on AWS Elastic Beanstalk and used AWS Relational Database Service for PostgreSQL. 
GitHub Actions were used in order to automatically deploy the project to AWS. 

The AWS services used have been terminated as to not incur a usage fee, but can be created again fairly quickly to start hosting again. 