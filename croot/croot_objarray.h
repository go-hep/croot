#ifndef CROOT_CROOT_OBJARRAY_H
#define CROOT_CROOT_OBJARRAY_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TObjArray */

CROOT_API
int64_t
CRoot_ObjArray_GetSize(CRoot_ObjArray self);

CROOT_API
int64_t
CRoot_ObjArray_GetEntries(CRoot_ObjArray self);

CROOT_API
CRoot_Object
CRoot_ObjArray_At(CRoot_ObjArray self, int64_t idx);

CROOT_API
const char*
CRoot_ObjArray_GetName(CRoot_ObjArray self);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_OBJARRAY_H */
