echo "Ejecutando ChemsLexer..."
antlr4 -Dlanguage=Go -o parser ChemsLexer.g4
echo "Ejecutando Chems..."
antlr4 -Dlanguage=Go -o parser Chems.g4