#include <Windows.h>
#include <stdio.h>

// [[CIPERTEXT]]

// [[KEY]]

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

BOOL SysFunc032(IN PBYTE pRc4Key, IN PBYTE pPayload, IN DWORD dwRc4KeySize, IN DWORD sPayloadSize) {

    NTSTATUS STATUS = NULL;

    USTRING Data = {
        .Length = sPayloadSize,
        .MaximumLength = sPayloadSize,
        .Buffer = pPayload
    };

    USTRING Key = {
        .Length = dwRc4KeySize,
        .MaximumLength = dwRc4KeySize,
        .Buffer = pRc4Key
    };

    fnSystemFunction032 SystemFunction032 = (fnSystemFunction032)GetProcAddress(LoadLibraryA("Advapi32"), "SystemFunction032");

    if ((STATUS = SystemFunction032(&Data, &Key)) != 0x0) {
        printf("[!] SystemFunction032 FAILED With Error: 0x%0.8X \n", STATUS);
        return FALSE;
    }

    return TRUE;
}

int main() {
    if (!SysFunc032(key, cipertext, sizeof(key), sizeof(cipertext))) {
        return -1;
    }

    printf("shellcode : \"%s\" \n", cipertext);
}