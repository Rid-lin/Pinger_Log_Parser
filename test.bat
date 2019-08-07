@Echo Off
Set oldipforsite=%~dp0
For /D %%a In ("%oldipforsite:~0,-1%.txt") Do Set oldipforsite=%%~na

Echo %oldipforsite%