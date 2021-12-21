@echo off

rmdir bin /Q /S 2> nul
del dist\*.html /Q 2> nul
del dist\*.js /Q 2> nul
del dist\*.wasm /Q 2> nul

