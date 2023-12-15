const data = require("./config.json");
const fs = require('node:fs');

for (let itemsKey in data.items) {
    let item = data.items[itemsKey];
    fs.writeFileSync('./config/types/' + item.type_name + ".json", JSON.stringify(item, null, 2));
}