const sum = (a: number, b: number) => {
    return a + b;
}

const div = (a: number, b: number) => {
    if (b === 0) {
        throw new Error('на ноль нельзя делить!!!');
    }

    return a / b;
}

const sub = (a: number, b: number) => {
    return a - b;
}

const mul = (a: number, b: number) => {
    return a * b;
}

const result = (actions: 'sum' | 'sub' | 'div' | 'mul') => {
    if (actions === 'sum') {
        return (a: number, b: number): number => {
            return sum(a, b);
        }
    }
}

export {sum, div, sub, mul};


