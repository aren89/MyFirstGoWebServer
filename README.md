# MyFirstGoWebServer

## Description:
This is my first go web server. 
I am creating a clean arch represent real life web application.
It's an application to hire new persons for a generic job.
There is a kafka module that consumes job offers and modules to CRUD operations on other resource, like person or job applications.
It's also dockerized to have a different build between api and kafka.

It's a WIP project.

## Requirements:
  - [docker-compose](https://docs.docker.com/compose/install/)
  - To connect with a client on mongodb replica you need to add in /etc/hosts mongodb1, mongodb2, mongodb3 at 127.0.0.1

## Run
```sh
$ docker-compose up -d --build
```
## Send a job offer kafka message
```sh
$ docker-compose exec kafka kafka-console-producer\
      --broker-list kafka:9092\
      --topic com.github.JobOffer \
      --property "parse.key=true" \
      --property "key.separator=_"
$ {"id":"asdjsdenc"}_{"company":"coolCompany","description":"super cool job for you","role":"scrum master"}
```

## Create person
```sh
curl --location --request POST 'http://127.0.0.1:8080/api/persons/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "Marco",
    "lastName": "Rossi",
    "email": "noreply@noreply.com",
    "age": 30,
    "YearsOfExperienceWorking": 5
}'
```
## Get persons with at least 5 year of experience
```sh
curl --location --request GET 'http://127.0.0.1:8080/api/persons?yearsOfExperienceWorking=5' \
--header 'Content-Type: application/json' 
```