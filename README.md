# DPLL SAT SOLVER
This is a simple SAT solver I implemented to explore the boolean satisfiability problem and understand why it's so difficult.
I followed this blog fairly closely: https://homes.cs.washington.edu/~emina/blog/2017-06-23-a-primer-on-sat.html
I ran into some issues with the bcp/unit propogation. I had gotten that for each unit literal you could remove any clause that contained it however.
I missed that for each clause with the negative of the unit literal you could remove the negative. And if the clause no longer contained any literals you had a contradiction.
This turned out to be a key part of the algorithm that I missed and so I was very confused about the speed of my solver for a long time.

Also included in this repo is a verifier and a simple 3SAT generator.

I used participle (a grammar engine) to parse CNF formulas.

## Testing
```shell
make test
make bench
```

## Building
```shell
make build
```

## Running
Generate 3SAT instance
```shell
./generator 100 50 > mycnf.cnf
```
100 is the number of clauses, 50 is the number of variables to choose from

Solve the CNF
```shell
./solver mycnf.cnf
./solver samples/4x3.cnf
./solver samples/80x50.cnf
```

Instead, you can run the solver in a REPL mode by running it with no arguments