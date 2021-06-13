@echo off
set /p Input= Enter the path to the txt file (Without extension) that you wanna convert to Lua:
rename %Input%.txt %Input%.lua
pause