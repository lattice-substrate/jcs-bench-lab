// JavaScriptCore (jsc) validation script.
// Reads CSV lines from stdin via readline(), outputs JSON summary to stdout.
// JSC's readline() blocks on EOF, so the orchestrator appends a sentinel line
// "__ORACLE_EOF__" to signal end of input.

var dv = new DataView(new ArrayBuffer(8));
function numberFromBits(hex) {
  dv.setBigUint64(0, BigInt('0x' + hex), false);
  return dv.getFloat64(0, false);
}

var result = {
  engine: 'jsc',
  version: 'JavaScriptCore',
  total: 0,
  passed: 0,
  failed: 0,
  divergences: []
};

var lineNum = 0;
while (true) {
  var line = readline();
  if (line === '__ORACLE_EOF__') break;
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
