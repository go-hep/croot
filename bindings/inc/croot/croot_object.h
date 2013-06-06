#ifndef CROOT_CROOT_OBJECT_H
#define CROOT_CROOT_OBJECT_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TObject */

CROOT_API
const char*
CRoot_Object_ClassName(CRoot_Object self);

CROOT_API
CRoot_Object
CRoot_Object_Clone(CRoot_Object self,
                   const char *newname);

CROOT_API
CRoot_Object
CRoot_Object_FindObject(CRoot_Object self, 
                        const char *name);

CROOT_API
const char*
CRoot_Object_GetName(CRoot_Object self);

CROOT_API
const char*
CRoot_Object_GetTitle(CRoot_Object self);

CROOT_API
CRoot_Bool
CRoot_Object_InheritsFrom(CRoot_Object self, 
                          const char *classname);

CROOT_API
void
CRoot_Object_Print(CRoot_Object self,
                   CRoot_Option *option);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_OBJECT_H */
