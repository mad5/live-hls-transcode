const readline = require('readline');

// Set your sleep duration in seconds
const sleepDuration = 5;

// Create a readline interface
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

async function infiniteLoop() {
  while (true) {
    //console.log('Currently sleeping...');

    // Wait for the specified duration in seconds
    await sleep(sleepDuration * 1000);

    //console.log('Waking up!');

  }
}

infiniteLoop();