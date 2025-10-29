# Maze Solver with a Genetic Algorithm

## 1. Task

The goal of this project is to solve a 2D maze using a Genetic Algorithm (GA).

- The maze is given as a matrix.
- The mouse starts at the blue cell.
- The goal (food) is the red cell.
- The objective is to find a **valid** and **short** path from start to goal, automatically (no hardcoded solution).

In the matrix:
- `1` = free cell (walkable)
- `1000` = wall / blocked

The only allowed moves are up, down, left, right.
<img width="1570" height="536" alt="output0" src="https://github.com/user-attachments/assets/9a52720b-d815-45f9-9d96-5e7c87fac15e" />

---

## 2. Approach

I use a Genetic Algorithm to evolve candidate paths through the maze.

**Representation**
- An individual = one possible path.
- A path is a sequence of moves (`U`, `D`, `L`, `R`) starting from the start cell.

**Fitness**
- Higher score for:
  - reaching the goal,
  - doing it in fewer steps,
  - not looping / revisiting cells.
- If the goal is not reached, being closer to the goal gives a better score.
- Stepping into a wall or outside the maze is heavily penalized.

This makes shorter, clean solutions beat long or messy ones.

**GA operations**
- **Selection:** tournament selection (pick the best out of a random sample).
- **Crossover:** combine parts of two parent paths to create new paths.
- **Mutation:** cut a path and regrow the rest with valid random moves.
- **Elitism:** keep the top ~10% every generation so I don’t lose the current best path.

---

## 3. Normalization

After generating, crossing, or mutating a path, I normalize it:

1. I simulate the moves from the start.
2. I stop the path if it would leave the maze, hit a wall, or after it reaches the goal.
3. I remove loops (if the path goes somewhere and comes back to the same cell, that detour is cut out).

I apply this normalization to every individual in every generation.  
That means the GA is always evolving valid, loop free, goal directed paths.

---

## 4. Result and Visualization

The GA is able to evolve a path from the start to the goal and make it shorter over generations.

**Maze + best path:** The final path is drawn as a glowing green line through the center of the cells, ending at the goal.
<img width="1570" height="536" alt="output" src="https://github.com/user-attachments/assets/de8f2c87-4f95-4ecc-8194-c6b1a5086b14" />
