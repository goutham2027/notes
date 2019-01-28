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
  return [r, g, b]
    .map(decCh => Math.max(0, Math.min(255, decCh)).toString(16))
    .map(hexCh => (hexCh.length === 1 ? `0${hexCh}` : hexCh))
    .join('');
}

export function hexToRgb(hexString: string): { r: number; g: number; b: number } {
  if (hexString.length === 3) {
    let hr = hexString[0];
    let hg = hexString[1];
    let hb = hexString[2];
    return hexToRgb(`${hr}${hr}${hg}${hg}${hb}${hb}`);
  }

  let [r, g, b] = [0, 2, 4].map(offset => hexString.substring(offset, offset + 2)).map(hexCh => parseInt(hexCh, 16));

  return { r, g, b };
}
