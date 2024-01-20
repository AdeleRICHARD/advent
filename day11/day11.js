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
        operation : action[2].split("= ")[1],
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
// On calcule le modulo de tous les singes
// On va pour chaque singe multiplier la valeur du test précédent avec le test d'après
// Ce qui te donne le max de multiplication possible 
// donc on fait modulo dessus ensuite et ça ne dépassera jamais ce nombre
var max = monkeys.reduce((val, monkey) => val * monkey.test, 1);
console.log("before loop ", monkeys);
while (count < 10000){
    monkeys.forEach((monkey) => {
        while (monkey.startingItems.length !== 0){
            firstItem = monkey.startingItems.shift();
            monkey.timesInspected++;
            let newItem = eval(`let old = ${firstItem}; ${monkey.operation};`);
            //Part 1
            //newItem /= 3;
            //newItem = Math.floor(newItem);

            // Part 2
            
            let itemModulo = newItem % max;
            if ( newItem % Number(monkey.test) == 0 ){
                // True
                monkeys[monkey.trueTest].startingItems.push(itemModulo);
            } else {
                // False
                monkeys[monkey.falseTest].startingItems.push(itemModulo);
            }
        };
    });
    count++;
    console.log(count);
}
// Logique finale pour déterminer les singes les plus actifs
let maxInspections = monkeys.map(monkey => monkey.timesInspected).sort((a, b) => b - a);
let monkeyBusinessLevel = maxInspections[0] * maxInspections[1];
console.log(monkeys);
console.log(`Level of monkey business: ${monkeyBusinessLevel}`);
// 32338289412 too high
