function getInput(): string {
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`;
}

enum Obstacle {
  Tree,
  Snow,
}

const obstacles = getInput()
  .split("\n")
  .map((x) =>
    x.split("").map((x) => (x === "." ? Obstacle.Snow : Obstacle.Tree)),
  );

const colLength = obstacles[0].length;
let treeCount = 0;

obstacles.forEach((row, i) => {
  if (row[(i * 3) % colLength] === Obstacle.Tree) {
    treeCount++;
  }
});

console.log("Tree Count", treeCount);
