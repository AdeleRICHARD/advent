// Read the file first
const fs = require('fs');
const getOperation = (operation) => {
    operation = operation.match(/([+\-*/])\s+(\d+|old)/);
    if(operation){
        return operation.slice(1);
    } else {
       console.log("No operation found");
    }
}

const getNumberInString = (string) => {
    let numbers = string.match(/\d+(,\s*\d+)*/g);
    if(numbers){
        return numbers[0];
    } else {
        console.log("No number found");
    }
}
const input = fs.readFileSync('./Monkeys.txt', 'utf-8').split('\n\n');

// Part 1
// Parse each monkey action into an object
const monkeys = input.map((line) => {
    if(line === ''){
        return;
    }
    const action = line.split('\n').map(line => line.trim());
    
    return{
        monkeyId : getNumberInString(action[0]),
        startingItems : getNumberInString(action[1]),
        operation : getOperation(action[2]),
        test : getNumberInString(action[3]),
        trueTest : getNumberInString(action[4]),
        falseTest : getNumberInString(action[5]),
    };
});


console.log(monkeys);
