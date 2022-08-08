# Go-URL-Shortener

#### The purpose of this project is to learn Go, Redis and especially how to use Docker.

## Some comments

#### After starting this app, go to
#### [localhost:3000/api/v1](http://localhost:3000)
#### You can use Postman or other testing tools to send post request 
#### A json format response should be returned, with generated short url, along with
#### expected expired time, current rate limit, and expected reset time for rate limit

<img width="700" alt="image" src="https://user-images.githubusercontent.com/83926585/183337760-383c55f9-edb6-4ccd-9fa7-cff2e06b11cc.png">

#### As shown above, you can use the short to access the url before the expired time (expiry)
#### The default rate limit is 10 times and that limit resets every 30 minutes

### Familiarized myself with: 
> - Golang, Fiber framework
> - Redis database
> - Dockerize a project
> - Testing with Postman
