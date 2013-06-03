#ifndef CROOT_CROOT_CINT_H
#define CROOT_CROOT_CINT_H 1

#ifdef __cplusplus
extern "C" {
#endif

typedef void* CRoot_Cint_TagInfo;

CROOT_API
CRoot_Cint_TagInfo
CRoot_Cint_TagInfo_new();

CROOT_API
void
CRoot_Cint_TagInfo_delete(CRoot_Cint_TagInfo ti);

CROOT_API
void
CRoot_Cint_TagInfo_SetTagName(CRoot_Cint_TagInfo self, const char* tagname);

CROOT_API
void
CRoot_Cint_TagInfo_SetTagType(CRoot_Cint_TagInfo self, char tagtype);

CROOT_API
void
CRoot_Cint_TagInfo_SetTagNum(CRoot_Cint_TagInfo self, short tagnum);

CROOT_API
const char*
CRoot_Cint_TagInfo_GetTagName(CRoot_Cint_TagInfo self);

CROOT_API
char
CRoot_Cint_TagInfo_GetTagType(CRoot_Cint_TagInfo self);

CROOT_API
short
CRoot_Cint_TagInfo_GetTagNum(CRoot_Cint_TagInfo self);

CROOT_API
int
CRoot_Cint_TagInfo_GetLinkedTagNum(CRoot_Cint_TagInfo self);

CROOT_API
int
CRoot_Cint_Defined_TagName(const char* tagname, int noerror);

#ifdef __cplusplus
typedef void (*CRoot_Cint_incsetup)(void);
#else  /* __cplusplus */
typedef void (*CRoot_Cint_incsetup)();
#endif /* __cplusplus */

CROOT_API
int
CRoot_Cint_TagTable_Setup(int tagnum, int size, int cpplink, int isabstract,
                          const char* comment, 
                          CRoot_Cint_incsetup setup_memvar,
                          CRoot_Cint_incsetup setup_memfunc);

CROOT_API
int
CRoot_Cint_Tag_MemVar_Setup(int tagnum);

CROOT_API
int
CRoot_Cint_MemVar_Setup(void *p, int type, int reftype,
                        int constvar, 
                        int tagnum,
                        int typenum,
                        int statictype,
                        int var_access,
                        const char* expr,
                        int definemacro,
                        const char* comment);
 
CROOT_API
int
CRoot_Cint_Tag_MemVar_Reset();

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_CINT_H */
