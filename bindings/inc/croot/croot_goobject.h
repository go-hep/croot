#ifndef CROOT_CROOT_GOOBJECT_H
#define CROOT_CROOT_GOOBJECT_H 1

#ifdef __cplusplus
extern "C" {
#endif


CROOT_API
CRoot_GoObject
CRoot_GoObject_New(void *ptr, const char *type);

CROOT_API
void*
CRoot_GoObject_Ptr(CRoot_GoObject self);

CROOT_API
const char*
CRoot_GoObject_Type(CRoot_GoObject self);

CROOT_API
int32_t
CRoot_GoObject_Size(CRoot_GoObject self);

CROOT_API
void
CRoot_GoObject_SetPtr(CRoot_GoObject self, void *ptr);

CROOT_API
void
CRoot_GoObject_SetSize(CRoot_GoObject self, int32_t size);

CROOT_API
int32_t
CRoot_GoObject_CnvFromC(CRoot_GoObject_Converter self, CRoot_GoObject obj, void* address);

CROOT_API
int32_t
CRoot_GoObject_CnvToC(CRoot_GoObject_Converter self, CRoot_GoObject obj, void* address);

CROOT_API
CRoot_GoObject_Converter
CRoot_GoObject_Converter_Create(const char *fullname);

CROOT_API
CRoot_GoObject_Converter
CRoot_GoObject_Converter_Get(const char *fullname);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_GOOBJECT_H */
