/*
    ** AUTHOR: 0xAlcidius

    Note:
    The AES algorithm uses AES-CBC. The first 16 bytes of the ciphertext are composed 
    of the initialization vector (IV), followed up by the ciphertext. The last byte of the
    ciphertext resembles the amount of bytes that are used for padding. The RemovePadding
    function uses this byte exclude X amount of bytes from the end of the payload.
*/

#include <Windows.h>
#include <bcrypt.h>
#include <stdio.h>

#pragma comment(lib,"bcrypt.lib")

#define KEYSIZE 32
#define IVSIZE 16

// [[CIPERTEXT]]

// [[KEY]]

// [[FILENAME]]

PBYTE AddPadding(IN BYTE item[], IN DWORD dwItemSize, IN DWORD dwModulus, OUT DWORD* dwPaddedSize) {
	DWORD dwPaddingLength = dwModulus - (dwItemSize % dwModulus);
	*dwPaddedSize = dwItemSize + dwPaddingLength;

	unsigned char* newItem = malloc(*dwPaddedSize);
	if (newItem == NULL) {
		return NULL;
	}

	memcpy(newItem, item, dwItemSize);

	memset(newItem + dwItemSize, 0xFF, dwPaddingLength);

	return newItem;
}

PBYTE RemovePadding(IN BYTE item[], IN DWORD dwItemSize, IN DWORD dwBytesToRemove, OUT DWORD* dwNewItemSize) {
	*dwNewItemSize = dwItemSize - dwBytesToRemove;
	PBYTE pbyNewItem = malloc(*dwNewItemSize);
	if (!pbyNewItem) {
		return NULL;
	}

	memcpy(pbyNewItem, item, *dwNewItemSize);
	return pbyNewItem;
}

PBYTE AESDeseal(OUT DWORD* dwSize) {
	NTSTATUS status;
	BCRYPT_ALG_HANDLE hAlg = NULL;
	BCRYPT_KEY_HANDLE hKey = NULL;

	BYTE byIv[IVSIZE];
	BYTE byCiphertext[sizeof(cipertext) - IVSIZE - 1];
	BYTE byCiphertextPadding[1];

	memcpy(byIv, cipertext, IVSIZE);
	memcpy(byCiphertext, cipertext + IVSIZE, sizeof(cipertext) - IVSIZE - 1);
	memcpy(byCiphertextPadding, cipertext + sizeof(cipertext) - 1, 1);

	status = BCryptOpenAlgorithmProvider(&hAlg, BCRYPT_AES_ALGORITHM, NULL, 0);
	if (!BCRYPT_SUCCESS(status)) {
		return 1;
	}

	status = BCryptSetProperty(hAlg, BCRYPT_CHAINING_MODE, (PUCHAR)BCRYPT_CHAIN_MODE_CBC, sizeof(BCRYPT_CHAIN_MODE_CBC), 0);
	if (!BCRYPT_SUCCESS(status)) {
		BCryptCloseAlgorithmProvider(hAlg, 0);
		return 1;
	}

	DWORD dwKeySize = 0;
	unsigned char* bKey = AddPadding(key, sizeof(key), KEYSIZE, &dwKeySize);

	status = BCryptGenerateSymmetricKey(hAlg, &hKey, NULL, 0, bKey, KEYSIZE, 0);
	if (!BCRYPT_SUCCESS(status)) {
		free(bKey);
		return 1;
	}

	free(bKey);

	BYTE bPaddedPlaintext[sizeof(byCiphertext)];
	DWORD dwPaddedPlaintextSize = 0;

	status = BCryptDecrypt(hKey, byCiphertext, sizeof(byCiphertext), NULL, byIv, IVSIZE, bPaddedPlaintext, sizeof(bPaddedPlaintext), &dwPaddedPlaintextSize, 0);
	if (!BCRYPT_SUCCESS(status)) {
		return 1;
	}

	PBYTE pbyPlaintext = RemovePadding(bPaddedPlaintext, dwPaddedPlaintextSize, *byCiphertextPadding, dwSize);

	return pbyPlaintext;
}

DWORD WritePayload(LPWSTR filename, PBYTE pbyPlaintext, DWORD dwSize) {
	HANDLE hFile = CreateFile(filename, GENERIC_WRITE, 0, NULL, CREATE_ALWAYS, FILE_ATTRIBUTE_NORMAL, NULL);

	if (hFile == INVALID_HANDLE_VALUE) {
		return -1;
	}

	DWORD bytesWritten;
	if (!WriteFile(hFile, pbyPlaintext, dwSize, &bytesWritten, NULL)) {
		CloseHandle(hFile);
		return -1;
	}
	CloseHandle(hFile);
	return 0;
}

DWORD PrintPayload(PBYTE pbyPlaintext) {
	printf("payload : \"%s\" \n", pbyPlaintext);
}

int main() {
	DWORD dwSize = 0;
	PBYTE pbyPlaintext = AESDeseal(&dwSize);

	WritePayload(filename, pbyPlaintext, dwSize);

	PrintPayload(pbyPlaintext);

	free(pbyPlaintext);
}
