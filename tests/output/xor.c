#include <Windows.h>
#include <stdio.h>

unsigned char ciphertext[] = {
	0x18, 0xc, 0x2, 0x19, 
	0x4c, 0x5, 0x1, 0x4d, 
	0x53, 0xc, 0x10, 0x52, 
	0x6, 0x1b, 0x4e, 0x27, 
	0x44
};

unsigned char key[] = {
	0x4b, 0x65, 0x65, 0x70, 
	0x20, 0x69, 0x74, 0x20, 
	0x73, 0x65, 0x63, 0x72, 
	0x65, 0x74, 0x21
};

LPCWSTR lpFilename = L"output/payload.txt";

BOOL XORDeseal(IN PBYTE pKey, IN PBYTE pPayload, IN DWORD dwKeySize, IN DWORD dwPayloadSize) {
	for (SIZE_T i = 0, j = 0; i < dwPayloadSize; i++, j++) {
		if (j >= dwKeySize) {
			j = 0;
		}
		pPayload[i] = pPayload[i] ^ pKey[j];
	}

	pPayload[dwPayloadSize] = '\0';
	return TRUE;
}

DWORD WritePayload() {
	HANDLE hFile = CreateFileW(lpFilename, GENERIC_WRITE, 0, NULL, CREATE_ALWAYS, FILE_ATTRIBUTE_NORMAL, NULL);

	if (hFile == INVALID_HANDLE_VALUE) {
		return -1;
	}

	DWORD bytesWritten;
	if (!WriteFile(hFile, ciphertext, sizeof(ciphertext), &bytesWritten, NULL)) {
		CloseHandle(hFile);
		return -1;
	}
	CloseHandle(hFile);
	return 0;
}

DWORD PrintPayload() {
	printf("payload : \"%s\" \n", ciphertext);
}

int main() {
	if (!XORDeseal(key, ciphertext, sizeof(key), sizeof(ciphertext))) {
		return -1;
	}

    WritePayload();
	PrintPayload();
	return 0;
}
