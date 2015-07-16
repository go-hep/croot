#ifndef CROOT_CROOT_ROOT_H
#define CROOT_CROOT_ROOT_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TROOT */
CROOT_API
CRoot_File
CRoot_ROOT_GetFile(CRoot_ROOT self,
                   const char *name);

CROOT_API
CRoot_ObjArray
CRoot_ROOT_GetListOfClasses(CRoot_ROOT self);

CROOT_API
CRoot_Interpreter
CRoot_ROOT_GetInterpreter(CRoot_ROOT self);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_ROOT_H */
