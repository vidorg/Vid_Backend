echo Parsing yaml...
python ./gen_swagger.py -m main.go -o ./docs/api.yaml -e go

echo Parsing html...
python ./gen_swagger_html.py -i ./docs/api.yaml -o ./docs/api.html
