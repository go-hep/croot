#ifndef GOEDM_GOEDMDICT_H
#define GOEDM_GOEDMDICT_H 1

#ifndef __APPLE__
# include <stdint.h> // for (u)intXXX_t
#endif

namespace golang {

  struct gostring {
    int   Len;
    char *Data; //[Len]
  };

  template < class T >
	struct goslice{
		int Len;
		int Cap;
		T *Data; //[Len]
	};
}

namespace {
	golang::goslice< double> i_1;
	golang::goslice< float> i_2;
	golang::goslice< int> i_3;
	golang::goslice< unsigned int> i_4;
	golang::goslice< long> i_5;
	golang::goslice< unsigned long> i_6;
	golang::goslice< golang::gostring> i_7;
}

#endif
