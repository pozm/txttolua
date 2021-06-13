const fs = require('fs');
const path = require('path');

function main() { // i feel like i now need a main function; it just doesn't feel right without..

    const pathTo = process.argv[2]; // argc doesn't exist so it's thrown in here

    if (!fs.existsSync(pathTo ?? '')) {
        console.log('i require a parameter passed in. OR a valid path.');
    }

    const oldP = path.resolve(pathTo);
    const lastidx = oldP.lastIndexOf('\\');
    const newP = oldP.slice(0,lastidx) + oldP.slice(lastidx,oldP.lastIndexOf('.')) + '.lua';

    fs.renameSync(oldP, newP);

}

main();
