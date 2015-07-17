
#ifndef GOEDM_GOEDMDICT_H
#define GOEDM_GOEDMDICT_H 1

namespace golang {
  struct gostring {
    int   Len;
    char *Data; //[Len]
  };

  template <typename T>
  struct goslice {
    int  Len;
    int  Cap;
    T   *Data; //[Len]
  };

}

namespace tmp {
  struct dict {
    golang::goslice<double> m_1;
    golang::goslice<float> m_2;
    golang::goslice<int> m_3;
    golang::gostring m_4;
  };
}

#endif
