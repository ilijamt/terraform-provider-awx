if (process.argv.length < 4) {
    console.error("config-merge.js [config dir] [api dir]");
    process.exit(1);
}

const configDir = process.argv[2];
const apiDir = process.argv[3];

console.log(`Config directory: ${configDir}`);
console.log(`API directory: ${apiDir}`);

const fs = require('fs');
const path = require('path');
const _ = require('lodash');

function readJSONFiles(directoryPath) {
    let files = fs.readdirSync(directoryPath);
    const dataMap = new Map();

    files.forEach((file) => {
        if (file.endsWith('.json')) {
            const filePath = path.join(directoryPath, file);
            const fileContent = fs.readFileSync(filePath, 'utf8');

            try {
                const jsonData = JSON.parse(fileContent);
                const typeName = jsonData.type_name;

                if (typeName) {
                    dataMap.set(typeName, jsonData);
                } else {
                    console.log(`File ${file} does not contain 'type_name' property.`);
                }
            } catch (err) {
                console.error(`Error reading ${file}: ${err}`);
            }
        }
    });

    return dataMap;
}


const configDefault = require(`${configDir}/default.json`);
const targetDefault = require(`${apiDir}/config/default.json`);
const configTypes = readJSONFiles(`${configDir}/types`);
const configCredentials = readJSONFiles(`${configDir}/credentials`);
let apiCustomConfigTypes = new Map();
try {
    apiCustomConfigTypes = readJSONFiles(`${apiDir}/config/types`);
} catch (e) {
    // nothing to do here
}

apiCustomConfigTypes.forEach(function (value, key, map) {
    if (!configTypes.has(key)) {
        configTypes.set(key, value)
    } else {
        const data = _.mergeWith(
            configTypes.get(key),
            value,
            function (o, i) {
                if (_.isArray(o)) {
                    return o.concat(i);
                }
            }
        );
        configTypes.set(key, data);
    }
})

const generatedConfig = Object.assign(
    {},
    configDefault,
    targetDefault,
    {
        items: Array.from(configTypes.values()),
        credentials: Array.from(configCredentials.values())
    }
)

fs.writeFileSync(`${apiDir}/config.json`, JSON.stringify(generatedConfig, null, 2));
