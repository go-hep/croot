#include "croot/croot.h"

#include "RVersion.h"
#include "TInterpreter.h"

CRoot_Interpreter CRoot_gInterpreter;

CRoot_Bool
CRoot_Interpreter_LoadText(CRoot_Interpreter self, const char *code)
{
	CRoot_Bool ok = 1;
#if ROOT_VERSION_CODE > ROOT_VERSION(6,0,0)
	ok = (CRoot_Bool)
#endif
	(((TInterpreter*)self)->LoadText(code));
    return ok;
}
