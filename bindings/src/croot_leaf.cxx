#include "croot/croot.h"

#include "TLeaf.h"
#include "TLeafI.h"
#include "TLeafF.h"
#include "TLeafD.h"
#include "TLeafO.h"

/* TLeaf */
CRoot_Branch
CRoot_Leaf_GetBranch(CRoot_Leaf self)
{
  return (CRoot_Branch)(((TLeaf*)self)->GetBranch());
}

int
CRoot_Leaf_GetLenStatic(CRoot_Leaf self)
{
  return ((TLeaf*)self)->GetLenStatic();
}

CRoot_Leaf
CRoot_Leaf_GetLeafCount(CRoot_Leaf self)
{
  return (CRoot_Leaf)(((TLeaf*)self)->GetLeafCount());
}

const char*
CRoot_Leaf_GetTypeName(CRoot_Leaf self)
{
  return ((TLeaf*)self)->GetTypeName();
}

void*
CRoot_Leaf_GetValuePointer(CRoot_Leaf self)
{
  return ((TLeaf*)self)->GetValuePointer();
}

/* TLeafI */
double
CRoot_LeafI_GetValue(CRoot_LeafI self, int idx)
{
  return ((TLeafI*)self)->GetValue(idx);
}

/* TLeafF */
double
CRoot_LeafF_GetValue(CRoot_LeafF self, int idx)
{
  return ((TLeafF*)self)->GetValue(idx);  
}

/* TLeafD */
double
CRoot_LeafD_GetValue(CRoot_LeafD self, int idx)
{
  return ((TLeafD*)self)->GetValue(idx);    
}

/* TLeafO */
double
CRoot_LeafO_GetValue(CRoot_LeafO self, int idx)
{
  return ((TLeafO*)self)->GetValue(idx);    
}

// EOF

