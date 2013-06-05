#ifndef CROOT_CROOT_FILE_H
#define CROOT_CROOT_FILE_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TFile */

typedef enum {
    kDefault =0,
    kLocal   =1,
    kNet     =2,
    kWeb     =3,
    kFile    =4
} CRoot_FileType;

CROOT_API
CRoot_File
CRoot_File_Open(const char *name, 
                CRoot_Option *option,
                const char *ftitle,
                int32_t compress,
                int32_t netopt);

CROOT_API
CRoot_Bool
CRoot_File_cd(CRoot_File self, const char *path);

CROOT_API
void
CRoot_File_Close(CRoot_File self, CRoot_Option *option);

CROOT_API
int
CRoot_File_GetFd(CRoot_File self);

CROOT_API
CRoot_Object
CRoot_File_Get(CRoot_File self, const char *namecycle);

CROOT_API
CRoot_Bool
CRoot_File_IsOpen(CRoot_File self);

CROOT_API
CRoot_Bool
CRoot_File_ReadBuffer(CRoot_File self,
                      char *buf, int64_t pos, int32_t len);

CROOT_API
CRoot_Bool
CRoot_File_ReadBuffers(CRoot_File self,
                       char *buf, int64_t *pos, int32_t *len, int32_t nbuf);

CROOT_API
int32_t
CRoot_File_WriteBuffer(CRoot_File self,
                       const char *buf, int32_t len);

CROOT_API
int32_t
CRoot_File_Write(CRoot_File self, 
                 const char *name, int32_t opt, int32_t bufsiz);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_FILE_H */
