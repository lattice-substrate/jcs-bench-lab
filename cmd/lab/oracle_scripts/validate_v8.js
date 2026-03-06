'use strict';
const readline = require('readline');

const dv = new DataView(new ArrayBuffer(8));
function numberFromBits(hex) {
  dv.setBigUint64(0, BigInt('0x' + hex), false);
  return dv.getFloat64(0, false);
}

const result = {
  engine: 'v8',
  version: process.version + ' V8=' + process.versions.v8,
  total: 0,
  passed: 0,
  failed: 0,
  divergences: []
};

const rl = readline.createInterface({ input: process.stdin });
let lineNum = 0;
rl.on('line', line => {
  lineNum++;
  const comma = line.indexOf(',');
  if (comma < 0) return;
  const hex = line.substring(0, comma);
  const expected = line.substring(comma + 1);
  const got = String(numberFromBits(hex));
  result.total++;
  if (got === expected) {
    result.passed++;
  } else {
    result.failed++;
    if (result.divergences.length < 100) {
      result.divergences.push({ line: lineNum, hex, expected, got });
    }
  }
});
rl.on('close', () => {
  process.stdout.write(JSON.stringify(result) + '\n');
});
