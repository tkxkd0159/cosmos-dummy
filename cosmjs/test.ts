let a = [1,2,3, {"b":5}]
let b = [...a];
b[0] = 10;
(b[3] as {b: number}).b = 30;
console.log(a, b)

function chnage(a: any) {
    a[3] = 1000
    a[0] = 777
}


chnage(b)
console.log(a, b)