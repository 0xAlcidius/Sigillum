#include <Windows.h>
#include <stdio.h>

// [[CIPERTEXT]]

// [[KEY]]

// [[FILENAME]]

BOOL XORDeseal(IN PBYTE pKey, IN PBYTE pPayload, IN DWORD dwKeySize, IN DWORD dwPayloadSize) {
	for (SIZE_T i = 0, j = 0; i < dwPayloadSize; i++, j++) {
		if (j >= dwKeySize) {
			j = 0;
		}
		pPayload[i] = pPayload[i] ^ pKey[j];
	}
	return TRUE;
}

DWORD WritePayload(LPWSTR filename) {
	HANDLE hFile = CreateFile(filename, GENERIC_WRITE, 0, NULL, CREATE_ALWAYS, FILE_ATTRIBUTE_NORMAL, NULL);

	if (hFile == INVALID_HANDLE_VALUE) {
		return -1;
	}

	DWORD bytesWritten;
	if (!WriteFile(hFile, cipertext, sizeof(cipertext), &bytesWritten, NULL)) {
		CloseHandle(hFile);
		return -1;
	}
	CloseHandle(hFile);
	return 0;
}

DWORD PrintPayload() {
	printf("payload : \"%s\" \n", cipertext);
}

int main() {
	if (!XORDeseal(key, cipertext, sizeof(key), sizeof(cipertext))) {
		return -1;
	}

	PrintPayload();
    WritePayload(filename);
	return 0;
}