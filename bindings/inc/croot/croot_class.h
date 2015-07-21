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

/* TDataMember */

CROOT_API
CRoot_Int
CRoot_DataMember_GetArrayDim(CRoot_DataMember self);

CROOT_API
CRoot_Class
CRoot_DataMember_GetClass(CRoot_DataMember self);

CROOT_API
CRoot_DataType
CRoot_DataMember_GetDataType(CRoot_DataMember self);

CROOT_API
const char*
CRoot_DataMember_GetFullTypeName(CRoot_DataMember self);

CROOT_API
CRoot_Long
CRoot_DataMember_GetOffset(CRoot_DataMember self);

CROOT_API
const char*
CRoot_DataMember_GetTypeName(CRoot_DataMember self);

CROOT_API
CRoot_Bool
CRoot_DataMember_IsaPointer(CRoot_DataMember self);

CROOT_API
CRoot_Bool
CRoot_DataMember_IsBasic(CRoot_DataMember self);

CROOT_API
CRoot_Bool
CRoot_DataMember_IsEnum(CRoot_DataMember self);

CROOT_API
CRoot_Bool
CRoot_DataMember_IsPersistent(CRoot_DataMember self);

CROOT_API
CRoot_STLType
CRoot_DataMember_IsSTLContainer(CRoot_DataMember self);

/* TDataType */

CROOT_API
const char*
CRoot_DataType_GetFullTypeName(CRoot_DataType self);

CROOT_API
CRoot_DataTypeKind
CRoot_DataType_GetType(CRoot_DataType self);

CROOT_API
const char*
CRoot_DataType_GetTypeName(CRoot_DataType self);

CROOT_API
CRoot_Int
CRoot_DataType_Size(CRoot_DataType self);

CROOT_API
CRoot_Long
CRoot_DataType_Property(CRoot_DataType self);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_CLASS_H */
