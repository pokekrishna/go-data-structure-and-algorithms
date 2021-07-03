# Data Structures & Algorithms Using Go Lang
Implementation of common algorithms and data structures.

- Implementations are my solutions and represent how I understand the problem

## Data Structures
- Classification of the data structures is taken from [this wiki.](https://en.wikipedia.org/wiki/List_of_data_structures)

## Generate Random Numbers
You may need a list of random numbers for sorting and especially for benchmarking the algorithms.
Use the following script to generate sample data sets.
```Shell
for i in `seq 1000`;\
do od -An -N2 -i /dev/random | tr -s ' ' ' ' | tr -s '\n' ',' ;\
done
```
Currently, this script only produces _randomized_ list and not the sorted list too.

---
TODOs
- convert algorithms/... to use modules correctly. Currently it is reusing main package. This will help someone to use
the packages in her packages