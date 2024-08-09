# stress-tester
Stress testing is a form of deliberately intense or thorough testing a system in order to determine its stability. It involves testing beyond normal operational capacity, often to a breaking point, in order to observe the results. In this project you can find an example of a tool for stress testing. It is a CLI tool written in Go with the support of Cobra and uses multithreading to create and send concurrent requests to an endpoint.

## Simple Usage
#### docker run [name-of-your-app] --url=http://google.com.br --requests=1000 --concurrency=10
