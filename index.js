console.log('Loading the calculator function');

exports.handler = function(event, context, callback) {
    let body = JSON.parse(event.body)

    console.log('Received event:', JSON.stringify(body, null, 2));
    if (body.a === undefined || body.b === undefined || body.op === undefined) {
        callback("400 Invalid Input");
    }

    var res = {};
    res.a = Number(body.a);
    res.b = Number(body.b);
    res.op = body.op;
    res.titulo = process.env.NAME;

    if (isNaN(body.a) || isNaN(body.b)) {
        callback("400 Invalid Operand");
    }

    switch(body.op)
    {
        case "+":
        case "add":
            res.c = res.a + res.b;
            break;
        case "-":
        case "sub":
            res.c = res.a - res.b;
            break;
        case "*":
        case "mul":
            res.c = res.a * res.b;
            break;
        case "/":
        case "div":
            res.c = res.b===0 ? NaN : Number(body.a) / Number(body.b);
            break;
        default:
            callback("400 Invalid Operator");
            break;
    }
    callback(null, res);
};