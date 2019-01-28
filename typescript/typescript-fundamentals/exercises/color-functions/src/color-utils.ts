/* Hints
1) converting a string to an integer
parseInt("124", 10) == 124

2) converting an integer into its hexadecimal representation (string)
parseInt(124, 10).toString(16) === '7c'

3) converting a hexadecimal number (string) into an integer
parseInt('7c', 16) === 124

npm test color-functions
*/

export function rgbToHex(r: number, g: number, b: number): string {
  return getHex(r) + getHex(g) + getHex(b);
}

export function hexToRgb(hexString: string): { r: number; g: number; b: number } {
  let r = 0;
  let g = 0;
  let b = 0;

  if (hexString.length == 3) {
    let r_s = hexString[0];
    let g_s = hexString[1];
    let b_s = hexString[2];

    r = getInt(r_s);
    g = getInt(g_s);
    b = getInt(b_s);
  } else if (hexString.length == 6) {
    let r_s = hexString.slice(0, 2);
    let g_s = hexString.slice(2, 4);
    let b_s = hexString.slice(4, 6);

    r = getInt(r_s);
    g = getInt(g_s);
    b = getInt(b_s);
  }
  return { r, g, b };
}

function getHex(num: number): string {
  if (num >= 0 && num <= 15) {
    return '0' + parseInt(num.toString(), 10).toString(16);
  } else if (num >= 0 && num <= 255) {
    return parseInt(num.toString(), 10).toString(16);
  } else if (num < 0) {
    return '00';
  } else if (num > 255) {
    return 'ff';
  }
  return '00';
}

function getInt(hexString: string): number {
  if (hexString.length == 2) return parseInt(hexString, 16);
  else if (hexString.length == 1) {
    hexString += hexString;
    return parseInt(hexString, 16);
  }
  return 0;
}
