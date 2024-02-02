class Node {
  constructor(value) {
    this.x = value.x;
    this.y = value.y;
    this.letter = value.letter;
    this.exploredScore = 0;
    this.exploredAndHeuristicScore = 0;
  }
}

// Function to estimate the lower bound of the distance between start and end
function ManhattenDistance(start, end) {
  return Math.abs(start.x - end.x) + Math.abs(start.y - end.y);
}

// Example of line: abcccccccccccccccccccccccccccccccccccccccccccccccccccaaaaacccccccccccccaaaaaaaccccccccccccccccaaaaaaccccccccccccccccccccccccccaaaaacccaaaccccccccccccccccccccccccccccccaaaaa
// S = starting point
// E = ending point
// Need to find the shortest path from S to E
const fs = require('fs');
const paths = require('path');
const filePath = paths.join(__dirname, 'climbing.txt');
const climbingInput = fs.readFileSync(filePath, 'utf-8').split('\n');

let startingPointCoordinates = {};
let endingPointCoordinates = {};

let climbingPaths = [];

// First we need to find the starting and ending point
climbingInput.find((line, index)  => {
  climbingPaths.push(line.split(''));
  if (line.includes('S')) {
    startingPointCoordinates = new Node({x: index, y: line.indexOf('S'), letter: 'S'});
  }
  
  if (line.includes('E')) {
    endingPointCoordinates = new Node ({x: index, y: line.indexOf('E'), letter: 'E'});
  }

});

// Then we need to find the shortest path from S to E
function findPaths(start, end) {
  const visited = new Set();
  const queue = [[start, 0]];

  while (queue.length > 0) {
    const current = queue.shift();
    // x , y
    const currentMove = current[0].split(',').map(el => parseInt(el));
    const possibleMoves = getPossibleMoves(...currentMove).map(el => el.join(','));

    for (const move of possibleMoves) {
      if (move === end) {
        return current[1] + 1;
      }

      if (!visited.has(move)) {
        visited.add(move);
        queue.push([move, current[1] + 1]);
      }
    }
  }
}

function getPossibleMoves(rowI, colI) {
  const neighbours = [
    [rowI - 1, colI], // top
    [rowI, colI + 1], // right
    [rowI + 1, colI], // down
    [rowI, colI - 1], // left
  ].filter(cords => {
    return (cords[0] >= 0 || cords[1] >= 0)
      && climbingInput?.[cords[0]]?.[cords[1]]
      && canMove(climbingInput[rowI][colI], climbingInput[cords[0]][cords[1]])
  })

  return neighbours;
}

function canMove(from, to) {
  if (from === 'S') {
    from = 'a'
  }
  if (to === 'E') {
    to = 'z'
  }

  if (from.toLowerCase() !== from || to.toLowerCase() !== to) {
    return false;
  }

  if (to.charCodeAt(0) - from.charCodeAt(0) <= 1) {
    return true;
  }

  return false;
}

path = findPaths(startingPointCoordinates.x + ',' + startingPointCoordinates.y, endingPointCoordinates.x + ',' + endingPointCoordinates.y);
console.log(path);
