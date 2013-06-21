#include <stdlib.h>

#include <vector>

#include "TSchemaRule.h"
#include "TSchemaRuleSet.h"
#include "TClass.h"
#include "TError.h"
#include "TVirtualObject.h"

#include "go_croot_goslice.h"

void
go_croot_cnv_vector_to_goslice(char *tgt, TVirtualObject *obj) {
  std::vector<double> *pers = (std::vector<double>*)(((char*)obj->GetObject())+8);
  go_slice *slice = (go_slice*)tgt;
  slice->Len = pers->size();
  slice->Cap = slice->Len;
  slice->Data = (void*)malloc(slice->Len * sizeof(double));
  double *data = (double*)slice->Data;
  for (int i = 0; i < slice->Len; i++) {
	data[i] = (*pers)[i];
  }
}

void 
go_croot_register_read_vector(char *tgt, TVirtualObject *obj) 
{
  go_slice *slice = (go_slice*)tgt;
  
  return;
}
