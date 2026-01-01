const fs = require('fs');
const path = require('path');

function parseConfig(filePath) {
  try {
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const config = JSON.parse(fileContent);
    return config;
  } catch (error) {
    if (error.code === 'ENOENT') {
      console.error(`Error: File not found at ${filePath}`);
    } else {
      console.error(`Error: Unable to parse file at ${filePath}. Error: ${error.message}`);
    }
    return null;
  }
}

function parseJsonInput(input) {
  try {
    const data = JSON.parse(input);
    return data;
  } catch (error) {
    console.error(`Error: Invalid JSON input. Error: ${error.message}`);
    return null;
  }
}

module.exports = {
  parseConfig,
  parseJsonInput,
};