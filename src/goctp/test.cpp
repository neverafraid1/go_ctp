#include "ctpwrapper/TraderApi.h"
int main()
{
    TraderApi p;
    int i = 2;
    p.ReqConnect("", "tcp://180.168.212.79:31205", "", "8000", "41008389",
    "33905075", (THOST_TE_RESUME_TYPE)i, (THOST_TE_RESUME_TYPE)i, "" , "");
    return 0;
}