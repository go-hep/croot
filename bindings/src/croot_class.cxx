#include "croot/croot.h"

#include "TClass.h"

CRoot_Class
CRoot_Class_GetClass(const char *name)
{
  return (CRoot_Class)(TClass::GetClass(name));
}
