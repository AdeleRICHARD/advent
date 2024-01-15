// Read the file first
const fs = require('fs');
const input = fs.readFileSync('./Monkeys.txt', 'utf-8').split('\n\n');

// Part 1
// Parse each monkey action into an object
// use pop ?
const monkeys = input.map((line) => {
    const action = line.split('\n').map(line => line.trim());
    console.log(action[0]);
    return{
        monkeyId : action[0].split(' ')[1],
        startingItems : action[1].split(':')[1].trim().split(','),
        operation : action[2].split('=')[1].trim().split('+'),
        test : action[3].split(':')[1].trim().split(' '),
        trueTest : action[4].split(':')[1].split(' '),
        falseTest : action[5].split(':')[1].split(' '),
    };
});

console.log(monkeys);
