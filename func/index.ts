import { DivFn, MulFn, SubFn, SumFn } from "./types";

const sum: SumFn = (a, b) => {
    return a + b;
}

const div: DivFn = (a: number, b: number) => {
    if (b === 0) {
        throw new Error('на ноль нельзя делить!!!');
    }

    return a / b;
}

const sub: SubFn = (a: number, b: number) => {
    return a - b;
}

const mul: MulFn = (a: number, b: number) => {
    return a * b;
}

const result = (actions: 'sum' | 'sub' | 'div' | 'mul') => {
    switch(actions) {
        case 'sum': (a: number, b: number) => {
            return sum(a, b);
        }
        case 'sub': (a: number, b: number) => {
            return sub(a, b);
        }
        case 'mul': (a: number, b: number) => {
            return mul(a, b);
        }
        case 'div': (a: number, b: number) => {
            return div(a, b);
        }
        default: return 'Unknown action!!!';
    }
}

export {sum, div, sub, mul};


