echo "(Optimizador) Ejecutando ChemsLexer..."
antlr4 -Dlanguage=Go -o parser ChemsLexer.g4
echo "(Optimizador) Ejecutando Chems..."
antlr4 -Dlanguage=Go -o parser Chems.g4