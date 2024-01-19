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

const path = require('path');
const filePath = path.join(__dirname, 'Monkeys.txt');
const input = fs.readFileSync(filePath, 'utf-8').split('\n\n');

// Part 1
// Parse each monkey action into an object
const monkeys = input.map((line) => {
    if(line === ''){
        return;
    }
    const action = line.split('\n').map(line => line.trim());
    
    return{
        monkeyId : getNumberInString(action[0]),
        startingItems : getNumberInString(action[1]).map(Number),
        operation : getOperation(action[2]),
        test : getNumberInString(action[3]),
        trueTest : getNumberInString(action[4]),
        falseTest : getNumberInString(action[5]),
        timesInspected : 0
    };
});

function getNumberInString(operation) {
    let numbers = operation.match(/\d+(,\s*\d+)*/g);
    if(numbers){
        if (numbers[0].includes(',')){
            return numbers[0].split(',');
        }
        return numbers[0] = [numbers[0]];
    } else {
        console.log("No number found");
    }
}
function operation(operation, number1, number2){
    if (number2 === 'old') {
        number2 = number1;
    }
    switch(operation){
        case '+':
            return number1 + number2;
        case '-':
            return number1 - number2;
            case '/':
            return number1 / number2;
        case '*':
            return number1 * number2;
    }
}

// Process each monkey action
let count = 0;
console.log("before loop ", monkeys);
while (count < 20){
    monkeys.forEach((monkey) => {
        while (monkey.startingItems.length !== 0){
            firstItem = monkey.startingItems.shift();
            // Worry lvl down
            monkey.timesInspected++;
            //console.log(firstItem);
            // Operation
            let newItem = operation(monkey.operation[0], firstItem, monkey.operation[1]);
            newItem /= 3;
            // Test
            if ( newItem % monkey.test === 0 ){
                // True
                monkeys[monkey.trueTest].startingItems.push(newItem);
            } else {
                // False
                monkeys[monkey.falseTest].startingItems.push(newItem);
            }
        };
    });
    count++;
    console.log(count);
}
console.log(monkeys);
// 12375 too low
// 16256 too low

