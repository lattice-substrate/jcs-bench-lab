// SpiderMonkey shell (js / js115 / js128) validation script.
// Reads CSV lines from stdin via readline(), outputs JSON summary to stdout.

const dv = new DataView(new ArrayBuffer(8));
function numberFromBits(hex) {
  dv.setBigUint64(0, BigInt('0x' + hex), false);
  return dv.getFloat64(0, false);
}

var smVersion = 'unknown';
if (typeof getBuildConfiguration === 'function') {
  // The shell reports its version via the C-level versionString or version().
  try { smVersion = version(); } catch(e) {}
}

var result = {
  engine: 'spidermonkey',
  version: 'SpiderMonkey ' + smVersion,
  total: 0,
  passed: 0,
  failed: 0,
  divergences: []
};

var lineNum = 0;
var line;
while ((line = readline()) !== null) {
  lineNum++;
  var comma = line.indexOf(',');
  if (comma < 0) continue;
  var hex = line.substring(0, comma);
  var expected = line.substring(comma + 1);
  var got = String(numberFromBits(hex));
  result.total++;
  if (got === expected) {
    result.passed++;
  } else {
    result.failed++;
    if (result.divergences.length < 100) {
      result.divergences.push({ line: lineNum, hex: hex, expected: expected, got: got });
    }
  }
}
print(JSON.stringify(result));
