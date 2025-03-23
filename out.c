#include <Windows.h>
#include <stdio.h>

unsigned char cipertext[] = {
	0x17, 0x14, 0x3, 0x4f, 
	0x15, 0xa, 0x16, 0x11, 
	0x1, 0x1d, 0x58, 0xc, 
	0x1d, 0x55, 0xc, 0xe, 
	0x14, 0x9, 0x1c
};

unsigned char key[] = {
	0x73, 0x75, 0x64, 0x6f, 
	0x78, 0x65
};

LPWSTR filename = L"payload.txt";

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
