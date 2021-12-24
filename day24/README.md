# Day 24

I solved this days puzzle (both part 1 & 2) with a whiteboard & a spreadsheet.

I reversed engineered the machine-code style language, and realised that the program is made up of 14 main blocks, where each block starts with taking the next digit of the input (Termed `w`), doing some intermediary computations with `x` and `y`, and then preserving the `z` value across blocks (The `x` & `y` variables are zeroed across blocks, and the `w` value changes from input).

The entire program revolves around a 'stack' of numbers that - in order to produce a final `z` value of `0` - must balance out with 7 "push"es and 7 "pop"s from the stack. There's nothing we can do about the pushes to the stack (There is no way to avoid them regardless of `w` input), but the "pop"s we can only trigger if we meet certain criteria. In order to empty the stack (and end up with `z=0`) we must ensure that we trigger **every** pop by meeting the criteria.

I did some algebra to construct the possible z-values and w-value constraints (based on other cells) that ensure we end up with a final `z=0`. Each w value (input in the puzzle) will be constrained to exactly one other cell in the puzzle (They act as pairs).

For my input, the following z & w values were calculated. The `z` value column shows the possible algebraic expression that can be used to calculate the possible `z` values once that block is finished. Any numbers that are `wX` (e.g. `w1`, `w2`, ...) will always be constrained between `1` and `9` (inclusive) and are what we are trying to calculate.

| Block # | z value         | w value                                     | Type   |
| ------- | --------------- | ------------------------------------------- | ------ |
| 1       | `w1+12`         | Unknown constraint                          | "push" |
| 2       | `26(z1)+w2+7`   | Unknown constraint                          | "push" |
| 3       | `26(z2)+w3+8`   | Unknown constraint                          | "push" |
| 4       | `26(z3)+w4+8`   | Unknown constraint                          | "push" |
| 5       | `26(z4)+w5+15`  | Unknown constraint                          | "push" |
| 6       | `26(z3)+w4+8`   | `w6=((26(z4)+w5+15)%26)-16` -> `w6=w5-1`    | "pop"  |
| 7       | `26(z6)+w7+8`   | Unknown constraint                          | "push" |
| 8       | `26(z3)+w4+8`   | `w8=((26(z6)+w7+8)%26)-11` -> `w8=w7-3`     | "pop"  |
| 9       | `26(z2)+w3+8`   | `w9=((26(z3)+w4+8)%26)-13` -> `w9=w4-5`     | "pop"  |
| 10      | `26(z9)+w10+13` | Unknown constraint                          | "push" |
| 11      | `26(z2)+w3+8`   | `w11=((26(z9)+w10+13)%26)-8` -> `w11=w10+5` | "pop"  |
| 12      | `26(z1)+w2+7`   | `w12=((26(z2)+w3+8)%26)-1` -> `w12=w3+7`    | "pop"  |
| 13      | `w1+12`         | `w13=((26(z1)+w2+7)%26)-4` -> `w13=w2+3`    | "pop"  |
| 14      | `0`             | `w14=(w1+12)%26-14` -> `w14=w1-2`           | "pop"  |

This leaves us with the following constraints (Where `w1` is the first digit in the model #, and `w14` is the last).

- `w1` must be `2` larger than `w14`
- `w4` must be `5` larger than `w9`
- `w5` must be `1` larger than `w6`
- `w7` must be `3` larger than `w8`
- `w11` must be `5` larger than `w10`
- `w12` must be `7` larger than `w3`
- `w13` must be `3` larger than `w2`

## Part 1

In order to solve part 1, we want to maximize every digit, while obeying our constraints.

Above, the constrained cells are listed on the right-hand side (`w14`, `w9`, `w6`, etc.), so we can fill in the _unconstrained_ cells right away (`w1`, `w4`, `w5`, etc.).

| 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  | 12  | 13  | 14  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 9   |     |     | 9   | 9   |     | 9   |     |     |     | 9   | 9   | 9   |     |

We now simply have to fill in the remaining following the constraints.

- `w14` must be `7` (2 less than `w1`)
- `w9` must be `4` (5 less than `w4`)
- `w6` must be `8` (1 less than `w5`)
- etc.

| 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  | 12  | 13  | 14  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 9   | 6   | 2   | 9   | 9   | 8   | 9   | 6   | 4   | 4   | 9   | 9   | 9   | 7   |

## Part 2

In order to solve part 2, we want to minimize every digit, while obeying our constraints.

In this scenario, our constrained cells are actually on the left-hand side (`w1`, `w4`, `w5`, etc.), so we can fill in the _unconstrained_ cells right away (`w14`, `w9`, `w6`, etc.).

| 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  | 12  | 13  | 14  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
|     | 1   | 1   |     |     | 1   |     | 1   | 1   | 1   |     |     |     | 1   |

We now simply have to fill in the remaining following the constraints.

- `w1` must be `3` (2 more than `w4`)
- `w4` must be `6` (5 more than `w9`)
- `w5` must be `2` (1 more than `w6`)
- etc.

| 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  | 12  | 13  | 14  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 3   | 1   | 1   | 6   | 2   | 1   | 4   | 1   | 1   | 1   | 6   | 8   | 4   | 1   |
