#include <Windows.h>
#include <bcrypt.h>
#include <stdio.h>

#pragma comment(lib,"bcrypt.lib")

#define KEYSIZE 32
#define IVSIZE 16

// [[CIPERTEXT]]

// [[KEY]]

// [[FILENAME]]

unsigned char* addPadding(IN BYTE item[], IN DWORD dwItemSize, IN DWORD modulus, OUT DWORD *dwPaddedSize) {
    DWORD dwPaddingLength = modulus - (dwItemSize % modulus);
    *dwPaddedSize = dwItemSize + dwPaddingLength;

    unsigned char *newItem = malloc(*dwPaddedSize);
    if (newItem == NULL) {
        return NULL;
    }

    memcpy(newItem, item, dwItemSize);

    memset(newItem + dwItemSize, 0xFF, dwPaddingLength);

    return newItem;
}

int main(void) {
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
    unsigned char *bKey = addPadding(key, sizeof(key), KEYSIZE, &dwKeySize);

    status = BCryptGenerateSymmetricKey(hAlg, &hKey, NULL, 0, bKey, KEYSIZE, 0);
    if (!BCRYPT_SUCCESS(status)) {
        return 1;
    }

    BYTE plaintext[sizeof(byCiphertext)] = {0};
    DWORD plaintextSize = 0;

    status = BCryptDecrypt(hKey, byCiphertext, sizeof(byCiphertext), NULL, byIv, IVSIZE, plaintext, sizeof(plaintext), &plaintextSize, 0);
    if (!BCRYPT_SUCCESS(status)) {
        return 1;
    }

    printf("plaintext: %s\n", plaintext);
}
