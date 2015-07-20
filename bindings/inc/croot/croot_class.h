#ifndef CROOT_CROOT_CLASS_H
#define CROOT_CROOT_CLASS_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TClass */

CROOT_API
CRoot_Class
CRoot_Class_GetClass(const char *);

CROOT_API
CRoot_Class
CRoot_Class_GetBaseClass(CRoot_Class self, const char* name);

CROOT_API
int
CRoot_Class_GetClassSize(CRoot_Class self);

CROOT_API
CRoot_Version
CRoot_Class_GetClassVersion(CRoot_Class self);

CROOT_API
CRoot_DataMember
CRoot_Class_GetDataMember(CRoot_Class self, const char *name);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_CLASS_H */
