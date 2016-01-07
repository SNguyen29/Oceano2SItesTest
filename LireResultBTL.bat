cd src
src.exe --files=data/csp*.btl
ncdump output/OS_CASSIOPEE_BTL.nc > output/FichierResult.txt
notepad++.exe output/FichierResult.txt