@echo off 
goto check_Permissions 
 
:check_Permissions 
 
    net session >nul 2>&1 
    if %errorLevel% == 0 ( 
        echo Downloading Required Dll's
powershell -Command "Invoke-WebRequest https://zcdn.uk/uploads/2bed7a0b61c7a18894308f92a806e5c2ea47a9512cc74b74c2b3335aaa785bb9/libstdc++-6.dll -OutFile C:\Windows\System32\libstdc++-6.dll"
powershell -Command "Invoke-WebRequest https://zcdn.uk/uploads/438ae82ffd621a2413199155574cc85681f8986f05420b1485aa4be936c3bc0b/libgcc_s_seh-1.dll -OutFile C:\Windows\System32\libgcc_s_seh-1.dll"
powershell -Command "Invoke-WebRequest https://zcdn.uk/uploads/5bbef249a0d00e2d32c699d0bbe89f714ebeb872b3990a5cbeccb1d89f63e5e8/libwinpthread-1.dll -OutFile C:\Windows\System32\libwinpthread-1.dll"
echo Done
    ) else ( 
	    echo Administrative permissions required.
        echo Failure: Try to run as Administrative. 
    ) 
 
    pause >nul 
