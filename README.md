# RPGen
RPG generator written in Go
```
python3 roll.py

# roll 10 d20's
> 20

# roll 10 d20+6's
> 20 6

# roll 10 d20's + d10's
> 20 0 10

# generate town / character list (Golang required)
> gen

# parse generated character list for a match. Space seperated queries,
# Ctrl+C to stop searching. Queries are done per-line and is case-sensitive.
# You can read out.txt after it is generated to read the generated town
# to search manually.
> char Human
> char Elf Male
> char Orc Female
...

```
