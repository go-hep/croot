#include "croot/croot.h"

#include "TClass.h"
#include "TDataMember.h"

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

CRoot_DataMember
CRoot_Class_GetDataMember(CRoot_Class self, const char *name)
{
	return (CRoot_DataMember)((TClass*)self)->GetDataMember(name);
}


