export function convert(bytes: number): string {
	if (bytes >= 1125899906842624)
		return (bytes / 1125899906842624).toFixed(2) + " PiB";

	if (bytes >= 1099511627776)
		return (bytes / 1099511627776).toFixed(2) + " TiB";

	if (bytes >= 1073741824)
		return (bytes / 1073741824).toFixed(2) + " GiB";
	
	if (bytes >= 1048576)
		return (bytes / 1048576).toFixed(2) + " MiB";
	
	if (bytes >= 1024)
		return (bytes / 1024).toFixed(2) + " KiB";
	
	return bytes + " Byte";
}
