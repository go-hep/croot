#include "croot/croot.h"

#include "TClass.h"
#include "TDataMember.h"
#include "TDataType.h"

CRoot_Class
CRoot_Class_GetClass(const char *name)
{
  return (CRoot_Class)(TClass::GetClass(name));
}

CRoot_Class
CRoot_Class_GetBaseClass(CRoot_Class self, const char* name)
{
	return (CRoot_Class)((TClass*)self)->GetBaseClass(name);
}

CRoot_Int
CRoot_Class_GetClassSize(CRoot_Class self)
{
	return CRoot_Int(((TClass*)self)->GetClassSize());
}

CRoot_Version
CRoot_Class_GetClassVersion(CRoot_Class self)
{
	return CRoot_Version(((TClass*)self)->GetClassVersion());
}

// TDataMember ----

CRoot_DataMember
CRoot_Class_GetDataMember(CRoot_Class self, const char *name)
{
	return (CRoot_DataMember)((TClass*)self)->GetDataMember(name);
}


CRoot_Int
CRoot_DataMember_GetArrayDim(CRoot_DataMember self)
{
	return CRoot_Int(((TDataMember*)self)->GetArrayDim());
}

CRoot_Class
CRoot_DataMember_GetClass(CRoot_DataMember self)
{
	return (CRoot_Class)((TDataMember*)self)->GetClass();
}

CRoot_DataType
CRoot_DataMember_GetDataType(CRoot_DataMember self)
{
	return (CRoot_DataType)((TDataMember*)self)->GetDataType();
}

const char*
CRoot_DataMember_GetFullTypeName(CRoot_DataMember self)
{
	return ((TDataMember*)self)->GetFullTypeName();
}

CRoot_Long
CRoot_DataMember_GetOffset(CRoot_DataMember self)
{
	return CRoot_Long(((TDataMember*)self)->GetOffset());
}

const char*
CRoot_DataMember_GetTypeName(CRoot_DataMember self)
{
	return ((TDataMember*)self)->GetTypeName();
}

CRoot_Bool
CRoot_DataMember_IsaPointer(CRoot_DataMember self)
{
	return CRoot_Bool(((TDataMember*)self)->IsaPointer());
}

CRoot_Bool
CRoot_DataMember_IsBasic(CRoot_DataMember self)
{
	return CRoot_Bool(((TDataMember*)self)->IsBasic());
}

CRoot_Bool
CRoot_DataMember_IsEnum(CRoot_DataMember self)
{
	return CRoot_Bool(((TDataMember*)self)->IsEnum());
}

CRoot_Bool
CRoot_DataMember_IsPersistent(CRoot_DataMember self)
{
	return CRoot_Bool(((TDataMember*)self)->IsPersistent());
}

CRoot_STLType
CRoot_DataMember_IsSTLContainer(CRoot_DataMember self)
{
	return CRoot_STLType(((TDataMember*)self)->IsSTLContainer());
}

// TDataType ---

const char*
CRoot_DataType_GetFullTypeName(CRoot_DataType self)
{
	return ((TDataType*)self)->GetFullTypeName();
}

CRoot_DataTypeKind
CRoot_DataType_GetType(CRoot_DataType self)
{
	return CRoot_DataTypeKind(((TDataType*)self)->GetType());
}

const char*
CRoot_DataType_GetTypeName(CRoot_DataType self)
{
	return ((TDataType*)self)->GetTypeName();
}

CRoot_Int
CRoot_DataType_Size(CRoot_DataType self)
{
	return CRoot_Int(((TDataType*)self)->Size());
}

CRoot_Long
CRoot_DataType_Property(CRoot_DataType self)
{
	return CRoot_Long(((TDataType*)self)->Property());
}
