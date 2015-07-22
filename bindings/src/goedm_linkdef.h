#if defined(__ROOTCLING__)
#pragma link C++ class golang::gostring;
#pragma link C++ class golang::goslice<double>;
#pragma link C++ class golang::goslice<float>;
#pragma link C++ class golang::goslice<int>;
#pragma link C++ class golang::goslice<golang::gostring>;
#endif

#ifdef __CINT__
#pragma link off all globals;
#pragma link off all classes;
#pragma link off all functions;
#pragma link C++ nestedclasses;

#pragma link C++ class golang::gostring+;
#pragma link C++ class golang::goslice< double>+;
#pragma link C++ class golang::goslice< float>+;
#pragma link C++ class golang::goslice< int>+;
#pragma link C++ class golang::goslice< golang::gostring>+;
#endif
