# Generated from trgen 0.21.2
set -e
if [ -f transformGrammar.py ]; then python3 transformGrammar.py ; fi
dotnet restore Test.csproj
dotnet build Test.csproj
exit 0
