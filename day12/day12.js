const fs = require('fs');

const path = require('path');
const filePath = path.join(__dirname, 'climbing.txt');
// Example of line: abcccccccccccccccccccccccccccccccccccccccccccccccccccaaaaacccccccccccccaaaaaaaccccccccccccccccaaaaaaccccccccccccccccccccccccccaaaaacccaaaccccccccccccccccccccccccccccccaaaaa
// S = starting point
// E = ending point
// Need to find the shortest path from S to E
const climbingInput = fs.readFileSync(filePath, 'utf-8').split('\n');
// Find the starting point
let startingPoint = "";

let squares = [];

climbingInput.find((line, index)  => {
  squares.push(line.split(''));
  if (line.includes('S')) {
    startingPoint = index;
  }
});

console.log(startingPoint);
console.log(squares);


