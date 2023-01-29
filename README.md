# Assignment

In this exercise, you are going to build a FlightAvailability service:

- The service should be developed using Go.
- For a given airport location and departure date the service will report available flights to a given destination.
- Design a simple domain model and use a straightforward algorithm to find the best matching flights.
- This service will have a gRPC or REST api (you decide) and will query a PostgreSQL database. **Note:** Only 1 endpoint
  is requested.

You should use the free available database from [Demonstration Database](https://postgrespro.com/education/demodb).

Do not model every entity from that database; choose a minimal type system such that you can at least respond with the
available flights as described above. See also
Postgres [Pro Standard : Documentation: 10: J.2. Schema Diagram](https://postgrespro.com/docs/postgrespro/10/apjs02.html)

# Notes

- You're allowed to use open-source packages as you see fit.
- Anything not in the description is up to you.
- We are not expected a production ready solution ; limit your time spent to at most 4 hours.

# How to run :

I wrote a `Makefile` to download demo database to import into `postgres`. Then I use docker compose to run the project.
I implemented both of `grpc` and `rest`. These are the ports :

```
grpc : 50051
rest : 1323
```

Requirements to run `Makefile`:

- unzip
- tar
- make
- curl

Run this command :

```
make up
```